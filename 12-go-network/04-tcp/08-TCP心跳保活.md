# TCP 心跳保活

> **04-tcp · TCP 协议与编程**  
> 上篇：[07-Go-TCP客户端](./07-Go-TCP客户端.md)；与 **应用层心跳** 勿混。

---

## 一、两类「心跳」

| 类型 | 谁发起 | 典型目的 |
|------|--------|----------|
| **TCP Keepalive** | 内核按套接字选项定时发 **探测段** | 发现 **死连接**、NAT 会话保活（弱） |
| **应用层 ping/pong** | 业务协议 | 语义级存活、版本协商、权限校验 |

**TCP keepalive** 不能替代 **应用心跳**：对端进程死锁但 OS 仍回 ACK 时，**TCP 层仍可能「看起来活着」**。

---

## 二、Linux 行为（概念）

受 **`tcp_keepalive_time` / `intvl` / `probes`** 等 **全局 sysctl** 影响；`setsockopt` 的 **`TCP_KEEPIDLE`/`TCP_KEEPINTVL`/`TCP_KEEPCNT`**（存在性与语义查 `man 7 tcp`）。

---

## 三、Go 常用 API

- **`TCPConn.SetKeepAlive(true)`**  
- **`TCPConn.SetKeepAlivePeriod(d)`**（Go 1.2+，具体行为见文档）  
- **`net.Dialer{KeepAlive: ...}`**：在 `Dial` 时一并设置。

---

## 四、与 `ReadDeadline` 组合

- **长连接空闲**：可设 **较短 `ReadDeadline`**，超时则发 **应用 ping** 或 **回收连接**。  
- **仅依赖 keepalive**：默认间隔往往 **很长**，不适合「秒级发现故障」。

---

## 五、极简总结

- **TCP keepalive** = 内核级 **死连接探测**。  
- **应用心跳** = 业务级 **端到端存活**。  
- 高可用服务 **两者常并存**，参数不同。

---

## 导航

- 上一篇：[07-Go-TCP客户端](./07-Go-TCP客户端.md)  
- 下一篇：[09-TCP长连接与短连接](./09-TCP长连接与短连接.md)
