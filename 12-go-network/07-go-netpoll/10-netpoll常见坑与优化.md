# netpoll 常见坑与优化

> **07-go-netpoll · Go netpoll 高并发核心**  
> 与 **06** 的 [08-网络超时与Deadline](../06-go-net-internals/08-网络超时与Deadline底层实现.md)、[10-连接关闭与资源泄漏排查](../06-go-net-internals/10-连接关闭与资源泄漏排查.md) 对照读，工程问题覆盖更全。

---

## 1. 常见坑

| 坑 | 后果 | 规避 |
|----|------|------|
| **从不设 deadline** | 慢客户端拖死大量 **G** 在 `Read` | 每连接、每请求设 **读超时**；配合 **Context** |
| **Close 与 Read 并发无协调** | `use of closed connection`、panic 在业务层 | 单写者关流；或 **sync** 约定生命周期 |
| **误把「G 阻塞」当「线程安全」** | 多 goroutine 同时 `Write` 同一 `conn` | **`net.Conn` 并发写不安全**；要加锁或单写协程 |
| **epoll fd / 进程 fd 泄漏** | 进程 `ulimit` 打满 | 确保 **`Close`**、defer、连接池上限 |
| **忙轮询自写** | CPU 飙高 | 不要绕过 `net` 自己 **非阻塞 + 死循环** |

---

## 2. 常见优化（ syscall / 协议 / 系统参数）

- **`TCP_NODELAY`**：降低小报文延迟（权衡与吞吐）。  
- **`SO_REUSEADDR` / `SO_REUSEPORT`**：监听与多进程/滚动发布策略（平台语义不同，查文档）。  
- **批量 accept / 批量读**：在业务层减少系统调用次数（与框架相关）。  
- **Go 版本升级**：调度、timer、netpoll 实现持续演进，**跟 release note**。

---

## 3. 与 GMP 新特性（点到为止）

- **异步抢占**等调度变化会影响「长循环是否会让出」等行为；**网络路径**仍以 **poll + gopark** 为主轴理解即可。

---

## 4. 回到索引

- [02-netpoll是什么](./02-netpoll是什么.md)  
- [08-netpoll与GMP调度深度联动](./08-netpoll与GMP调度深度联动.md)  
- 调优专题：`12-tuning-and-issues`
