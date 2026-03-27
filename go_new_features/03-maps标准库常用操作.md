# 03 - maps 标准库常用操作

> 对应包：`maps`（Go 标准库）

## 1. 常用能力

- `maps.Clone(m)`：拷贝 map
- `maps.Copy(dst, src)`：把 src 拷到 dst
- `maps.Equal(a, b)`：比较两个 map 是否相等（键值都相等）

## 2. 示例

```go
m1 := map[string]int{"a": 1}
m2 := maps.Clone(m1)
same := maps.Equal(m1, m2) // true
```

## 3. 经验建议

map 是引用类型，直接赋值会共享底层；需要“独立副本”时用 `maps.Clone/Copy`。

