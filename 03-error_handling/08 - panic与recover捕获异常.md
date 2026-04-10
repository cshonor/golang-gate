# 08 - panic 与 recover 捕获异常

## 1. 定位

在错误处理体系里，`panic/recover` 是**最后兜底**，不是常规业务分支。`defer` 里如何合并 `Close` 等错误见 [09 - defer 与错误处理协同](./09%20-%20defer与错误处理协同.md)。

## 2. 原则

- 可预期错误：返回 `error`
- 不可恢复错误：`panic`
- 边界兜底：`defer + recover`

## 3. 典型模板

```go
defer func() {
    if r := recover(); r != nil {
        // 记录日志/报警
    }
}()
```

---

## 4. 关键规则（别混）

1. **`recover()` 只有在 `defer` 函数里才会生效**。在普通代码里直接调用 `recover()` 永远拿不到当前 panic。
2. `panic` 会导致当前 goroutine 退栈，退栈过程中会执行该 goroutine 上注册的 `defer`（LIFO）。
3. 是否“吞掉” panic 是策略问题：框架边界常把 panic 转成 **500 + 日志**；库/业务层一般不应吞（否则隐藏 bug）。

---

## 4.1 `error` vs `panic` 判断清单（写代码时对照）

### 直接用 `error`（99% 业务代码）

- **外部输入不可信**：参数校验、HTTP 请求、用户输入、配置内容（可提示修正）。
- **外部依赖可失败**：文件/网络/DB/缓存/调用第三方（超时、NotFound、权限、连接失败）。
- **可预期分支**：业务规则不满足、幂等冲突、资源不存在、重试/降级可处理。
- **库代码 / 可复用组件**：优先 `return error`，把“怎么处理”留给调用方。

### 才考虑 `panic`（少数，且通常在边界）

- **不变量被破坏（bug）**：本不可能发生的状态发生了（内部逻辑错误、越界、空指针等）。
- **启动阶段无法继续**：关键依赖缺失且你决定“宁可启动失败也不要带病运行”（常见在 `main/init` 或启动流程）。
- **框架/边界兜底**：HTTP/RPC/middleware 最外层用 `defer + recover` 防止单次请求把进程带崩；记录日志/指标后返回 500。

### 反向自检（两句就够）

- **用户/环境能导致它发生吗？** 能 → `error`
- **这是程序员写错导致的不可能状态吗？** 是 → `panic`（并在边界 recover）

### recover 放哪

- **只放在边界**（HTTP handler/middleware、goroutine worker 起点、main 入口），不要在业务深处吞 panic。

---

## 5. 可运行示例：`recover` 必须在 `defer` 中

```go
package main

import "fmt"

func main() {
	// 这行不会捕获到下面的 panic
	fmt.Println("recover outside defer:", recover()) // <nil>

	defer func() {
		fmt.Println("recover in defer:", recover()) // "boom"
	}()

	panic("boom")
}
```

---

## 6. 可运行示例：把 panic 转成 error（边界兜底）

适合：在 **goroutine 边界 / 框架边界** 统一兜底，避免整个进程被拖死，同时把“异常路径”统一成 `error` 返回值。

```go
package main

import (
	"fmt"
)

func safeRun(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	fn()
	return nil
}

func main() {
	err := safeRun(func() {
		panic("db pool corrupted")
	})
	fmt.Println(err) // panic: db pool corrupted
}
```

注意：这种做法通常只放在**边界**，业务逻辑仍优先 `return error`。

---

## 6.1 可运行示例：panic + defer + recover 的执行顺序（打印版）

下面用打印顺序把「**panic 发生后发生了什么**」跑一遍：

```go
package main

import "fmt"

func main() {
	fmt.Println("main: start")
	defer fmt.Println("main: defer 1 (runs last)")
	defer func() {
		fmt.Println("main: defer 2 (recover start)")
		if r := recover(); r != nil {
			fmt.Println("main: recovered:", r)
		}
		fmt.Println("main: defer 2 (recover end)")
	}()

	foo()
	fmt.Println("main: after foo (only prints if recovered)")
}

func foo() {
	fmt.Println("foo: enter")
	defer fmt.Println("foo: defer 1")
	defer fmt.Println("foo: defer 2")
	bar()
	fmt.Println("foo: after bar (won't reach)")
}

func bar() {
	fmt.Println("bar: enter")
	defer fmt.Println("bar: defer 1")
	panic("boom")
	// unreachable
}
```

你会观察到：

1. `panic("boom")` 之后，开始退栈；
2. 先执行 `bar` 的 defer，再执行 `foo` 的 defer（**LIFO**）；
3. 最后回到 `main` 的 defer，`recover()` 捕获后，程序继续执行到 `main` 的后续打印。

---

## 7. 可运行示例：HTTP handler 兜底模板（最常见）

```go
package main

import (
	"log"
	"net/http"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic recovered: %v", rec)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/boom", func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	})
	_ = http.ListenAndServe(":8080", RecoverMiddleware(mux))
}
```

实践建议：日志里带 `trace_id` / `request_id`（见 [11 - 错误与日志集成](./11%20-%20错误与日志集成.md) 与 `10-context_pkg` 相关篇）。

---

## 8. goroutine 中的 panic：你以为的“不会影响我”常常是错的

- `panic` 发生在哪个 goroutine，就只会在该 goroutine 上退栈执行 defer。
- **但是**：如果没有任何一层 recover，运行时会打印栈并让整个进程崩溃（这就是为什么边界要兜底）。

常见模式：启动 goroutine 时用一层 `defer + recover` 做“护栏”，并把错误上报（channel / log / metric）。示意：

```go
go func() {
	defer func() {
		if r := recover(); r != nil {
			// TODO: log/metric + 告警 + 触发降级
		}
	}()
	// ... work ...
}()
```

---

## 9. 延伸阅读

- [09 - defer 与错误处理协同](./09%20-%20defer与错误处理协同.md)
- [10 - 错误处理最佳实践与反模式](./10%20-%20错误处理最佳实践与反模式.md)

