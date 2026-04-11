# 02 - context 是什么、作用、核心接口

> **极简、面试必过、开发必用**：最精简、最好记、直接能背的版本；含关键细节、面试常问、代码示例。

---

## 1. 一句话定义（背）

**Context = 请求生命周期控制器**

负责：**取消、超时、请求级数据传递**，全链路一层层传递。

---

## 2. Context 解决的 3 个核心问题

1. **取消**：上游主动取消，下游 goroutine 尽快退出  
2. **超时/截止时间**：防止请求无限等待、资源耗尽  
3. **请求级元数据透传**：traceID、userID、requestID 等**小数据**

---

## 3. 核心接口（只记 4 个方法）

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool) // 什么时候超时
    Done() <-chan struct{}                  // 取消信号通道（关闭即取消）
    Err() error                             // 取消原因：Canceled / DeadlineExceeded
    Value(key any) any                      // 存取值：仅存请求级小数据
}
```

### 每个方法作用（极简版）

| 方法 | 作用 |
|------|------|
| **Deadline()** | 有没有截止时间？什么时候到点？ |
| **Done()** | 监听取消信号（`<-ctx.Done()`） |
| **Err()** | 为什么结束：主动取消 / 超时 |
| **Value()** | 请求级键值读取（不要存大对象） |

---

## 4. 最常用的创建方式（速记）

```go
context.Background()                    // 根 ctx，全局入口常用，永不取消
context.TODO()                          // 占位：还不知道传谁（与 Background 语义不同，见 03）
ctx, cancel := context.WithCancel(parent) // 可主动取消；务必在合适路径调用 cancel()
ctx, cancel := context.WithTimeout(parent, time.Second) // 超时自动取消；同样要 defer cancel()
```

说明：`WithCancel` / `WithTimeout` 都会返回 **`CancelFunc`**，**在 goroutine 或函数结束前调用 `cancel()`** 可尽早释放子树资源（避免泄漏）。`WithDeadline` 见 [06-WithDeadline 精确到时间点的取消控制.md](./06-WithDeadline%20精确到时间点的取消控制.md)（与 [05-WithTimeout](./05-WithTimeout%20超时控制.md) 对照）。

---

## 5. 最经典使用代码（必会）

### 监听取消/超时

```go
func do(ctx context.Context) {
    select {
    case <-ctx.Done():
        // 取消了/超时了
        fmt.Println(ctx.Err()) // context.Canceled 或 DeadlineExceeded
        return
    case <-time.After(1 * time.Second):
        // 正常执行（示例：真实代码里常是 IO、channel、业务逻辑）
        fmt.Println("done")
    }
}
```

### 启动 goroutine 必须传 ctx

```go
go do(ctx)
```

让下游能感知同一请求的取消/超时，**不要**在子 goroutine 里用 `context.Background()` 替代业务 ctx（除非该任务与请求生命周期无关）。

---

## 6. 工程规范（面试高频）

1. **函数第一个参数一般是 `ctx`**：`func xxx(ctx context.Context, ...)`  
2. **不要把 Context 塞进结构体当长期字段**：生命周期混乱、容易泄漏、难以测试。短期包装（如 `http.Request` 携带）是框架约定，与“业务 struct 里藏 ctx 字段”不同。  
3. **不要用 `WithValue` 传业务参数或大对象**：只传**元数据**（traceID、userID、requestID）；业务数据用函数参数、返回值或依赖注入。  
4. **上层取消，下层一起退出**：派生 ctx 形成树，**父取消 → 子 `Done()` 都会收到**（见 [03-Background()、TODO()、上下文树.md](./03-Background()、TODO()、上下文树.md)）。  
5. **Context 值可安全地在多个 goroutine 中并发读取**；`CancelFunc` 可多次调用（幂等）。

---

## 7. 面试必问 3 问（直接背答案）

### Q1：Context 用来干嘛？

控制请求生命周期：**取消、超时、请求级数据透传**。

### Q2：Context 怎么实现取消协作？

通过 `Done()` 返回的 channel：**关闭该 channel 即表示取消（或超时）**，下游用 `select` 或单独 `<-ctx.Done()` 感知。

### Q3：Context 里适合存什么？

只能存**请求级小元数据**（traceID、requestID、userID）。**不要**存大对象、**不要**把 `WithValue` 当通用参数传递机制。

---

## 8. 超级总结（一张脑图）

- **作用**：取消、超时、透传小数据  
- **核心**：`Done()` 监听取消  
- **规范**：函数首参传 ctx、不滥用结构体字段、逐层传递  
- **原理**：树形派生，**父取消 → 子同步取消**

---

## 延伸阅读（同目录）

| 主题 | 文档 |
|------|------|
| 取消链路与 goroutine | [01-Goroutine链与Context取消.md](./01-Goroutine链与Context取消.md) |
| Background / TODO / 树 | [03-Background()、TODO()、上下文树.md](./03-Background()、TODO()、上下文树.md) |
| WithCancel | [04-WithCancel 手动取消.md](./04-WithCancel%20手动取消.md) |
| WithTimeout | [05-WithTimeout 超时控制.md](./05-WithTimeout%20超时控制.md) |
| WithDeadline | [06-WithDeadline 精确到时间点的取消控制.md](./06-WithDeadline%20精确到时间点的取消控制.md) |
| WithValue | [07-WithValue 数据透传.md](./07-WithValue%20数据透传.md) |
| 中间件 | [08-context在中间件中的实战.md](./08-context在中间件中的实战.md) |
| 陷阱与反模式 | [09-context常见陷阱与反模式.md](./09-context常见陷阱与反模式.md) |
| trace-id 透传 | [10-中间件实战-trace-id透传.md](./10-中间件实战-trace-id透传.md) |
