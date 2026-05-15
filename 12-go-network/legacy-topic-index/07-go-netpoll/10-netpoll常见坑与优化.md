# netpoll 常见坑与优化

> **07-go-netpoll · Go netpoll 高并发核心**  
> 对照：[06/08 Deadline](../06-go-net-internals/08-网络超时与Deadline底层实现.md)、[06/10 关闭泄漏](../06-go-net-internals/10-连接关闭与资源泄漏排查.md)。

---

## 一、常见坑

| 坑 | 后果 | 规避 |
|----|------|------|
| **从不设 deadline** | 大量 **G 卡在 `Read`**，像泄漏 | **每连接/每请求读超时**；`Context` + 主动 `Close` |
| **`Close` 与 `Read` 无协调** | `use of closed network connection` 风暴 | **单写者关流**；`errgroup`/`WaitGroup` |
| **多 goroutine 同时 `Write` 同一 `TCPConn`** | 字节流交错 | **单写协程** 或 **互斥** |
| **fd 泄漏** | `too many open files` | **`defer conn.Close()`** 不够，还要 **错误路径关闭** |
| **自写非阻塞 + 死循环 `Read`** | CPU 飙高 | **不要绕过 `net`** |

---

## 二、常见优化（ syscall / 内核 / 协议）

- **`TCP_NODELAY`**：小包延迟 vs 吞吐（Nagle）。  
- **`TCP_KEEPALIVE` + Go `SetKeepAlive*`**：见 [04-tcp/08](../04-tcp/08-TCP心跳保活.md)。  
- **`SO_REUSEPORT`**：多进程/多 `Listen` 水平扩展（**平台语义先查文档**）。  
- **批量业务读**：减少 **小 `Read` 次数**（`bufio`、协议层一次读够头）。

---

## 三、Go 版本与 runtime

- **1.14+ 异步抢占**：改变 **纯计算循环** 行为；**网络**仍以 **poll** 为主。  
- **timer 优化** 多版本持续迭代：**线上问题先看 release note**。

---

## 四、与 06 的闭环

- **deadline 设置错误** → **G 醒不来/醒太早** → 回到 **[06/08](../06-go-net-internals/08-网络超时与Deadline底层实现.md)** 复查。  
- **连接泄漏** → **[06/10](../06-go-net-internals/10-连接关闭与资源泄漏排查.md)** + **`pprof`**。

---

## 五、极简总结

- **netpoll 不是魔法**：**背压、超时、关闭顺序** 写错一样炸。  
- **优化先测量**：`trace`/`pprof`/`ss` 三件套。

---

## 导航

- 上一篇：[09-pollDesc等待队列与唤醒原理](./09-pollDesc等待队列与唤醒原理.md)  
- 架构延伸：[10-server-architecture](../10-server-architecture/01-acceptor-worker模型.md)
