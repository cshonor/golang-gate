# 02 - unsafe.Pointer 指针转换

## 1. 基本规则

- `*T` 可以转 `unsafe.Pointer`
- `unsafe.Pointer` 可以转 `*T`
- 这是“告诉编译器我知道自己在做什么”

## 2. 典型写法

```go
var x int64 = 10
p := &x
up := unsafe.Pointer(p)
p2 := (*int64)(up)
```

## 3. 风险点

- 把 `unsafe.Pointer` 转成错误类型会读错数据
- 指针指向对象生命周期结束后继续用会崩溃

## 4. 实践建议

- 转换链尽量短
- 不要跨层长期保存“转换后的不透明指针”

