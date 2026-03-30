# 03 - errgroup 并发错误处理

`golang.org/x/sync/errgroup` = `WaitGroup` + 错误传播 +（可选）Context 取消。

---

## 1. 最常用写法

```go
var g errgroup.Group
for _, t := range tasks {
	t := t
	g.Go(func() error { return do(t) })
}
return g.Wait()
```

---

## 2. WithContext

`g, ctx := errgroup.WithContext(parent)`：某任务失败可触发取消，适合“多任务并发下载/查询”。

---

## 延伸阅读

- `03-error_handling/`（错误包装与链）
