# React 模式介绍

## 1. 什么是 React 模式

React 模式是一种基于图（Graph）的智能代理执行模式，用于构建能够自主决策、调用工具并处理结果的智能代理。它实现了一个循环执行流程：

1. **思考**：代理根据对话历史生成思考和决策
2. **行动**：如果需要，代理调用外部工具
3. **观察**：代理接收并处理工具执行结果
4. **循环**：根据处理结果决定是否继续执行或结束

React 模式的核心思想是让代理能够在对话和工具调用之间自主循环，直到完成任务或达到最大迭代次数。

## 2. 核心组件

### 2.1 状态管理

React 模式使用 `State` 结构体管理代理的运行状态：

```go
type State struct {
    Messages []Message                // 对话历史消息
    HasReturnDirectly bool            // 是否直接返回工具调用结果
    ReturnDirectlyToolCallID string   // 直接返回的工具调用ID
    ToolGenActions map[string]*AgentAction // 工具生成的动作映射
    AgentName string                  // Agent的名称
    RemainingIterations int           // 剩余迭代次数
}
```

### 2.2 图结构

React 模式基于图结构构建，主要包含两个核心节点：

1. **聊天模型节点** (`ChatModel`)：负责生成思考和决策
2. **工具节点** (`ToolNode`)：负责执行工具调用

### 2.3 配置选项

```go
type reactConfig struct {
    model model.ToolCallingChatModel  // 工具调用聊天模型
    toolsConfig *compose.ToolsNodeConfig // 工具节点配置
    toolsReturnDirectly map[string]bool // 标记哪些工具需要直接返回结果
    agentName string                   // Agent名称
    maxIterations int                  // 最大迭代次数
    beforeChatModel, afterChatModel []func(context.Context, *ChatModelAgentState) error // 聊天模型前后处理函数
}
```

## 3. 工作流程

### 3.1 基本执行流程

1. **初始化**：创建 React 图，设置状态生成函数和最大迭代次数
2. **添加节点**：将聊天模型节点和工具节点添加到图中
3. **添加边**：建立节点之间的连接关系
4. **执行循环**：
   - 聊天模型生成思考和决策
   - 如果需要调用工具，执行工具调用
   - 处理工具执行结果
   - 决定是否继续循环或结束

### 3.2 详细流程图

```
START
  │
  ▼
ChatModel
  │
  ├─┬─ 有工具调用？ ── 是 ── 工具节点
  │ │
  │ ▼
  │ 工具执行
  │  │
  │  ├─┬─ 需要直接返回？ ── 是 ── END
  │  │ │
  │  │ ▼
  │  │ 转换节点
  │  │  │
  │  │  ▼
  │  │ END
  │  │
  │  └── 否 ── 聊天模型（循环）
  │
  └── 否 ── END
```

## 4. 状态管理机制

### 4.1 状态初始化

```go
genState := func(ctx context.Context) *State {
    return &State{
        ToolGenActions: map[string]*AgentAction{},
        AgentName:      config.agentName,
        RemainingIterations: func() int {
            if config.maxIterations <= 0 {
                return 20 // 默认最大迭代次数
            }
            return config.maxIterations
        }(),
    }
}
```

### 4.2 状态更新

- **迭代次数管理**：每次执行聊天模型前减少剩余迭代次数
- **工具调用跟踪**：记录工具调用ID和是否需要直接返回
- **工具生成动作**：存储工具执行过程中生成的动作

### 4.3 工具生成动作

```go
// 发送工具生成的动作到状态中
func SendToolGenAction(ctx context.Context, toolName string, action *AgentAction) error {
    return compose.ProcessState(ctx, func(ctx context.Context, st *State) error {
        st.ToolGenActions[toolName] = action
        return nil
    })
}
```

## 5. 工具调用机制

### 5.1 工具配置

React 模式支持配置多种工具，并可以指定哪些工具需要直接返回结果：

```go
toolsReturnDirectly map[string]bool // 标记哪些工具需要直接返回结果
```

### 5.2 工具调用处理

1. 聊天模型生成工具调用请求
2. 工具节点执行工具调用
3. 根据配置决定是否直接返回结果或继续循环
4. 处理工具执行结果并更新状态

### 5.3 直接返回机制

对于某些工具，可能需要直接返回结果而不继续循环：

