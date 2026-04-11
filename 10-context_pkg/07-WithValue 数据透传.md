# 07 - WithValue：请求级数据透传【完整版详细笔记】

---

## 1. 函数定义与基本作用

```go
func WithValue(parent Context, key, val any) Context
```

- **作用**：在 `parent` 上绑定一对 **key / value**，返回**新的**子 `Context`（**不修改**父 ctx）。  
- **特点**：  
  - **不可变**：每次 `WithValue` 都是新节点，旧 ctx 上的调用方看不到新挂的数据。  
  - **向下可见**：子链路上 `Value(key)` 可沿父链查找；父 ctx 对象本身**不包含**子节点里后挂的键（查找从**当前** ctx 节点向上走）。  
  - **并发**：多个 goroutine **读**同一 `Context` 是安全的（只读共享）；不要通过 `WithValue` 塞**可变**共享对象当「隐式全局状态」。  

**本质心智模型**：在父链前头**挂一个新节点**，`Value` 查找时**自当前向根**扫描，命中第一个相等的 `key`（`==` 比较）即返回。

---

## 2. 适用场景（严格范围）

只放**请求级、贯穿全链路、非业务核心逻辑**的**小元数据**：

- traceID / spanID（链路追踪）  
- requestID  
- 已通过鉴权的 **userID**、租户 ID 等（仍建议敏感信息最小化）  
- 权限/角色摘要、客户端类型等**日志与审计公共字段**

一句话：

> 跨函数、跨 goroutine、跨中间件传递**「跟这次请求生命周期绑定」的小信息**。

---

## 3. 绝对不能做的事（高频坑 + 线上事故点）

1. **不要塞业务 DTO**：`order`、`query`、大结构体等应走**函数参数**或依赖注入。  
2. **不要塞大对象、大切片、大 buffer**：放大 GC 与内存占用，且生命周期跟请求绑定不清晰。  
3. **不要把 ctx 当全局 map / 服务定位器**：依赖隐晦、测试与重构成本高。  
4. **不要指望「更新」某个 key**：只能再 `WithValue` 包一层得到**新** ctx，并继续**向下**传新 ctx。  
5. **不要替代正常函数参数**：能显式传参就不要藏进 `Value`。

**后果**：可读性差、排查难、内存与耦合失控；可变共享数据还可能带来**数据竞争**（若多处无锁写同一结构体）。

---

## 4. Key 为什么要自定义类型？（核心原理）

`key` 使用 **`==` 可比较** 的类型；若都用裸 `string`，不同包容易撞字符串：

- A 包 `"user_id"`，B 包也 `"user_id"` → 覆盖、串值。

**推荐：包内私有类型 + 非导出常量**

```go
// 自定义 key 类型，避免与其他包的 string key 冲突
type ctxKey int

const (
	_ ctxKey = iota // 占位，避免零值被误用
	traceIDKey
	userIDKey
)
```

写入与读取：

```go
ctx = context.WithValue(ctx, traceIDKey, "trace_123456")

val := ctx.Value(traceIDKey)
traceID, ok := val.(string)
if !ok {
	// 未设置或类型不符
}
```

**优点**：包级隔离、外部无法伪造同类型未导出常量（若 key 类型与常量均未导出）、减少魔法字符串。

也可用 `type ctxKey string` + `const traceIDKey ctxKey = "trace_id"`，但**同一包内**仍要统一常量，避免手写重复字面量。

---

## 5. 取值模板（标准写法）

```go
func GetTraceID(ctx context.Context) (string, bool) {
	val := ctx.Value(traceIDKey)
	if val == nil {
		return "", false
	}
	traceID, ok := val.(string)
	return traceID, ok
}
```

注意：

- `ctx.Value(key)` 可能为 **`nil`**（未设置或父链上无此 key）。  
- 必须**类型断言 + `ok`**，避免强转 panic。  
- 对业务关键路径，**缺省要有明确策略**（降级、拒绝、打日志）。

---

## 6. 底层原理（简单理解）

`Value` 不是一张全局 `map`，实现上多为**父指针 + 当前节点 kv** 组成的**链**（概念上「从子到根的单向链表」）：

```text
当前 ctx 节点 → 父 → 父 → … → 根
```

每次 `WithValue`：

1. 分配新节点，保存 `key`、`val`；  
2. 父指针指向 `parent`；  
3. 查找时从**当前**节点沿父链**线性**扫描，**第一个** `key == 给定 key` 的 `val` 胜出，否则 `nil`。

**特点**：写少读多、不可变、无「删改 key」API；链过长时查找为 **O(深度)**，因此要**少挂、挂小、控层数**。

---

## 7. WithValue 在中间件中的典型用法

```go
func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 示例：可从 Header 继承或生成；此处用简单字符串演示
		traceID := r.Header.Get("X-Request-ID")
		if traceID == "" {
			traceID = "gen-" + strconv.FormatInt(time.Now().UnixNano(), 10)
		}

		ctx := context.WithValue(r.Context(), traceIDKey, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
```

需补充：`import` 含 `context`、`net/http`、`strconv`、`time`，且 `traceIDKey` 与包内定义一致。后续 handler / service / DAO 使用**同一 `r.Context()`（或继续向下传递的 ctx）**即可读到 traceID。更完整链路见 [08-context在中间件中的实战.md](./08-context在中间件中的实战.md)、[10-中间件实战-trace-id透传.md](./10-中间件实战-trace-id透传.md)。

---

## 8. 面试标准答案（完整版）

- `WithValue` 用于在请求全链路**只读透传小元数据**（traceID、requestID 等）。  
- 实现上多为**父链 + 当前节点**；查找沿链向上，**子链可见父链上的值**，新包一层只影响**持有新 ctx** 的下游。  
- **key** 应用**自定义类型**（常为包内私有类型）避免冲突；取值要 **`nil` 检查 + 类型断言**。  
- **禁止**业务大对象、可变共享状态、把 ctx 当 map；**不能**替代函数参数。  
- 滥用会导致依赖不透明与性能问题；**超时与取消**仍应优先用 `WithTimeout` / `WithCancel` 等（见 [05-WithTimeout 超时控制.md](./05-WithTimeout%20超时控制.md)）。

---

## 延伸阅读

- [02-context是什么、作用、核心接口.md](./02-context是什么、作用、核心接口.md)  
- [03-Background()、TODO()、上下文树.md](./03-Background()、TODO()、上下文树.md)  
- [05-WithTimeout 超时控制.md](./05-WithTimeout%20超时控制.md)  
- [06-WithDeadline 精确到时间点的取消控制.md](./06-WithDeadline%20精确到时间点的取消控制.md)  
- [09-context常见陷阱与反模式.md](./09-context常见陷阱与反模式.md)
