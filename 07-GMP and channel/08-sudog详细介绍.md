# sudog 详细介绍

`sudog` 是 Go runtime 里封装**阻塞中 Goroutine** 的核心结构，是 channel、部分同步原语与 GMP 协作的底层组件之一。

## 1. 本质与作用

### 本质

`sudog` 定义在 `runtime` 包中，表示**等待队列上的一个节点**：把某个 **G** 与「等谁、等什么、数据在哪」等信息绑在一起，挂进 `sendq` / `recvq` 等 **waitq**。

### 作用

- G 因 **channel 收发**、**`sync.Mutex`**、**`sync.WaitGroup`** 等阻塞时，常以 `sudog` 形式进入对应等待队列（如 channel 的 `sendq` / `recvq`）。
- 保存等待状态、数据指针、关联对象等，供后续 **唤醒与继续调度**。
- **用户态**完成阻塞/唤醒链路，避免每次阻塞都进内核（与「G 让出 M」的调度模型配合）。

## 2. 核心结构（源码级）

`runtime/runtime2.go` 中核心字段示意（完整定义以当前版本为准）：

```go
type sudog struct {
	g    *g             // 被封装的 Goroutine
	next *sudog
	prev *sudog         // 双向链表，串成 waitq

	elem unsafe.Pointer // 待发送/待接收的数据地址（channel 等场景）

	c *hchan // 关联的 channel（channel 场景）

	// 超时、唤醒条件、其他辅助字段...
}
```

### 字段说明

| 字段 | 作用 |
|------|------|
| `g` | 指向被阻塞的 G，是 sudog 的语义核心 |
| `next` / `prev` | 双向链表，挂入 `waitq`（`first`/`last` 管理） |
| `elem` | 数据指针，channel 路径上用于拷贝收发 |
| `c` | 关联的 `hchan`，channel 阻塞场景 |

## 3. 工作流程（以 channel 为例）

### 阻塞：G → sudog → 入队

当 channel 操作无法立刻完成（如无缓冲对端未就绪、有缓冲满/空等）：

1. 构造 `sudog`，令 `sudog.g` 指向当前 G。
2. 将 `sudog` 挂入 `sendq` 或 `recvq`。
3. G 进入 **`_Gwaiting`**，从当前 M 上剥离，M 继续调度其他 G。

### 唤醒：sudog 出队 → G 可运行

当条件满足（对端就绪、缓冲区有空间/有数据等）：

1. 从 `waitq` 取出对应 `sudog`（常为队首策略，细节见 `chan.go`）。
2. 通过 `sudog.g` 拿到 G，完成必要的数据拷贝（channel 场景）。
3. G 置为 **`_Grunnable`**，放入 P 本地队列或经全局路径调度。
4. `sudog` 往往回收到 **缓存池**，减少分配开销。

## 4. 特性（面试常提）

1. **双向链表**：`next`/`prev` 支持 O(1) 入队/出队（在已有节点前提下）。
2. **对象复用**：runtime 维护 `sudog` 池，降低频繁分配。
3. **用户态路径**：与 GMP 协作，多数逻辑在用户态完成。
4. **通用等待节点**：除 channel 外，也用于 `sync.Cond`、`select` 等路径（以源码为准）。

## 5. 与 G、channel 的关系

| 组件 | 关系 |
|------|------|
| **G** | sudog 封装的对象；阻塞主体仍是 G |
| **channel** | `sendq`/`recvq` 是 **sudog 链表**（`waitq`），不是裸 G 指针队列 |
| **GMP** | sudog 是 G **离开运行栈、进入等待结构**时的关键中间表示 |

## 6. 易混淆点

1. **sudog ≠ G**：sudog 是排队用的**节点**；同一时刻一个 G 在一条等待链上通常只对应一个有效 sudog（逻辑上不重复入队）。
2. **临时性**：唤醒后 sudog 多回收到池中，不长期悬挂。
3. **waitq 里是 sudog**：因此「channel 等的是 sudog，sudog 里才有 G」。
4. **超时 / select**：可在 sudog 及相关路径上携带超时与唤醒条件（实现见 `runtime`）。

## 7. 一句话

**sudog 是 runtime 里表示「阻塞 G 在等待队列中的节点」，是 channel 等同步机制与调度器之间的桥梁。**

---

## 复习速记

| 考点 | 一句话 |
|------|--------|
| 是什么 | 等待队列上的节点，包着 G |
| 放哪 | `sendq` / `recvq` 等 `waitq`（sudog 链表） |
| 为啥快 | 用户态协作 + sudog 池复用 |

## 延伸阅读

- channel 阻塞与 `hchan`：[10-channel阻塞原理.md](./10-channel阻塞原理.md)
- G 与 sudog 总览：[03-G、M、P、sudog 四个结构体的关系.md](./03-G、M、P、sudog%20四个结构体的关系.md)
