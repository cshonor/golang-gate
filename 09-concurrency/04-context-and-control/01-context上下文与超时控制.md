# 04 - context 上下文与超时控制

> 已迁移自旧目录的 context 笔记；这里补齐与并发的结合点。

---

## 1. 为什么需要 context

- **取消**：请求结束后停止后台任务。
- **超时**：避免无限等待。
- **链路传递**：在多层调用中传递截止时间与取消信号。

---

## 2. 并发里最常用的组合

- `ctx, cancel := context.WithTimeout(parent, d)` + `defer cancel()`
- goroutine 内：

```go
select {
case <-ctx.Done():
	return ctx.Err()
case v := <-ch:
	_ = v
}
```

---

## 3. 注意点

- `cancel()` 一定要调用（释放计时器等资源）。
- 不要把 context 存进结构体长期保存（一般只在请求链路里传）。

---

## 延伸阅读

- [../01-basics/03-select并发模式与实战.md](../01-basics/03-select并发模式与实战.md)
