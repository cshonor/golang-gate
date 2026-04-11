# `pollDesc` / `internal/poll`：`net` 与 netpoll 的桥梁

> **06-go-net-internals · Go net 包源码级理解**  
> 读完本篇再读 [07-go-netpoll/02-netpoll是什么](../07-go-netpoll/02-netpoll是什么.md)，可把 **「`conn.Read` 一行代码」→「为何不占满 OS 线程」** 串成闭环。  
> **命名提醒**：口语里的 **pollDesc** 常泛指 **`internal/poll.FD` 与 `runtime` 侧 `runtime_pollDesc` 等协作结构**；**字段名以 `$GOROOT/src/internal/poll`、`src/runtime/netpoll*.go` 为准**。

---

## 一、为什么必须有这一层？

- **`net` 包**描述 **`TCPConn`/`Listener`/地址** 等业务对象。  
- **Linux 上高并发**要求 socket **非阻塞** + **多路复用**；**「等就绪」**不能等价于 **「让 OS 线程睡在 `recv` 上」**。  
- 因此 **`netFD` 内嵌 `internal/poll.FD`**：负责 **`EAGAIN` 时的等待、向 runtime 注册/注销、与 deadline 协作**。

> **记忆**：**`net`** 管对象；**`internal/poll`** 管 **「等不等、怎么等」**；**`runtime netpoll`** 管 **「多路之一就绪后叫醒哪个 G」**。

---

## 二、分层心智（必背）

```text
*TCPConn
  └─ net.conn
        └─ netFD
              └─ internal/poll.FD
                    ├─ Sysfd / 句柄
                    ├─ 读写锁、closing 等状态
                    └─ runtime 注册：runtime_pollOpen / Wait / Close...
```

一次 **`FD.Read`** 典型路径：

1. **循环 `syscall.Read`（或等价）**。  
2. 若返回 **`EAGAIN`**（非阻塞且无数据）→ 进入 **`pollWait`**。  
3. **`runtime_pollWait`**：把当前 **G** 与 **该 fd 的读/写事件** 关联并 **`gopark`**。  
4. **netpoll** 在 **`epoll_wait`/…** 返回后 **标记就绪**，把 **G** 置为 **`Grunnable`**。  
5. G 再次被调度，重试 `Read`，读到数据或错误返回。

---

## 三、「pollDesc」到底指哪张表？

- **`internal/poll.FD`** 里持有 **`pollDesc`**（或等价指针），其 **`runtime_pollDesc`** 由 **`runtime`** 分配，用于 **把用户态 fd 与 netpoll 事件表项绑定**。  
- **不要**把网上简化版伪结构当成稳定 ABI；升级 Go 版本后字段常变。

---

## 四、生命周期

```text
Dial / Listen / Accept 成功
    → 创建 netFD，初始化 poll.FD（runtime_pollOpen）
Read / Write
    → 快路径：syscall 直接成功
    → 慢路径：EAGAIN → runtime_pollWait → gopark
事件就绪 / deadline 到期
    → netpoll / timer 路径唤醒 G
Close
    → runtime_pollClose → close(fd) → 等待方带着错误返回
```

泄漏与关闭顺序见 [10](./10-连接关闭与资源泄漏排查.md)。

---

## 五、源码阅读顺序（建议）

1. `src/net/net.go`、`src/net/tcpsock*.go`：`conn` / `netFD`。  
2. `src/internal/poll/fd_unix.go`（或平台文件）：**`FD.Read`/`Write`**、`readLock`。  
3. `src/internal/poll/fd_poll_runtime.go`：**`pollDesc` / `runtime_pollWait`**。  
4. `src/runtime/netpoll_*.go`：**`netpoll`/`epoll`**。  
5. 回到 **07**：[05-Go的Read与Write为什么看起来阻塞](../07-go-netpoll/05-Go的Read与Write为什么看起来阻塞.md)、[09-pollDesc等待队列与唤醒原理](../07-go-netpoll/09-pollDesc等待队列与唤醒原理.md)。

---

## 六、与相邻篇目

| 篇目 | 关系 |
|------|------|
| [03-TCPConn结构](./03-TCPConn结构.md) | 谁持有 `netFD` |
| [05/06 Dial、Listen](./05-net.Dial底层流程.md) | 何时 `runtime_pollOpen` |
| [08 Deadline](./08-网络超时与Deadline底层实现.md) | 如何把等待打断成 `timeout` |

---

## 导航

- 上一篇：[06-net.Listen底层流程](./06-net.Listen底层流程.md)  
- 下一篇：[08-网络超时与Deadline底层实现](./08-网络超时与Deadline底层实现.md)
