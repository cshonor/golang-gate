# 03 - uintptr 与内存地址计算

## 1. 关键区别

- `unsafe.Pointer`：GC 认可的指针语义
- `uintptr`：纯整数，不是“受 GC 追踪”的指针

## 2. 地址计算常见写法

```go
base := unsafe.Pointer(&arr[0])
addr := unsafe.Pointer(uintptr(base) + uintptr(i)*unsafe.Sizeof(arr[0]))
```

## 3. 风险与边界

- 不要把 `uintptr` 长期保存后再转回指针
- 中间若对象移动/释放，地址可能失效

## 4. 面试结论

> `uintptr` 只用于“临时地址计算”，计算完立刻转回 `unsafe.Pointer` 并使用。

