# Chapter 08 — HTTP 协议编程

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 8 章。  
> **HTTP** 是现代互联网与微服务的事实坐标系；Go 的 **`net/http`** 把 **URL、连接池、TLS、服务端并发模型** 绑在一起。本章建立 **Client / Server** 双端心智，并与 **TLS（第 7 章）**、**协议演进（第 5 章）** 对齐。

**性质**：实战章；深度排错见 [`09-http-internals`](../../legacy-topic-index/09-http-internals/)。

---

## 8.1 URL 与资源（`net/url`）

**URL** 不只是「地址」，更是服务侧路由、缓存键、签名字段的一部分。**`url.URL`** 对应 **RFC 3986** 的主要组件：

| 字段 | 含义 |
|------|------|
| **Scheme** | `http` / `https` 等 |
| **Host** | authority（主机 + 可选端口） |
| **Path** | 路径；解析时会对 **path 段**做百分号解码 |
| **RawQuery** | 原始查询串；**`Query()`** 返回 `url.Values`（`map[string][]string`）便于读写 |

**查询参数**：`u.Query()` → `q.Set` / `q.Del` → **`u.RawQuery = q.Encode()`**（注意 Encode 会排序键，利于缓存与签名稳定）。

**国际化域名（IDNA）**：DNS 仍吃 **ASCII**；含非 ASCII 的注册域名在链路上常以 **Punycode（`xn--`）** 出现。需要显式转换时（自定义 SNI、手写 Resolver 等）用 **`golang.org/x/net/idna`**；**`net/http` 拨号**路径上 Go 会按规则处理 Unicode 主机名，但**不要**假设所有工具链行为完全一致——关键路径显式规范化更稳。

**路径与 `path.Clean`**：`path.Clean` 能折叠 `..`、`.`，**不等于**文件系统安全：若要把 URL 路径映射到磁盘，必须在 **`Join` 根目录后**校验结果仍落在根下（例如 `filepath.Clean` + `strings.HasPrefix` 对根前缀的严格比较），否则仍有穿越风险。

```go
// 需 import: net/url, path
u, err := url.Parse("https://api.example.com/api/../v1/search?q=golang&page=1")
if err != nil {
	return
}
q := u.Query()
q.Set("q", "network")
q.Del("page")
u.RawQuery = q.Encode()
u.Path = path.Clean(u.Path)
_ = u.String()
```

---

## 8.2 HTTP 特性与版本演进

**无状态**：每个请求自带语义；水平扩展友好，但 **会话、登录态** 要在应用层用 **Cookie / Session / JWT** 等显式建模。

| 版本 | 要点 | 连接模型 |
|------|------|----------|
| **HTTP/1.0** | 首部、状态码 | 默认短连接 |
| **HTTP/1.1** | **Keep-Alive**、分块、Host | 持久连接；**队头阻塞**在同连接上仍存在 |
| **HTTP/2** | 二进制帧、多路复用、**HPACK** 头压缩 | 单连接多流 |
| **HTTP/3** | **QUIC（UDP）**、流级丢包恢复 | 弱化 TCP 全局队头阻塞；弱网体验常更好 |

**Upgrade**：**`Connection: Upgrade`** 等机制支撑 **WebSocket**（见原书第 15 章与本仓库后续笔记）。  
**HTTP/3**：在 UDP 上由 QUIC 提供加密与多路复用；与「TCP 上全局重传阻塞所有流」的问题解耦——具体栈行为随实现演进，抓包与基准测试比背口号更重要。

---

## 8.3 简单用户代理（`Response` 与 Body）

**`http.Response.Body`** 是 **`io.ReadCloser`**：不读满或不关闭，**连接可能无法回到 `Transport` 空闲池**，在高 QPS 下表现为 fd 泄漏或新建连接暴涨。

**黄金顺序**：判 **`err`** → **`defer resp.Body.Close()`** → 若只关心状态码：仍建议 **`io.Copy(io.Discard, resp.Body)`** 把剩余字节 drain 掉（除非确定已读完或 `ContentLength==0` 等安全情形）。

```go
// 需 import: fmt, io, net/http
respHead, err := http.Head(url)
if err != nil {
	return
}
defer respHead.Body.Close()
fmt.Println("Content-Length:", respHead.ContentLength)

resp, err := http.Get(url)
if err != nil {
	return
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK {
	// 错误响应体也应 drain，便于复用连接
	_, _ = io.Copy(io.Discard, resp.Body)
	return
}
_, _ = io.Copy(io.Discard, resp.Body)
```

**不要用已弃用的 `io/ioutil`**：读全文用 **`io.ReadAll`**，丢弃用 **`io.Copy(io.Discard, …)`**。

---

## 8.4 配置 HTTP 请求（`http.Request`）

**`http.Get` / `Post`** 适合脚本；生产 API 调用应 **`http.NewRequest` / `http.NewRequestWithContext`**，以便：

