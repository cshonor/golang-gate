# Go UDP 服务端

> **05-udp-multicast · UDP 与组播**

## 内容大纲

- ListenPacket 与 UDPConn；ReadFrom 拿远端地址
- 每包一个 goroutine 与单循环：背压与公平性
- 缓冲区大小：ReadFrom 切片容量与丢包
- 错误处理：截断、ICMP 不可达（平台差异）
- 与 SetReadBuffer / SetWriteBuffer 调优

## 正文

（待补充）
