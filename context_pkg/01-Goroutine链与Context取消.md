# 01 - Goroutine 链与 Context 取消

在 Go 并发里，**Goroutine 链**和 **Context** 经常一起出现：前者描述「一次业务里层层派生的执行结构」，后者提供**可向下传播的取消、超时与请求级数据**。本文从定义、典型形态、泄漏风险到 Context 的**取消方向**，做一次系统整理。

> 配套阅读：[02-context是什么、作用、核心接口.md](./02-context是什么、作用、核心接口.md)、[04-WithCancel 手动取消.md](./04-WithCancel%20手动取消.md)、[03-Background()、TODO()、上下文树.md](./03-Background()、TODO()、上下文树.md)

---

## 一、什么是 Goroutine？

Goroutine 是由 **Go runtime** 调度的轻量级执行单元（常见说法：用户态协程），跑在 GMP 模型上。

### 1. 核心特性（面试口径）

- **栈小且可扩**：初始栈很小，可按需增长；与固定 MB 级 OS 线程栈对比，更易海量创建。
- **创建与切换成本低**：调度多在用户态完成（细节见 `GMP and channel` 笔记）。
- **并发规模大**：实际仍受内存、FD、业务逻辑约束，不是「无限开」。

### 2. 启动方式

```go
doTask()   // 当前 goroutine 同步执行
go doTask() // 新 goroutine 异步执行
```

---

## 二、什么是 Goroutine 链？

**Goroutine 链**：在同一条业务链路（例如一次 HTTP 请求）里，由 **父 goroutine 启动子 goroutine**，子再启动下一层，形成**有依赖、有传递关系**的执行结构。

### 1. 典型场景（HTTP）

```text
客户端请求
      ↓
[Goroutine A]  入口（解析、路由）
      ↓ 启动
[Goroutine B]  业务逻辑
      ↓ 启动
[Goroutine C]  访问数据库
      ↓ 启动
[Goroutine D]  访问缓存（如 Redis）
```

整条链路可记为 **A → B → C → D**（名字是示意，真实框架里可能是池化 worker，但「请求级子任务串联」的直觉相同）。

### 2. 代码示例 A：同步调用（仅说明调用关系，不是四条 goroutine）

下面 **没有** `go` 关键字，实际是**同一条 goroutine** 上的函数调用栈，用来对照「谁在调谁」；**不等价**于并发链。

```go
package main

import (
	"fmt"
	"time"
)

func taskD() {
	fmt.Println("运行 taskD -> 缓存")
	time.Sleep(100 * time.Millisecond)
}

func taskC() {
	fmt.Println("运行 taskC -> 数据库")
	taskD()
	time.Sleep(150 * time.Millisecond)
}

func taskB() {
	fmt.Println("运行 taskB -> 业务")
	taskC()
	time.Sleep(200 * time.Millisecond)
}

func taskA() {
	fmt.Println("运行 taskA -> 入口")
	taskB()
	time.Sleep(250 * time.Millisecond)
}

func main() {
	fmt.Println("主 goroutine 启动")
	taskA()
	fmt.Println("主 goroutine 退出")
}
```

### 3. 代码示例 B：真正的多 goroutine 链（每层 `go`）

要与「链」和 Context 取消结合，通常每层会 `go` 起下游，并用 `WaitGroup` / channel / HTTP handler 生命周期保证不提前退出 `main`：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func taskD(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine D: 缓存")
	time.Sleep(100 * time.Millisecond)
}

func taskC(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine C: 数据库")
	wg.Add(1)
	go taskD(wg)
	time.Sleep(150 * time.Millisecond)
}

func taskB(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine B: 业务")
	wg.Add(1)
	go taskC(wg)
	time.Sleep(200 * time.Millisecond)
}

