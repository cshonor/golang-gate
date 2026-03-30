# 02 - slices 标准库常用操作

> 对应包：`slices`（Go 标准库）

## 1. 常用能力清单

- 排序：`slices.Sort` / `slices.SortFunc`
- 查找：`slices.Index` / `slices.Contains`
- 比较：`slices.Equal`
- 删除：`slices.Delete`
- 插入：`slices.Insert`
- 克隆：`slices.Clone`

## 2. 示例（记住“可读性提升”）

```go
nums := []int{3, 1, 2}
slices.Sort(nums) // [1,2,3]

ok := slices.Contains(nums, 2) // true
```

## 3. 面试表述（工程价值）

> `slices` 把常见 slice 操作标准化，减少手写 for 循环与边界 bug，同时让代码更短更清晰。

