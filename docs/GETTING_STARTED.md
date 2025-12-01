# Eino ADK 入门阶段源码阅读文档

## 1. 概述

本文档旨在帮助开发者快速入门 Eino ADK 源码，重点介绍四个核心文件：

1. `interface.go` - 核心接口定义
2. `react.go` - ReAct 模式实现
3. `agent_tool.go` - 代理转工具功能
4. `runctx.go` - 运行上下文管理

通过阅读本文档，你将了解 ADK 的核心概念、设计思路和代码结构，为进一步学习打下坚实基础。

## 2. interface.go - 核心接口定义

### 2.1 文件定位

`interface.go` 是 ADK 的核心文件，定义了所有代理的基础接口和数据结构。

### 2.2 核心功能

- 定义 `Agent` 接口，所有代理必须实现
- 定义消息处理相关的数据结构
- 定义代理事件和动作
- 提供事件生成和处理的辅助函数

### 2.3 主要数据结构

#### 2.3.1 Agent 接口

```go
type Agent interface {
    Name(ctx context.Context) string
    Description(ctx context.Context) string
    Run(ctx context.Context, input *AgentInput, options ...AgentRunOption) *AsyncIterator[*AgentEvent]
}
```

这是所有代理的基础接口，包含三个方法：
- `Name` 和 `Description`：返回代理的名称和描述
- `Run`：执行代理，返回一个异步迭代器，用于接收代理执行过程中产生的事件

#### 2.3.2 消息相关结构

```go
type Message = *schema.Message
type MessageStream = *schema.StreamReader[Message]

type MessageVariant struct {
    IsStreaming bool
    Message       Message
    MessageStream MessageStream
    Role          schema.RoleType
    ToolName      string
}
```

- `Message`：单条消息
- `MessageStream`：消息流，用于流式输出
- `MessageVariant`：消息变体，同时支持流式和非流式消息

#### 2.3.3 代理事件和动作

```go
type AgentEvent struct {
    AgentName string
    RunPath []RunStep
    Output *AgentOutput
    Action *AgentAction
    Err error
}

type AgentAction struct {
    Exit bool
    Interrupted *InterruptInfo
    TransferToAgent *TransferToAgentAction
    BreakLoop *BreakLoopAction
    CustomizedAction any
}
```

- `AgentEvent`：代理执行过程中产生的事件
- `AgentAction`：代理可以执行的动作，如退出、中断、转移到其他代理等

### 2.4 关键函数

```go
func EventFromMessage(msg Message, msgStream MessageStream, role schema.RoleType, toolName string) *AgentEvent
func NewTransferToAgentAction(destAgentName string) *AgentAction
func NewExitAction() *AgentAction
```

这些函数用于创建和处理代理事件和动作。

### 2.5 阅读建议

1. 首先理解 `Agent` 接口的设计理念
2. 掌握消息处理机制，特别是流式和非流式消息的区别
3. 理解代理事件和动作的作用
4. 注意 `AsyncIterator` 的使用，这是 ADK 异步处理的核心

## 3. react.go - ReAct 模式实现

### 3.1 文件定位

`react.go` 实现了 ReAct（Reasoning and Acting）模式，这是一种让 LLM 思考并执行工具的模式。

### 3.2 核心功能

- 实现 ReAct 模式的核心逻辑
- 管理代理状态
- 处理工具调用和结果反馈
- 支持最大迭代次数限制

### 3.3 主要数据结构

#### 3.3.1 State 结构体

```go
type State struct {
    Messages []Message
    HasReturnDirectly        bool
    ReturnDirectlyToolCallID string
    ToolGenActions map[string]*AgentAction
    AgentName string
    AgentToolInterruptData map[string]*agentToolInterruptInfo
    RemainingIterations int
}
```

`State` 结构体用于管理 ReAct 代理的状态，包括：
- 消息历史
- 工具调用信息
- 剩余迭代次数
- 代理工具中断数据

#### 3.3.2 reactConfig 结构体

```go
type reactConfig struct {
    model model.ToolCallingChatModel
    toolsConfig *compose.ToolsNodeConfig
    toolsReturnDirectly map[string]bool
    agentName string
    maxIterations int
}
```

