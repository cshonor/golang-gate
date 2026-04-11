# TCPConn 结构

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- 内嵌 `conn`：fd、netFD、pollDesc（概念）
- syscall 封装：`read`/`write` 与 `errno` 映射
- SetLinger、SetNoDelay、SetKeepAlive 等到 syscall 的映射
- 与 `netFD` 生命周期：dup、SetBlocking
- 阅读指引：`src/net/tcpsock.go`、`src/net/fd_unix.go`（按本机 Go 版本）

## 与 netpoll 的桥梁（扩写）

`TCPConn` 最终落在 **`net.netFD`**，读写阻塞/超时/关闭都经 **`internal/poll.FD`** 与 **`runtime` netpoll** 协作。系统拆解见 **[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)**，工程侧见 **[08](./08-网络超时与Deadline底层实现.md)**、**[09](./09-网络错误分类与处理.md)**、**[10](./10-连接关闭与资源泄漏排查.md)**。

## 正文

（待补充）
