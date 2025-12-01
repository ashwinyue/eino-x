/*
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package adk

import (
	"context"
	"errors"
	"io"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

// ErrExceedMaxIterations 表示超过最大迭代次数的错误
var ErrExceedMaxIterations = errors.New("exceeds max iterations")

// State 表示Agent的状态信息
// 包含对话消息、工具调用状态、迭代次数等

type State struct {
	// Messages 存储对话历史消息
	Messages []Message

	// HasReturnDirectly 标记是否直接返回工具调用结果
	HasReturnDirectly bool
	// ReturnDirectlyToolCallID 直接返回的工具调用ID
	ReturnDirectlyToolCallID string

	// ToolGenActions 存储工具生成的动作映射
	ToolGenActions map[string]*AgentAction

	// AgentName Agent的名称
	AgentName string

	// RemainingIterations 剩余迭代次数
	RemainingIterations int
}

// SendToolGenAction 发送工具生成的动作到状态中
// 参数：
//   - ctx: 上下文
//   - toolName: 工具名称
//   - action: 工具生成的动作
//
// 返回：
//   - error: 执行过程中的错误
func SendToolGenAction(ctx context.Context, toolName string, action *AgentAction) error {
	return compose.ProcessState(ctx, func(ctx context.Context, st *State) error {
		st.ToolGenActions[toolName] = action

		return nil
	})
}

// popToolGenAction 从状态中获取并移除指定工具的动作
// 参数：
//   - ctx: 上下文
//   - toolName: 工具名称
//
// 返回：
//   - *AgentAction: 工具动作，如果不存在则返回nil
func popToolGenAction(ctx context.Context, toolName string) *AgentAction {
	var action *AgentAction
	err := compose.ProcessState(ctx, func(ctx context.Context, st *State) error {
		action = st.ToolGenActions[toolName]
		if action != nil {
			delete(st.ToolGenActions, toolName)
		}

		return nil
	})

	if err != nil {
		panic("impossible")
	}

	return action
}

// reactConfig 表示React Agent的配置信息
// 包含模型、工具配置、直接返回工具等

type reactConfig struct {
	// model 工具调用聊天模型
	model model.ToolCallingChatModel

	// toolsConfig 工具节点配置
	toolsConfig *compose.ToolsNodeConfig

	// toolsReturnDirectly 标记哪些工具需要直接返回结果
	toolsReturnDirectly map[string]bool

	// agentName Agent名称
	agentName string

	// maxIterations 最大迭代次数
	maxIterations int

	// beforeChatModel 调用聊天模型前的处理函数列表
	// afterChatModel 调用聊天模型后的处理函数列表
	beforeChatModel, afterChatModel []func(context.Context, *ChatModelAgentState) error
}

// genToolInfos 生成工具信息列表
// 参数：
//   - ctx: 上下文
//   - config: 工具节点配置
//
// 返回：
//   - []*schema.ToolInfo: 工具信息列表
//   - error: 生成过程中的错误
func genToolInfos(ctx context.Context, config *compose.ToolsNodeConfig) ([]*schema.ToolInfo, error) {
	toolInfos := make([]*schema.ToolInfo, 0, len(config.Tools))
	for _, t := range config.Tools {
		tl, err := t.Info(ctx)
		if err != nil {
			return nil, err
		}

		toolInfos = append(toolInfos, tl)
	}

	return toolInfos, nil
}

// reactGraph 定义React图的类型别名
// 输入为消息列表，输出为单条消息

type reactGraph = *compose.Graph[[]Message, Message]

// sToolNodeOutput 定义工具节点输出的类型别名
// 表示消息列表的流读取器

type sToolNodeOutput = *schema.StreamReader[[]Message]

// sGraphOutput 定义图输出的类型别名
// 表示消息流

type sGraphOutput = MessageStream

// getReturnDirectlyToolCallID 获取直接返回的工具调用ID
// 参数：
//   - ctx: 上下文
//
// 返回：
//   - string: 工具调用ID
//   - bool: 是否需要直接返回
func getReturnDirectlyToolCallID(ctx context.Context) (string, bool) {
	var toolCallID string
	var hasReturnDirectly bool
	handler := func(_ context.Context, st *State) error {
		toolCallID = st.ReturnDirectlyToolCallID
		hasReturnDirectly = st.HasReturnDirectly
		return nil
	}

	_ = compose.ProcessState(ctx, handler)

	return toolCallID, hasReturnDirectly
}

// newReact 创建一个React图实例
// 参数：
//   - ctx: 上下文
//   - config: React配置信息
//
// 返回：
//   - reactGraph: React图实例
//   - error: 创建过程中的错误
//
// 该函数构建一个React Agent图，包含聊天模型节点、工具节点和相应的分支条件
func newReact(ctx context.Context, config *reactConfig) (reactGraph, error) {
	// genState 生成Agent状态的函数
	// 初始化状态，设置工具动作映射、Agent名称和剩余迭代次数
	genState := func(ctx context.Context) *State {
		return &State{
			ToolGenActions: map[string]*AgentAction{},
			AgentName:      config.agentName,
			RemainingIterations: func() int {
				// 如果未设置最大迭代次数，默认使用20次
				if config.maxIterations <= 0 {
					return 20
				}
				return config.maxIterations
			}(),
		}
	}

	// 定义节点名称常量
	const (
		chatModel_ = "ChatModel" // 聊天模型节点名称
		toolNode_  = "ToolNode"  // 工具节点名称
	)

	// 创建新的图实例，设置状态生成函数
	g := compose.NewGraph[[]Message, Message](compose.WithGenLocalState(genState))

	// 生成工具信息列表
	toolsInfo, err := genToolInfos(ctx, config.toolsConfig)
	if err != nil {
		return nil, err
	}

	// 为聊天模型配置工具
	chatModel, err := config.model.WithTools(toolsInfo)
	if err != nil {
		return nil, err
	}

	// 创建工具节点
	toolsNode, err := compose.NewToolNode(ctx, config.toolsConfig)
	if err != nil {
		return nil, err
	}

	// modelPreHandle 聊天模型节点的预处理函数
	// 检查剩余迭代次数，执行beforeChatModel处理函数
	modelPreHandle := func(ctx context.Context, input []Message, st *State) ([]Message, error) {
		// 检查是否超过最大迭代次数
		if st.RemainingIterations <= 0 {
			return nil, ErrExceedMaxIterations
		}
		// 减少剩余迭代次数
		st.RemainingIterations--

		// 创建聊天模型Agent状态，执行所有beforeChatModel处理函数
		s := &ChatModelAgentState{Messages: append(st.Messages, input...)}
		for _, b := range config.beforeChatModel {
			err = b(ctx, s)
			if err != nil {
				return nil, err
			}
		}
		// 更新状态中的消息
		st.Messages = s.Messages

		return st.Messages, nil
	}

	// modelPostHandle 聊天模型节点的后处理函数
	// 执行afterChatModel处理函数
	modelPostHandle := func(ctx context.Context, input Message, st *State) (Message, error) {
		// 创建聊天模型Agent状态，执行所有afterChatModel处理函数
		s := &ChatModelAgentState{Messages: append(st.Messages, input)}
		for _, a := range config.afterChatModel {
			err = a(ctx, s)
			if err != nil {
				return nil, err
			}
		}
		// 更新状态中的消息
		st.Messages = s.Messages
		return input, nil
	}

	// 添加聊天模型节点到图中，配置预处理和后处理函数
	_ = g.AddChatModelNode(chatModel_, chatModel,
		compose.WithStatePreHandler(modelPreHandle), compose.WithStatePostHandler(modelPostHandle), compose.WithNodeName(chatModel_))

	// toolPreHandle 工具节点的预处理函数
	// 检查工具调用是否需要直接返回结果
	toolPreHandle := func(ctx context.Context, input Message, st *State) (Message, error) {
		// 获取最新的消息
		input = st.Messages[len(st.Messages)-1]
		// 检查是否有工具需要直接返回结果
		if len(config.toolsReturnDirectly) > 0 {
			for i := range input.ToolCalls {
				toolName := input.ToolCalls[i].Function.Name
				// 如果工具配置为直接返回，设置相应标记
				if config.toolsReturnDirectly[toolName] {
					st.ReturnDirectlyToolCallID = input.ToolCalls[i].ID
					st.HasReturnDirectly = true
				}
			}
		}

		return input, nil
	}

	// 添加工具节点到图中，配置预处理函数
	_ = g.AddToolsNode(toolNode_, toolsNode,
		compose.WithStatePreHandler(toolPreHandle), compose.WithNodeName(toolNode_))

	// 添加从START到聊天模型节点的边
	_ = g.AddEdge(compose.START, chatModel_)

	// toolCallCheck 检查聊天模型输出是否包含工具调用
	// 如果包含工具调用，跳转到工具节点；否则结束
	toolCallCheck := func(ctx context.Context, sMsg MessageStream) (string, error) {
		defer sMsg.Close()
		for {
			// 接收消息流中的下一个消息
			chunk, err_ := sMsg.Recv()
			if err_ != nil {
				// 如果流结束，返回END
				if err_ == io.EOF {
					return compose.END, nil
				}
				return "", err_
			}

			// 如果消息包含工具调用，返回工具节点名称
			if len(chunk.ToolCalls) > 0 {
				return toolNode_, nil
			}
		}
	}

	// 创建聊天模型节点的分支条件
	// 根据是否有工具调用，决定跳转到工具节点或结束
	branch := compose.NewStreamGraphBranch(toolCallCheck, map[string]bool{compose.END: true, toolNode_: true})
	_ = g.AddBranch(chatModel_, branch)

	// 根据是否配置了直接返回工具，设置不同的边和分支
	if len(config.toolsReturnDirectly) == 0 {
		// 如果没有直接返回工具，添加从工具节点到聊天模型节点的边
		// 形成循环：聊天模型 → 工具 → 聊天模型
		_ = g.AddEdge(toolNode_, chatModel_)
	} else {
		// 如果有直接返回工具，添加工具节点到结束的转换节点
		const (
			toolNodeToEndConverter = "ToolNodeToEndConverter" // 工具节点到结束的转换节点名称
		)

		// cvt 将工具节点输出转换为图输出
		// 只返回指定工具调用ID的结果
		cvt := func(ctx context.Context, sToolCallMessages sToolNodeOutput) (sGraphOutput, error) {
			// 获取直接返回的工具调用ID
			id, _ := getReturnDirectlyToolCallID(ctx)

			// 创建流转换器，只返回匹配ID的消息
			return schema.StreamReaderWithConvert(sToolCallMessages,
				func(in []Message) (Message, error) {
					for _, chunk := range in {
						if chunk != nil && chunk.ToolCallID == id {
							return chunk, nil
						}
					}
					// 没有匹配的消息，返回ErrNoValue
					return nil, schema.ErrNoValue
				}), nil
		}

		// 添加转换节点到图中
		_ = g.AddLambdaNode(toolNodeToEndConverter, compose.TransformableLambda(cvt),
			compose.WithNodeName(toolNodeToEndConverter))
		// 添加从转换节点到END的边
		_ = g.AddEdge(toolNodeToEndConverter, compose.END)

		// checkReturnDirect 检查是否需要直接返回结果
		// 根据HasReturnDirectly标记，决定跳转到转换节点或聊天模型节点
		checkReturnDirect := func(ctx context.Context,
			sToolCallMessages sToolNodeOutput) (string, error) {

			_, ok := getReturnDirectlyToolCallID(ctx)

			// 如果需要直接返回，跳转到转换节点；否则返回聊天模型节点
			if ok {
				return toolNodeToEndConverter, nil
			}

			return chatModel_, nil
		}

		// 创建工具节点的分支条件
		// 根据是否需要直接返回，决定跳转到转换节点或聊天模型节点
		branch = compose.NewStreamGraphBranch(checkReturnDirect,
			map[string]bool{toolNodeToEndConverter: true, chatModel_: true})
		_ = g.AddBranch(toolNode_, branch)
	}

	// 返回构建好的React图实例
	return g, nil
}