`reactConfig` 用于配置 ReAct 代理，包括：
- 聊天模型
- 工具配置
- 直接返回结果的工具列表
- 代理名称
- 最大迭代次数

### 3.4 关键函数

#### 3.4.1 newReact 函数

```go
func newReact(ctx context.Context, config *reactConfig) (reactGraph, error)
```

这是创建 ReAct 代理的核心函数，主要步骤：
1. 初始化状态生成函数
2. 创建图结构
3. 添加聊天模型节点
4. 添加工具节点
5. 构建节点之间的边和分支

#### 3.4.2 工具生成动作相关函数

```go
func SendToolGenAction(ctx context.Context, toolName string, action *AgentAction) error
func popToolGenAction(ctx context.Context, toolName string) *AgentAction
```

这些函数用于处理工具生成的动作，允许工具影响代理的执行流程。

### 3.5 代码流程

ReAct 模式的执行流程如下：

1. 接收输入消息
2. 将消息传递给聊天模型
3. 聊天模型生成思考内容或工具调用
4. 如果是工具调用，执行工具并获取结果
5. 将工具结果反馈给聊天模型
6. 重复步骤 3-5，直到生成最终答案或达到最大迭代次数

### 3.6 阅读建议

1. 理解 ReAct 模式的核心思想
2. 掌握 `State` 结构体的设计和使用
3. 分析 `newReact` 函数的实现，特别是图结构的构建
4. 理解工具调用和结果反馈的流程

## 4. agent_tool.go - 代理转工具功能

### 4.1 文件定位

`agent_tool.go` 提供了将代理包装为工具的功能，使代理可以被其他代理调用。

### 4.2 核心功能

- 将 `Agent` 接口转换为 `tool.BaseTool` 接口
- 处理代理输入和输出的转换
- 支持完整聊天历史作为输入
- 支持自定义输入 schema

### 4.3 主要数据结构

#### 4.3.1 AgentToolOptions 结构体

```go
type AgentToolOptions struct {
    fullChatHistoryAsInput bool
    agentInputSchema       *schema.ParamsOneOf
}
```

`AgentToolOptions` 用于配置代理工具的行为：
- `fullChatHistoryAsInput`：是否将完整聊天历史作为输入
- `agentInputSchema`：自定义输入 schema

#### 4.3.2 agentTool 结构体

```go
type agentTool struct {
    agent Agent
    fullChatHistoryAsInput bool
    inputSchema            *schema.ParamsOneOf
}
```

`agentTool` 是 `tool.BaseTool` 接口的实现，包装了一个 `Agent` 实例。

### 4.4 关键函数

#### 4.4.1 NewAgentTool 函数

```go
func NewAgentTool(_ context.Context, agent Agent, options ...AgentToolOption) tool.BaseTool
```

这是创建代理工具的入口函数，接受一个 `Agent` 实例和可选的配置选项。

#### 4.4.2 InvokableRun 方法

```go
func (at *agentTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error)
```

这是 `tool.BaseTool` 接口的核心方法，实现了工具的调用逻辑：
1. 解析输入参数
2. 准备代理输入
3. 执行代理
4. 处理代理输出
5. 返回结果

### 4.5 代码流程

代理转工具的执行流程如下：

1. 调用 `NewAgentTool` 创建代理工具
2. 其他代理调用该工具时，执行 `InvokableRun` 方法
3. `InvokableRun` 方法解析输入，准备 `AgentInput`
4. 调用代理的 `Run` 方法执行代理
5. 处理代理返回的 `AgentEvent`，提取结果
6. 返回结果给调用者

### 4.6 阅读建议

1. 理解代理和工具的区别和联系
2. 掌握 `NewAgentTool` 函数的实现
3. 分析 `InvokableRun` 方法的执行流程
4. 理解输入输出转换的逻辑

## 5. runctx.go - 运行上下文管理

### 5.1 文件定位

`runctx.go` 管理代理运行时的上下文信息，包括会话状态、运行路径等。

### 5.2 核心功能

- 管理代理运行时的上下文信息
- 提供会话状态管理
- 支持运行路径追踪
- 处理中断信息

