# 06 - WithDeadline：精确到时间点的取消控制【完整版详细笔记】

> 与 [05 - WithTimeout：超时控制](./05-WithTimeout%20超时控制.md) 成对阅读：相对时长 vs 绝对时刻；`defer cancel()` 纪律一致。本篇在目录中为 **06**，中间件实战见 **08**。

---

## 1. 作用与定义

```go
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```

- **作用**：创建一个在**给定绝对时刻**到达后自动取消的子上下文。  
- **关系**：`WithTimeout(parent, d)` 在语义上等价于 **`WithDeadline(parent, time.Now().Add(d))`**，工程上常直接用 `WithTimeout` 写相对超时；需要**对齐外部时间点**或**精算剩余时间**时用 `WithDeadline`。  
- **取消触发条件**（满足任一即可，下游从 `Done()` 感知）：  
  1. **到达** `deadline`（若与父 deadline 取整后仍有效）  
  2. **手动** `cancel()`  
  3. **父**上下文先取消 / 先到期 → **级联**

---

## 2. WithDeadline vs WithTimeout（面试必分）

| 特性 | `WithTimeout` | `WithDeadline` |
|------|----------------|----------------|
| **参数** | `time.Duration`（再活多久） | `time.Time`（活到何时） |
| **语义** | 相对时长 | 绝对时刻 |
| **关系** | 常实现为「`Now` + `d`」再调 `WithDeadline` | 更直接表达「与墙钟 / SLA 对齐」的截止点 |

**一句话：** `WithTimeout` 是相对时间；`WithDeadline` 是**绝对时间点**（实现上同属一类 timer + `cancel` 机制）。

---

## 3. 什么时候用 WithDeadline？

适合**已有明确截止时刻**或**要复用父级 deadline 做减法**的场景：

1. **全链路 SLA**：上游传入「请求必须在时刻 T 前结束」，下游继续用**同一 T** 或子区间。  
2. **令牌 / 会话到期**：与 JWT `exp`、服务端会话绝对过期时间对齐。  
3. **剩余时间精算**：父 `ctx` 已有 deadline，本层用 `ctx.Deadline()` 算 **`remaining`**，再 `WithDeadline(parent, time.Now().Add(remaining))` 或 `WithTimeout(parent, remaining)` 交给下游。  
4. **对齐整点 / 批窗口**：例如「本窗口在 00:00:00 准时收口」（仍注意**单调时钟**与业务时钟区别，见第 8 节）。

---

## 4. 标准模板（依然必须 `defer cancel`）

```go
deadline := time.Now().Add(3 * time.Second) // 示例：也可用任意 time.Time

ctx, cancel := context.WithDeadline(parentCtx, deadline)
defer cancel()

err := doSomething(ctx)
if errors.Is(err, context.DeadlineExceeded) {
	// 到达截止时间（或父链上更早的 deadline 先到）
}
```

---

## 5. 核心方法：`Deadline()`

```go
Deadline() (deadline time.Time, ok bool)
```

- **`ok == true`**：当前 ctx 链路上能确定一个**有效截止时间**（通常来自 `WithDeadline` / `WithTimeout` 派生）。  
- **`ok == false`**：无截止时间（如 `Background()` / `TODO()`），`deadline` 无意义。

**典型：看还剩多久**

```go
if d, ok := ctx.Deadline(); ok {
	remaining := time.Until(d)
	_ = remaining // 用于限流子调用时长、打日志等
}
```

---

## 6. 传播与覆盖规则（易踩坑）

在**已有 deadline** 的父 `ctx` 上再 `WithDeadline(parent, t)`：

- **有效 deadline** = **不晚于**父级已有 deadline 的那个时刻（**取更早**）。  
- 若子传入的 `t` **晚于**父级已有 deadline，**不会**把整条链「放宽」；子仍会在**父级到点**时一并取消。  

**记忆：** 下游只能把死线**提前**或保持不变（从整条链视角），**不能推翻上游更紧的 SLA**。

---

## 7. 错误判断

与 `WithTimeout` 相同，**因到点而结束**时，`ctx.Err()` 一般为 **`context.DeadlineExceeded`**（与主动 `cancel` 的 `Canceled` 区分）：

```go
if errors.Is(err, context.DeadlineExceeded) {
	// 处理超时 / 到点
}
```

---

## 8. 高频坑点

1. **时区与语义**：`time.Time` 要带正确**时区 / 语义**（UTC 或业务约定）；不要混用「本地墙钟」与「单调计时」需求——**超时取消依赖的是 `time` 包与调度**，极端**闰秒 / NTP 大步调整**下墙钟可能跳变，生产环境要知晓该边界。  
2. **忘记 `defer cancel()`**：与 `WithTimeout` 一样，应尽早释放内部 timer 等资源。  
3. **试图在子 ctx「延长」父 deadline**：做不到；只能更紧或不变。  
4. **`deadline` 已在过去**：子 ctx 可能**立即**处于已超时状态，`Done()` 已关闭或很快关闭。

---

## 9. 面试一句话

> `WithDeadline` 用**绝对时间点**控制取消，与 `WithTimeout` 共享同一套级联与 `defer cancel()` 规范；多级嵌套时**有效 deadline 取更早者**，下游不能放宽上游 SLA。

---

## 延伸阅读

- [04-WithCancel 手动取消.md](./04-WithCancel%20手动取消.md)  
- [05-WithTimeout 超时控制.md](./05-WithTimeout%20超时控制.md)  
- [03-Background()、TODO()、上下文树.md](./03-Background()、TODO()、上下文树.md)  
- [09-context常见陷阱与反模式.md](./09-context常见陷阱与反模式.md)
