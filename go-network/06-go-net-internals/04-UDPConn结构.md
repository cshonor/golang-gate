# UDPConn 结构

> **06-go-net-internals · Go net 包源码级理解**

## 内容大纲

- ReadFrom/WriteTo 与 `conn` 共性
- 已连接 UDP vs 无连接：`DialUDP` 后 `Write` 行为
- 缓冲区与丢包：应用层队列
- 与 `syscall`：recvfrom/sendto 路径
- Multicast 相关方法在结构上的挂点

## 正文

（待补充）
