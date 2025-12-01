# Eino 源码学习计划

## 1. 概述

Eino 是一个用于构建 LLM 应用的 Golang 框架，提供了组件抽象、组合框架、API 和工作流等功能。本学习计划将帮助你系统地学习 Eino 的源码，从核心概念到高级功能，逐步深入理解框架的设计和实现。

## 2. 学习路径

### 2.1 核心数据模型（schema）

**重要性**：★★★★★

schema 包定义了 Eino 的核心数据模型，是其他所有模块的基础。学习 schema 包可以帮助你理解框架中数据的表示和流动方式。

**学习内容**：
- 消息模型（Message）
- 工具调用（ToolCall）
- 流式处理（Stream）
- 多模态内容（MultiContent）

**学习顺序**：
1. `schema/message.go` - 核心消息定义
2. `schema/stream.go` - 流式处理
3. `schema/tool.go` - 工具调用
4. `schema/document.go` - 文档处理

### 2.2 组件抽象（components）

**重要性**：★★★★☆

components 包定义了各种组件的抽象接口，如 ChatModel、Tool、Retriever 等。学习 components 包可以帮助你理解框架中组件的设计原则和使用方式。

**学习内容**：
- 组件类型和接口
- 回调机制
- 组件选项

**学习顺序**：
1. `components/types.go` - 组件类型定义
2. `components/model/interface.go` - 聊天模型接口
3. `components/tool/interface.go` - 工具接口
4. `components/retriever/interface.go` - 检索器接口
5. `components/embedding/interface.go` - 嵌入接口
6. `components/indexer/interface.go` - 索引器接口
7. `components/prompt/interface.go` - 提示模板接口
8. `components/document/interface.go` - 文档处理接口

### 2.3 组合框架（compose）

**重要性**：★★★★★

compose 包提供了强大的组合框架，用于构建复杂的工作流和图。学习 compose 包可以帮助你理解框架中如何组合和编排各种组件。

**学习内容**：
- 链式调用（Chain）
- 图编排（Graph）
- 工作流（Workflow）
- 状态管理

**学习顺序**：
1. `compose/chain.go` - 简单链式调用
2. `compose/graph.go` - 图编排
3. `compose/workflow.go` - 工作流
4. `compose/state.go` - 状态管理
5. `compose/branch.go` - 分支处理
6. `compose/parallel.go` - 并行执行
7. `compose/checkpoint.go` - 检查点和恢复

### 2.4 工作流实现（flow）

**重要性**：★★★☆☆

flow 包提供了预构建的工作流实现，如 ReAct 代理、多查询检索器等。学习 flow 包可以帮助你理解如何使用组合框架构建复杂的工作流。

**学习内容**：
- ReAct 代理
- 多查询检索器
- 多代理管理

**学习顺序**：
1. `flow/agent/react/react.go` - ReAct 代理实现
2. `flow/retriever/multiquery/multi_query.go` - 多查询检索器
3. `flow/agent/multiagent/host/compose.go` - 多代理管理

### 2.5 回调机制（callbacks）

**重要性**：★★★☆☆

callbacks 包提供了回调机制，用于监控和扩展组件的行为。学习 callbacks 包可以帮助你理解如何在框架中添加自定义的监控和扩展。

**学习内容**：
- 回调接口
- 回调处理器
- 切面注入

**学习顺序**：
1. `callbacks/interface.go` - 回调接口定义
2. `callbacks/handler_builder.go` - 回调处理器构建
3. `callbacks/aspect_inject.go` - 切面注入

### 2.6 内部工具（internal）

**重要性**：★★☆☆☆

internal 包包含了框架的内部工具和辅助函数，如泛型工具、序列化、并发管理等。学习 internal 包可以帮助你理解框架的内部实现细节。

**学习内容**：
- 泛型工具
- 序列化
- 并发管理
- 流处理

**学习顺序**：
1. `internal/generic/generic.go` - 泛型工具
2. `internal/serialization/serialization.go` - 序列化
3. `internal/core/interrupt.go` - 中断处理
4. `internal/channel.go` - 通道工具

## 3. 学习方法

### 3.1 循序渐进

按照学习路径的顺序，从核心概念到高级功能，逐步深入理解框架的设计和实现。每个模块学习完成后，建议编写简单的示例代码来巩固所学知识。

### 3.2 重点关注接口

Eino 框架采用了面向接口的设计，重点关注各种组件的接口定义和使用方式，而不是具体的实现细节。通过理解接口，可以更好地理解框架的设计原则和扩展方式。

### 3.3 结合示例代码

Eino 提供了丰富的示例代码，可以帮助你理解框架的使用方式。学习源码时，建议结合示例代码一起阅读，这样可以更好地理解各种组件和功能的实际应用场景。

