# `net.Dial` 底层流程（客户端建连）

> **06-go-net-internals · Go net 包源码级理解**  
> 上篇：[04-UDPConn结构](./04-UDPConn结构.md)；内核握手见 [03-linux](../03-linux-networking/04-三次握手内核做了什么.md)。

---

## 一、总览（TCP）

`net.Dial("tcp", "host:port")` 在概念上经历：

1. **解析地址**：字符串 → **`TCPAddr`**；若为域名则 **DNS**（可能很慢）。  
2. **`socket`**：创建 fd，**通常设为非阻塞** 以接入 netpoll。  
3. **`connect`**：发起 **三次握手**；非阻塞下可能 **立刻 `EINPROGRESS`**，由 **「fd 可写 + SO_ERROR」** 判定最终结果。  
4. **封装 `netFD`/`TCPConn`**：初始化 **`internal/poll.FD`**，注册到 **netpoll**。

> 精确函数名/分支以 **`src/net/dial.go`、`src/net/sock_*.go`** 为准；双栈、Happy Eyeballs 等会拉长调用图。

---

## 二、与 `net.Dialer` 的关系

生产推荐 **`(&net.Dialer{...}).DialContext`**：

- **`Timeout`**：限制 **连接建立** 阶段。  
- **`KeepAlive`**：打开/配置 **TCP keepalive**（见 [04-tcp/08](../04-tcp/08-TCP心跳保活.md)）。  
- **`Control`**：在 **`syscall.RawConn`** 上执行 **`setsockopt`** 等黑科技。

---

## 三、超时从哪里来？

- **不是**简单「只依赖内核 `SO_RCVTIMEO`」一条路径；Go 常在 **`connect`/`pollWait` 路径** 上叠 **`Context` + deadline**。  
- 读写到 **`Conn`** 后，仍靠 **`SetReadDeadline`/`SetWriteDeadline`**（见 [08](./08-网络超时与Deadline底层实现.md)）。

---

## 四、极简总结

- **`Dial`** = **解析 → socket → connect（握手）→ netFD/poll 初始化**。  
- **非阻塞 fd + netpoll** 让 **`Dial`/`Read` 阻塞 G 而非卡死线程池**（与 **07** 呼应）。

---

## 导航

- 上一篇：[04-UDPConn结构](./04-UDPConn结构.md)  
- 下一篇：[06-net.Listen底层流程](./06-net.Listen底层流程.md)  
- 桥梁：[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)
