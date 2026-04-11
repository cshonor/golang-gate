# 02 - defer 执行顺序规则

## 1. 核心顺序规则

多个 `defer` 严格遵循 **后进先出（LIFO，栈结构）** 执行：**最后注册的 defer 最先执行，最先注册的 defer 最后执行**。

```go
package main

import "fmt"

func main() {
	defer fmt.Println("A") // 第1个注册，最后执行
	defer fmt.Println("B") // 第2个注册，中间执行
	defer fmt.Println("C") // 第3个注册，最先执行
	fmt.Println("main logic")
}
// 输出：
// main logic
// C
// B
// A
```

## 2. 底层原理

- Go 运行时为每个 Goroutine 维护一个 **defer 链表（栈结构）**。
- 每个 Goroutine（`g` 结构体）都有一个 `_defer` 指针，指向当前 Goroutine 的 defer 链表头。
- 遇到 `defer` 时，将调用打包为 `_defer` 结构体，**头插法**加入链表（最新 defer 始终在链表头部）。
- 函数退出前，从链表头部开始 **依次弹出并执行** 所有 defer，天然形成 LIFO 顺序。

```go
// src/runtime/runtime2.go
type _defer struct {
	siz       int32     // 参数+返回值总字节数
	started   bool      // 是否已开始执行
	heap      bool      // 是否在堆上分配
	openDefer bool      // 是否为开放编码（open-coded）优化
	sp        uintptr   // defer 注册时的栈指针（用于栈帧校验）
	pc        uintptr   // defer 注册时的程序计数器（返回地址）
	fn        *funcval  // 延迟执行的函数（含地址与参数）
	_panic    *_panic   // 关联的 panic 结构体（panic 时执行）
	link      *_defer   // 链表指针，指向下一个 defer 节点

	// 开放编码优化专用字段
	fd   unsafe.Pointer // 函数元数据指针
	varp uintptr        // 栈帧变量指针
	// ... 其他内部字段
}
```

### `_defer` 的生命周期（注册 → 执行）

**注册（`deferproc`）**

- 编译器将 `defer fn()` 转为 `runtime.deferproc` 调用。
- 分配 `_defer` 结构体（栈 / 堆），拷贝参数，头插链表。

**执行（`deferreturn`）**

- 函数 `return` 前，runtime 调用 `deferreturn`。
- 遍历 `g._defer` 链表，从头部开始执行所有 `_defer` 节点的 `fn`。

**释放**

- 执行完毕后，`_defer` 节点被回收（栈上自动释放，堆上由 GC 回收）。

## 3. 为什么 LIFO 很重要

**完全匹配「资源申请 → 释放」的自然逻辑**：**后申请的资源必须先释放**，避免死锁、依赖错误或资源泄漏。

### 示例：嵌套锁安全释放

```go
import "sync"

var mu1, mu2 sync.Mutex

func nestedLock() {
	mu1.Lock()            // 先申请锁1
	defer mu1.Unlock()    // 最后释放锁1

	mu2.Lock()            // 后申请锁2
	defer mu2.Unlock()    // 先释放锁2

	// 业务逻辑...
}
// 执行顺序：lock1 → lock2 → 业务 → unlock2 → unlock1
```

### 示例：多层文件/目录操作

```go
import "os"

// 先创建外层目录，再创建内层文件；删除时必须先删文件，再删目录
func createAndClean() {
	os.MkdirAll("outer/inner", 0755)
	defer os.RemoveAll("outer") // 最后删外层

	f, _ := os.Create("outer/inner/file.txt")
	defer f.Close()                     // 先关文件
	defer os.Remove("outer/inner/file.txt") // 先删文件
}
```

## 4. 关键边界与实践

### （1）执行时机

-所有 defer 统一在：函数 return 之后、RET 指令返回调用方之前 执行。

-函数正常返回、panic、runtime 退出，均会触发 defer 执行。

RET 指令：底层汇编的 “函数返回” 指令，真正把结果还给调用者。

3 种触发场景：

1.正常 return
2.发生 panic（崩溃）
3.runtime.Goexit() 主动退出

执行顺序：LIFO（后进先出）—— 最后 defer 的最先执行。

### （2）与 return、返回值的关系

`return` 分三步：

1. 计算并赋值给 **具名返回值**；
2. 按 LIFO 执行所有 defer（可修改具名返回值）；
3. 执行 RET 指令，返回结果。

```go
func namedReturn() (result int) {
	defer func() { result++ }() // defer 修改返回值
	return 1                    // 先赋值 result=1，再执行 defer → result=2
}
// 调用 namedReturn() → 返回 2
```

### （3）实践建议

- **资源申请后立即 defer 释放**，严格按「先申请、后 defer」的顺序书写。
- 复杂场景（多层锁、事务、文件）给 defer 加注释，明确释放顺序。
- 循环内 **慎用 defer**（累积到函数结束才执行，易导致资源泄漏）。

---

## 复习速记

| 要点 | 内容 |
|------|------|
| 执行顺序 | LIFO：后注册先执行 |
| 底层结构 | 每 G 一条 defer 链表，头插、从头弹出 |
| 与 return | 先具名赋值 → defer → RET |
| 实践 | 申请后即 defer；循环里慎用 |