- **`context.Context`**：超时、取消、链路元数据  
- **Header**：`Authorization`、`Content-Type`、`User-Agent` 等  
- **Body**：`io.Reader`（注意 `ContentLength`、可重放与 `GetBody`）

```go
// 需 import: context, net/http, time
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.example.com/v1/data", nil)
if err != nil {
	return
}
req.Header.Set("Authorization", "Bearer <token>")
req.Header.Set("Content-Type", "application/json")

client := &http.Client{Timeout: 10 * time.Second} // 见 8.5；勿长期用 DefaultClient
resp, err := client.Do(req)
```

---

## 8.5 `http.Client` 与 `http.Transport`（生产级）

**避免裸用 `http.DefaultClient`**：默认 **无总超时**，故障时会挂死大量 goroutine。

**常用旋钮**（按业务调参，下列为示意）：

```go
// 需 import: net, net/http, time
transport := &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
client := &http.Client{
	Transport: transport,
	Timeout:   15 * time.Second, // 含 Body 读取；与 Context 可组合使用，取更严者
}
```

**`Client.Timeout`** 与 **`Request.Context`** 同时存在时，以**先触发者**为准；复杂场景可只保留 Context 并在 Transport 层设各阶段超时。

---

## 8.6 代理（Proxy）

企业网常见 **HTTP(S) 代理**。优先 **`HTTP_PROXY` / `HTTPS_PROXY` / `NO_PROXY`**，代码里用 **`http.ProxyFromEnvironment`** 与运维约定一致。

**显式代理**：`Transport.Proxy = http.ProxyURL(proxyURL)`。若代理需要 **Basic 认证**，推荐把凭据写进 **proxy URL 的 userinfo**（`http://user:pass@host:port`），由 `net/http` 在 **CONNECT** 路径上处理；手写 **`Proxy-Authorization`** 仅在你确认代理实现要求时再考虑。

```go
// 需 import: net/http, net/url
proxyURL, err := url.Parse("http://user:pass@proxy.corp.com:8080")
if err != nil {
	return
}
tr := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
client := &http.Client{Transport: tr}
```

---

## 8.7 客户端 HTTPS 与信任链

HTTPS 依赖 **X.509** 校验（见 [`chapter07`](../../02-general-network-stage/chapter07-security/note.md)）。对接 **私有 CA / 自签名（仅联调）** 时，把根证书装入 **`tls.Config.RootCAs`**，而不是在生产打开 **`InsecureSkipVerify`**。

```go
// 需 import: crypto/tls, crypto/x509, net/http, os
caPEM, err := os.ReadFile("corp-root-ca.pem")
if err != nil {
	return
}
pool := x509.NewCertPool()
if !pool.AppendCertsFromPEM(caPEM) {
	return
}
tr := &http.Transport{
	TLSClientConfig: &tls.Config{
		RootCAs:    pool,
		MinVersion: tls.VersionTLS12,
	},
}
client := &http.Client{Transport: tr}
```

---

## 8.8 服务端（`ServeMux`、静态文件、中间件、HTTPS）

**模型**：`ListenAndServe` 每请求 **一 goroutine**（实现细节可演进，心智上仍是高并发友好）。

**`http.ServeMux`**（Go 1.22+ 支持 **方法路由** 与 **路径变量**；老代码大量 `HandleFunc` 仍兼容）：静态资源常用 **`http.StripPrefix`** + **`http.FileServer`**。

**中间件**：`http.Handler` 包装 `next.ServeHTTP(w, r)`，统一日志、超时、鉴权、trace id 等。

```go
// 需 import: log, net/http
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

mux := http.NewServeMux()
mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))
mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, Go Web!"))
})

// 联调可用 :8443；绑定 :443 在 Unix 上通常要特权或 capability
log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", loggingMiddleware(mux)))
```

---

## 8.9 本章小结

1. **错误优先**：网络失败是常态，`err` 与 **`resp.StatusCode`** 都要处理。  
2. **Body 生命周期**：**`Close` + drain**，否则连接池与健康度遭殃。  
3. **超时分层**：**`Client.Timeout`**、**`Context`**、`Transport` 各阶段超时组合使用。  
4. **TLS**：信任根可配置；**跳过校验仅限受控环境**。  
5. **URL**：查询串用 **`Values.Encode`**；路径映射磁盘要 **防穿越**。

**背诵版**：**URL 规范化 → Request + Context → 专用 Client/Transport → 读透 Body → TLS 信任锚正确。**  

**前后章节**：[`chapter07` TLS](../../02-general-network-stage/chapter07-security/note.md) · [`chapter05` 应用层协议](../../01-foundation-stage/chapter05-application-protocols/note.md) · [`09-http-internals`](../../legacy-topic-index/09-http-internals/)
