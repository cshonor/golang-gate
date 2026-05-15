# netpoll 是什么

> **07-go-netpoll · Go netpoll 高并发核心**  
> 桥梁篇：[06-go-net-internals/07-pollDesc](../06-go-net-internals/07-pollDesc核心结构与原理.md)。

---

## 一、定义

**netpoll** 是 **`runtime` 内部** 对 **「一批网络 fd 的就绪事件」** 的统一轮询与分发机制：把 **内核多路复用**（Linux **`epoll`**、BSD/macOS **`kqueue`**、Windows **IOCP/AFD** 等）封装成 **`netpollinit` / `netpollopen` / `netpoll` / `netpollBreak`** 一类入口，供 **`internal/poll`** 在 **`EAGAIN`** 时 **挂起 G**，在 **就绪/超时/关闭** 时 **唤醒 G**。

---

## 二、它不是什么？

- **不是**用户态可调用的 `package netpoll`。  
- **不是**替代 `net` 的第三方库；**业务仍只用 `net`**。  
- **不负责**替你解决 **协议分帧、业务幂等**。

---

## 三、最小心智模型

```text
internal/poll: runtime_pollWait(mode)
        ↓
runtime:     gopark（G 睡眠）
        ↓
kernel:      epoll_wait / ...
        ↓
runtime:     netpoll → ready(G)
```

---

## 四、为什么要绑定 `internal/poll`？

- **跨平台**：`net` 不想复制 `#ifdef` 地狱，**fd 等待语义**集中在 **`internal/poll`**。  
- **与 timer 协同**：**deadline** 与 **可读** 都是 **「唤醒条件」**。

---

## 五、源码入口（随版本检索）

- `src/runtime/netpoll_*.go`  
- `src/internal/poll/fd_poll_runtime.go`

---

## 六、极简总结

- **netpoll** = **runtime 的网络事件循环 + G 唤醒器**。  
- **与 epoll**：Linux 上 **多数时间在 `epoll_wait`**，但 **语义归属 runtime**。

---

## 导航

- 上一篇：[01-Go为什么高并发](./01-Go为什么高并发.md)  
- 下一篇：[03-epoll-kqueue-IOCP支持](./03-epoll-kqueue-IOCP支持.md)
