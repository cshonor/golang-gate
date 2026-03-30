# TCPConn 结构

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- 内嵌 `conn`：fd、netFD、pollDesc（概念）
- syscall 封装：`read`/`write` 与 `errno` 映射
- SetLinger、SetNoDelay、SetKeepAlive 等到 syscall 的映射
- 与 `netFD` 生命周期：dup、SetBlocking
- 阅读指引：`src/net/tcpsock.go`、`src/net/fd_unix.go`（按本机 Go 版本）

## 正文

（待补充）
