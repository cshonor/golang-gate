# 08 - defer 最佳实践与反模式

把 [01](./01-defer%20基本用法.md)～[07](./07-defer%20与错误处理协同.md) 的规则收束成**可执行的工程习惯**与**面试速查**。与错误处理总览对照：[03-error_handling/10](../03-error_handling/10%20-%20错误处理最佳实践与反模式.md)。

---

## 1. 最佳实践（建议照做）

1. **资源申请成功后立刻 `defer`**，减少分支遗漏（文件、锁、事务）。
2. **`Close` / `Commit` 等返回 `error` 的清理**：不要裸 `defer r.Close()`；用 `defer func(){ … }()` 处理或合并进返回 `err`（见 [07](./07-defer%20与错误处理协同.md)）。
3. **需要改返回值的场景**：优先 **命名返回** + defer 内赋值，并写清注释，避免读者误解（与 [04](./04-defer%20与%20return%20执行流程.md) 一致）。
4. **锁**：`defer mu.Unlock()` 紧跟 `Lock()`，标准写法。
5. **`recover`**：只放在 defer 内；边界捕获后记录日志/指标，业务层仍优先返回 `error`（见 [03-error_handling/08](../03-error_handling/08%20-%20panic与recover捕获异常.md)）。
6. **与错误链**：向上传递原因时用 `fmt.Errorf("...: %w", err)`（[05](../03-error_handling/05%20-%20错误包装与错误链.md)）。
7. **热点路径**：defer 有成本（见 [06](./06-defer%20底层实现原理.md)）；先保证正确，再 profile 决定是否手写清理。

---

## 2. 反模式（避免）

| 反模式 | 后果 | 更好方向 |
|--------|------|----------|
| 循环里 `defer Close` | 句柄/内存堆积到函数结束 | 抽函数、`func(){...}()` 包一层，或显式 `Close` |
| `defer f.Close()` 从不看返回值 | 静默丢错 | [07](./07-defer%20与错误处理协同.md) 中的模式 A/B |
| defer 里做重逻辑、RPC、大锁 | 延迟退出、难测 | defer 只做轻量清理 |
| 深层嵌套多个 defer | 难读、顺序难想 | 合并清理或抽 `cleanup()` |
| 误用 `recover` 处理业务错误 | 控制流混乱 | 可预期失败用 `error` |
| 忽略参数预计算 | 闭包读到错误值 | 见 [03](./03-defer%20参数预计算陷阱.md) |

---

## 3. 按场景的「标准骨架」

### 3.1 文件 IO

```go
f, err := os.Open(path)
if err != nil {
	return fmt.Errorf("open: %w", err)
}
defer func() {
	if cerr := f.Close(); cerr != nil {
		// log 或合并进 err，按团队约定
	}
}()
```

### 3.2 互斥锁

```go
mu.Lock()
defer mu.Unlock()
```

### 3.3 HTTP Handler 内局部资源

与 3.1 相同；**边界**处统一打日志、映射状态码见 [03-error_handling/11](../03-error_handling/11%20-%20错误与日志集成.md)。

### 3.4 并发与 channel

`defer` 常用于 `wg.Done()`、`close(ch)`（注意 `close` 只关一次）；错误收集常用 `errgroup`，错误模型见 `03-error_handling` 与 `09-concurrency` 相关篇。

---

## 4. 面试速记（可与 `defer_traps.md` 一起背）

- 注册顺序与执行顺序：**LIFO**
- defer 参数：**注册时求值**（[03](./03-defer%20参数预计算陷阱.md)）
- 命名返回：`defer` 可改最终返回值（[04](./04-defer%20与%20return%20执行流程.md)）
- 循环 + defer：**大坑**（[05](./05-defer%20在循环中的风险.md)）
- 清理阶段的 `error`：**要有策略**（[07](./07-defer%20与错误处理协同.md)）

---

## 延伸阅读

- 总览速背：[defer_traps.md](./defer_traps.md) · [06 - 底层实现](./06-defer%20底层实现原理.md)
