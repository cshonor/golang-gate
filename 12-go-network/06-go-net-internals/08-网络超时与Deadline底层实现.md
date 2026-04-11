# 网络超时与 Deadline 底层实现

> **06-go-net-internals · Go net 包源码级理解**  
> 前置：[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)。**Deadline** = **绝对时间点**：到达后，**正在等 I/O 的 G** 必须被叫醒，**`Read`/`Write` 返回错误**（常见 **`i/o timeout`** / **`os.ErrDeadlineExceeded`**）。

---

## 一、用户 API

- **`SetDeadline(t)`**：同时约束后续 **读+写**（并影响已在等待的调用）。  
- **`SetReadDeadline` / `SetWriteDeadline`**：分别约束。  
- **零值 `time.Time{}`**：在 `net` 中通常表示 **清除** 对应 deadline（以当前版本文档为准）。

`go doc net.Conn`：**deadline 过期后**，可用 **`errors.Is(err, os.ErrDeadlineExceeded)`**；同时 **`net.Error.Timeout()` 也可能为 true**（文档提醒：**并非所有 Timeout 都来自 deadline**）。

---

## 二、实现直觉（不要等同于「内核 SO_RCVTIMEO 包办一切」）

1. **`SetReadDeadline` 进入 `netFD` → `poll.FD`**，把 **绝对时间** 存进 **poll 状态**（具体字段名随版本变）。  
2. 当 `Read` 因 **`EAGAIN`** 走 **`runtime_pollWait`** 时，等待条件不仅是 **「fd 可读」**，还包括 **「deadline 到期」**（由 **runtime 定时器子系统** 与 netpoll 协同）。  
3. **到期**：等待结束，`Read` 返回 **`timeout` 语义错误**；若已用 **`context` 取消**，那是另一条链，需 **`errors.Is`** 区分。

> **纠偏**：Go **确实会用内核/运行时定时能力**；更准确的说法是——**应用层通常不直接靠 `SO_RCVTIMEO` 一条路径完成所有语义**，而是 **`poll` + timer + 返回错误包装** 组合实现 **`SetXxxDeadline`** 的可移植行为。

---

## 三、伪代码心智（非真实源码）

```text
SetReadDeadline(t):
  pd.storeDeadline(READ, t)
  pd.updatePollerOrTimer()   // 注册/更新下一次唤醒时间

Read():
  for {
    n, err = syscall.Read(fd, buf)
    if err != EAGAIN { return }
    runtime_pollWait(pd, 'r') // 内部合并：可读 OR 到期 OR closed
    if deadlinePassed { return 0, timeoutErr }
  }
```

---

## 四、工程用法（速查）

| 场景 | 建议 |
|------|------|
| HTTP 服务端读 body | **`ReadHeaderTimeout`/`ReadTimeout`** 或 conn 级 **读 deadline** |
| 慢客户端防护 | **每连接周期性 `SetReadDeadline(now.Add(...))`** |
| 与 `Context` | **`ctx` 取消不会自动关 conn**；要 **`Close` 或上层封装** |

---

## 五、与 07 章衔接

- **定时器到期** 与 **对端数据到达** 都要走 **「唤醒 G → 重试/返回」** 框架，见 [08-netpoll与GMP调度深度联动](../07-go-netpoll/08-netpoll与GMP调度深度联动.md)。

---

## 导航

- 上一篇：[07-pollDesc核心结构与原理](./07-pollDesc核心结构与原理.md)  
- 下一篇：[09-网络错误分类与处理](./09-网络错误分类与处理.md)
