# 03 - 错误创建：`errors.New` 与 `fmt.Errorf`

在 [01](./01%20-%20错误介绍与error接口本质.md) 里已从「接口与选型」角度概括过本节主题；这里把 **创建错误** 单独成篇，方便对照标准库行为与工程习惯。语法速查仍可见 [02 - 错误基本语法](./02%20-%20错误基本语法.md)。

---

## 1. `errors.New`

- 返回一个**仅含固定字符串**的 `error`，实现类型为标准库内部类型（可视为**不透明**）。
- **无格式化参数**：文案在编译期固定，适合 **包级哨兵**：`var ErrXxx = errors.New("...")`（设计模式见 [07](./07%20-%20Sentinel%20Error哨兵错误设计.md)）。

```go
var ErrEmptyID = errors.New("empty id")

func Parse(id string) error {
	if id == "" {
		return ErrEmptyID
	}
	return nil
}
```

---

## 2. `fmt.Errorf`（不含 `%w`）

- 先按格式串拼接出**一条字符串**，再作为 `error` 返回；适合要把 **动态值**（路径、ID、计数）写进消息的场景。
- 使用 **`%v` 打印下层 `error`** 时，**不会**把下层挂进可 `Unwrap` 的链，上层无法用 `errors.Is` / `As` 追溯到原因。

```go
return fmt.Errorf("load %s failed: %v", path, err) // 仅人类可读，链断
```

---

## 3. `fmt.Errorf` + `%w`（Go 1.13+）

- `%w` 专门用于包装 **`error`**，生成带 `Unwrap()` 的包装错误，**错误链** 成立，后续可用 `errors.Is` / `errors.As`（见 [05](./05%20-%20错误包装与错误链.md)、[06](./06%20-%20errors.Is与errors.As使用.md)）。

```go
return fmt.Errorf("load %s: %w", path, err)
```

---

## 4. `error` 接口与返回值在「实现侧」长什么样

Go 里 **`error` 是接口**，只要实现了 `Error() string` 即可：

```go
type error interface {
	Error() string
}
```

`fmt.Errorf` 返回的具体类型是标准库**内部实现**（包外不可引用类型名），随版本可能微调；从**语义**上可以这样理解：

### 4.1 无 `%w` 时

- 本质是「**一条格式化好的字符串**」对应的 `error`，通常**没有**可供 `errors.Unwrap` 继续走的下层。
- 教学用简化模型（与 `errors.New` 类似：**只有 `Error()`**）：

```go
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

### 4.2 有 `%w` 时

- 除 `Error()` 外，还需支持 **错误链**，因此实现上带有 **`Unwrap() error`**（见 `errors` 包的链式约定）。
- 教学用简化模型：

```go
type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}
```

**结论**：`fmt.Errorf` 的返回值**一定满足** `error`；使用 **`%w`** 时还会在语义上满足 **`Unwrap()`**，从而支持 `errors.Is` / `As` 沿链查找。

---

## 5. 可运行示例：有无 `%w` 时 `errors.Is` 的差异

下面用**同一个**底层哨兵 `ErrRoot`，对比包装方式对 `errors.Is` 的影响。

```go
package main

import (
	"errors"
	"fmt"
)

var ErrRoot = errors.New("root cause")

func failWithV() error {
	return fmt.Errorf("context: %v", ErrRoot)
}

func failWithW() error {
	return fmt.Errorf("context: %w", ErrRoot)
}

func main() {
	ev := failWithV()
	ew := failWithW()

	fmt.Println("without %w, errors.Is(ev, ErrRoot):", errors.Is(ev, ErrRoot)) // false
	fmt.Println("with    %w, errors.Is(ew, ErrRoot):", errors.Is(ew, ErrRoot)) // true
}
```

- **无 `%w`**：新错误只是把 `ErrRoot` 的**字符串**拼进消息，链断，`Is` 失败。
- **有 `%w`**：保留底层，`Is` 仍能识别 `ErrRoot`。

---

## 6. 选型速查

| 需求 | 用法 |
|------|------|
| 固定语义、配合 `errors.Is` | `errors.New` / 包级 `var Err…` |
| 动态信息、不需要保留链 | `fmt.Errorf("... %v ...", x, err)` |
| 上传且保留判定能力 | `fmt.Errorf("...: %w", err)` |

---

## 7. 与后续章节的关系

- 需要 **字段、错误码**：在 `errors.New` / `fmt.Errorf` 之外定义 **自定义类型**（见 [04](./04%20-%20自定义错误类型.md)）。
- **多层包装**、避免断链：见 [05](./05%20-%20错误包装与错误链.md)。
