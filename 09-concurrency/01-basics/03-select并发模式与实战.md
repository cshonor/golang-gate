# 03 - select 并发模式与实战

本篇迁移自旧的 `select` 笔记，并按新目录结构修正标题编号。

---

## 1. select 是什么

`select` 用于同时等待多个 channel 事件，是 Go 并发编排的“控制台”。

## 2. 3 个最常用套路

### 2.1 超时控制

```go
select {
case v := <-ch:
	_ = v
case <-time.After(200 * time.Millisecond):
	return errors.New("timeout")
}
```

生产环境常用 `context.WithTimeout` 替代 `time.After`（更好组合、便于统一取消）。

### 2.2 取消退出

```go
select {
case <-ctx.Done():
	return ctx.Err()
case v := <-ch:
	_ = v
}
```

### 2.3 扇入/限流

- 扇入：合并多个输入 channel
- 限流：worker pool + buffered channel

---

## 延伸阅读

- [../04-context-and-control/01-context上下文与超时控制.md](../04-context-and-control/01-context上下文与超时控制.md)
- [../05-channel-and-practice/02-channel经典并发模式.md](../05-channel-and-practice/02-channel经典并发模式.md)
