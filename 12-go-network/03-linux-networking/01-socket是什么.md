# socket 是什么

> **03-linux-networking · Linux 网络底层**

## 内容大纲

- BSD socket 抽象：端点 = IP + 端口 + 协议
- fd 统一抽象：socket 也是文件描述符（一切皆文件）
- TCP/UDP/UNIX domain 套接字差异与典型 syscall 序列
- 内核对象：`struct sock` / `struct socket`（概念级，不必背字段）
- 与 Go：`net.Dial`/`Listen` 在 Linux 上对应的 syscall 链条

## 正文

（待补充）
