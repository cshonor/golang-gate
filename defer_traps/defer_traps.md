# defer 易错点（线上 bug 常见来源）

## 1. 一句话总纲（背）

`defer` 的作用是 **“函数退出时执行”**，但它的陷阱来自两点：

- **参数在注册 defer 时就求值**  
- **defer 的执行顺序是 LIFO（后进先出）**

---

## 2. 5 大高频坑（面试必问）

### 坑 1：循环里 defer，资源堆积，直到函数结束才释放

```go
for _, name := range files {
    f, _ := os.Open(name)
    defer f.Close() // ❌ 循环 1w 次：1w 个文件句柄一直占着
}
```

正确做法：把循环体抽成函数，或显式 close。

```go
for _, name := range files {
    func() {
        f, err := os.Open(name)
        if err != nil { return }
        defer f.Close() // ✅ 这个匿名函数结束就释放
        // ...
    }()
}
```

---

### 坑 2：defer 参数提前求值，别指望它“读到最后的变量值”

```go
x := 1
defer fmt.Println(x) // 输出 1，不是 2
x = 2
```

如果你要“最后的值”，用闭包：

```go
x := 1
defer func() { fmt.Println(x) }() // 输出 2
x = 2
```

---

### 坑 3：defer + 命名返回值，容易被“偷偷改返回值”

```go
func f() (ret int) {
    defer func() { ret++ }()
    return 1
} // 返回 2
```

面试表述：`return` 先把返回值写到命名返回变量，再执行 defer，最后真正返回。

---

### 坑 4：锁/资源释放写法不统一，容易忘（线上最常见）

```go
mu.Lock()
defer mu.Unlock() // ✅ 最稳：拿锁后第一时间 defer 释放
```

同理：`Open` 后立即 `defer Close()`，`Begin` 后立即 `defer Rollback()`（提交成功再取消/覆盖）。

---

### 坑 5：recover 只有在 defer 里才生效

```go
func bad() {
    recover() // ❌ 无效
}
```

正确兜底：

```go
func safeGo(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                // 记录日志/指标，避免进程挂
            }
        }()
        fn()
    }()
}
```

---

## 3. 面试速记（背诵）

- `defer`：函数退出执行，**LIFO**
- defer 参数：**注册时求值**
- 命名返回：defer 能改返回值（慎用）
- 循环 defer：会堆积资源（抽函数/显式 close）
- recover：必须在 defer 里才有效

