# 02 - goroutine 基础与常见坑

## 1. 循环变量捕获

```go
for _, v := range xs {
	v := v
	go func() { fmt.Println(v) }()
}
```

## 2. goroutine 泄漏

- 用 `context` + `select` 统一退出。

## 3. defer

- 循环里 defer 可能导致资源延迟释放。