### 3.4 调试和测试

Eino 框架包含了丰富的测试用例，可以帮助你理解框架的各种功能和边界情况。学习源码时，建议运行测试用例，观察框架的行为和输出，这样可以更好地理解框架的内部实现细节。

## 4. 文档结构

为了帮助你系统地学习 Eino 的源码，我们将编写以下文档：

### 4.1 核心数据模型（schema）

- **概述**：schema 包的功能和作用
- **目录结构**：schema 包的目录结构
- **核心概念**：消息模型、工具调用、流式处理等
- **主要接口和实现**：Message、Stream、ToolCall 等
- **学习路径**：schema 包的学习顺序
- **最佳实践**：如何使用 schema 包
- **示例代码**：schema 包的使用示例

### 4.2 组件抽象（components）

- **概述**：components 包的功能和作用
- **目录结构**：components 包的目录结构
- **核心概念**：组件类型、接口、回调机制等
- **主要接口和实现**：ChatModel、Tool、Retriever 等
- **学习路径**：components 包的学习顺序
- **最佳实践**：如何使用和扩展组件
- **示例代码**：components 包的使用示例

### 4.3 组合框架（compose）

- **概述**：compose 包的功能和作用
- **目录结构**：compose 包的目录结构
- **核心概念**：链式调用、图编排、工作流等
- **主要接口和实现**：Chain、Graph、Workflow 等
- **学习路径**：compose 包的学习顺序
- **最佳实践**：如何构建复杂的工作流和图
- **示例代码**：compose 包的使用示例

### 4.4 工作流实现（flow）

- **概述**：flow 包的功能和作用
- **目录结构**：flow 包的目录结构
- **核心概念**：ReAct 代理、多查询检索器、多代理管理等
- **主要接口和实现**：react.go、multi_query.go、compose.go 等
- **学习路径**：flow 包的学习顺序
- **最佳实践**：如何使用预构建的工作流
- **示例代码**：flow 包的使用示例

### 4.5 回调机制（callbacks）

- **概述**：callbacks 包的功能和作用
- **目录结构**：callbacks 包的目录结构
- **核心概念**：回调接口、回调处理器、切面注入等
- **主要接口和实现**：interface.go、handler_builder.go、aspect_inject.go 等
- **学习路径**：callbacks 包的学习顺序
- **最佳实践**：如何添加自定义的监控和扩展
- **示例代码**：callbacks 包的使用示例

## 5. 学习进度

| 模块 | 预计学习时间 | 完成状态 |
|------|--------------|----------|
| schema | 1-2 天 | ☐ |
| components | 2-3 天 | ☐ |
| compose | 3-4 天 | ☐ |
| flow | 2-3 天 | ☐ |
| callbacks | 1-2 天 | ☐ |
| internal | 1-2 天 | ☐ |

## 6. 资源推荐

