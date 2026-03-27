# Go 错误处理：最重要（面试必背版）

## 1. 一句话总纲（背这句）

Go 错误处理的核心是：**把错误当返回值传递，并在“合适的层”补充上下文、做分类处理。**

你写错误处理，永远围绕 3 件事：

1. **返回**：函数签名里把 `error` 返回出去  
2. **包装**：在上层补上下文（告诉你“在哪一步失败”）  
3. **判断**：用 `errors.Is / errors.As` 做分流（重试/降级/返回给用户/打日志）

---

## 2. 最常见写法（标准模板）

```go
func LoadUser(id string) (*User, error) {
    u, err := repo.GetUser(id)
    if err != nil {
        return nil, fmt.Errorf("load user id=%s: %w", id, err) // 包装上下文
    }
    return u, nil
}
```

**为什么要 `%w`**：它把“原始错误”挂到错误链里，后续才能 `errors.Is/As` 判断。

---

## 3. 错误分类（面试常问：怎么设计错误？）

### 3.1 Sentinel error（哨兵错误）+ `errors.Is`

```go
var ErrNotFound = errors.New("not found")

if errors.Is(err, ErrNotFound) {
    // 走 404 / 忽略 / 返回空
}
```

适合：**业务明确的一类错误**（not found、permission denied…）。

### 3.2 自定义错误类型 + `errors.As`

```go
type ValidationError struct {
    Field string
    Msg   string
}
func (e *ValidationError) Error() string { return e.Field + ": " + e.Msg }

var ve *ValidationError
if errors.As(err, &ve) {
    // 根据 ve.Field 做更细分处理
}
```

适合：**需要携带结构化信息**，并在上层做策略判断。

---

## 4. 分层原则（写得像“高级工程师”）

### 4.1 谁负责打日志？

一般建议：**最顶层（入口）统一打日志**，中间层只包装上下文并返回。

原因：中间层打日志会导致 **重复日志**、噪音大、还可能漏关键信息（请求 ID、用户信息）。

### 4.2 上下文要写“动作”，不要只写“失败”

坏例子：

- `return fmt.Errorf("failed: %w", err)`

好例子：

- `return fmt.Errorf("read config %q: %w", path, err)`
- `return fmt.Errorf("query user id=%s: %w", id, err)`

---

## 5. panic / recover（什么时候用？）

### 5.1 结论（面试一句话）

> **能返回 error 就返回 error**；`panic` 用于“程序无法继续”的 bug/不变量破坏；`recover` 常用于框架边界兜底（防止整个进程崩）。

### 5.2 常见可接受场景

- 标准库/框架内部：遇到不可能发生的状态（逻辑 bug）  
- goroutine 的最外层：`defer` + `recover` 兜底，保证服务不中断

---

## 6. 面试速记（背诵）

- 错误链：`fmt.Errorf("...: %w", err)`  
- 判断：`errors.Is`（判等/哨兵），`errors.As`（类型提取）  
- 分层：中间层不乱打日志，只包装上下文；入口统一记录  
- panic：不可恢复的 bug；recover：框架边界兜底  

