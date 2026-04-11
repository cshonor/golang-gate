# poll 等待队列与唤醒原理（承接 06）

> **07-go-netpoll · Go netpoll 高并发核心**  
> 上篇：[08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)。  
> 与 channel 的 **`sudog` + `waitq`** 心智类似：都是 **「把等某件事的 G 挂起来，再由 runtime 在某个时刻叫醒」**（见 [07-GMP/08-sudog](../../07-GMP%20and%20channel/08-sudog核心数据结构与作用详解.md)）。

---

## 一、等待发生时（简化）

1. **`pollWait`** 发现 **fd 未就绪** 且 **未关闭/未超时**。  
2. 将 **当前 G** 关联到 **该 `pollDesc`/`FD` 的等待结构**（实现细节：可能与 **`sudog` 池** 复用）。  
3. **`gopark`**：**G 离开运行队列**；**M** 继续调度其他 G。

---

## 二、唤醒发生时（简化）

1. **内核**通过 **`epoll_wait`/…** 报告 **fd 可读/可写**。  
2. **`netpoll`** 扫描就绪列表，定位 **`pollDesc`**。  
3. **`pdReady`/`netpollready`**：把关联的 **G** 标记为 **`Grunnable`** 并放入队列。  
4. **G** 恢复执行后，**`Read`/`Write`** 路径重试 syscall，直到 **成功/EAGAIN/错误**。

**deadline 到期** 与 **对端 `RST`** 等，也会走 **「结束等待 → 返回错误」** 的同类框架。

---

## 三、与 06 的双向闭环

| 06 | 07（本文） |
|----|------------|
| [07-pollDesc](../06-go-net-internals/07-pollDesc核心结构与原理.md) 讲 **对象与 API 边界** | 本文讲 **G 睡/醒与队列** |

---

## 四、面试自检

- **`EAGAIN` 之后，是 M 睡还是 G 睡？**（默认网络路径：**G 睡**）  
- **为什么 `Read`「阻塞」不会拖死所有 M？**（**G 让出 P，M 跑别的 G**）  
- **deadline 与可读，谁先到谁赢？**（**都能 `ready` G**，返回错误由 `poll` 路径判定）

---

## 五、极简总结

- **等待队列** = **fd ↔ G** 的胶水。  
- **netpoll** = **批量 `ready`** 的性能关键。

---

## 导航

- 上一篇：[08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)  
- 下一篇：[10-netpoll常见坑与优化](./10-netpoll常见坑与优化.md)
