# 05 - defer 在循环中的风险

循环里误用 `defer`，是线上 **`too many open files`**、连接池耗尽、内存上涨的常见来源之一。参数预计算与闭包捕获叠加见 [03](./03-defer%20参数预计算陷阱.md)；`return` 与 `defer` 顺序见 [04](./04-defer%20与%20return%20执行流程.md)。

---

## 1. 核心风险（一句话）

**在循环体里直接 `defer`，所有 `defer` 会堆积到外层函数返回前才执行**，导致文件句柄、DB 连接、锁、内存等资源长时间占用，甚至拖垮进程。

---

## 2. 典型错误（必踩坑）

```go
// ❌ 错误：循环内 defer，所有 Close 等外层函数结束才执行
func processFiles(paths []string) {
	for _, p := range paths {
		f, _ := os.Open(p)
		defer f.Close() // 风险：文件一直开着，直到本函数退出才关闭
		// 处理文件...
	}
}
```

**后果**：

- 循环成千上万次时，可能**同时持有大量打开句柄**，触发系统/进程限制（常见现象是 `too many open files`，具体阈值因 OS/配置而异）。
- 数据库连接、网络连接、锁等同样会堆积。
- `defer` 记录增长，带来额外内存与调度开销。

---

## 3. 正确做法（三种方案）

### 方案 1：匿名函数封装（推荐，保留 `defer` 语义）

每次迭代用 **IIFE**（立即执行函数）包一层，让 `defer` 绑定在**子函数**的返回路径上：

```go
// ✅ 正确：匿名函数包裹，defer 在每次迭代结束执行
func processFiles(paths []string) {
	for _, p := range paths {
		func() {
			f, err := os.Open(p)
			if err != nil {
				log.Print(err)
				return
			}
			defer f.Close()

			// 处理文件...
		}()
	}
}
```

### 方案 2：显式 `Close`（简单直接）

不用 `defer`，用完立刻关，并处理 `Close` 的错误（与 [07](./07-defer%20与错误处理协同.md)、[03-error_handling/09](../03-error_handling/09%20-%20defer与错误处理协同.md) 协同）：

```go
// ✅ 正确：显式关闭，及时释放
func processFiles(paths []string) {
	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			log.Print(err)
			continue
		}

		// 处理文件...

		if err := f.Close(); err != nil {
			log.Print("close err:", err)
		}
	}
}
```

### 方案 3：抽成独立小函数（最规范、易复用）

```go
// ✅ 最佳：抽离为独立函数
func processFile(p string) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	// 处理文件...
	return nil
}

func processFiles(paths []string) {
	for _, p := range paths {
		if err := processFile(p); err != nil {
			log.Print(err)
		}
	}
}
```

---

## 4. 面试一句话总结

> **循环里直接 `defer` 会把资源释放推迟到外层函数退出，是句柄/连接堆积的高频坑；用 IIFE、显式关闭或抽函数，让释放发生在每次迭代结束。**

---

## 5. 延伸：循环 + `defer` + 闭包变量（双重坑）

循环变量只有一个槽位，所有 `defer` 闭包共享它，执行时往往已是最终值（详解见 [03](./03-defer%20参数预计算陷阱.md) §四）。

```go
// ❌ 常见输出：3 3 3
for i := 0; i < 3; i++ {
	defer func() { fmt.Println(i) }()
}

// ✅ 传参快照（输出顺序 2 1 0，因 defer LIFO）
for i := 0; i < 3; i++ {
	defer func(n int) { fmt.Println(n) }(i)
}

// ✅ 每次迭代新局部变量
for i := 0; i < 3; i++ {
	n := i
	defer func() { fmt.Println(n) }()
}
```

---

## 6. 避坑口诀

- **循环别裸 defer，资源堆积必出事**
- **IIFE 包一层，defer 跟着迭代走**
- **长循环抽函数，清晰不泄漏**

---

## 7. 一页速记 + 5 道高频题（复习用）

### 速记

| 主题 | 结论 |
|------|------|
| 循环 + `defer` | 默认堆积到**外层函数退出** |
| 修法 | IIFE / 显式 Close / 抽函数 |
| + 闭包 | 注意循环变量捕获（传参或 `n:=i`） |
| `return` vs `defer` | 见 [04](./04-defer%20与%20return%20执行流程.md) |
| 实参求值 | 见 [03](./03-defer%20参数预计算陷阱.md) |

### 5 道高频题（简答）

1. **问**：`for { defer f.Close() }` 何时执行 `Close`？  
   **答**：外层函数返回前；循环中不会每次迭代结束就执行。

2. **问**：如何保留 `defer` 又避免堆积？  
   **答**：`func(){ ...; defer ... }()` 或抽 `processOne()`。

3. **问**：`for i:=0;i<3;i++ { defer func(){ println(i) }() }` 打印？  
   **答**：常见为 `3 3 3`（共享循环变量）。

4. **问**：`defer fmt.Println(i)` 与 `defer func(){ fmt.Println(i) }()` 在循环里区别？  
   **答**：前者 `i` 在注册 `defer` 时求值（每次迭代快照）；后者闭包执行时再读 `i`（易踩共享变量坑）。

5. **问**：不用 `defer` 时 `Close` 要注意什么？  
   **答**：错误路径也要关；`Close` 的错误不要默认吞（见错误处理专题）。

---

## 延伸阅读

- [02 - defer 执行顺序规则](./02-defer%20执行顺序规则.md) · [08 - 最佳实践与反模式](./08-defer%20最佳实践与反模式.md)
