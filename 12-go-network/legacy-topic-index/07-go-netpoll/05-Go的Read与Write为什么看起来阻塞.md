# Go 的 `Read`/`Write` 为什么「看起来阻塞」？

> **07-go-netpoll · Go netpoll 高并发核心**  
> 上篇：[04-Go的IO模型到底是什么](./04-Go的IO模型到底是什么.md)；桥梁：[06/07-pollDesc](../06-go-net-internals/07-pollDesc核心结构与原理.md)。

---

## 一、用户看到的现象

调用 **`conn.Read`** 后，如果暂时没有数据，**当前 goroutine 不再往下执行**——这就是日常说的 **「阻塞」**。

---

## 二、runtime 实际做了什么？

1. **`internal/poll.FD.Read`** 调 **`syscall.Read`**。  
2. **非阻塞 fd** 无数据 → **`EAGAIN`**。  
3. 进入 **`pollWait`** → **`runtime_pollWait`**：  
   - 记录 **当前 G** 与该 **fd 事件** 的关联；  
   - **`gopark`**：**G 变成等待态**，从 **P 的可运行集合** 消失。  
4. **M** 继续 **`findrunnable` → `schedule`** 跑别的 G（常见路径；若卡在 **阻塞 syscall/cgo** 则另论）。  
5. **数据就绪**：**`netpoll`** 把 G **`ready`** → G 再次运行 → **重试 `Read`** → 返回字节数。

---

## 三、`Write` 对称心智

发送缓冲满时同样可能 **`EAGAIN`** → **等待可写** → 唤醒后继续写（注意：**短写**与 **`SetWriteDeadline`**）。

---

## 四、和 `SetDeadline` 的关系

- **deadline 到期** 也会 **`ready` G**，`Read`/`Write` 以 **`timeout` 类错误**返回（见 [06/08](../06-go-net-internals/08-网络超时与Deadline底层实现.md)）。

---

## 五、极简总结

- **阻塞的是 G，不是「每个连接一个永远睡死的 OS 线程模型」**（对比经典 BIO 线程池）。  
- **关键 syscall**：**`read`/`write` + `epoll_wait`（批量）** 的分工。

---

## 导航

- 上一篇：[04-Go的IO模型到底是什么](./04-Go的IO模型到底是什么.md)  
- 下一篇：[06-Goroutine与netpoll调度](./06-Goroutine与netpoll调度.md)
