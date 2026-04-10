# 07 - Sentinel Error（哨兵错误）设计

**哨兵错误**：包内（或跨包约定）**导出的** `var ErrXxx = errors.New("...")`，用 **`errors.Is(err, ErrXxx)`** 识别，而不是比较 `err.Error()` 字符串。前置：[03](./03%20-%20错误创建-errors.New与fmt.Errorf.md)、[06](./06%20-%20errors.Is与errors.As使用.md)。

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

## 2. 设计习惯

1. **命名**：`Err` 前缀，如 `ErrNotFound`、`ErrPermission`。
2. **导出**：需要给**其他包**判断时再导出；仅包内使用可小写 `errXxx`（较少见）。
3. **文案**：简短、稳定；用户可见文案放在更上层映射，不必与哨兵字符串混为一谈。
4. **包装上传**：中间层用 `fmt.Errorf("...: %w", err)`，保留链，`Is` 仍成立。

---

## 3. 常见坑

| 坑 | 说明 |
|----|------|
| 用 `err == ErrXxx` | 包装后相等性失效，应用 `errors.Is` |
| 同一语义多个 `errors.New` | 两次 `New` 不相等，`Is` 失败；**同一包共用一个 var** |
| 滥用哨兵 | 每种细枝末节都建 `Err` 会导致 API 膨胀；**可恢复分支**可用自定义类型 + `As`（见 [04](./04%20-%20自定义错误类型.md)） |

---

## 4. 替代与补充

- **私有错误 + 判断函数**：`func IsNotFound(err error) bool { return errors.Is(err, errNotFound) }`，不导出哨兵，仍可集中演进。
- **需要结构化信息**：哨兵只表达「类别」；字段用结构体错误 + `errors.As`。

---

## 延伸阅读

- [05 - 错误包装与错误链](./05%20-%20错误包装与错误链.md) · [10 - 最佳实践与反模式](./10%20-%20错误处理最佳实践与反模式.md)
