# netpoll 源码核心流程（怎么读）

> **07-go-netpoll · Go netpoll 高并发核心**  
> 上篇：[06-Goroutine与netpoll调度](./06-Goroutine与netpoll调度.md)。

---

## 一、建议带着问题读

一次 **`Accept` → 返回 `Conn` → 第一次 `Read`**，栈上依次出现哪些 **`runtime`/`internal/poll`/`net`** 符号？

---

## 二、Linux 上常见符号链（概念）

| 阶段 | 方向 | 关键词 |
|------|------|--------|
| 初始化 | runtime 启动 | `netpollinit` |
| 打开 fd | `poll.FD` 初始化 | `runtime_pollOpen` |
| 等待 | `EAGAIN` | `runtime_pollWait` |
| 轮询 | 调度路径 | `epollwait` / `netpoll` |
| 就绪 | 内核事件 | `netpollready` / `pdReady`（名称随版本） |
| 关闭 | `Close` | `runtime_pollClose` |

> **不要背函数名**：升级 Go 后常改名；掌握 **「open/wait/ready/close」四段**。

---

## 三、`internal/poll` 边界

- **`FD`** 持有 **`pd`**（`pollDesc`）指针，**`Read/Write`** 在 **`EAGAIN`** 时进入 **`pollWait`**。  
- **`net`** 不直接 `epoll_ctl`；**`poll` 包**通过 **`runtime`** 提供的 **`runtime_poll*`** 完成。

---

## 四、与 06 的对照表

| 06 | 07（本文） |
|----|------------|
| [07-pollDesc](../06-go-net-internals/07-pollDesc核心结构与原理.md) 讲 **对象关系** | 本文讲 **runtime 侧事件泵** |
| [08 Deadline](../06-go-net-internals/08-网络超时与Deadline底层实现.md) | `netpoll` 与 **timer** 的交汇在 **08/09** 深化 |

---

## 五、极简总结

- 读源码：**`FD.Read` → `pollWait` → `runtime.netpoll*`** 三条锚点。  
- **版本差异大**：抓 **语义** 不背 **行号**。

---

## 导航

- 上一篇：[06-Goroutine与netpoll调度](./06-Goroutine与netpoll调度.md)  
- 下一篇：[08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)
