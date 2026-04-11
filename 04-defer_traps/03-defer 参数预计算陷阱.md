# 03 - defer 参数预计算陷阱

`defer` 最经典、面试必考的坑之一：**参数在注册 `defer` 时求值**，与「闭包在执行时再读变量」不是一回事。执行顺序总览见 [02 - defer 执行顺序规则](./02-defer%20执行顺序规则.md)；循环里 `defer` 见 [05](./05-defer%20在循环中的风险.md)。

---

## 一、核心结论（一句话）

> **`defer` 里被调函数的实参，在「注册 defer」的那一刻就会求值并固定下来；闭包 `defer func(){ ... }()` 里没有把变量做成实参时，函数体里对外层变量的读取发生在 defer 真正执行时，因此能读到后续修改后的值。**

---

## 二、两种写法对比

### 1. 陷阱写法：直接传参 → 值被“焊死”在注册时

```go
package main

import "fmt"

func main() {
	x := 1
	// 注册 defer 时：先求值实参 x → 得到 1，相当于 defer fmt.Println(1)
	defer fmt.Println(x)
	x = 2
}
// 输出：1
```

**流程拆解**：

1. `x := 1`
2. `defer fmt.Println(x)`：对 `fmt.Println` 的实参 `x` **立即求值**，把「调用 `fmt.Println(1)`」这条 defer 记下来
3. `x = 2`：不影响已固定的实参
4. 函数退出前执行 defer：打印 `1`

### 2. 常见正确写法：闭包 → 执行时再读 `x`

```go
package main

import "fmt"

func main() {
	x := 1
	// 注册时没有把 x 作为“被 defer 调用的函数的实参”固定下来
	defer func() { fmt.Println(x) }()
	x = 2
}
// 输出：2
```

**流程拆解**：

1. `x := 1`
2. `defer func() { fmt.Println(x) }()`：入栈的是**无参闭包**；闭包体里对 `x` 的读取发生在 **defer 执行阶段**
3. `x = 2`
4. 退出前执行闭包：此时读到的 `x` 为 `2`

---

## 三、底层直觉（不必背 runtime 字段名）

Go 处理 `defer` 可以粗分为两阶段：

1. **注册阶段**：为 `defer f(a, b, ...)` 计算并保存 **`a,b,...` 的快照**（值/指针等按表达式语义），把「稍后调用 `f` + 这些实参」压入当前 goroutine 的 defer 链。
2. **执行阶段**：函数返回路径上按 **LIFO** 弹出并执行。

因此：

- **`defer f(x)`**：`x` 作为实参，属于注册阶段求值。
- **`defer func(){ f(x) }()`**：注册阶段只保存闭包；`x` 在闭包执行时再求值（仍受「循环变量只有一个」等规则约束，见下一节）。

---

## 四、延伸陷阱：循环里的 defer 闭包（高频）

### 错误：闭包捕获同一个循环变量

```go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }()
	}
}
// 常见输出：3 3 3（不是 0 1 2）
```

原因：`i` 在整个循环里只有一个变量槽；所有 defer 闭包都引用它，执行 defer 时循环已结束，`i` 常为 `3`。

### 方案 A：传参 → 用注册时的快照隔离每次循环

```go
for i := 0; i < 3; i++ {
	defer func(n int) { fmt.Println(n) }(i)
}
// 输出顺序：2 1 0（defer LIFO）
```

### 方案 B：每次循环新建局部变量

```go
for i := 0; i < 3; i++ {
	n := i
	defer func() { fmt.Println(n) }()
}
// 输出顺序：2 1 0
```

更多循环资源问题见 [05](./05-defer%20在循环中的风险.md)。

---

## 五、实战：`time.Since` 统计耗时（真实踩坑）

### 错误：`time.Since(start)` 在注册时就求值

```go
package main

import (
	"fmt"
	"time"
)

func doSomething() {
	start := time.Now()
	// ❌ time.Since(start) 作为实参在 defer 注册时就计算，接近 0
	defer fmt.Printf("耗时: %v\n", time.Since(start))

	time.Sleep(1 * time.Second)
}
// 典型输出：耗时: 0s（错误）
```

### 正确：放进闭包，执行时再算

```go
package main

import (
	"fmt"
	"time"
)

func doSomething() {
	start := time.Now()
	defer func() { fmt.Printf("耗时: %v\n", time.Since(start)) }()

	time.Sleep(1 * time.Second)
}
```

---

## 六、避坑速查表

| 写法 | 关键求值时机 | 后续修改外层变量是否影响结果 | 典型场景 |
|------|----------------|------------------------------|----------|
| `defer f(x)` | 注册 `defer` 时对 `x` 求值 | 一般 **不影响**（已被快照） | 想把当时的值固定下来 |
| `defer func() { f(x) }()` | 执行 defer 时再读 `x` | 一般 **影响**（读最新值） | 耗时统计、读最终状态 |
| 循环 + 闭包捕获循环变量 | 执行时再读 `i` | 常 **全部读到最终值** | 用传参或 `n := i` 修复 |

---

## 七、和「事务 / recover」类代码的关系

`defer func() { ... tx ... }()` 里直接引用外层的 `tx`、`err`，通常就是希望 **执行时再读取最新状态**；这与本节闭包写法一致。错误处理与 `defer` 协同见 [07](./07-defer%20与错误处理协同.md)、[03-error_handling/09](../03-error_handling/09%20-%20defer与错误处理协同.md)。

---

## 延伸阅读

- [04 - defer 与 return 执行流程](./04-defer%20与%20return%20执行流程.md) · [08 - 最佳实践与反模式](./08-defer%20最佳实践与反模式.md)
