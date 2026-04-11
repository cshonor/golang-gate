# 04 - WithCancel：手动取消｜完整版详细笔记

## 1. 底层定义与作用

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

- 接收一个**父上下文 `parent`**，派生一个新的子上下文；
- 返回：
  1. 子 **`ctx`**：继承父上下文的截止时间、取消信号、`Value` 键值；
  2. **`cancel` 取消函数**：`func()`，**手动触发该子上下文及其后代**的取消。
- 核心作用：**手动控制一组 goroutine 的生命周期**，主动下发取消信号，便于资源回收与协程退出。

---

## 2. 使用场景（详细拆解）

1. **请求链路中断**：客户端断开、连接关闭，主动取消下游所有协程；  
2. **任务提前终止**：业务报错、前置校验失败，无需再跑后续逻辑；  
3. **后台常驻协程管控**：启停消费者、定时任务、异步 worker；  
4. **多协程统一退出**：一处调用 `cancel()`，所有监听 `ctx.Done()` 的协程协作退出。

---

## 3. 执行机制（理解上下文树）

1. **继承**：子 ctx 挂在父链上——**父 ctx 取消，子 ctx 会跟着取消**（见 [03-Background()、TODO()、上下文树.md](./03-Background()、TODO()、上下文树.md)）。  
2. **传播**：对该子 ctx 调用 `cancel()`，会取消**该子 ctx 及其所有后代**，**不会**取消父 ctx、**不会**取消「兄弟」子树（彼此独立的 `WithCancel` 分支）。  
3. **取消流程（心智模型）**：  
   - 调用 `cancel()` → 关闭该 ctx 的 `Done()` channel；  
   - 阻塞在 `<-ctx.Done()` 或 `select` 里的 goroutine 被唤醒；  
   - `ctx.Err()` 变为 **`context.Canceled`**（与超时区分见 [05-WithTimeout 超时控制.md](./05-WithTimeout%20超时控制.md)）。

---

## 4. 标准完整代码模板（逐行解析）

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到取消信号，协程退出，原因：", ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("处理任务：%d\n", job)
		}
	}
}

func main() {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)

	// 规范：拿到 cancel 后尽快 defer，保证任意 return / panic 路径上都会释放子树资源
	defer cancel()

	jobs := make(chan int, 5)
	go worker(ctx, jobs)

	jobs <- 1
	jobs <- 2

	// 业务上「手动」触发取消（与 defer 再调一次 cancel 均为幂等，无害）
	cancel()

	// 留一点时间让 worker 打印退出日志；main 返回时 defer cancel() 仍会执行一次
	time.Sleep(200 * time.Millisecond)
	close(jobs)
}
```

### 关键要点解析

1. **`defer cancel()`**：强规范——函数退出（含 panic 后的 defer 阶段）仍会执行，避免子树泄漏。  
2. **必须监听 `<-ctx.Done()`**：协程内用 `select` 协作退出；否则 `cancel()` 无法让 goroutine 停下。  
3. **`cancel` 可多次调用**：幂等，不会重复关闭同一 channel 引发 panic。

---

## 5. 核心底层原理（面试够用版）

1. `WithCancel` 内部对应**可取消的 context 实现**（实现细节随 Go 版本演进，不必背源码行号）。  
2. **`Done()`** 对外暴露**只读** `<-chan struct{}`：关闭表示已取消；调用方只监听，不自行 `close`。  
3. 调用 `cancel()`：关闭 channel、设置 `err`，并**沿子节点向下**传播取消（父、兄弟不受影响）。

---

## 6. 高频坑点 + 问题详解

### 坑 1：忘记 `defer cancel()`

- **后果**：子树与关联 goroutine / 定时器可能**晚释放或不释放**，表现为泄漏或资源占用偏高。  
- **对策**：`WithCancel` / `WithTimeout` 等**成对出现**：`ctx, cancel := ...` 后紧跟 `defer cancel()`。

### 坑 2：只 `cancel`，下游不监听 `ctx.Done()`

- **后果**：取消信号发出，业务仍卡在阻塞 IO / 死循环里，协程不退出。  
- **要点**：Context 是**协作式取消**，不是操作系统式「强杀 goroutine」。

### 坑 3：误以为 `cancel` 会关掉「父」上下文

- **事实**：子 ctx 的 `cancel` **只作用于该子树**；**不能**通过子 `cancel` 去取消父 ctx。  
- **记忆**：取消沿派生方向**向下**传播；兄弟分支互不影响。

### 坑 4：`main` 末尾 `select {}` 死阻塞

- **后果**：`main` 永不返回，**`defer cancel()` 要等到进程退出才跑**（演示时容易误解「defer 已经执行过」）。  
- **对策**：用 `time.Sleep`、`sync.WaitGroup` 或正常 return，保证 defer 能走完。

---

## 7. 面试标准答案（完整版）

1. `WithCancel` 从父 ctx 派生**可手动取消**的子 ctx，并返回 `CancelFunc`。  
2. 调用 `cancel()` 会关闭 `Done`，当前及**后代** ctx 收到信号，`Err()` 为 `context.Canceled`。  
3. 取消是**协作式**的：下游必须监听 `ctx.Done()`（或传递 ctx 到支持取消的 API）。  
4. 开发上应 **`defer cancel()`**，避免泄漏，保证链路资源可回收。

---

## 8. 关键区分补充

| 项 | 说明 |
|----|------|
| **`Background()`** | 根节点，**不**携带业务取消、**不**超时；不是 `WithCancel` 的「关闭对象」。 |
| **`WithCancel` 子 ctx** | 具备手动取消能力，是业务里**管控一组 goroutine 生命周期**的基础 building block。 |

---

## 延伸阅读

- [02-context是什么、作用、核心接口.md](./02-context是什么、作用、核心接口.md)  
- [05-WithTimeout 超时控制.md](./05-WithTimeout%20超时控制.md)  
- [08-context常见陷阱与反模式.md](./08-context常见陷阱与反模式.md)
