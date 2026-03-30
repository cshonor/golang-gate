# 05 - errors.Is 与 errors.As 使用

## 1. `errors.Is`

用于判断“这条错误链里是否包含某个目标错误”。

```go
if errors.Is(err, ErrNotFound) { ... }
```

## 2. `errors.As`

用于提取某个错误类型实例。

```go
var ve *ValidationError
if errors.As(err, &ve) { ... }
```

## 3. 记忆法

- Is：判等/判类别
- As：取类型/取细节

