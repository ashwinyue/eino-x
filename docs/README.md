# Eino ADK 源码学习指南

## 1. 概述

ADK（Agent Development Kit）是 Eino 框架的核心组件之一，用于开发和管理智能代理（Agent）。它提供了代理的核心接口、工作流管理、工具集成和状态管理等功能，是构建复杂 AI 应用的基础。

## 2. 目录结构

```
adk/
├── prebuilt/           # 预构建的代理实现
│   ├── deep/           # Deep 代理，用于深度推理和任务分解
│   ├── planexecute/    # 计划执行代理，用于执行复杂的计划
│   └── supervisor/     # 监督者代理，用于管理和协调其他代理
├── agent_tool.go       # 将代理包装为工具
├── chatmodel.go        # 聊天模型相关
├── flow.go             # 代理工作流管理
├── interface.go        # 核心接口定义
├── react.go            # ReAct 模式实现
├── runctx.go           # 运行上下文管理
├── runner.go           # 代理运行器
├── workflow.go         # 工作流实现
└── 其他辅助文件和测试文件
```

## 3. 核心概念

### 3.1 Agent 接口

`interface.go` 中定义了 `Agent` 接口，这是所有代理的基础：

```go
type Agent interface {
    Name(ctx context.Context) string
    Description(ctx context.Context) string
    Run(ctx context.Context, input *AgentInput, options ...AgentRunOption) *AsyncIterator[*AgentEvent]
}
```

- `Name` 和 `Description`：返回代理的名称和描述
- `Run`：执行代理，返回一个异步迭代器，用于接收代理执行过程中产生的事件

### 3.2 消息处理

ADK 支持两种消息模式：

- **非流式消息**：一次性返回完整消息
- **流式消息**：分块返回消息，适合实时交互场景

`MessageVariant` 结构体用于处理这两种消息模式：

```go
type MessageVariant struct {
    IsStreaming bool
    Message       Message
    MessageStream MessageStream
    Role          schema.RoleType
    ToolName      string
}
```

### 3.3 代理事件和动作

- **AgentEvent**：代理执行过程中产生的事件，包含输出和动作
- **AgentAction**：代理可以执行的动作，如退出、中断、转移到其他代理等

```go
type AgentAction struct {
    Exit bool
    Interrupted *InterruptInfo
    TransferToAgent *TransferToAgentAction
    BreakLoop *BreakLoopAction
    CustomizedAction any
}
```

### 3.4 ReAct 模式

`react.go` 实现了 ReAct（Reasoning and Acting）模式，这是一种让 LLM 思考并执行工具的模式：

1. LLM 接收输入并生成思考内容
2. 如果需要执行工具，生成工具调用
3. 执行工具并获取结果
4. 将结果反馈给 LLM，继续思考或生成最终答案

### 3.5 代理工作流

`flow.go` 管理代理的工作流，支持：

- 父子代理关系
- 代理之间的转移
- 历史记录重写
- 子代理管理

### 3.6 运行上下文

`runctx.go` 管理代理运行时的上下文信息，包括：

- 会话状态
- 运行路径
- 中断信息
- 其他运行时数据

## 4. 核心组件详解

### 4.1 代理工具转换

`agent_tool.go` 提供了将代理包装为工具的功能，使代理可以被其他代理调用：

```go
func NewAgentTool(_ context.Context, agent Agent, options ...AgentToolOption) tool.BaseTool
```

这允许代理之间的嵌套调用，构建更复杂的代理系统。

### 4.2 工作流代理

`workflow.go` 实现了工作流代理，支持：

- 定义复杂的工作流
- 状态管理
- 条件分支
- 并行执行

### 4.3 预构建代理

`prebuilt/` 目录包含了一些预构建的代理实现：

#### 4.3.1 Deep 代理

`deep/` 目录实现了 Deep 代理，用于深度推理和任务分解：

- 支持复杂任务的分解和执行
- 内置 TODO 管理功能
- 支持子代理调用
- 提供任务工具生成功能

#### 4.3.2 PlanExecute 代理

`planexecute/` 目录实现了计划执行代理：

- 用于执行复杂的计划
- 支持计划的生成和执行
- 提供计划管理功能

#### 4.3.3 Supervisor 代理

`supervisor/` 目录实现了监督者代理：

- 用于管理和协调其他代理
- 支持代理之间的通信和转移
- 提供代理生命周期管理

## 5. 学习路径

### 5.1 入门阶段

1. 首先阅读 `interface.go`，理解核心接口定义
2. 阅读 `react.go`，理解 ReAct 模式的实现
3. 查看 `agent_tool.go`，了解代理如何转换为工具
4. 学习 `runctx.go`，了解运行上下文管理

### 5.2 进阶阶段

1. 学习 `flow.go`，理解代理工作流管理
2. 阅读 `workflow.go`，了解复杂工作流的实现
3. 研究 `prebuilt/` 目录下的预构建代理：
   - 先学习 `planexecute/`，了解计划执行代理
   - 再学习 `supervisor/`，了解监督者代理
   - 最后学习 `deep/`，了解深度推理代理
