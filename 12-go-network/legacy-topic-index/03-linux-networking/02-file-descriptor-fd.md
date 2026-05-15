# file descriptor（fd）

> **03-linux-networking · Linux 网络底层**  
> 上篇：[01-socket是什么](./01-socket是什么.md)

---

## 一、核心定义

**文件描述符（File Descriptor, fd）** 是内核为 **「进程打开的资源」** 分配的 **小整数句柄**。进程通过 fd 告诉内核「我要操作哪一个打开对象」。

- **本质**：进程 **打开文件表（file descriptor table）** 的下标；内核用其找到 **`struct file`** 等内核对象。  
- **惯例**：`0` = stdin，`1` = stdout，`2` = stderr（可被重定向/关闭后再复用编号）。  
- **上限**：受 **`RLIMIT_NOFILE`**（`ulimit -n`）等限制；**默认值因发行版/容器而异**，高并发服务常需调到 **数万～数十万** 量级并配合内核参数。

---

## 二、核心特性

### 1. 进程隔离

每个进程有 **自己的 fd 表**。进程 A 的 fd `3` 与进程 B 的 fd `3` **毫无关系**。

### 2. 资源统一抽象

在 Unix/Linux 上，fd 可代表：

- 普通文件、目录  
- **socket**、pipe、FIFO  
- 设备节点、`eventfd`、`timerfd`、`epollfd` …

因此上层可用 **`read`/`write`/`poll`/`epoll_ctl`** 等统一接口组合复杂程序（具体能力视类型而定）。

### 3. 生命周期

1. **创建**：`open` / `socket` / `accept` / `pipe` … 返回新 fd。  
2. **使用**：`read` / `write` / `setsockopt` …  
3. **关闭**：`close(fd)` 递减引用；**未关闭** 则内核对象与端口/缓冲可能长期占用 → **`too many open files`**、连接堆积。

---

## 三、fd 与 Socket

- **每个 TCP 连接**（`accept` 返回）通常对应 **一个独立 fd**。  
- **Go `net.Conn`** 底层持有 **fd**（经 `netFD` / `internal/poll` 封装），`Read`/`Write` 即对该 fd 的系统调用路径。  
- **泄漏排查**：`lsof -p <pid>`、`/proc/<pid>/fd`、Go **`runtime/pprof`** 的 goroutine/block 与业务 **`Close`** 审计。

---

## 四、与 IO 模型

- **阻塞 IO**：`read` 在无数据时 **阻塞线程**（若未 `O_NONBLOCK`）。  
- **非阻塞**：`EAGAIN` / `EWOULDBLOCK`，配合 **`select`/`poll`/`epoll`**（见 [02-io-models](../02-io-models/06-IO多路复用-epoll原理.md)）。  
- **`FD_CLOEXEC`**：`exec` 后子进程不继承该 fd，避免 **敏感 fd 泄漏**（Go runtime 常见相关处理，具体以版本为准）。

---

## 五、极简总结

- **fd** = 进程访问打开对象的 **整数句柄**；socket 也是一种打开对象。  
- **高并发** 必查 **`ulimit -n`** 与 **泄漏**。  
- **网络性能** 与 **fd 数量、缓冲、状态** 强相关，后续篇目展开。

---

## 导航

- 上一篇：[01-socket是什么](./01-socket是什么.md)  
- 下一篇：[03-socket创建流程](./03-socket创建流程.md)