- **官方文档**：[Eino 官方文档](https://www.cloudwego.io/zh/docs/eino/)
- **示例代码**：[Eino 示例代码](https://github.com/cloudwego/eino-examples)
- **扩展库**：[Eino 扩展库](https://github.com/cloudwego/eino-ext)
- **GitHub 仓库**：[Eino GitHub 仓库](https://github.com/cloudwego/eino)

## 7. 总结

Eino 是一个功能强大的 LLM 应用框架，学习其源码可以帮助你深入理解框架的设计和实现，掌握构建复杂 AI 应用的技能。按照本学习计划，循序渐进地学习各个模块，可以帮助你系统地掌握 Eino 的核心概念和高级功能。

祝你学习愉快！

# Eino 源码文档

## 1. 核心数据模型（schema）

### 1.1 概述

schema 包定义了 Eino 的核心数据模型，包括消息、工具调用、流式处理和多模态内容等。这些数据模型是 Eino 框架中数据表示和流动的基础，被其他所有模块所依赖。

### 1.2 目录结构

```
schema/
├── message.go        # 核心消息定义
├── stream.go         # 流式处理
├── tool.go           # 工具调用
├── document.go       # 文档处理
├── serialization.go  # 序列化
├── select.go         # 选择器
└── 其他辅助文件和测试文件
```

### 1.3 核心概念

#### 1.3.1 消息模型

消息是 Eino 框架中最基本的数据单元，用于表示用户输入、模型输出和工具调用结果等。消息模型支持文本和多模态内容，以及流式处理。

```go
type Message struct {
    Role RoleType `json:"role"`
    Content string `json:"content"`
    UserInputMultiContent []MessageInputPart `json:"user_input_multi_content,omitempty"`
    AssistantGenMultiContent []MessageOutputPart `json:"assistant_output_multi_content,omitempty"`
    ToolCalls []ToolCall `json:"tool_calls,omitempty"`
    // 其他字段...
}
```

#### 1.3.2 工具调用

工具调用用于表示模型请求执行外部工具的操作，包括工具名称、参数和调用 ID 等。

```go
type ToolCall struct {
    Index *int `json:"index,omitempty"`
    ID string `json:"id"`
    Type string `json:"type"`
    Function FunctionCall `json:"function"`
    Extra map[string]any `json:"extra,omitempty"`
}
```

#### 1.3.3 流式处理

流式处理用于表示模型实时生成的输出，支持将多个流块合并为完整的消息。

```go
func ConcatMessages(msgs []*Message) (*Message, error) {
    // 合并多个消息流块
}
```

### 1.4 主要接口和实现

#### 1.4.1 Message

Message 是 Eino 框架中最核心的数据结构，用于表示各种类型的消息，包括用户输入、模型输出和工具调用结果等。

**主要方法**：
- `Format`：格式化消息内容
- `String`：返回消息的字符串表示
- `ConcatMessages`：合并多个消息流块

**使用示例**：

```go
// 创建用户消息
userMsg := schema.UserMessage("你好，Eino！")

// 创建系统消息
systemMsg := schema.SystemMessage("你是一个 helpful assistant。")

// 创建工具调用结果
toolMsg := schema.ToolMessage("工具调用结果", "call-123")
```

#### 1.4.2 Stream

Stream 用于表示流式数据，支持实时处理和合并流块。

**主要类型**：
- `StreamReader[T]`：流式数据读取器
- `StreamWriter[T]`：流式数据写入器
- `AsyncIterator[T]`：异步迭代器

**使用示例**：

```go
// 读取流式数据
for {
    msg, err := stream.Recv()
    if errors.Is(err, io.EOF) {
        break
    }
    // 处理消息
}
```

### 1.5 学习路径

1. **入门阶段**：
   - 阅读 `schema/message.go`，理解核心消息定义
   - 学习 `schema/stream.go`，了解流式处理
   - 查看 `schema/tool.go`，了解工具调用

2. **进阶阶段**：
   - 学习 `schema/document.go`，了解文档处理
   - 查看 `schema/serialization.go`，了解序列化
   - 运行测试用例，观察消息的合并和处理

3. **实践阶段**：
   - 尝试创建不同类型的消息
   - 实现简单的流式处理
   - 测试消息合并功能

### 1.6 最佳实践

1. **使用类型安全的消息创建函数**：
   - 使用 `schema.UserMessage`、`schema.SystemMessage` 等函数创建消息，而不是直接实例化 `Message` 结构体

2. **合理处理流式数据**：
   - 对于流式输出，使用 `schema.ConcatMessages` 函数合并流块
   - 对于需要完整消息的组件，确保先合并流块再传递

3. **使用多模态内容**：
   - 对于需要处理图像、音频等多模态内容的场景，使用 `UserInputMultiContent` 和 `AssistantGenMultiContent` 字段

### 1.7 示例代码

```go
// 创建文本消息
userMsg := schema.UserMessage("What is the capital of France?")

// 创建多模态消息
multiModalMsg := &schema.Message{
    Role: schema.User,
    UserInputMultiContent: []schema.MessageInputPart{
        {Type: schema.ChatMessagePartTypeText, Text: "What is in this image?"},
        {Type: schema.ChatMessagePartTypeImageURL, Image: &schema.MessageInputImage{
            MessagePartCommon: schema.MessagePartCommon{
                URL: toPtr("https://example.com/cat.jpg"),
            },
            Detail: schema.ImageURLDetailHigh,
        }},
    },
}

// 合并消息流
msgs := []*schema.Message{
    {Role: schema.Assistant, Content: "Hello"},
    {Role: schema.Assistant, Content: " World"},
}
concatedMsg, _ := schema.ConcatMessages(msgs)
// concatedMsg.Content = "Hello World"
```

## 2. 组件抽象（components）

### 2.1 概述

components 包定义了 Eino 框架中各种组件的抽象接口，如 ChatModel、Tool、Retriever 等。这些组件接口是构建 LLM 应用的基础，提供了统一的使用方式和扩展点。

### 2.2 目录结构

```
components/
├── model/          # 聊天模型组件
├── tool/           # 工具组件
├── retriever/      # 检索器组件
├── embedding/      # 嵌入组件
├── indexer/        # 索引器组件
├── prompt/         # 提示模板组件
├── document/       # 文档处理组件
└── types.go        # 组件类型定义
```

### 2.3 核心概念

#### 2.3.1 组件类型

Eino 框架支持多种组件类型，每种组件类型都有其特定的功能和接口定义。

```go
type Component string

const (
    ComponentOfPrompt      Component = "ChatTemplate"
    ComponentOfChatModel   Component = "ChatModel"
    ComponentOfEmbedding   Component = "Embedding"
    ComponentOfIndexer     Component = "Indexer"
    ComponentOfRetriever   Component = "Retriever"
    ComponentOfLoader      Component = "Loader"
    ComponentOfTransformer Component = "DocumentTransformer"
    ComponentOfTool        Component = "Tool"
)
```

#### 2.3.2 组件接口

每种组件类型都有其特定的接口定义，包括输入输出类型、选项类型和流式处理方式等。

**ChatModel 接口示例**：

```go
type BaseChatModel interface {
    Generate(ctx context.Context, input []*schema.Message, options ...ChatModelOption) (*schema.Message, error)
    GenerateStream(ctx context.Context, input []*schema.Message, options ...ChatModelOption) (*schema.StreamReader[*schema.Message], error)
}
```

#### 2.3.3 回调机制

组件支持回调机制，可以在组件执行前后触发回调函数，用于监控和扩展组件的行为。

```go
type Checker interface {
    IsCallbacksEnabled() bool
}
```

### 2.4 主要接口和实现

#### 2.4.1 ChatModel

ChatModel 组件用于与 LLM 模型进行交互，支持生成文本和流式输出。

**主要方法**：
- `Generate`：生成文本输出
- `GenerateStream`：生成流式输出

**使用示例**：

```go
// 创建 ChatModel 实例
model, _ := openai.NewChatModel(ctx, config)

// 生成文本输出
message, _ := model.Generate(ctx, []*schema.Message{
    schema.SystemMessage("你是一个 helpful assistant。"),
    schema.UserMessage("法国的首都是什么？")})

// 生成流式输出
stream, _ := model.GenerateStream(ctx, []*schema.Message{
    schema.SystemMessage("你是一个 helpful assistant。"),
    schema.UserMessage("法国的首都是什么？")})
```

#### 2.4.2 Tool

Tool 组件用于执行外部工具，支持同步和异步执行。

**主要方法**：
- `Call`：同步执行工具
- `CallStream`：异步执行工具，返回流式结果

**使用示例**：

```go
// 创建 Tool 实例
tool := tool.NewFunctionTool(func(ctx context.Context, input string) (string, error) {
    // 工具实现
    return "结果", nil
})

// 调用工具
result, _ := tool.Call(ctx, map[string]any{"input": "参数"})
```

#### 2.4.3 Retriever

Retriever 组件用于从数据源检索相关信息，支持多种检索策略。

**主要方法**：
- `Retrieve`：检索相关信息

**使用示例**：

```go
// 创建 Retriever 实例
retriever, _ := vectorstore.NewRetriever(ctx, config)

// 检索信息
docs, _ := retriever.Retrieve(ctx, "查询文本")
```

### 2.5 学习路径

1. **入门阶段**：
   - 阅读 `components/types.go`，理解组件类型定义
   - 学习 `components/model/interface.go`，了解 ChatModel 接口
   - 查看 `components/tool/interface.go`，了解 Tool 接口

2. **进阶阶段**：
   - 学习 `components/retriever/interface.go`，了解 Retriever 接口
   - 查看 `components/embedding/interface.go`，了解 Embedding 接口
   - 学习 `components/prompt/interface.go`，了解 Prompt 接口

3. **实践阶段**：
   - 尝试实现简单的组件
   - 使用组件构建简单的应用
   - 测试组件的回调机制

### 2.6 最佳实践

1. **面向接口编程**：
   - 始终通过接口使用组件，而不是具体实现，这样可以提高代码的灵活性和可测试性

2. **合理使用选项**：
   - 使用选项模式配置组件，而不是冗长的构造函数
   - 为组件定义合理的默认值

3. **实现回调机制**：
   - 对于需要监控的组件，实现 `Checker` 接口，启用回调机制
   - 在组件执行的关键节点触发回调

### 2.7 示例代码

```go
// 创建 ChatModel 实例
model, _ := openai.NewChatModel(ctx, config)

// 创建 Tool 实例
tool := tool.NewFunctionTool(func(ctx context.Context, input string) (string, error) {
    return "结果", nil
})

// 创建 Retriever 实例
retriever, _ := vectorstore.NewRetriever(ctx, config)

// 使用组件构建应用
// ...
```

## 3. 组合框架（compose）

### 3.1 概述

compose 包提供了强大的组合框架，用于构建复杂的工作流和图。组合框架处理了类型检查、流式处理、并发管理、切面注入和选项分配等复杂问题，让用户可以专注于业务逻辑。

### 3.2 目录结构

```
compose/
├── chain.go         # 链式调用
├── graph.go         # 图编排
├── workflow.go      # 工作流
├── branch.go        # 分支处理
├── parallel.go      # 并行执行
├── state.go         # 状态管理
├── checkpoint.go    # 检查点和恢复
└── 其他辅助文件和测试文件
```

### 3.3 核心概念

#### 3.3.1 链式调用（Chain）

链式调用是一种简单的有向图，只能向前流动，适合构建线性的工作流。

```go
chain, _ := compose.NewChain[map[string]any, *schema.Message]().
           AppendChatTemplate(prompt).
           AppendChatModel(model).
           Compile(ctx)

result, _ := chain.Invoke(ctx, map[string]any{"query": "查询文本"})
```

#### 3.3.2 图编排（Graph）

图编排是一种强大的有向图，可以是循环或非循环的，适合构建复杂的工作流。

```go
graph := compose.NewGraph[map[string]any, *schema.Message]()

_ = graph.AddChatTemplateNode("node_template", chatTpl)
_ = graph.AddChatModelNode("node_model", chatModel)
_ = graph.AddToolsNode("node_tools", toolsNode)
_ = graph.AddEdge("node_template", "node_model")
_ = graph.AddBranch("node_model", branch)

compiledGraph, _ := graph.Compile(ctx)
result, _ := compiledGraph.Invoke(ctx, map[string]any{"query": "查询文本"})
```

#### 3.3.3 工作流（Workflow）

工作流是一种非循环图，支持字段级别的数据映射，适合构建复杂的数据处理流程。

```go
wf := compose.NewWorkflow[[]*schema.Message, *schema.Message]()
wf.AddChatModelNode("model", m).AddInput(compose.START)
wf.AddLambdaNode("lambda1", compose.InvokableLambda(lambda1)).
    AddInput("model", compose.MapFields("Content", "Input"))

runnable, _ := wf.Compile(ctx)
result, _ := runnable.Invoke(ctx, []*schema.Message{schema.UserMessage("启动工作流")})
```

### 3.4 主要接口和实现

#### 3.4.1 Chain

Chain 用于构建简单的链式工作流，支持顺序执行多个组件。

**主要方法**：
- `AppendChatTemplate`：添加 ChatTemplate 组件
- `AppendChatModel`：添加 ChatModel 组件
- `AppendTool`：添加 Tool 组件
- `Compile`：编译 Chain
- `Invoke`：执行 Chain

**使用示例**：

```go
chain, _ := compose.NewChain[map[string]any, *schema.Message]().
           AppendChatTemplate(prompt).
           AppendChatModel(model).
           Compile(ctx)

result, _ := chain.Invoke(ctx, map[string]any{"query": "查询文本"})
```

#### 3.4.2 Graph

Graph 用于构建复杂的有向图工作流，支持分支、并行和循环等复杂逻辑。

**主要方法**：
- `AddChatTemplateNode`：添加 ChatTemplate 节点
- `AddChatModelNode`：添加 ChatModel 节点
- `AddToolsNode`：添加 Tools 节点
- `AddEdge`：添加边
- `AddBranch`：添加分支
- `Compile`：编译 Graph
- `Invoke`：执行 Graph

**使用示例**：

```go
graph := compose.NewGraph[map[string]any, *schema.Message]()

_ = graph.AddChatTemplateNode("node_template", chatTpl)
_ = graph.AddChatModelNode("node_model", chatModel)
_ = graph.AddToolsNode("node_tools", toolsNode)
_ = graph.AddEdge("node_template", "node_model")
_ = graph.AddBranch("node_model", branch)
_ = graph.AddEdge("node_tools", "node_model")

compiledGraph, _ := graph.Compile(ctx)
result, _ := compiledGraph.Invoke(ctx, map[string]any{"query": "查询文本"})
```

#### 3.4.3 Workflow

Workflow 用于构建字段级别的数据处理流程，支持复杂的数据映射和转换。

**主要方法**：
- `AddChatModelNode`：添加 ChatModel 节点
- `AddLambdaNode`：添加 Lambda 节点
- `AddInput`：添加输入
- `MapFields`：映射字段
- `Compile`：编译 Workflow
- `Invoke`：执行 Workflow

**使用示例**：

```go
type Input1 struct {
    Input string
}

type Output1 struct {
    Output string
}

lambda1 := func(ctx context.Context, input Input1) (Output1, error) {
    return Output1{Output: input.Input + " processed"}, nil
}

wf := compose.NewWorkflow[[]*schema.Message, *schema.Message]()
wf.AddChatModelNode("model", m).AddInput(compose.START)
wf.AddLambdaNode("lambda1", compose.InvokableLambda(lambda1)).
    AddInput("model", compose.MapFields("Content", "Input"))

runnable, _ := wf.Compile(ctx)
result, _ := runnable.Invoke(ctx, []*schema.Message{schema.UserMessage("启动工作流")})
```

### 3.5 学习路径

1. **入门阶段**：
   - 阅读 `compose/chain.go`，理解链式调用
   - 学习 `compose/graph.go`，了解图编排
   - 查看 `compose/workflow.go`，了解工作流

2. **进阶阶段**：
   - 学习 `compose/branch.go`，了解分支处理
   - 查看 `compose/parallel.go`，了解并行执行
   - 学习 `compose/state.go`，了解状态管理

3. **实践阶段**：
   - 尝试构建简单的链式工作流
   - 实现复杂的图编排
   - 构建字段级别的工作流

### 3.6 最佳实践

1. **选择合适的组合方式**：
   - 对于简单的线性流程，使用 Chain
   - 对于复杂的分支和循环流程，使用 Graph
   - 对于需要字段级映射的流程，使用 Workflow

2. **合理设计节点和边**：
   - 为节点取有意义的名称
   - 清晰地定义节点之间的依赖关系
   - 避免构建过于复杂的图

3. **使用类型安全的输入输出**：
   - 为组合框架指定明确的输入输出类型
   - 使用类型安全的字段映射

### 3.7 示例代码

```go
// 构建链式工作流
chain, _ := compose.NewChain[map[string]any, *schema.Message]().
           AppendChatTemplate(prompt).
           AppendChatModel(model).
           Compile(ctx)

// 执行链式工作流
result, _ := chain.Invoke(ctx, map[string]any{"query": "查询文本"})

// 构建图编排
graph := compose.NewGraph[map[string]any, *schema.Message]()

_ = graph.AddChatTemplateNode("node_template", chatTpl)
_ = graph.AddChatModelNode("node_model", chatModel)
_ = graph.AddToolsNode("node_tools", toolsNode)
_ = graph.AddEdge(compose.START, "node_template")
_ = graph.AddEdge("node_template", "node_model")
_ = graph.AddBranch("node_model", branch)
_ = graph.AddEdge("node_tools", compose.END)

// 执行图编排
compiledGraph, _ := graph.Compile(ctx)
result, _ := compiledGraph.Invoke(ctx, map[string]any{"query": "查询文本"})
```

## 4. 工作流实现（flow）

### 4.1 概述

flow 包提供了预构建的工作流实现，如 ReAct 代理、多查询检索器和多代理管理等。这些工作流实现了常见的 LLM 应用模式，可以直接使用或作为参考进行定制。

### 4.2 目录结构

```
flow/
├── agent/                # 代理实现
│   ├── react/           # ReAct 代理
│   └── multiagent/      # 多代理管理
├── retriever/            # 检索器实现
│   ├── multiquery/      # 多查询检索器
│   └── router/          # 检索器路由器
└── indexer/             # 索引器实现
    └── parent/          # 父索引器
```

### 4.3 核心概念

#### 4.3.1 ReAct 代理

ReAct 代理实现了 ReAct（Reasoning and Acting）模式，结合了推理和行动能力，能够使用工具解决复杂问题。

```go
// ReAct 代理配置
type reactConfig struct {
    model          model.BaseChatModel
    toolsConfig    *toolsConfig
    agentName      string
    maxIterations  int
}

// 创建 ReAct 代理
graph, _ := newReact(ctx, config)
```

#### 4.3.2 多查询检索器

多查询检索器通过生成多个查询来提高检索的召回率，适合处理模糊或复杂的查询。

```go
// 多查询检索器配置
type Config struct {
    Name            string
    Description     string
    ChatModel       model.BaseChatModel
    Retriever       retriever.BaseRetriever
    QueryCount      int
    // 其他配置...
}

// 创建多查询检索器
multiQueryRetriever, _ := multiquery.New(ctx, config)
```

#### 4.3.3 多代理管理

多代理管理用于管理和协调多个代理，支持代理之间的通信和转移。

```go
// 多代理配置
type Config struct {
    Name            string
    Description     string
    ChatModel       model.BaseChatModel
    SubAgents       map[string]adk.Agent
    // 其他配置...
}

// 创建多代理管理器
hostAgent, _ := host.New(ctx, config)
```

### 4.4 主要接口和实现

#### 4.4.1 ReAct 代理

ReAct 代理实现了 ReAct 模式，能够使用工具解决复杂问题。

**主要组件**：
- ChatModel 节点：用于生成思考和工具调用
- Tools 节点：用于执行工具调用
- 分支逻辑：用于决定是继续思考还是返回最终答案

**使用示例**：

```go
// 创建 ReAct 代理配置
config := &reactConfig{
    model:          chatModel,
    toolsConfig:    toolsConfig,
    agentName:      "react-agent",
    maxIterations:  10,
}

// 创建 ReAct 代理
graph, _ := newReact(ctx, config)

// 执行 ReAct 代理
result, _ := graph.Invoke(ctx, map[string]any{"query": "复杂问题"})
```

#### 4.4.2 多查询检索器

多查询检索器通过生成多个查询来提高检索的召回率。

**主要组件**：
- ChatModel：用于生成多个查询
- Retriever：用于执行检索
- 结果合并：用于合并多个检索结果

**使用示例**：

```go
// 创建多查询检索器配置
config := &multiquery.Config{
    Name:            "multi-query-retriever",
    Description:     "多查询检索器",
    ChatModel:       chatModel,
    Retriever:       retriever,
    QueryCount:      3,
}

// 创建多查询检索器
multiQueryRetriever, _ := multiquery.New(ctx, config)

// 执行检索
docs, _ := multiQueryRetriever.Retrieve(ctx, "查询文本")
```

### 4.5 学习路径

1. **入门阶段**：
   - 阅读 `flow/agent/react/react.go`，理解 ReAct 代理实现
   - 学习 `flow/retriever/multiquery/multi_query.go`，了解多查询检索器

2. **进阶阶段**：
   - 学习 `flow/agent/multiagent/host/compose.go`，了解多代理管理
   - 查看 `flow/retriever/router/router.go`，了解检索器路由器

3. **实践阶段**：
   - 尝试使用 ReAct 代理解决复杂问题
   - 实现自定义的多查询检索器
   - 构建简单的多代理系统

### 4.6 最佳实践

1. **合理配置参数**：
   - 为 ReAct 代理配置合适的最大迭代次数
   - 为多查询检索器配置合适的查询数量

2. **扩展预构建工作流**：
   - 基于预构建工作流进行扩展，而不是从头开始实现
   - 合理使用选项模式进行定制

3. **测试和优化**：
   - 测试工作流的性能和准确性
   - 根据测试结果优化工作流的配置和实现

### 4.7 示例代码

```go
// 创建 ReAct 代理配置
config := &reactConfig{
    model:          chatModel,
    toolsConfig:    toolsConfig,
    agentName:      "react-agent",
    maxIterations:  10,
}

// 创建 ReAct 代理
graph, _ := newReact(ctx, config)

// 执行 ReAct 代理
result, _ := graph.Invoke(ctx, map[string]any{"query": "法国的首都是什么？"})

// 创建多查询检索器配置
multiQueryConfig := &multiquery.Config{
    Name:            "multi-query-retriever",
    Description:     "多查询检索器",
    ChatModel:       chatModel,
    Retriever:       retriever,
    QueryCount:      3,
}

// 创建多查询检索器
multiQueryRetriever, _ := multiquery.New(ctx, multiQueryConfig)

// 执行检索
docs, _ := multiQueryRetriever.Retrieve(ctx, "Eino 框架的核心功能是什么？")
```

## 5. 回调机制（callbacks）

### 5.1 概述

callbacks 包提供了回调机制，用于监控和扩展组件的行为。通过回调机制，可以在组件执行前后触发自定义的逻辑，如日志记录、指标收集和跟踪等。

### 5.2 目录结构

```
callbacks/
├── interface.go        # 回调接口定义
├── handler_builder.go  # 回调处理器构建
├── aspect_inject.go    # 切面注入
└── 其他辅助文件和测试文件
```

### 5.3 核心概念

#### 5.3.1 回调接口

回调接口定义了回调函数的签名和触发时机，包括组件执行前后、错误处理和流式处理等。

```go
// 回调处理器接口
type Handler interface {
    OnStart(ctx context.Context, info *RunInfo, input CallbackInput) context.Context
    OnEnd(ctx context.Context, info *RunInfo, output CallbackOutput) context.Context
    OnError(ctx context.Context, info *RunInfo, err error) context.Context
    // 其他回调方法...
}
```

#### 5.3.2 回调处理器

回调处理器实现了回调接口，用于处理回调事件。可以使用 HandlerBuilder 来构建回调处理器。

```go
// 创建回调处理器
handler := callbacks.NewHandlerBuilder().
  OnStartFn(func(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
      log.Infof("组件开始执行: %v", info)
      return ctx
  }).
  OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
      log.Infof("组件执行结束: %v", info)
      return ctx
  }).
  Build()
```

#### 5.3.3 切面注入

切面注入用于将回调处理器注入到组件中，支持全局注入、组件类型注入和节点级注入等。

```go
// 全局注入回调处理器
compiledGraph.Invoke(ctx, input, compose.WithCallbacks(handler))

// 组件类型注入
compiledGraph.Invoke(ctx, input, compose.WithChatModelOption(model.WithTemperature(0.5)))

// 节点级注入
compiledGraph.Invoke(ctx, input, compose.WithCallbacks(handler).DesignateNode("node_1"))
```

### 5.4 主要接口和实现

#### 5.4.1 Handler 接口

Handler 接口定义了回调函数的签名和触发时机，是回调机制的核心。

**主要方法**：
- `OnStart`：组件开始执行时触发
- `OnEnd`：组件执行结束时触发
- `OnError`：组件执行出错时触发
- `OnStartWithStreamInput`：处理流式输入时触发
- `OnEndWithStreamOutput`：处理流式输出时触发

**使用示例**：

```go
// 实现 Handler 接口
type MyHandler struct {}

func (h *MyHandler) OnStart(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
    log.Infof("OnStart: %v", info)
    return ctx
}

func (h *MyHandler) OnEnd(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
    log.Infof("OnEnd: %v", info)
    return ctx
}

// 其他方法实现...
```

#### 5.4.2 HandlerBuilder

HandlerBuilder 用于构建回调处理器，提供了链式 API 来配置回调函数。

**主要方法**：
- `OnStartFn`：配置 OnStart 回调函数
- `OnEndFn`：配置 OnEnd 回调函数
- `OnErrorFn`：配置 OnError 回调函数
- `Build`：构建回调处理器

**使用示例**：

```go
// 创建回调处理器
handler := callbacks.NewHandlerBuilder().
  OnStartFn(func(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
      log.Infof("OnStart: %v", info)
      return ctx
  }).
  OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
      log.Infof("OnEnd: %v", info)
      return ctx
  }).
  OnErrorFn(func(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
      log.Errorf("OnError: %v, err: %v", info, err)
      return ctx
  }).
  Build()
```

### 5.5 学习路径

1. **入门阶段**：
   - 阅读 `callbacks/interface.go`，理解回调接口定义
   - 学习 `callbacks/handler_builder.go`，了解如何构建回调处理器

2. **进阶阶段**：
   - 学习 `callbacks/aspect_inject.go`，了解切面注入
   - 查看测试文件，了解回调机制的使用

3. **实践阶段**：
   - 实现自定义的回调处理器
   - 测试回调机制的各种注入方式
   - 构建带有监控的应用

### 5.6 最佳实践

1. **合理设计回调函数**：
   - 回调函数应该简洁高效，避免阻塞
   - 不要在回调函数中修改组件的状态

2. **使用合适的注入方式**：
   - 对于全局监控，使用全局注入
   - 对于特定组件类型，使用组件类型注入
   - 对于特定节点，使用节点级注入

3. **结合日志和指标**：
   - 在回调函数中记录关键日志
   - 收集组件执行的指标，如执行时间、调用次数等

### 5.7 示例代码

```go
// 创建回调处理器
handler := callbacks.NewHandlerBuilder().
  OnStartFn(func(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
      log.Infof("组件开始执行: %s", info.Name)
      return ctx
  }).
  OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
      log.Infof("组件执行结束: %s, 耗时: %v", info.Name, info.EndTime.Sub(info.StartTime))
      return ctx
  }).
  OnErrorFn(func(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
      log.Errorf("组件执行出错: %s, 错误: %v", info.Name, err)
      return ctx
  }).
  Build()

// 使用回调处理器
compiledGraph.Invoke(ctx, input, compose.WithCallbacks(handler))

// 组件类型注入
compiledGraph.Invoke(ctx, input, compose.WithChatModelOption(model.WithTemperature(0.5)))

// 节点级注入
compiledGraph.Invoke(ctx, input, compose.WithCallbacks(handler).DesignateNode("node_1"))
```

## 6. 总结

Eino 是一个功能强大的 LLM 应用框架，提供了组件抽象、组合框架、API 和工作流等功能。通过系统地学习 Eino 的源码，你可以深入理解框架的设计和实现，掌握构建复杂 AI 应用的技能。

本学习计划和文档涵盖了 Eino 框架的核心模块，包括核心数据模型、组件抽象、组合框架、工作流实现和回调机制等。按照学习路径逐步深入，结合示例代码和实践，可以帮助你更好地理解和使用 Eino 框架。

祝你学习愉快！
