# 01 - 错误介绍与 error 接口本质

本笔记说明 **Go 错误处理的设计哲学**，并补上工程里必须吃透的底层视角：**`error` 是接口**、接口值的内存布局，以及 **`errors.New` 与 `fmt.Errorf` 的差异与选型**。创建方式展开见 [03 - 错误创建](./03%20-%20错误创建-errors.New与fmt.Errorf.md)；语法与写法见 [02 - 错误基本语法](./02%20-%20错误基本语法.md)。

---

## 一、核心设计理念

Go 不用传统的 `try` / `catch` / `finally`，而用 **「错误作为返回值」的显式处理**。

### 1. 设计原则

- **显式处理**：错误是普通返回值，由调用方判断，而不是隐式抛出、远距离捕获。
- **就地处理**：失败在发生处附近处理或包装后上传，控制流不依赖长栈回溯。
- **可组合性**：错误是实现 `error` 接口的值，可包装、传递、用 `errors.Is` / `As` 判定（见 [05](./05%20-%20错误包装与错误链.md)、[06](./06%20-%20errors.Is与errors.As使用.md)）。

### 2. 为什么适合服务端与长期维护

- 线性控制流更易读、易调试；显式 `err` 减少「静默失败」。
- 与熔断、降级、日志、监控结合时，边界更清晰。

---

## 二、`error` 接口的本质

### 1. 接口定义

```go
type error interface {
	Error() string
}
```

任何实现了 `Error() string` 的具体类型，都可以赋给 `error` 变量并作为返回值传递。**判等、类型断言、`Is`/`As` 等行为，都建立在这套「接口 + 具体类型」模型上。**

### 2. 接口值的「类型 + 数据」

一个接口变量在运行时持有两部分（概念上常称为 **动态类型** 与 **动态值**）：

- **动态类型**：当前承载的具体类型（如 `*errors.errorString`、你自定义的 `*MyError`）。
- **动态值**：该类型的一个值（可能为 `nil`，见下文「坑」）。

因此：

- `fmt.Println(err)` 会调用具体类型的 `Error()`（经接口分派）。
- **不能用** `err.Error()` 的字符串去**稳定判断**错误类别（文案会变、国际化、包装层会改串）；应使用哨兵错误 + `errors.Is`，或自定义类型 + `errors.As`（见 [06](./06%20-%20errors.Is与errors.As使用.md)、[07](./07%20-%20Sentinel%20Error哨兵错误设计.md)、[10](./10%20-%20错误处理最佳实践与反模式.md)）。

### 3. 与 `nil` 相关的经典坑（值接收者 vs 指针）

若具体类型为指针，且你把「类型的 `nil`」塞进接口，接口值**不等于** `nil`：

```go
package main

import "fmt"

type MyErr struct{ msg string }

func (e *MyErr) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.msg
}

func mayFail(ok bool) error {
	var e *MyErr // nil 指针
	if !ok {
		return e // 返回 (*MyErr)(nil)，但接口里已带上类型信息
	}
	return nil
}

func main() {
	err := mayFail(false)
	fmt.Println(err == nil) // false — 动态类型是 *MyErr，动态值为 nil
}
```

**习惯**：失败时直接 `return nil`；需要返回「无错误」时不要通过「nil 的具体类型指针」再赋给 `error`。

### 4. 和 `panic` / `recover` 的边界

- **`error`**：可预期的失败（参数非法、文件不存在、网络超时等），属于正常控制流。
- **`panic` / `recover`**：严重、不应继续的逻辑；不能替代常规错误返回（见 [08](./08%20-%20panic与recover捕获异常.md)）。

---

## 三、`errors.New` 与 `fmt.Errorf`：实现差异与选型

### 1. `errors.New`

- 返回标准库内部的**静态字符串错误**（实现上通常是带固定文案的小类型），**无格式化**。
- 适合：**固定错误语义**、与包级 `var ErrXxx = errors.New(...)` 哨兵配合（见 [07](./07%20-%20Sentinel%20Error哨兵错误设计.md)）。

```go
var ErrNotFound = errors.New("not found")

func Lookup(id string) error {
	if id == "" {
		return ErrNotFound
	}
	return nil
}
```

### 2. `fmt.Errorf`（无 `%w`）

- 先做 **格式化**，再得到 `error`，适合要把 **路径、ID、数值** 等编进消息的场景。
- **默认用 `%v` 打印下层 `err` 会丢失错误链**；需要保留链时用 `%w`（见 [05](./05%20-%20错误包装与错误链.md)）。

```go
return fmt.Errorf("read %s: %v", path, err) // 仅文案，上层无法用 Is/As 追溯下层
```

### 3. `fmt.Errorf` + `%w`（Go 1.13+）

- 会构造**可 `Unwrap` 的包装错误**，保留下层，供 `errors.Is` / `errors.As` 遍历错误链。

```go
return fmt.Errorf("read %s: %w", path, err)
```

### 4. 选型小结

| 场景 | 更合适的 API |
|------|----------------|
| 包级固定语义、与 `errors.Is` 配合 | `errors.New` / 自定义变量 |
| 需要动态信息、单行消息 | `fmt.Errorf("...", ...)` |
| 上传错误且保留判等/类型能力 | `fmt.Errorf("...: %w", err)` |

更多创建与返回习惯见 [02](./02%20-%20错误基本语法.md)、[03](./03%20-%20错误创建-errors.New与fmt.Errorf.md)；自定义类型见 [04](./04%20-%20自定义错误类型.md)。

---

## 四、显式错误 vs 传统 try/catch（简要）

| 特性 | Go 显式错误 | 传统 try/catch |
|------|-------------|----------------|
| 控制流 | 线性 | 抛出/捕获打断流程 |
| 可见性 | 倾向显式处理 | 易被忽略 |
| 可组合性 | 错误是值，易包装 | 依赖异常体系 |

---

## 五、实践摘要

1. 把 **`error` 当接口** 理解：类型信息决定 `Is`/`As` 与 `nil` 比较是否合理。
2. **尽早** `if err != nil`；上传时尽量 **带上下文 + `%w`**。
3. **不要用** `err.Error() == "..."` 做分支；用哨兵或类型断言式 API。
4. 避免无说明的 `_ = err`（见 [10](./10%20-%20错误处理最佳实践与反模式.md)）。

---

## 六、可运行小例子：接口视角看 `errors.New`

```go
package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	err := errors.New("oops")
	fmt.Println(err.Error())

	// 观察动态类型（不同版本标准库具体类型名可能略有差异，但都是非导出实现类型）
	fmt.Println(reflect.TypeOf(err)) // 例如 *errors.errorString
}
```

理解本节之后，按顺序阅读 **02 → 03 → 04 → 05** 即可完成「创建 → 自定义 → 包装 → 判定」的闭环。
