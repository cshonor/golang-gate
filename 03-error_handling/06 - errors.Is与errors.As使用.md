# 06 - errors.Is 与 errors.As 使用

前置：[05 - 错误包装与错误链](./05%20-%20错误包装与错误链.md)（`%w` 与 `Unwrap`）。哨兵错误设计见 [07](./07%20-%20Sentinel%20Error哨兵错误设计.md)；自定义类型见 [04](./04%20-%20自定义错误类型.md)。

> **一句话**：**`Is` 比「是否与目标错误匹配」**（常见是**哨兵**相等，或下文 **§6** 的**自定义 `Is`**）；**`As` 比「链上是否出现过某 `*T`」**并取出实例。**二者内部都会自动沿链 `Unwrap`**，无需你手写循环（与 **§4** 单次 `Unwrap` 手撕相对）。

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

## 3. 对照：`Is` vs `As`（背表 + 口诀）

| 函数 | 主要在比什么 | 典型场景 | 写法 |
|------|----------------|----------|------|
| `errors.Is` | **目标错误是否匹配**（默认 **`==`**，链上某层；可配合 **自定义 `Is`**，见 §6） | `os.ErrNotExist`、`io.EOF`、包级哨兵 `var ErrX = errors.New(...)` | `errors.Is(err, target)` |
| `errors.As` | **是否存在某动态类型 `*T`**（类型断言语义） | 自定义 `PathError`、`*UserErr` 等，需要 **读字段** | `var u *T; errors.As(err, &u)` |

**底层共性**：两者都会**内部循环**，沿错误链反复 **`Unwrap`**，直到匹配或链结束；区别在每一层做的是 **相等/自定义 `Is`**（`Is`）还是 **类型断言**（`As`）。

**口诀**：

- **Is 找「是不是这个失败（哨兵/匹配语义）」，As 找「是不是这个类型并取出」。**
- **哨兵 / 标准库 `ErrXxx` 成败 → `Is`；要拿结构体字段 → `As`。**

**高频踩坑**：

- 经 `%w` 包装后，**`err == ErrX` 常为 `false`**，必须用 **`errors.Is`** 沿链找哨兵。
- 需要 **`Code` / `Field` 等字段**时，用 **`errors.As`** 取出 `*T`；**不要**指望用 `Is`「拿出字段」（`Is` 的第二个参数是你要比对的那个 `target`，不是接收结果的槽位）。
- 自定义类型若实现了 **§6** 的 **`Is(target) bool`**，`errors.Is` 的「相等」语义会被扩展，与单纯「同一变量」略有不同，面试常考。

---

## 4. `errors.Unwrap`：只解一层（不递归）

- `errors.Unwrap(err)` **每次只解开一层**，得到**直接内层**的 `error`；**不会**自动循环，**不会**一次剥完整条链。
- 想走完整条链：要么**自己写循环**反复 `Unwrap`，要么直接用 **`errors.Is` / `errors.As`**（二者内部会按链遍历，见下节）。标准库**没有**名为 `OnWrapper` 之类的 API。

### 4.1 三层 `%w` 示例

```go
rootErr := errors.New("db connect fail")
layer1 := fmt.Errorf("query user: %w", rootErr)
layer2 := fmt.Errorf("handle request: %w", layer1)
// 语义链：layer2 → layer1 → rootErr

step1 := errors.Unwrap(layer2) // → layer1
step2 := errors.Unwrap(step1)  // → rootErr
step3 := errors.Unwrap(step2)  // → nil（`errors.New` 无下层）
```

### 4.2 `Unwrap` 单层 vs `Is` / `As` 自动遍历

| API | 行为 |
|-----|------|
| `errors.Unwrap` | **手动、单次**；需自行 `for` 才能遍历整条链 |
| `errors.Is` / `errors.As` | 实现里**循环**调用 `Unwrap`（及 `Is` 匹配规则），直到命中或链结束 |

手动打印整条链：

```go
for e := err; e != nil; e = errors.Unwrap(e) {
	fmt.Println(e)
}
```

---

## 5. 可运行综合示例（`Is` + `As` + 多层 `%w`）

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

## 6. 高级：自定义 `Is(target error) bool`（匹配规则）

`errors.Is` 在遍历链时，除 **`==` 相等** 外，若某一环的具体类型实现了：

```go
func (e *T) Is(target error) bool
```

则会调用该方法，由你定义「**怎样算与 `target` 匹配**」。这是 **`interface { Is(error) bool }` 式的方法**，**不是**在业务里再声明一个名叫 `Is` 的独立接口类型；详见 [`errors.Is` 文档](https://pkg.go.dev/errors#Is)。

内部 `errors.New` / `fmt` 的包装类型见 [05](./05%20-%20错误包装与错误链.md) §2.3；**自定义业务错误**可在同一类型上按需实现 `Error()`、`Unwrap()`、`Is()`。

### 6.1 示例：按错误码匹配，而非整结构体相等

```go
package main

import (
	"errors"
	"fmt"
)

type BizErr struct {
	Code int
	Msg  string
}

func (e *BizErr) Error() string {
	return fmt.Sprintf("[%d]%s", e.Code, e.Msg)
}

// 自定义：只要 Code 相同即视为与 target「同类」匹配
func (e *BizErr) Is(target error) bool {
	t, ok := target.(*BizErr)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func main() {
	e1 := &BizErr{Code: 500, Msg: "server down"}
	wrapErr := fmt.Errorf("wrap: %w", e1)

	// 文案不同，但 Code 同为 500 → Is 为 true
	target := &BizErr{Code: 500, Msg: "任意文案"}
	fmt.Println(errors.Is(wrapErr, target)) // true
}
```

### 6.2 与 `Unwrap` 串联（提要）

`errors.Is` 会**先沿链**逐层 `Unwrap`（含 `fmt` 的 `%w`、以及你自己实现的 `Unwrap() error`），再在每一层尝试 **`==`** 与 **自定义 `Is`**。因此多层包装后仍可能命中你在某层实现的 `Is`。

---

## 7. `errors.As` 能自定义吗？

**不能。** `errors.As` 固定为：**沿链 `Unwrap` + 类型匹配并写入指针**，没有像 `Is` 那样的「每类型自定义 `As`」扩展点。

---

## 8. 谁实现什么（速查）

| 来源 | `Error()` | `Unwrap()` | 自定义 `Is` |
|------|-----------|------------|-------------|
| `errors.New` | ✓（内部 `errorString` 等） | 一般无下层 | 一般无 |
| `fmt.Errorf("...: %w", err)` | ✓（内部包装） | ✓ | 由内层错误决定 |
| 自建 `*MyErr` | 你实现 | 可选自实现 | 可选，用于定制 `Is` 语义 |

---

## 9. 记忆法

- **`Unwrap`**：**一次只剥一层**；走全链需自写循环或改用 `Is`/`As`。
- **Is**：判**是不是**某个错误（哨兵、标准库 `ErrXxx`）；链上类型可实现 **`Is(target) bool`** 自定义匹配。
- **As**：**拿出来**某种具体类型；**不能**自定义匹配逻辑。
- 包装只用 **`%w`**，`%v` 会断链（见 [05](./05%20-%20错误包装与错误链.md)）。

---

## 延伸阅读

- [07 - Sentinel Error 哨兵错误设计](./07%20-%20Sentinel%20Error哨兵错误设计.md) · [10 - 最佳实践与反模式](./10%20-%20错误处理最佳实践与反模式.md)
