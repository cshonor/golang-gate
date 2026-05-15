# `Listener` 接口（服务端监听抽象）

> **06-go-net-internals · Go net 包源码级理解**  
> 上篇：[01-net.Conn接口](./01-net.Conn接口.md)

---

## 一、核心定义

`net.Listener` 描述 **被动等待连接** 的一侧：**`Accept`** 返回 **`Conn`**，**`Addr`** 报告监听地址，**`Close`** 释放监听 fd。

```go
type Listener interface {
	Accept() (Conn, error)
	Close() error
	Addr() Addr
}
```

---

## 二、服务端主路径

1. **`net.Listen("tcp", ":8080")`** → 得到 **`Listener`**（TCP 下常为 **`*TCPListener`**）。  
2. 循环 **`Accept()`**：阻塞直到 **全连接队列** 有可交付连接（内核已完成握手，见 [03-linux](../03-linux-networking/07-半连接队列与全连接队列.md)）。  
3. 对每个 **`Conn`** 启动业务（goroutine / worker pool）。  
4. 关停时 **`Close()`**：`Accept` 返回 **`ErrClosed`/`use of closed network listener`** 一类错误，跳出循环。

---

## 三、与系统调用的对应

- **`Listen`**：`socket` → `bind` → `listen`。  
- **`Accept`**：`accept` 返回 **新连接 fd**，包装为 **`netFD`/`TCPConn`** 并纳入 **poll + netpoll**（见 [06-net.Listen](./06-net.Listen底层流程.md)、[07-pollDesc](./07-pollDesc核心结构与原理.md)）。

---

## 四、与 `http.Server` 的关系

- **`http.Server`** 内部持有 **`net.Listener`**；**`Shutdown`** 会关闭 listener 并处理在途请求——**`Close` listener  alone** 不等于完整优雅退出。

---

## 五、极简总结

- **`Listener`** = **守门人**：只负责 **接客（Accept）** 与 **关门（Close）**。  
- **`Accept` 返回的 `Conn`** 才是读写主战场。

---

## 导航

- 上一篇：[01-net.Conn接口](./01-net.Conn接口.md)  
- 下一篇：[03-TCPConn结构](./03-TCPConn结构.md)
