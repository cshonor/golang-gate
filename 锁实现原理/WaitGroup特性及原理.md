# 并发：`sync.WaitGroup` 特性及原理

> 课表：**并发_waitgroup特性及原理**（同步原语，常与「锁/并发工具」一起讲）

## 1. 它解决什么问题

- 需要等待 **一组 goroutine 全部结束**（或到达某同步点）再继续。
- 比「每起一个 goroutine 就开一个 channel 收信号」更省事，语义是 **计数器 + 等待**。

## 2. API 与语义

| 方法 | 作用 |
|------|------|
| `Add(delta int)` | `delta` 一般为正：增加待等待数量；也可为负（少用，需小心与 Wait 竞态） |
| `Done()` | 等价 `Add(-1)`，表示「本任务结束」 |
| `Wait()` | 阻塞直到计数器为 **0** |

**典型用法**：

```go
var wg sync.WaitGroup
for i := 0; i < n; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // work
    }()
}
wg.Wait()
```

## 3. 实现直觉（原理层）

- 内部维护 **计数 + 等待队列**（概念上类似条件变量）：计数非 0 时 `Wait` 阻塞；`Done` 减到 0 时唤醒等待者。
- 与 **`Mutex`** 的关系：实现会用到锁保护内部状态，但 **对外语义不是互斥临界区**，而是「等全部任务结束」。
- 与 **channel** 对比：channel 偏「传递数据/信号」；WaitGroup 偏 **纯同步、不传业务数据**。

## 4. 易错点（面试常考）

1. **`Add` 必须在 goroutine 启动前完成**（或保证不会与 `Wait` 产生计数竞态）。反例：在 goroutine 里 `Add(1)` 而主协程已经 `Wait()`，可能永远等不到。
2. **`Done` 次数不能多于 `Add` 总增量**，否则 panic（负数计数）。
3. **WaitGroup 不可复制**（结构体内含运行时状态），应传指针或让闭包捕获同一个 `wg`。
4. **复用 WaitGroup**：官方允许在 `Wait` 返回、计数归零后再次 `Add`，但需明确生命周期，避免与上一轮 goroutine 重叠。

## 5. 自检问题

- 为什么「先 `Add` 再 `go`」比「进 goroutine 再 `Add`」更安全？
- 若要用 WaitGroup 等「多个生产者 + 一个消费者」结束，channel 方案怎么写？各适合什么场景？

## 6. 延伸阅读（源码）

- `src/sync/waitgroup.go`
