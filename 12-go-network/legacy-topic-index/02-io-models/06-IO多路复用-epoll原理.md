# IO 多路复用：`epoll` 原理

> **02-io-models · IO 模型全解**  
> 上篇：[05-IO多路复用-poll](./05-IO多路复用-poll.md) · 关联：[07-pollDesc 桥梁](../06-go-net-internals/07-pollDesc核心结构与原理.md)

---

## 一、核心定义（Linux）

`epoll` 是 Linux 上常用的高性能多路复用机制，围绕 **「一个 epoll 实例 + 多 fd 的注册表 + 就绪通知」** 工作。相对 `select`/`poll`，在 **大量连接、少量活跃** 的场景下，**减少无谓扫描与重复传递 fd 集合** 的开销。

### 三个系统调用

1. **`epoll_create1` / `epoll_create`**：创建 epoll 实例，返回 `epollfd`。  
2. **`epoll_ctl`**：增删改对某 fd 的监听（读/写/ET/LT/oneshot 等）。  
3. **`epoll_wait`**：阻塞等待，返回一批**已就绪**的 `(fd, events)`。

### 典型流程

1. `epoll_create1` 得 `epfd`。  
2. 对每个连接 `sockfd`，`epoll_ctl(EPOLL_CTL_ADD, ...)` 挂到 `epfd`。  
3. 循环 `epoll_wait`，拿到就绪列表。  
4. 对每个就绪 fd：`read`/`write`（fd 一般为非阻塞）。  
5. 处理完继续 `epoll_wait`。

---

## 二、相对 `select`/`poll` 的常见优势（概念层）

1. **fd 注册与兴趣列表**：`epoll_ctl` 维护内核侧结构，**不必每次把「全量 fd 集合」整包拷进内核**（与旧接口对比的直觉）。  
2. **`epoll_wait` 直接给出就绪子集**：应用遍历的是 **m 个就绪**，而不是每次都扫 **n 个全量**（活跃少时差别大）。  
3. **规模**：受机器内存、ulimit 等约束，**远高于「1024 个 fd_set」那种教学例子**。  
4. **边缘触发 ET**（见下一篇）：可进一步减少重复通知次数，但编程要求更严。

> 具体内核数据结构（红黑树、就绪链表等）随内核版本演进，面试答 **「兴趣表 + 就绪队列/回调路径」** 即可，不必背死某一版源码细节。

---

## 三、与 Go 的关系

Linux 上 **`runtime` 的 `netpoll`** 常以 **`epoll`** 为后端之一：**非阻塞网络 fd** 就绪时唤醒阻塞在 `Read`/`Write` 上的 **G**。更细的 **`pollDesc`、deadline、错误语义** 见 **`06-go-net-internals`**；调度与 **`netpoll` 循环** 见 **`07-go-netpoll`**。

---

## 四、典型场景

- Nginx、Redis、大量自研 C/C++/Go 网关。  
- **Reactor** 模型的 Linux 默认优等生。

---

## 五、极简总结

- `epoll`：**ctl 注册 + wait 取就绪**，适合 **高并发、低活跃率**。  
- Go：**netpoll + 非阻塞 fd**，对上仍是同步式 API 体验。  
- 下一篇：**LT / ET** 触发语义。

---

## 导航

- 上一篇：[05-IO多路复用-poll](./05-IO多路复用-poll.md)  
- 下一篇：[07-epoll-LT与ET模式](./07-epoll-LT与ET模式.md)
