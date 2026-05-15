# `UDPConn` 结构（UDP 端点实体）

> **06-go-net-internals · Go net 包源码级理解**  
> 上篇：[03-TCPConn结构](./03-TCPConn结构.md)；UDP 协议见 [05-udp-multicast](../05-udp-multicast/01-UDP特点.md)。

---

## 一、核心定义

**`*net.UDPConn`** 封装 **UDP socket**：可 **无连接**（`WriteTo`/`ReadFrom`）也可 **`DialUDP` 后 `Write`**（「已连接 UDP」语义，内核可缓存对端地址）。

- **无 TCP 状态机**：没有 **`ESTABLISHED`** 等内核 TCP 状态。  
- **面向报文**：一次 `ReadFrom` 通常对应一帧 UDP payload（受 MTU/内核分片策略影响，应用仍要做上限校验）。

---

## 二、与 `TCPConn` 的对比

| 维度 | TCPConn | UDPConn |
|------|---------|---------|
| 协议语义 | 可靠有序字节流 | 尽力而为报文 |
| 粘包 | 有（需应用分帧） | 一般无（仍要防超大报文） |
| 典型 API | `Read`/`Write` | `ReadFrom`/`WriteTo` 更常见 |
| 底层 | 同样 **`netFD` + poll** | 同样 **`netFD` + poll** |

---

## 三、与 netpoll 的桥梁

**`UDPConn.ReadFrom` 阻塞** 与 TCP **`Read`** 类似：无报文时 **`EAGAIN`** → **`internal/poll` 等待** → **G 挂起** → 报文到达 **唤醒**。

详见：[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)。

---

## 四、极简总结

- **`UDPConn`** 也是 **`netFD`** 一族；**轻**在协议，**不轻**在「等数据仍要走 poll」。  
- **可靠性、有序性** 必须在应用层补。

---

## 导航

- 上一篇：[03-TCPConn结构](./03-TCPConn结构.md)  
- 下一篇：[05-net.Dial底层流程](./05-net.Dial底层流程.md)
