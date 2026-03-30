# 01 - sync.Mutex 互斥锁原理与实战

与 `08-atomic and lock/` 的「锁实现原理」互补：这里偏怎么用与常见坑。

---

## 1. 基本模板

```go
mu.Lock()
defer mu.Unlock()
// 临界区
```

---

## 2. 常见坑

- 锁内做慢 IO
- 多锁顺序不一致（死锁）
