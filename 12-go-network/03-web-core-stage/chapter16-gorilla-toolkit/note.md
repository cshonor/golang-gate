# Chapter 16 — Gorilla Web 工具包

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 16 章。  
> **Gorilla** 不是「全家桶框架」，而是一组与 **`net/http`** 对齐、可组合的库：**路由（`mux`）**、**横切中间件（`handlers`）**、**表单解码（`schema`）**、**安全 Cookie**、轻量 **JSON-RPC（`rpc`）** 等——在标准库 **`ServeMux`** 不够用时，企业里很常见。

**性质**：叠读 [`chapter08` HTTP](../chapter08-http/note.md)、[`chapter14` REST](../../04-advanced-protocol-stage/chapter14-rest/note.md)；与 [`chapter15` WebSocket](../../04-advanced-protocol-stage/chapter15-websocket/note.md) 同属「Web 栈扩展」。

---

## 16.1 中间件模式（`http.Handler` 洋葱模型）

**横切关注点**：日志、鉴权、限流、trace、**panic 恢复**——用 **`func(http.Handler) http.Handler`** 包装，保证链上**始终调用** **`next.ServeHTTP(w, r)`**（除非短路返回 4xx/5xx）。

**顺序**：**鉴权 / 限流**尽量靠近业务前；**Recovery、访问日志**常放**最外层**，以便 panic 也能被记录（具体顺序依安全与审计要求）。

```go
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
// 外层先执行：Logging(Auth(business))
```

**易错**：中间件里**忘记** `next.ServeHTTP` → 客户端**挂起**。

---

## 16.2 标准库 `ServeMux`

**优势**：零依赖、开销低，适合 **health、webhook、静态资源**。

**局限（历史痛点）**：老版本对 **路径参数**、**方法路由**表达弱，Handler 内易充斥 **`strings.Split`** 与 **`if r.Method`**。

**Go 1.22+**：**`http.ServeMux`** 支持 **方法前缀**（`"GET /items/{id}"`）等（以当前 Go 文档为准），部分场景可**延后**引入 `mux`。

---

## 16.3 自定义多路复用器（按 Host 等分发）

```go
type HostSwitch struct{ M map[string]http.Handler }

func (h HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ha, ok := h.M[r.Host]; ok {
		ha.ServeHTTP(w, r)
		return
	}
	http.Error(w, "unknown host", http.StatusForbidden)
}
```

**多租户**：按 **`Host`**、**`X-Tenant-ID`** 分发到不同子 `Handler`，在网关后很常见。

---

## 16.4 `gorilla/mux`

- **路径变量**：`/articles/{category}/{id:[0-9]+}`、`mux.Vars(r)`。  
- **方法 / Scheme / Header / Query** 链式约束：`.Methods("GET").Schemes("https")`。  
- **`StrictSlash(true)`**：减少 **`/api` vs `/api/`** 404 困扰。  
- **代价**：功能强于默认 `ServeMux` 时，**匹配成本**可能更高——热点路径要 **benchmark**；极简单场景不必上 `mux`。

---

## 16.5 为何关心 Gorilla（与标准库对照）

| 维度 | 纯 `net/http` | Gorilla 常见组合 |
|------|----------------|------------------|
| REST 路径参数 | 手写解析多 | **`mux` 声明式** |
| 中间件 | 手写或自研链 | **`handlers`** 现成件 |
| 表单 → 结构体 | 手写 `ParseForm` | **`schema`** |
| 签名 Cookie | 易踩坑 | **`securecookie`** |

---

## 16.6 `gorilla/handlers`

**`LoggingHandler`**：Apache **Common / Combined** 格式，对接 **ELK** 等。

**`RecoveryHandler`**（API 以模块文档为准）：**接收 `http.Handler`、返回包装后的 `http.Handler`**——**不要**写成 **`RecoveryHandler()(h)`** 这类双调用（除非你所用版本明确提供**选项构造器**再返回 `func(http.Handler) http.Handler`，以 **pkg.go.dev 为准**）。

```go
// 典型顺序：Recovery 包最外，其次 Logging，再到路由
r := mux.NewRouter()
r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte("ok")) })

logged := handlers.LoggingHandler(os.Stdout, r)
safe := handlers.RecoveryHandler(logged)
log.Fatal(http.ListenAndServe(":8080", safe))
```

另：**`ProxyHeaders`** 在反代后修正 **`RemoteAddr` / `Scheme`** 等，生产Behind Nginx 时常用。

---

## 16.7 其它 Gorilla 组件（鸟瞰）

除本章详述外，生态还有 **会话、CSRF** 等（按项目**按需**引入，避免「为了 Gorilla 而 Gorilla」）。

---

## 16.8 `gorilla/rpc`（JSON-RPC over HTTP）

轻量 **HTTP + JSON-RPC** 管理面协议；与 **gRPC** 相比：**无 `protoc` 流水线**、**浏览器友好**，但 **性能 / 强类型契约**通常不如 **Protobuf + gRPC**。

服务端心智（包名以 `github.com/gorilla/rpc` / `github.com/gorilla/rpc/json` 为准）：**`rpc.NewServer()`** → **`RegisterCodec(json.NewCodec(), "application/json")`** → **`RegisterService`** → **`http.Handle("/rpc", s)`**。

---

## 16.9 `gorilla/schema`

把 **`url.Values`**（`PostForm` / `Query`）**解码到 struct**，靠 **`schema:"name"`** 标签；字段需**导出**，标签拼写错误会导致**静默不填**——重要字段解码后做**校验**。

```go
if err := r.ParseForm(); err != nil { /* 400 */ }
var u struct {
	Name string `schema:"full_name"`
	Age  int    `schema:"age"`
}
if err := schema.NewDecoder().Decode(&u, r.PostForm); err != nil {
	http.Error(w, "bad form", http.StatusBadRequest)
	return
}
```

---

## 16.10 `gorilla/securecookie`

**HMAC 签名**防篡改；可选 **block key** 做**加密**，防读取。

**密钥**：**持久化**（环境变量 / KMS / 配置文件），**禁止**每次进程启动 **`GenerateRandomKey`** 否则全站会话失效；**禁止**把示例密钥复制进生产。

**4KB 限制**：Cookie 有大小上限；**加密+Base64** 膨胀明显——**大对象放服务端**（Redis/DB），Cookie 只存 **session id** 或短引用。

---

## 16.11 本章小结

1. **中间件**：`Handler` 包装 + **明确顺序** + **勿断链**。  
2. **路由**：标准库够用则轻；复杂 REST 再 **`mux`**。  
3. **生产件**：**`handlers`**（日志、Recovery、代理头）、**`schema`**、**`securecookie`** 按域引入。  
4. **RPC 子集**：`gorilla/rpc` 适合**小管理接口**；高性能内网优先 **gRPC**。

**背诵版**：**洋葱中间件；mux 声明路由；handlers 补生产面；securecookie 管密钥与体积。**

**前后章节**：[`chapter08` HTTP](../chapter08-http/note.md) · [`chapter14` REST](../../04-advanced-protocol-stage/chapter14-rest/note.md) · [`chapter15` WebSocket](../../04-advanced-protocol-stage/chapter15-websocket/note.md)
