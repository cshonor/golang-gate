# net.Conn 接口

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- 接口方法：Read/Write/Close、LocalAddr/RemoteAddr、SetDeadline 族
- 实现者：TCPConn、UDPConn、TLSConn、UnixConn 等
- 接口语义：多返回值 err、短读/短写与 `io.Copy` 行为
- 与 `io.Reader`/`Writer`：如何在中间件链中组合
- 调试：包装 Conn 记录字节流与耗时

## 正文

（待补充）
