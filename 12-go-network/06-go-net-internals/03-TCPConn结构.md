# `TCPConn` 结构（TCP 连接实体）

> **06-go-net-internals · Go net 包源码级理解**  
> 上篇：[02-Listener接口](./02-Listener接口.md)

---

## 一、核心定义

**`*net.TCPConn`** 是 **`net.Conn` 在 TCP 上的具体实现类型**（对用户一般以 **`net.Conn` 接口值** 使用）。

- 对外：实现 **`Read`/`Write`/`Close`、Deadline、地址`**。  
- 对内：持有 **`net.conn`**，其核心是 **`netFD`**（文件描述符 + poll 状态 + 地址元数据）。

> 源码位置随版本变化；阅读锚点：**`src/net/tcpsock.go`**、**`src/net/net.go`**（`conn` / `netFD`）。

---

## 二、结构关系（概念图）

```text
TCPConn
  └─ 内嵌 net.conn
        └─ netFD
              └─ internal/poll.FD   ← 与 runtime netpoll 的边界
                    └─ fd（os 句柄）
```

- **`netFD`**：**`net` 包内** 对「系统 fd + 生命周期 + 超时」的封装。  
- **`internal/poll.FD`**：**等待可读/可写、注册到 netpoll、处理 `EAGAIN`** 的枢纽（口语上常与 **pollDesc** 混称，见 [07](./07-pollDesc核心结构与原理.md)）。

---

## 三、常见 `TCPConn` 专有方法

- **`SetNoDelay` / `SetKeepAlive` / `SetLinger` / `SetReadBuffer`…**：对 **`setsockopt`** 的封装，与 [03-linux 08](../03-linux-networking/08-socket缓冲区SO_RCVBUF与SO_SNDBUF.md) 呼应。

---

## 四、与 netpoll 的关系（一句话）

**`TCPConn.Read` 无数据且 fd 非阻塞** → **`internal/poll` 等待路径** → **`runtime_pollWait` + `gopark`** → **G 挂起**；就绪后 **netpoll 唤醒 G** 继续读（详见 **07** 目录）。

---

## 五、极简总结

- **`TCPConn`** = **`net.Conn` 的 TCP 肉身**。  
- **`netFD` + `poll.FD`** 是 **`net` ↔ `runtime`** 的桥梁。  
- **调优 syscall** 多从 **`TCPConn`** 往下打。

---

## 导航

- 上一篇：[02-Listener接口](./02-Listener接口.md)  
- 下一篇：[04-UDPConn结构](./04-UDPConn结构.md)  
- 桥梁：[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)
