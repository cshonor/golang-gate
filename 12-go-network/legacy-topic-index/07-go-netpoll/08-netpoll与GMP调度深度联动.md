# netpoll 与 GMP 调度深度联动

> **07-go-netpoll · Go netpoll 高并发核心**  
> 前置：[06-go-net-internals/07-pollDesc](../06-go-net-internals/07-pollDesc核心结构与原理.md)、本目录 [06-Goroutine与netpoll调度](./06-Goroutine与netpoll调度.md)。GMP 系统化：[07-GMP and channel](../../../07-GMP%20and%20channel/README.md)。

---

## 一、读路径闭环（必须能画）

1. **`conn.Read(buf)`**（业务 goroutine **G1** 运行在 **M** 上，持有 **P**）。  
2. **`netFD.Read` → `poll.FD.Read`**：`syscall.Read` 返回 **`EAGAIN`**。  
3. **`runtime_pollWait('r')`**：把 **G1** 记录到 **pollDesc 等待结构**，执行 **`gopark`**。  
4. **效果**：**G1 不再占用 P 的 runqueue**；**M** 通过 **`schedule`** 去拉 **G2** 继续跑——这就是 **「网络阻塞不等价于线程池全员睡眠」** 的核心。  
5. **网卡数据到达** → **`epoll_wait` 返回** → **`netpoll`** 遍历就绪表 → **`netpollready`** → **G1 变 `Grunnable`** → 进入 **P 本地队列或全局队列**。  
6. 某 **M** 再次 **`execute(G1)`**，`Read` 重试 syscall，返回 **`n>0`**。

> **边界**：若 **`Read` 内部走了长时间阻塞且不被 Go 包装的路径**（少见）或 **cgo 阻塞**，**M** 可能被占用较久——网络 `net` 路径通常不属此类。

---

## 二、`P`、`timer`、`netpoll` 的交界

- **网络就绪** 与 **deadline 到期** 都要 **`ready(G)`**。  
- **timer**：`runtime` 时间堆在 **`findrunnable`/`checkTimers`** 等路径与 **netpoll** 交错执行——读源码时关注 **「谁先唤醒」** 不如关注 **「最终 G 都会 runnable」**。

---

## 三、与「同步阻塞 API」的对照

| 模型 | 卡住的单位 |
|------|------------|
| BIO 一线程一 `recv` | **OS 线程** |
| Go `net` 默认 | **G**（常见） |

---

## 四、`findrunnable` 里 `netpoll` 的意义（口述）

当 **本地队列空**、**全局队列也暂时抢不到活**时，**先 `netpoll` 一下** 等价于 **「顺便把已到期的网络事件捞进来」**，减少 **纯空转** 与 **延迟**（实现细节随版本）。

---

## 五、极简总结

- **`gopark`/`ready`** 是 **G 与 netpoll** 的铆钉。  
- **理解 `P` 如何在没活时找网络事件**，面试就超过一半人。

---

## 导航

- 上一篇：[07-netpoll源码核心流程](./07-netpoll源码核心流程.md)  
- 下一篇：[09-pollDesc等待队列与唤醒原理](./09-pollDesc等待队列与唤醒原理.md)