func taskA(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine A: 入口")
	wg.Add(1)
	go taskB(wg)
	time.Sleep(250 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go taskA(&wg)
	wg.Wait()
}
```

---

## 三、核心问题：生命周期与泄漏

子 goroutine **不会**因为父函数返回而自动结束；若上游已不需要结果（用户断连、超时），下游仍在跑，会造成：

- CPU / 内存占用
- DB、Redis 连接池被拖住
- 雪崩式积压

因此需要**可传播的结束条件**；标准答案之一是 **Context**（另有 done channel、errgroup 等组合）。

---

## 四、Context：向下游传递「该停了」

`context.Context` 常用于三件事：

1. **取消（Cascade）**：从某个 context 根上发出取消，**所有派生自它的子 context** 都会进入 `Done`。
2. **超时 / 截止时间**：`WithTimeout` / `WithDeadline` 到期后自动 `cancel`。
3. **请求级 Value**：trace id 等（慎用，见 [06-WithValue 数据透传.md](./06-WithValue%20数据透传.md)）。

### 1. 取消方向（必须精确）

- **在父 context 上取消**（或根超时）：**整条子树**上监听该链路的 goroutine 都应退出。
- **仅取消某一个子 context**：只有拿**这个子 context**（及其再派生的后代）的 goroutine 会结束；**兄弟、祖先**若仍握着**未取消的父 context**，**不会**因此被取消。

口诀：**cancel 作用在「以该 ctx 为根的子树」，不是「杀整条进程里所有 goroutine」**。

### 2. 示例：C 只取消自己派生给 D 的 context

下面 **D** 使用 **`ctxD`**（C 用 `WithCancel(ctx)` 派生）；**C 调用 `cancelD()` 只会取消 `ctxD`**，**A/B 手里的 `ctx` 仍存活**（假设根上没人 cancel）。

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func D(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("D: 收到取消，退出")
	case <-time.After(1 * time.Second):
		fmt.Println("D: 正常跑完（不应在本例出现）")
	}
}

func C(ctx context.Context) {
	ctxD, cancelD := context.WithCancel(ctx)
	go D(ctxD)

	fmt.Println("C: 开始")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("C: cancel 下游 ctxD（不取消父 ctx）")
	cancelD()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("C: 结束")
}

func B(ctx context.Context) {
	go C(ctx)
	fmt.Println("B: 开始")
	select {
	case <-ctx.Done():
		fmt.Println("B: 收到取消")
	case <-time.After(300 * time.Millisecond):
		fmt.Println("B: 未收到取消，继续")
	}
}

func A(ctx context.Context) {
	go B(ctx)
	fmt.Println("A: 开始")
	select {
	case <-ctx.Done():
		fmt.Println("A: 收到取消")
	case <-time.After(300 * time.Millisecond):
		fmt.Println("A: 未收到取消，继续")
	}
}

func main() {
	ctx := context.Background()
	go A(ctx)
	time.Sleep(800 * time.Millisecond)
	fmt.Println("main: 结束")
}
```

**典型输出**（顺序可能因调度略有交错）：

- A、B 打印「未收到取消」——符合预期：它们监听的是**未被 cancel 的 `ctx`**。
- D 打印「收到取消」——因为 **`ctxD` 被 `cancelD()`**。

若希望「用户断开则整条请求链结束」，应在**根**（如 HTTP 请求的 `r.Context()`）上超时或 cancel，并把**同一个派生链**一路传给 B、C、D，而不是在中间另起互不关联的 context。

---

## 五、最佳实践（简）

1. **`ctx` 作第一个参数**：`func F(ctx context.Context, ...)`，链路一路传下去。
2. **`defer cancel()`**：对 `WithCancel` / `WithTimeout` 返回的 `cancel` 务必在拿到后 `defer`，避免泄漏 timer 或内部资源（具体见 04/05 篇）。
3. **在 I/O 与阻塞点 `select ctx.Done()`**：网络读写用带 `ctx` 的 API（如 `http.NewRequestWithContext`）。
4. **goroutine 内 panic**：在边界用 `recover` 或保证不泄漏（与错误处理目录呼应）；Context 不负责兜 panic。

---

## 六、小结

| 要点 | 结论 |
|------|------|
| Goroutine 链 | 业务上父子相继启动的并发结构；示意代码里要区分**同步调用栈** vs **`go` 真并发** |
| 泄漏根因 | 子 goroutine 默认不随父返回而结束 |
| Context | 取消/超时沿**派生树**传播；**子树 cancel 不自动取消兄弟与祖先** |
| 工程 | 根 ctx 对齐请求生命周期；中间应用 `WithCancel`/`WithTimeout` 划分子任务 |

更多 API 细节与中间件写法见本目录 [README.md](./README.md) 索引。
