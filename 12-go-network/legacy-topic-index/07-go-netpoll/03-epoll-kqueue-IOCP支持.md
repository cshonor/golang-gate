# epoll / kqueue / IOCP：netpoll 的平台后端

> **07-go-netpoll · Go netpoll 高并发核心**  
> 上篇：[02-netpoll是什么](./02-netpoll是什么.md)；理论见 [02-io-models](../02-io-models/06-IO多路复用-epoll原理.md)。

---

## 一、各平台用什么？

| OS | 常见后端 | 备注 |
|----|----------|------|
| **Linux** | `epoll` | 生产绝对主流 |
| **macOS / FreeBSD …** | `kqueue` | 心智与 epoll 类似：事件表 |
| **Windows** | **IOCP / AFD** 等演进路径 | 语义以 Windows runtime 实现为准 |

Go 通过 **`//go:build`** 拆分 **`netpoll_*.go`**；本机直接打开 **`$GOROOT/src/runtime`** 搜索 **`netpoll`** 即可定位。

---

## 二、用户需要直接调 epoll 吗？

**几乎从不**。`net` 已把 **非阻塞 + 注册 + 等待** 做完；自写 epoll 常见于 **极端定制** 或 **非 `net` fd** 的特殊集成。

---

## 三、LT / ET（Linux）

- **Go netpoll 路径以 LT 心智最安全**（事件未处理完仍可能再次通知）。  
- **ET** 必须 **读至 `EAGAIN`**，否则易「丢通知感」——与业务自研 epoll 强相关，**`net` 默认路径不必强行 ET**。

见：[02-io-models/07](../02-io-models/07-epoll-LT与ET模式.md)。

---

## 四、极简总结

- **netpoll** = **抽象**；**epoll/kqueue/IOCP** = **实现**。  
- **读 runtime 构建标签** 比背表格更重要。

---

## 导航

- 上一篇：[02-netpoll是什么](./02-netpoll是什么.md)  
- 下一篇：[04-Go的IO模型到底是什么](./04-Go的IO模型到底是什么.md)
