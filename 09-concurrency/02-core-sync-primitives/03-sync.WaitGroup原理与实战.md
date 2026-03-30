# 03 - sync.WaitGroup 原理与实战

---

## 模板

```go
wg.Add(1)
go func() {
	defer wg.Done()
	work()
}()
wg.Wait()
```

---

## 注意

- `Add` 先于启动 goroutine。
