# net.Dial 底层流程

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- 解析地址：DNS、IPv4/IPv6 选择、Happy Eyeballs（若启用）
- socket 创建：`socket` syscall、非阻塞设置
- connect 与 poll：可中断等待
- Dialer 字段：Timeout、KeepAlive、Control（raw hook）
- 错误链：`OpError`、`Addr`、`Err`  unwrap

## 与 netpoll 的桥梁（扩写）

`Dial` 成功后会初始化 **`netFD` / poll**，非阻塞 `connect` 与 **可写事件** 由 **`internal/poll` + netpoll** 收尾。精读 **[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)**，超时与错误见 **[08](./08-网络超时与Deadline底层实现.md)**、**[09](./09-网络错误分类与处理.md)**。

## 正文

（待补充）
