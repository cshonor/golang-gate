# pollDesc / `internal/poll`：net 与 netpoll 的桥梁

> **06-go-net-internals · Go net 包源码级理解**  
> 读完本篇再读 [07-go-netpoll](../07-go-netpoll/02-netpoll是什么.md)，可把 **「`net.Conn.Read` 一行代码」→「为何不占满 OS 线程」** 串成闭环。字段名与文件路径以本机 **`$GOROOT/src`** 为准。

---

## 1. 为什么需要这一层

- **`net` 包**只描述「连接、监听、地址」等 API；真正 **fd 上的阻塞/非阻塞、等待可读可写、超时** 必须落到 **`internal/poll`**（历史上常被口语称为 **pollDesc 机制**）再与 **`runtime` netpoll** 对接。  
- 没有这一层，**06** 里讲的 `TCPConn`、`netFD` 与 **07** 里的 `epoll` 唤醒 **G** 会断档。

---

## 2. 心智模型：`netFD` → `poll.FD` → `runtime_poll*`

典型链路（Linux，概念上）：

1. `*net.TCPConn` 内嵌 **`net.conn`**，持有 **`net.netFD`**。  
2. **`netFD`** 内嵌或组合 **`internal/poll.FD`**（封装了 **fd 号、锁、读写状态、与 runtime 的注册句柄** 等）。  
3. 当 `Read`/`Write` 遇到 **`EAGAIN`**（非阻塞 fd 暂无数据/不可写）时，进入 **`poll` 包的等待路径**（如 `waitRead` / `waitWrite`），最终调用 **`runtime_pollWait`** 一类入口，把 **当前 G** 与 **该 fd 的就绪事件** 绑定。  
4. **netpoll** 在调度点或阻塞路径上被唤醒后，把就绪的 **G** 重新推入可运行队列（详见 **07** 目录 [08-netpoll与GMP调度深度联动](../07-go-netpoll/08-netpoll与GMP调度深度联动.md)）。

> 记忆：**`net` 管「业务 fd」，`internal/poll` 管「等不等、怎么等」，`runtime netpoll` 管「多路之一就绪后叫醒谁」。**

---

## 3. 生命周期（简图）

```text
Dial / Listen / Accept 成功
    → 创建 netFD + 初始化 poll.FD（注册到 netpoll）
Read / Write
    → 快路径：直接 syscall 成功则返回
    → 慢路径：EAGAIN → runtime_pollWait → G park
就绪事件到达
    → runtime 标记可读/可写 → 唤醒对应 G
Close
    → 从 netpoll 注销 → 关闭 fd → 释放 poll 状态
```

与 **关闭、泄漏** 相关见本目录 [10-连接关闭与资源泄漏排查](./10-连接关闭与资源泄漏排查.md)。

---

## 4. 源码阅读顺序（建议）

1. `src/net/net.go`、`src/net/tcpsock*.go`：`TCPConn` / `netFD` 字段。  
2. `src/internal/poll/fd_*.go`：`FD.Read` / `Write` 与 `wait`。  
3. `src/runtime/netpoll_*.go`：平台多路复用封装。  
4. 回到 **07**：[05-Go的Read与Write为什么看起来阻塞](../07-go-netpoll/05-Go的Read与Write为什么看起来阻塞.md)、[09-pollDesc等待队列与唤醒原理](../07-go-netpoll/09-pollDesc等待队列与唤醒原理.md)。

---

## 5. 与相邻篇目

| 篇目 | 关系 |
|------|------|
| [03-TCPConn结构](./03-TCPConn结构.md) | 对象侧：谁在持有 `netFD` |
| [05-net.Dial底层流程](./05-net.Dial底层流程.md) | 创建侧：`poll` 初始化时机 |
| [06-net.Listen底层流程](./06-net.Listen底层流程.md) | 监听侧：`Accept` 新 fd 的 `poll` 注册 |
| [08-网络超时与Deadline](./08-网络超时与Deadline底层实现.md) | `SetDeadline` 如何打断 `pollWait` |

---

## 下一篇（工程向）

[08-网络超时与Deadline底层实现.md](./08-网络超时与Deadline底层实现.md)
