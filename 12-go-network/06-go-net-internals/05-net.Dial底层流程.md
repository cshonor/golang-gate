# net.Dial 底层流程

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- 解析地址：DNS、IPv4/IPv6 选择、Happy Eyeballs（若启用）
- socket 创建：`socket` syscall、非阻塞设置
- connect 与 poll：可中断等待
- Dialer 字段：Timeout、KeepAlive、Control（raw hook）
- 错误链：`OpError`、`Addr`、`Err`  unwrap

## 正文

（待补充）
