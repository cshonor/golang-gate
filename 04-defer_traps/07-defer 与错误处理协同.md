# 07 - defer 与错误处理协同

`defer` 在工程里**绝大多数**与「清理资源 + `error`」绑在一起：`Close`、事务回滚、与 `panic/recover` 同层出现。本章从 **defer 专题** 侧把这条线写清楚，并与 **`03-error_handling`** 打通（`%w`、`errors.Is`/`As`、`errors.Join`）。更系统的「错误处理分章」见 [09 - defer 与错误处理协同](../03-error_handling/09%20-%20defer与错误处理协同.md)。

**前置**：本目录 [01](./01-defer%20基本用法.md)～[05](./05-defer%20在循环中的风险.md)；错误处理侧建议已读 [05 - 错误包装](../03-error_handling/05%20-%20错误包装与错误链.md)、[06 - Is/As](../03-error_handling/06%20-%20errors.Is与errors.As使用.md)。

---

## 1. 典型坑：`defer f.Close()` 丢错

```go
f, err := os.Open(path)
if err != nil {
	return err
}
defer f.Close() // Close 的 error 被丢弃
```

`Close()` 可能失败（刷盘、网络 FS 等），调用方完全不知道。

**改进方向**：用 `defer func()` 接住 `Close` 的返回值，再选择 **打日志**、**写入命名返回 `err`** 或 **`errors.Join`**（见第 3 节）。

---

## 2. 模式 A：至少记录，不向上传

适合：业务已成功，关闭失败只作运维告警。

```go
defer func() {
	if cerr := f.Close(); cerr != nil {
		log.Printf("close failed: %v", cerr)
	}
}()
```

---

## 3. 模式 B：命名返回值，把关闭错误并进 `err`

与 [03-error_handling/09](../03-error_handling/09%20-%20defer与错误处理协同.md) 一致：仅在 **当前 `err == nil`** 时用 `Close` 覆盖，避免掩盖业务错误；需要保留两类失败时用 **`errors.Join`**（Go 1.20+）。

```go
func work(path string) (err error) {
	f, e := os.Open(path)
	if e != nil {
		return fmt.Errorf("open: %w", e)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close: %w", cerr)
		}
	}()
	// …
	return nil
}
```

链式判定仍可用 `errors.Is` / `As`（见 [06](../03-error_handling/06%20-%20errors.Is与errors.As使用.md)）。

---

## 4. 模式 C：`defer` + `recover` 与错误统一

在**边界**把 `panic` 转成 `error` 返回时，常与 `defer` 同层；recover **必须**写在 defer 里（详见本目录 [04](./04-defer%20与%20return%20执行流程.md) 与 [03-error_handling/08](../03-error_handling/08%20-%20panic与recover捕获异常.md)）。

```go
defer func() {
	if r := recover(); r != nil {
		err = fmt.Errorf("panic: %v", r)
	}
}()
```

注意：若同时有 **事务回滚**，要在 defer 里理清顺序（先 rollback，再决定是否重抛 panic），避免半提交状态。

---

## 5. 事务示意（命名 `err` + defer）

```go
func doTx(ctx context.Context, db *sql.DB) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	// SQL…
	return nil
}
```

---

## 6. 与本目录其它章的关系

| 话题 | 见 |
|------|-----|
| return 与 defer 时序、命名返回修改返回值 | [04](./04-defer%20与%20return%20执行流程.md) |
| 循环里 defer 堆积 | [05](./05-defer%20在循环中的风险.md) |
| 工程化清单与反模式总表 | [08](./08-defer%20最佳实践与反模式.md) |

---

## 延伸阅读

- [03-error_handling/09](../03-error_handling/09%20-%20defer与错误处理协同.md) · [03-error_handling/10](../03-error_handling/10%20-%20错误处理最佳实践与反模式.md)