4. 查看测试文件，了解如何使用和测试代理

### 5.3 实践阶段

1. 尝试实现一个简单的代理
2. 使用 ReAct 模式构建一个可以调用工具的代理
3. 实现代理之间的转移和通信
4. 构建一个复杂的工作流代理
5. 使用 Deep 代理实现复杂任务的分解和执行

## 6. 关键 API 和函数

### 6.1 Agent 相关

- `Agent` 接口：所有代理的基础
- `NewAgentTool`：将代理转换为工具
- `SetSubAgents`：设置子代理
- `AgentWithOptions`：创建带有选项的代理
- `NewChatModelAgent`：创建聊天模型代理

### 6.2 ReAct 模式

- `newReact`：创建 ReAct 代理
- `SendToolGenAction`：发送工具生成动作
- `popToolGenAction`：获取工具生成动作

### 6.3 工作流

- `flowAgent`：工作流代理实现
- `run`：执行代理工作流
- `genAgentInput`：生成代理输入

### 6.4 Deep 代理

- `deep.New`：创建 Deep 代理
- `Config`：Deep 代理配置
- `newWriteTodos`：创建 TODO 管理工具
- `newTaskToolMiddleware`：创建任务工具中间件

### 6.5 运行上下文

- `initRunCtx`：初始化运行上下文
- `getRunCtx`：获取运行上下文
- `addEvent`：添加事件到运行上下文

## 7. 最佳实践

1. **接口优先**：始终通过接口定义代理，而不是具体实现
2. **模块化设计**：将复杂代理拆分为多个简单代理，通过工作流组合
3. **使用 ReAct 模式**：对于需要调用工具的代理，优先使用 ReAct 模式
4. **状态管理**：合理使用运行上下文管理状态，避免全局状态
5. **测试驱动开发**：编写充分的测试用例，确保代理的正确性
6. **利用预构建代理**：对于常见场景，优先使用预构建代理，避免重复造轮子
7. **合理使用中间件**：通过中间件扩展代理功能，保持核心逻辑简洁

## 8. 示例代码

### 8.1 创建一个简单代理

```go
type SimpleAgent struct {
    name        string
    description string
}

func (sa *SimpleAgent) Name(ctx context.Context) string {
    return sa.name
}

func (sa *SimpleAgent) Description(ctx context.Context) string {
    return sa.description
}

func (sa *SimpleAgent) Run(ctx context.Context, input *AgentInput, options ...AgentRunOption) *AsyncIterator[*AgentEvent] {
    // 实现代理逻辑
}
```

### 8.2 使用 ReAct 模式

```go
config := &reactConfig{
    model:          chatModel,
    toolsConfig:    toolsConfig,
    agentName:      "react-agent",
    maxIterations:  10,
}

graph, err := newReact(ctx, config)
if err != nil {
    return nil, err
}
```

### 8.3 创建 Deep 代理

```go
config := &deep.Config{
    Name:          "deep-agent",
    Description:   "深度推理代理",
    ChatModel:     chatModel,
    Instruction:   "你是一个智能代理，能够分解和执行复杂任务",
    SubAgents:     subAgents,
    ToolsConfig:   toolsConfig,
    MaxIteration:  20,
}

deepAgent, err := deep.New(ctx, config)
if err != nil {
    return nil, err
}
```

## 9. 常见问题

### 9.1 如何处理流式消息？

使用 `MessageVariant` 结构体，它同时支持流式和非流式消息。对于流式消息，可以使用 `schema.ConcatMessageStream` 函数将流转换为完整消息。

### 9.2 如何实现代理之间的转移？

使用 `NewTransferToAgentAction` 函数创建转移动作，然后在代理的 `Run` 方法中返回包含该动作的 `AgentEvent`。

### 9.3 如何将代理转换为工具？

使用 `NewAgentTool` 函数，它将代理包装为工具，使其可以被其他代理调用。

### 9.4 如何使用 Deep 代理分解复杂任务？

1. 创建 Deep 代理实例，配置聊天模型和工具
2. 添加子代理（可选）
3. 调用 Deep 代理的 `Run` 方法，传入复杂任务
4. Deep 代理会自动分解任务并执行

## 10. 进一步学习资源

- [Eino 官方文档](https://www.cloudwego.io/zh/docs/eino/)
- [Eino 示例代码](https://github.com/cloudwego/eino-examples)
- [Eino 扩展库](https://github.com/cloudwego/eino-ext)

## 11. 总结

ADK 是 Eino 框架的核心组件，提供了构建智能代理的基础功能。通过学习 ADK 的源码，你可以深入理解代理的工作原理，掌握构建复杂 AI 应用的技能。

ADK 包含了多种预构建代理，从简单的计划执行代理到复杂的深度推理代理，可以满足不同场景的需求。建议按照学习路径逐步深入，从核心接口开始，然后学习各种功能模块，最后通过实践来巩固所学知识。

祝你学习愉快！