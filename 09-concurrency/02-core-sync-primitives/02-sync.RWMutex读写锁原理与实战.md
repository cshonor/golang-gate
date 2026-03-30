# 02 - sync.RWMutex 读写锁原理与实战

---

## 1. 基本模板

```go
mu.RLock()
// read
mu.RUnlock()

mu.Lock()
// write
mu.Unlock()
```

---

## 2. 注意

- 不要尝试从 `RLock` 升级到 `Lock`。