### 5.3 主要数据结构

#### 5.3.1 runContext 结构体

```go
type runContext struct {
    RunPath   []RunStep
    Session   *session
    RootInput *AgentInput
}
```

`runContext` 用于管理代理的运行上下文，包括：
- 运行路径：记录代理调用链
- 会话：管理会话状态和事件
- 根输入：原始输入信息

#### 5.3.2 session 结构体

```go
type session struct {
    events []*AgentEvent
    mu     sync.RWMutex
}
```

`session` 用于管理会话状态，包括事件历史。

### 5.4 关键函数

#### 5.4.1 initRunCtx 函数

```go
func initRunCtx(ctx context.Context, agentName string, input *AgentInput) (context.Context, *runContext)
```

这是初始化运行上下文的函数，主要步骤：
1. 创建 `runContext` 实例
2. 初始化运行路径
3. 创建会话
4. 保存根输入
5. 将上下文添加到 `context.Context` 中

#### 5.4.2 getRunCtx 函数

```go
func getRunCtx(ctx context.Context) *runContext
```

从 `context.Context` 中获取运行上下文。

#### 5.4.3 addEvent 方法

```go
func (s *session) addEvent(event *AgentEvent)
```

将事件添加到会话历史中。

### 5.5 代码流程

运行上下文管理的流程如下：

1. 代理执行前，调用 `initRunCtx` 初始化运行上下文
2. 运行上下文被添加到 `context.Context` 中，随请求传递
3. 代理执行过程中，通过 `getRunCtx` 获取上下文
4. 代理产生事件时，调用 `addEvent` 添加到会话历史
5. 代理之间转移时，更新运行路径

### 5.6 阅读建议

1. 理解运行上下文的作用和设计理念
2. 掌握 `runContext` 和 `session` 结构体的设计
3. 分析 `initRunCtx` 和 `getRunCtx` 函数的实现
4. 理解会话状态管理的机制

## 6. 代码关系图

```
┌─────────────────┐     ┌─────────────────┐
│   interface.go  │     │   runctx.go     │
│                 │     │                 │
│  ┌───────────┐  │     │  ┌───────────┐  │
│  │  Agent    │  │     │  │ runContext│  │
│  └───────────┘  │     │  └───────────┘  │
│        ▲        │     │        ▲        │
│        │        │     │        │        │
└────────┼────────┘     └────────┼────────┘
         │                       │
         ▼                       ▼
┌─────────────────┐     ┌─────────────────┐
│    react.go     │     │  agent_tool.go  │
│                 │     │                 │
│  ┌───────────┐  │     │  ┌───────────┐  │
│  │  State    │  │     │  │ agentTool │  │
│  └───────────┘  │     │  └───────────┘  │
└─────────────────┘     └─────────────────┘
```

## 7. 学习建议

1. **顺序阅读**：按照 `interface.go` → `react.go` → `agent_tool.go` → `runctx.go` 的顺序阅读
2. **重点理解**：
   - `Agent` 接口的设计理念
   - ReAct 模式的执行流程
   - 代理转工具的实现机制
   - 运行上下文的管理方式
3. **结合测试**：阅读每个文件时，参考对应的测试文件（如 `interface_test.go`），了解如何使用这些组件
4. **动手实践**：尝试实现一个简单的代理，加深对核心概念的理解
5. **绘制流程图**：自己绘制代码执行流程图，帮助理解复杂逻辑

## 8. 下一步学习

完成本入门阶段的学习后，你可以继续学习：

1. `flow.go` - 代理工作流管理
2. `workflow.go` - 工作流实现
3. `prebuilt/` 目录下的预构建代理
4. 其他辅助文件和测试文件

## 9. 总结

本文档详细介绍了 Eino ADK 入门阶段的四个核心文件，包括它们的功能、数据结构、关键函数和代码流程。通过阅读本文档，你应该已经掌握了 ADK 的核心概念和设计思路，为进一步学习打下了坚实基础。

建议你在阅读源码时，结合本文档的说明，逐行理解代码的含义和设计意图。同时，多参考测试文件，了解如何使用这些组件，这将有助于你更好地掌握 ADK 的使用方法。

祝你学习愉快！