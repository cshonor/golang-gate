# 06 - errors.Is 与 errors.As 使用

前置：[05 - 错误包装与错误链](./05%20-%20错误包装与错误链.md)（`%w` 与 `Unwrap`）。哨兵错误设计见 [07](./07%20-%20Sentinel%20Error哨兵错误设计.md)；自定义类型见 [04](./04%20-%20自定义错误类型.md)。

---

## 1. `errors.Is`：链上是否「就是」某个错误

用于判断**整条错误链**里是否出现过目标错误（含目标类型自定义了 `Is` 时的行为）。**不要**用 `err == ErrXxx` 代替（包装后外层是新值，`==` 常为 `false`）。

### 1.1 最小用法

```go
if errors.Is(err, ErrNotFound) {
	// 404、空结果等分支
}
```

### 1.2 与 `%w` 配合（典型）

```go
var ErrNotFound = errors.New("not found")

func load(id string) error {
	return fmt.Errorf("load id=%s: %w", id, ErrNotFound)
}

func main() {
	err := load("x")
	fmt.Println(errors.Is(err, ErrNotFound)) // true：沿链能找到 ErrNotFound
	fmt.Println(err == ErrNotFound)        // false：外层是包装错误，不是同一指针/值
}
```

### 1.3 标准库哨兵示例

```go
_, err := os.Open("不存在的文件.txt")
if errors.Is(err, os.ErrNotExist) {
	// 文件不存在
}
```

---

## 2. `errors.As`：从链上「取出」某具体类型的错误

第二个参数必须是**指向错误变量的指针**：若链上匹配的类型是 `*MyError`，则常见写法是先声明 `var ve *MyError`，再 `errors.As(err, &ve)`（即传入 `**MyError`，由 `As` 写入匹配到的 `*MyError`）。

```go
var ve *ValidationError
if errors.As(err, &ve) {
	_ = ve.Field // 使用具体字段
}
```

### 2.1 与 `%w` 包装后的自定义错误

```go
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return e.Msg
}

func validate() error {
	return &ValidationError{Field: "email", Msg: "invalid"}
}

func handler() error {
	return fmt.Errorf("request: %w", validate())
}

func main() {
	err := handler()
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println(ve.Field, ve.Msg) // email invalid
	}
}
```

### 2.2 易错点

- 必须传 **`&变量`**，且变量类型为 `*T`（例如 `var ve *ValidationError` 再 `errors.As(err, &ve)`）。传错层级的指针会导致 `As` 失败或不符合预期。
- 若链上**没有**该类型， `As` 返回 `false`，目标变量保持原样，使用前务必判断返回值。

---

## 3. 对照：`Is` vs `As`

| API | 典型输入 | 典型用途 |
|-----|----------|----------|
| `errors.Is(err, target)` | 哨兵：`var ErrX = errors.New(...)` 或与某 `error` **可比较相等** | 问「是不是这一类失败」 |
| `errors.As(err, &ptr)` | `ptr` 为 `**T`，指向 `*T` 变量 | 问「链上有没有 `*T`」并取出字段 |

---

## 4. 可运行综合示例（`Is` + `As` + 多层 `%w`）

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrNotFound = errors.New("not found")

type BizError struct {
	Code int
	Msg  string
}

func (e *BizError) Error() string { return e.Msg }

func deep() error {
	return fmt.Errorf("layer1: %w", fmt.Errorf("layer2: %w", ErrNotFound))
}

func deepBiz() error {
	return fmt.Errorf("biz: %w", &BizError{Code: 400, Msg: "bad"})
}

func main() {
	e1 := deep()
	fmt.Println("Is NotFound:", errors.Is(e1, ErrNotFound))

	e2 := deepBiz()
	var be *BizError
	if errors.As(e2, &be) {
		fmt.Println("As BizError Code:", be.Code)
	}

	// 标准库
	_, err := os.Open("nope.txt")
	fmt.Println("Is ErrNotExist:", errors.Is(err, os.ErrNotExist))
}
```

---

## 5. 记忆法

- **Is**：判**是不是**某个错误（尤其哨兵、标准库 `ErrXxx`）。
- **As**：**拿出来**某种具体类型，读字段、做分支。
- 包装只用 **`%w`**，`%v` 会断链（见 [05](./05%20-%20错误包装与错误链.md)）。

---

## 延伸阅读

- [07 - Sentinel Error 哨兵错误设计](./07%20-%20Sentinel%20Error哨兵错误设计.md) · [10 - 最佳实践与反模式](./10%20-%20错误处理最佳实践与反模式.md)
