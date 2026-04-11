# Goroutine 与 netpoll 调度

> **07-go-netpoll · Go netpoll 高并发核心**  
> 上篇：[05-Go的Read与Write为什么看起来阻塞](./05-Go的Read与Write为什么看起来阻塞.md)。

---

## 一、`netpoll` 在调度循环里的位置（概念）

- **`findrunnable`**：当 **P 本地/全局队列**暂时没 G 时，可能调用 **`netpoll(false)`** 拉一批 **就绪网络 G**（实现细节随版本调整，抓 **`netpoll` 调用点** 即可）。  
- **`startlockedm`/`sysmon`** 等路径也可能与 **netpoll 抢占式唤醒** 有关（浅读即可）。

---

## 二、批量就绪与公平性

一次 **`epoll_wait`** 可能返回 **多个 fd** → **多个 G 同时 `ready`**：

- **优点**：**吞吐高**（一次 syscall 处理多事件）。  
- **风险**：若业务 handler **极重**，可能放大 **尾延迟**——需要 **隔离（worker pool）** 与 **超时**。

---

## 三、与异步抢占（Go 1.14+）

- **异步抢占**让 **纯 Go 死循环** 更难霸占 P；**网络等待**仍以 **poll + gopark** 为主轴理解。  
- **不要**把抢占当成可以取消 **`Read` 阻塞**的银弹：**该设 deadline 仍要设**。

---

## 四、反模式

- **`runtime.LockOSThread`** 滥用：可能 **降低并行度**、与 **网络 poll 模型** 冲突于某些图形/实时场景——网络服务默认别玩。

---

## 五、极简总结

- **netpoll** 把 **「等网络」** 融入 **G 调度** 的 **找活干** 路径。  
- **批量唤醒** = **吞吐换尾延迟风险**，要工程折中。

---

## 导航

- 上一篇：[05-Go的Read与Write为什么看起来阻塞](./05-Go的Read与Write为什么看起来阻塞.md)  
- 下一篇：[07-netpoll源码核心流程](./07-netpoll源码核心流程.md)  
- 深化：[08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)
