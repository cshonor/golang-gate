# 09 - `defer` 与错误处理协同

`defer` 常用于 **释放资源**（`Close`）、**解锁**、`recover`。本节关注：**defer 里的失败如何与外层 `error` 合并或上报**，避免静默丢错。前置：[02](./02%20-%20错误基本语法.md)、[08](./08%20-%20panic与recover捕获异常.md)。与 **defer 专题**（循环、return 时序、更多示例）对照：[04-defer_traps/07](../04-defer_traps/07-defer%20与错误处理协同.md)。

---

## 1. 典型问题：`Close` 返回 `error`

```go
f, err := os.Open(path)
if err != nil {
	return err
}
defer f.Close() // Close 的错误被忽略
```

若 `Close` 失败（刷盘、网络句柄），调用方完全不知道。

---

## 2. 常见做法

### 2.1 命名返回值 + defer 合并（适合单点返回）

```go
func copyFile(dst, src string) (err error) {
	in, e := os.Open(src)
	if e != nil {
		return e
	}
	defer func() {
		if e := in.Close(); e != nil && err == nil {
			err = e
		}
	}()
	// ... 若中间 err 已非 nil，可按策略选择是否覆盖
	return nil
}
```

要点：**已有业务错误时**，是否用 `Close` 覆盖属于策略（有的团队只 `log` `Close` 错误）。

### 2.2 辅助函数 `defer closeErr(&err, f)`

把「若 `err==nil` 则吸收 `Close`」抽成小函数，减少重复（见 [10](./10%20-%20错误处理最佳实践与反模式.md) 反模式表）。

### 2.3 `errors.Join`（Go 1.20+）

需要**同时保留**多个失败时，可用 `errors.Join(err, closeErr)`，上层可用 `errors.Is` 遍历（可与 `06-go_new_features` 中 `errors.Join` 笔记对照）。

---

## 3. `defer` + `recover`

边界捕获 `panic` 时，常把 `recover` 结果转成 `error` 返回，便于与正常错误路径统一：

```go
defer func() {
	if r := recover(); r != nil {
		err = fmt.Errorf("panic: %v", r)
	}
}()
```

详见 [08](./08%20-%20panic与recover捕获异常.md)。

---

## 延伸阅读

- [10 - 错误处理最佳实践与反模式](./10%20-%20错误处理最佳实践与反模式.md) · [12 - 错误与单元测试](./12%20-%20错误与单元测试.md)
