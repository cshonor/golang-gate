# 网络超时与 Deadline 底层实现

> **06-go-net-internals · Go net 包源码级理解**  
> 前置：[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)。**Deadline 的本质**：为一次或一类 IO 等待设置**上限**，到期必须让阻塞在 `pollWait` 上的 **G** 醒来并带着 **`timeout` 类错误**返回用户代码。

---

## 1. 用户能调什么 API

- **`SetDeadline(t)`**：读+写共用一个绝对时刻。  
- **`SetReadDeadline` / `SetWriteDeadline`**：分别约束读、写。

语义要点：

- **零值 `time.Time{}`** 在 `net` 里常表示 **清除** 对应 deadline（以当前实现文档为准）。  
- **已过期的时刻** 会导致后续 IO **很快失败**（不必等到「下一次真的去等」才发现）。

---

## 2. 与 `internal/poll` 的关系（直觉）

1. 调用进入 `netFD` → `poll.FD` 的路径，把 **绝对时间** 记录到 **poll 状态**（常与 **timer** 或 **runtime 定时机制** 协作）。  
2. 当 `Read`/`Write` 因 **`EAGAIN`** 进入 **`runtime_pollWait`** 时，等待不仅是「等 fd 就绪」，还包括 **「等 deadline 到期」** 这条分支。  
3. **到期**：等待被解除，`Read`/`Write` 返回 **`net.Error` 且 `Timeout() == true`**（常见包装为 `i/o timeout`）。

实现细节随 Go 版本演进，阅读时以 **`SetReadDeadline` 调用链** 向下跟到 `internal/poll` 为准。

---

## 3. 常见工程用法

| 场景 | 建议 |
|------|------|
| 每个请求有 SLA | 对 **`conn`** 或 **`Response.Body`** 设 **读 deadline** |
| 防止写阻塞拖死 | 设 **写 deadline**；大 body 分块写 |
| 长连接心跳 | **短 read deadline** + 业务层超时重置，或与 `TCPKeepAlive` 组合 |

与 **Context** 取消配合时，注意：**取消 context 不会自动关 fd**；通常要自己 `Close` 或上层封装一并处理。

---

## 4. 与 07 章衔接

- netpoll 在 **定时器到期** 时同样要把 **G** 唤醒，路径与 **「对端发来数据」** 类事件共享「出等待 → 可运行」框架，见 [08-netpoll与GMP调度深度联动](../07-go-netpoll/08-netpoll与GMP调度深度联动.md)。

---

## 下一篇

[09-网络错误分类与处理.md](./09-网络错误分类与处理.md)
