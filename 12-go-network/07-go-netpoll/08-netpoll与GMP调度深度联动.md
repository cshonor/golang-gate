# netpoll 与 GMP 调度深度联动

> **07-go-netpoll · Go netpoll 高并发核心**  
> 前置：[06-go-net-internals/07-pollDesc核心结构与原理](../06-go-net-internals/07-pollDesc核心结构与原理.md)、本目录 [06-Goroutine与netpoll调度](./06-Goroutine与netpoll调度.md)。调度细节还可对照仓库 **[07-GMP and channel](../../07-GMP%20and%20channel/README.md)**。

---

## 1. 一条完整链路（读路径，概念）

1. 业务调用 **`conn.Read(buf)`**（同步 API）。  
2. **`netFD.Read`** → **`internal/poll.FD.Read`**：若 **`EAGAIN`**，进入 **等待**。  
3. **`runtime_pollWait`**（或等价路径）把 **当前 G** 与 **该 fd 的就绪事件** 关联，**`gopark`** —— **G 不再占用 P 去跑业务**，但 **M 往往可继续找别的 G**（取决于阻塞点与是否占着 syscall 等）。  
4. **多路复用**（`epoll_wait` / `kevent` / `GetQueuedCompletionStatus`…）在 **调度循环**或**专用路径**中被驱动。  
5. **事件就绪**：runtime 把对应 **G** 标记为可运行，放入 **P 的本地队列或全局队列**，最终被 **M** 取出执行，`Read` 返回用户态。

> 用户观感：**「我这行 Read 卡住了」**；runtime 实际：**「卡的是 G，不是一定要卡死整个线程池」**。

---

## 2. 与 P、timer 的交界

- **P** 在查找可运行 G 时，可能顺带 **poll 网络**（实现细节随版本；有的版本在 **`findrunnable`** 路径附近调用 **`netpoll`**）。  
- **Deadline** 到期与 **网络就绪** 一样，都要能把 **G** 从等待里捞出来，因此 **timer 子系统** 与 **netpoll** 在调度上有交集（读源码时关注「谁先把 G 叫醒」）。

---

## 3. 和「同步阻塞」模型的对比

| 模型 | OS 线程占用 |
|------|----------------|
| 一线程一阻塞 `Read` | 线程睡眠，并发连接数 ≈ 线程数 |
| Go：`Read` 阻塞 G | 多在同一组 M 上复用，**线程数远小于连接数**（常见形态） |

误区：**goroutine 不是零成本**；海量仍要 **fd 上限、内存、调度延迟** 与 **业务背压**。

---

## 4. 源码阅读锚点（随版本检索）

- `runtime/netpoll*.go`：平台实现与 **`netpoll`** 入口。  
- `runtime/proc.go`：`findrunnable`、`schedule` 与 **netpoll 调用关系**。  
- `runtime/netpoll.go`（若存在）：通用包装。  
- `internal/poll`：`Read` 里 **`pollWait`** 的触发点。

---

## 5. 相邻篇目

| 篇目 | 内容 |
|------|------|
| [09-pollDesc等待队列与唤醒原理](./09-pollDesc等待队列与唤醒原理.md) | 等待队列与 `sudog` 类比 |
| [07-netpoll源码核心流程](./07-netpoll源码核心流程.md) | 事件循环细读 |
| [05-Go的Read与Write为什么看起来阻塞](./05-Go的Read与Write为什么看起来阻塞.md) | 用户态观感解释 |

---

## 下一篇

[09-pollDesc等待队列与唤醒原理.md](./09-pollDesc等待队列与唤醒原理.md)
