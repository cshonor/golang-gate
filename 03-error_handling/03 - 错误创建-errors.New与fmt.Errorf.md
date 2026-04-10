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

- 先按格式串拼接出**一条字符串**，再作为错误返回；适合要把 **动态值**（路径、ID、计数）写进消息的场景。
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

## 4. 选型速查

| 需求 | 用法 |
|------|------|
| 固定语义、配合 `errors.Is` | `errors.New` / 包级 `var Err…` |
| 动态信息、不需要保留链 | `fmt.Errorf("... %v ...", x, err)` |
| 上传且保留判定能力 | `fmt.Errorf("...: %w", err)` |

---

## 5. 与后续章节的关系

- 需要 **字段、错误码**：在 `errors.New` / `fmt.Errorf` 之外定义 **自定义类型**（见 [04](./04%20-%20自定义错误类型.md)）。
- **多层包装**、避免断链：见 [05](./05%20-%20错误包装与错误链.md)。
