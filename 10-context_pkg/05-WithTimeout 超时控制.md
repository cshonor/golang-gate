# 05 - WithTimeout：超时控制【完整版详细笔记】

详细、深入、可直接用于学习和面试；含原理、示例、坑点、面试追问。**绝对时刻**与传播规则见专篇 [06-WithDeadline 精确到时间点的取消控制.md](./06-WithDeadline%20精确到时间点的取消控制.md)；文末仍保留极简对照。

---

## 1. 作用与定义

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

- **作用**：创建一个**到点自动取消**的子上下文。  
- **本质**：`WithTimeout(parent, d)` **等价于** `WithDeadline(parent, time.Now().Add(d))`，是**语法糖**。  
- **取消会在以下任一情况发生**（满足任一即 `Done()` 关闭，下游可感知）：  
  1. **时间到** → 自动取消（`Err()` 一般为 `DeadlineExceeded`）  
  2. **手动** `cancel()` → 提前取消（`Err()` 一般为 `Canceled`）  
  3. **父上下文**取消/超时 → **级联**到子（见 [03-Background()、TODO()、上下文树.md](./03-Background()、TODO()、上下文树.md)）

---

## 2. 什么时候必须用 WithTimeout（详细场景）

凡**可能阻塞、耗时不确定、依赖外部**的路径，都应有**显式上限**（超时或整体 deadline）：

1. **数据库**（MySQL、PostgreSQL 等）  
2. **缓存**（Redis 等）  
3. **HTTP 调第三方**、**gRPC / RPC**  
4. **消息队列**生产/消费、**网络 / 磁盘 I/O**（视 SLA）  
5. 任何「**可能卡住不返回**」的逻辑  

目的：**避免请求无限挂起、goroutine / 连接堆积、雪崩**。

---

## 3. 标准完整模板（带解释）

```go
// 1. 创建：例如 200ms 后自动取消（若父未更早取消）
ctx, cancel := context.WithTimeout(parentCtx, 200*time.Millisecond)

// 2. 必须 defer cancel()：提前结束或正常返回时尽快释放 timer 等资源
defer cancel()

// 3. 把 ctx 传给可能阻塞的下游（DB、HTTP Client、grpc.Dial 等）
err := doBusiness(ctx)

// 4. 区分「超时」与其它错误（见第 6 节）
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        // 典型：本 ctx 的 deadline 到了，或下游把该错误透传上来
        return err
    }
    return err
}
```

### 为什么必须 `defer cancel()`？

- `WithTimeout` 会在内部关联**定时器**等状态；**尽早** `cancel()` 可提前拆掉定时器，减轻高并发下的开销。  
- **规范**：`WithTimeout` / `WithDeadline` **与** `defer cancel()` **成对出现**（与 [04-WithCancel 手动取消.md](./04-WithCancel%20手动取消.md) 同一纪律）。

---

## 4. WithTimeout 底层原理（深入理解）

1. `WithTimeout(parent, d)` = `WithDeadline(parent, time.Now().Add(d))`。  
2. 标准库在实现上会在**到点**时触发 **`cancel`**（内部常基于 **timer**，如 `AfterFunc` 一类机制，细节随 Go 版本可查源码）。  
3. 到点或手动 `cancel` → 关闭（或完成）该 ctx 的 `Done()` 通道语义 → 取消沿子树传播。  
4. 监听 `<-ctx.Done()` 或使用支持 `Context` 的 API，即可协作退出。  

可记：**带自动到点的 `WithCancel`**。

---

## 5. WithTimeout 对比 `time.After`（工程上为何优先 ctx）

### `time.After` / 本地 `select` 的局限

- 主要约束**当前这一小段代码**，**不易**作为「整条调用链的契约」传递。  
- 与「**父请求已取消**」「**手动取消**」等**不天然统一**（要自己再接 `Done()`）。  
- 难以让 **DB / RPC / 子 goroutine** 默认遵守**同一份**截止时间。

### `WithTimeout` 的优势

- **可逐层传递**同一 `ctx`，下游同一套 `Done()` / `Err()` / `Deadline()`。  
- **超时 + 手动取消 + 父级取消**合一，利于级联与资源回收。  
- 与 **Go 官方**对 HTTP、gRPC 等生态的用法一致。

**一句话：** `time.After` 偏**局部**超时；`WithTimeout` 偏**链路级**截止时间协作。

---

## 6. 如何正确判断超时错误

必须用 **`errors.Is`**（错误常被 `fmt.Errorf("...: %w", err)` 包装，`==` 不可靠）：

```go
import "errors"

if errors.Is(err, context.DeadlineExceeded) {
    // 超时：降级、重试策略、日志、指标
}
```

常量：**`context.DeadlineExceeded`**。若需区分「主动取消」，再用 **`context.Canceled`** 与 `errors.Is`。

---

## 7. 完整可运行示例

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// 模拟慢调用：500ms 才「完成」
func slowAPI(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(500 * time.Millisecond):
		return nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	err := slowAPI(ctx)
	if err != nil && errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("调用失败：超时了")
		return
	}
	if err != nil {
		fmt.Println("其它错误:", err)
	}
}
```

典型输出：

```text
调用失败：超时了
```

说明：`slowAPI` 在 `ctx` 200ms 到期时走 `ctx.Done()`，返回的 `ctx.Err()` 为 **`DeadlineExceeded`**。

---

## 8. 高频坑点（面试必考）

1. **不写 `defer cancel()`**：timer / 内部状态可能**晚释放**，高并发下放大成资源压力。  
2. **下游不感知 `ctx`**：既不传 `ctx`，也不监听 `Done()`，**超时形同虚设**。  
3. **父已取消 / 已超时**：子再包一层 `WithTimeout` 可能**立即**进入已取消状态，要理解 [上下文树](./03-Background()、TODO()、上下文树.md)。  
4. **只在最外层包超时、内层不用传入的 `ctx`**：链路中间仍可能**长时间阻塞**。  
5. **超时时间拍脑袋**：过短误杀、过长起不到保护；应贴近 **P99** 与业务 SLA。

---

## 9. 面试标准答案（完整版）

- `WithTimeout` 为请求/子任务设置**绝对时长上限**，到点**自动取消**，防阻塞与堆积。  
- 实现上基于 **`WithDeadline` + timer**，到点触发 `cancel`，关闭 `Done()` 协作下游退出。  
- **必须** `defer cancel()`，与手动取消、父取消一起保证资源及时回收。  
- 相对 **`time.After`**，更适合**全链路传递**与**与取消模型统一**。  
- 判断超时用 **`errors.Is(err, context.DeadlineExceeded)`**。

---

## WithDeadline 补充（极简）

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```

- 与 `WithTimeout` **同一套** `defer cancel()`、`Done()`、`Err()`；差别在参数是 **`time.Time`**。  
- 详见 [06-WithDeadline 精确到时间点的取消控制.md](./06-WithDeadline%20精确到时间点的取消控制.md)。

---

## 延伸阅读

- [02-context是什么、作用、核心接口.md](./02-context是什么、作用、核心接口.md)  
- [04-WithCancel 手动取消.md](./04-WithCancel%20手动取消.md)  
- [06-WithDeadline 精确到时间点的取消控制.md](./06-WithDeadline%20精确到时间点的取消控制.md)  
- [07-WithValue 数据透传.md](./07-WithValue%20数据透传.md)（勿用 Value 替代超时）  
- [09-context常见陷阱与反模式.md](./09-context常见陷阱与反模式.md)