```go
// 检查工具调用是否需要直接返回结果
toolPreHandle := func(ctx context.Context, input Message, st *State) (Message, error) {
    // ...
    if config.toolsReturnDirectly[toolName] {
        st.ReturnDirectlyToolCallID = input.ToolCalls[i].ID
        st.HasReturnDirectly = true
    }
    // ...
}
```

## 6. 配置选项

### 6.1 基本配置

| 配置项 | 类型 | 描述 |
|--------|------|------|
| model | model.ToolCallingChatModel | 工具调用聊天模型 |
| toolsConfig | *compose.ToolsNodeConfig | 工具节点配置 |
| toolsReturnDirectly | map[string]bool | 标记哪些工具需要直接返回结果 |
| agentName | string | Agent名称 |
| maxIterations | int | 最大迭代次数 |
| beforeChatModel | []func | 调用聊天模型前的处理函数列表 |
| afterChatModel | []func | 调用聊天模型后的处理函数列表 |

### 6.2 处理函数

React 模式支持在聊天模型调用前后添加自定义处理函数：

- **beforeChatModel**：在调用聊天模型前执行，可用于修改输入消息或状态
- **afterChatModel**：在调用聊天模型后执行，可用于处理输出消息或状态

## 7. 错误处理

### 7.1 最大迭代次数限制

当超过最大迭代次数时，React 模式会返回 `ErrExceedMaxIterations` 错误：

```go
var ErrExceedMaxIterations = errors.New("exceeds max iterations")
```

### 7.2 状态检查

在每次执行聊天模型前，会检查剩余迭代次数：

```go
if st.RemainingIterations <= 0 {
    return nil, ErrExceedMaxIterations
}
```

## 8. 使用示例

### 8.1 创建 React 代理

```go
// 1. 创建聊天模型
chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
    Model: "gpt-4o",
})
if err != nil {
    return err
}

// 2. 创建工具配置
toolsConfig := &compose.ToolsNodeConfig{
    Tools: []compose.Tool{
        // 添加你的工具
    },
}

// 3. 创建 React 配置
reactConfig := &reactConfig{
    model: chatModel,
    toolsConfig: toolsConfig,
    agentName: "MyReactAgent",
    maxIterations: 10,
}

// 4. 创建 React 图
reactGraph, err := newReact(ctx, reactConfig)
if err != nil {
    return err
}

// 5. 运行 React 代理
result := reactGraph.Run(ctx, []Message{
    schema.UserMessage("Hello, please help me with this task..."),
})
```

### 8.2 配置直接返回工具

```go
reactConfig := &reactConfig{
    // ... 其他配置
    toolsReturnDirectly: map[string]bool{
        "search_tool": true, // search_tool 的结果将直接返回
    },
}
```

## 9. 最佳实践

1. **合理设置最大迭代次数**：根据任务复杂度调整，避免无限循环
2. **选择合适的工具**：根据任务需求选择必要的工具
3. **配置直接返回工具**：对于不需要后续处理的工具，配置为直接返回
4. **添加自定义处理函数**：根据需要在聊天模型前后添加处理逻辑
5. **监控状态变化**：通过状态管理跟踪代理执行过程

## 10. 与其他模式的比较

| 特性 | React 模式 | 普通聊天模式 |
|------|------------|--------------|
| 自主决策 | ✅ | ❌ |
| 工具调用 | ✅ | ❌ |
| 循环执行 | ✅ | ❌ |
| 状态管理 | ✅ | ❌ |
| 结果处理 | ✅ | ❌ |

React 模式提供了更强大的自主执行能力，适用于需要复杂决策和工具调用的场景。

## 11. 总结

React 模式是一种强大的智能代理执行模式，通过图结构和状态管理实现了自主思考、行动和观察的循环。它适用于构建需要复杂决策、多轮工具调用和结果处理的智能代理。

主要优势：

1. **自主循环执行**：无需外部干预，自主完成任务
2. **灵活的工具调用**：支持多种工具和调用方式
3. **强大的状态管理**：完整跟踪执行过程和状态变化
4. **可配置的执行流程**：支持自定义处理函数和配置选项
5. **基于图结构**：易于扩展和修改执行流程

React 模式为构建智能代理提供了一种高效、灵活的解决方案，适用于各种复杂的AI应用场景。