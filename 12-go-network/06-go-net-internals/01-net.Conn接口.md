# `net.Conn` 接口（网络连接顶层抽象）

> **06-go-net-internals · Go net 包源码级理解**  
> 前置：[01-io-fundamentals](../01-io-fundamentals/01-什么是IO.md)、[04-tcp](../04-tcp/06-Go-TCP服务端.md)。本篇回答：**业务代码操作的「连接」在类型系统里长什么样**。

---

## 一、核心定位

`net.Conn` 是 **面向流式网络连接** 的 **接口**：屏蔽 TCP / Unix 域套接字等差异，用同一套 **`Read`/`Write`/`Close` + 地址 + Deadline** 组织代码。

- **`net.Dial` / `Listener.Accept`** 返回的类型 **满足 `net.Conn`**（具体动态类型多为 **`*TCPConn`、`*UnixConn`** 等）。  
- **UDP 报文** 常用 **`net.PacketConn`**（另一套抽象）；`UDPConn` 也可实现 `Conn` 的语义子集，但典型用法不同。

---

## 二、接口方法（以本机 `go doc net.Conn` 为准）

```go
type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error

	LocalAddr() Addr
	RemoteAddr() Addr

	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}
```

文档还说明：**多个 goroutine 可同时调用 `Conn` 的方法**——但 **`*TCPConn` 上并发 `Write` 会把字节流交错打乱**，生产上仍常用 **单写协程** 或 **互斥**（见 [07-go-netpoll/10](../07-go-netpoll/10-netpoll常见坑与优化.md)）。

---

## 三、语义要点

1. **`Read`**：可能 **短读**（`n < len(b)` 且 `err == nil`），循环读满或按协议解析。  
2. **`Write`**：可能 **短写**；`n` 表示已写入字节数，`Write` 超时仍可能 **`n > 0`**（见接口注释）。  
3. **Deadline**：**绝对时间**；超时错误可用 **`errors.Is(err, os.ErrDeadlineExceeded)`** 判断（与 `net.Error.Timeout()` 并存，注意文档提示）。  
4. **`Close`**：唤醒阻塞中的 `Read`/`Write`，后续操作返回 **已关闭** 类错误。

---

## 四、与底层的关系

- `net.Conn` **不是**具体结构体；**`*TCPConn`** 内嵌 **`net.conn`**，再落到 **`netFD` → `internal/poll.FD` → runtime netpoll**（见 [03-TCPConn结构](./03-TCPConn结构.md)、[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)）。

---

## 五、极简总结

- **`net.Conn`** = 网络字节流的 **统一操作面**。  
- **超时 / 阻塞体验** 由 **`internal/poll` + netpoll** 实现，不是「魔法阻塞 OS 线程」。  
- **读写到边界**：短读短写 + `EOF` + deadline，决定代码是否健壮。

---

## 导航

- 下一篇：[02-Listener接口](./02-Listener接口.md)  
- 桥梁：[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)
