# 07 - Sentinel Error（哨兵错误）设计

**哨兵错误（Sentinel error）**：提前在包级只创建一次、用于“标记某一类失败”的固定 `error` 值。典型写法是 **`var ErrXxx = errors.New("...")`**，调用方用 **`errors.Is(err, ErrXxx)`** 判断（即使中间被 `%w` 多层包装），而不是用 `err.Error()` 字符串比较。

前置：[03](./03%20-%20错误创建-errors.New与fmt.Errorf.md)、[06](./06%20-%20errors.Is与errors.As使用.md)。

---

## 0. 哨兵是什么（定义）

一句话：**哨兵是“全局唯一的错误标记”，用来做分支判断。**

它的关键特征：

- **包级唯一实例**：只 `errors.New` 一次，后续只复用该变量（否则 `Is` 匹配不到，见 §3.1）。
- **表达“类别”而非细节**：它告诉调用方“是不是 not found / permission denied 这一类”，而不是携带字段。
- **可与错误链共存**：上层用 `%w` 继续补上下文，但不破坏 `errors.Is`。

---

## 1. 为什么需要哨兵

- **稳定判定**：文案调整、包装层增加上下文后，`Error()` 字符串会变，`==` 字符串不可靠。
- **API 合同**：调用方用 `errors.Is` 表达「是否属于这一类失败」（如 `io.EOF`、`sql.ErrNoRows` 风格）。

```go
var ErrNotFound = errors.New("not found")

func Lookup(id string) error {
	if id == "" {
		return ErrNotFound
	}
	return nil
}

// 上层
if errors.Is(err, ErrNotFound) {
	// 404 / 空结果
}
```

---

## 1.1 怎么使用哨兵（最常见套路）

### 步骤 1：定义哨兵（包级只 New 一次）

```go
var ErrNotFound = errors.New("not found")
```

### 步骤 2：在产生错误的地方返回哨兵（不要重复 New）

```go
func find(id string) error {
	if id == "" {
		return ErrNotFound
	}
	return nil
}
```

### 步骤 3：上层加上下文时用 `%w` 保留链

```go
func handler(id string) error {
	if err := find(id); err != nil {
		return fmt.Errorf("handler: id=%s: %w", id, err)
	}
	return nil
}
```

### 步骤 4：调用方用 `errors.Is` 沿链判断

```go
err := handler("")
if errors.Is(err, ErrNotFound) {
	// 例如：HTTP 404 / 返回空结果 / 不告警
}
```

要点：

- **哨兵用于判断分支**（是不是这一类失败）
- **字段/细节**用自定义错误类型 + `errors.As`（见 [04](./04%20-%20自定义错误类型.md)）

---

## 2. 设计习惯

1. **命名**：`Err` 前缀，如 `ErrNotFound`、`ErrPermission`。
2. **导出**：需要给**其他包**判断时再导出；仅包内使用可小写 `errXxx`（较少见）。
3. **文案**：简短、稳定；用户可见文案放在更上层映射，不必与哨兵字符串混为一谈。
4. **全局只 `New` 一次**：哨兵的关键是“**唯一实例**”——只在包级 `var` 创建一次，后续只复用该变量。
5. **包装上传**：中间层用 `fmt.Errorf("...: %w", err)`，保留链，`errors.Is` 仍成立。

---

## 3. 常见坑

| 坑 | 说明 |
|----|------|
| 用 `err == ErrXxx` | 包装后相等性失效，应用 `errors.Is` |
| 同一语义多个 `errors.New` | 两次 `New` 不相等，`Is` 失败；**同一包共用一个 var** |
| 滥用哨兵 | 每种细枝末节都建 `Err` 会导致 API 膨胀；**可恢复分支**可用自定义类型 + `As`（见 [04](./04%20-%20自定义错误类型.md)） |

### 3.1 反例：每次 `errors.New` 都是新对象，`Is` 匹配不到（可运行）

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found") // ✅ 哨兵：只 New 一次

func good() error { return ErrNotFound }

func bad() error { return errors.New("not found") } // ❌ 每次都是新对象

func main() {
	e1 := fmt.Errorf("wrap: %w", good())
	fmt.Println("good:", errors.Is(e1, ErrNotFound)) // true

	e2 := fmt.Errorf("wrap: %w", bad())
	fmt.Println("bad :", errors.Is(e2, ErrNotFound)) // false
}
```

---

## 4. 替代与补充

- **私有错误 + 判断函数**：`func IsNotFound(err error) bool { return errors.Is(err, errNotFound) }`，不导出哨兵，仍可集中演进。
- **需要结构化信息**：哨兵只表达「类别」；字段用结构体错误 + `errors.As`。

---

## 延伸阅读

- [05 - 错误包装与错误链](./05%20-%20错误包装与错误链.md) · [10 - 最佳实践与反模式](./10%20-%20错误处理最佳实践与反模式.md)
