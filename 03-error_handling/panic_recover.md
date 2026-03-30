# panic / recover（面试与工程都要会）

## 1. 一句话结论（背）

> **能返回 error 就返回 error**；`panic` 用于“程序无法继续”的 bug/不变量破坏；`recover` 多用于“边界兜底”，防止整个进程被拖死。

---

## 2. panic 是什么？

- `panic` 会立刻开始“向上退栈”
- 退栈过程中会执行每层函数注册的 `defer`
- 如果最终没人 `recover`，程序崩溃（打印栈）

---

## 3. recover 为什么必须在 defer 里？

`recover()` 只有在 **defer 函数** 的调用链上才会捕获当前 panic。

标准模板：

```go
func runSafely(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    fn()
    return nil
}
```

---

## 4. goroutine 边界兜底（服务必备）

panic 一旦发生在 goroutine 中，如果没兜底，可能导致整个进程退出（或把关键 goroutine 弄死）。

```go
func GoSafe(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                // 记录日志/指标，必要时报警
            }
        }()
        fn()
    }()
}
```

---

## 5. 面试追问：哪些场景“应该 panic”？

- 违反不变量（逻辑 bug）：例如内部状态不可能发生却发生了
- 编程错误：数组越界、空指针解引用（本质也是 bug）
- 初始化失败且无法继续：极少数场景（比如核心配置缺失）

业务可预期错误（参数不合法、资源不存在、网络超时）：

- ✅ 返回 `error`  
- ❌ 不要 `panic`

