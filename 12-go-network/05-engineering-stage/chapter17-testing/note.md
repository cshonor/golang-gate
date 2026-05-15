# Chapter 17 — 测试（网络与 HTTP）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 17 章。  
> 网络代码的测试是在**可控环境**里主动对抗「**分布式八大谬误**」（见 [`chapter01`](../../01-foundation-stage/chapter01-architectural-layers/note.md)）：**超时、断连、慢客户端、端口争用**必须可复现，否则 CI 里只会得到 **Flaky** 与误报。

**性质**：与 [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md)、[`chapter03` Socket](../../01-foundation-stage/chapter03-socket-programming/note.md) 叠读。

---

## 17.1 简单但脆弱的反模式

| 问题 | 后果 |
|------|------|
| **硬编码端口**（如 `:8080`） | CI **并行包**冲突、随机失败 |
| **依赖本机已起服务** | 不自包含、新人/沙箱无法跑 |
| **无 `Deadline`** | 挂死流水线 |
| **`t.Fatal` 在子 goroutine** | 违反 `testing` 约定，行为未定义 |
| **不回填 `defer conn.Close()` / `resp.Body.Close()`** | **fd** 或连接池泄漏 |

**回环地址**：`127.0.0.1` 与 **`[::1]`** 在部分环境策略不同；**`:0` 监听**后应以 **`Listener.Addr()`** 或 **`httptest.Server.URL`** 注入真实地址，避免写死。

```go
// 反面教材：固定端口 + 假定外部已监听
func TestDialFixed_Bad(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
}
```

---

## 17.2 `net/http/httptest`

**`httptest.ResponseRecorder`**：实现 **`http.ResponseWriter`**，把响应写入**内存**——适合 **Handler 单元测试**，无真实 socket。

```go
h := func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	_, _ = fmt.Fprint(w, "queued")
}
req := httptest.NewRequest(http.MethodPost, "/queue", nil)
rec := httptest.NewRecorder()
h(rec, req)
if rec.Code != http.StatusAccepted {
	t.Fatalf("code=%d", rec.Code)
}
```

**`httptest.NewServer` / `NewUnstartedServer`**：真实 **`net.Listener`**，端口 **`0` 动态分配**；客户端用 **`ts.URL`** 注入，**`defer ts.Close()`**；读完 **`Body`** 或 **`io.Discard`** 以便连接复用（同第 8 章）。

---

## 17.3 HTTP 之下：`net.Pipe` 与故障注入

**`net.Pipe()`**：一对内存 **`net.Conn`**，适合测**编解码、帧协议**，不经内核协议栈。

```go
client, server := net.Pipe()
defer client.Close()
defer server.Close()

	go func() {
		buf := make([]byte, 4)
		_, _ = io.ReadFull(server, buf)
		if string(buf) == "PING" {
			_, _ = server.Write([]byte("PONG"))
		}
	}()

	_ = client.SetDeadline(time.Now().Add(time.Second))
	_, _ = client.Write([]byte("PING"))
	got := make([]byte, 4)
	if err := io.ReadFull(client, got); err != nil {
		t.Fatal(err)
	}
	if string(got) != "PONG" {
		t.Fatalf("got %q", got)
	}
```

**注意**：**不要在非测试 goroutine 里调用 `t.Error/Fatal`**；上例服务端 goroutine只做 I/O，断言留在**测试 goroutine**。

**自定义 `Listen`**：`net.Listen("tcp", "127.0.0.1:0")` 取空闲端口；IPv6-only 环境可评估 **`[::1]:0`** 或双栈策略。

**故障注入**：对依赖 **`net.Conn` / `io.Reader`** 的代码，用**小包装类型**覆写 **`Read`/`Write`** 返回 **`os.ErrDeadlineExceeded`**、`io.ErrUnexpectedEOF` 等，验证错误路径（包装需满足接口其余方法，可用**嵌入真实 `Conn` 再覆写单方法**）。

---

## 17.4 标准库测试习惯

- **`t.Helper()`**：封装 `assertXxx` 时标记，失败栈指向**调用方**行号。  
- **`t.Parallel()`**：缩短 I/O 密集套件总时间；与 **表驱动子测** 并用时务必 **`tt := tt`**（或把 `tt` 作为参数传入子闭包），避免**经典循环变量捕获 bug**。  
- **`TestMain`**：包级一次初始化（全局 fake、环境探测）。

**表驱动 + `httptest.Server` + 并行**（注意 **`tt := tt`**）：

```go
ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/error" {
		http.Error(w, "boom", http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(w, "OK")
}))
defer ts.Close()

tests := []struct{ name, path string; code int }{
	{"ok", "/", http.StatusOK},
	{"err", "/error", http.StatusInternalServerError},
}
for _, tt := range tests {
	tt := tt
	t.Run(tt.name, func(t *testing.T) {
		t.Parallel()
		resp, err := http.Get(ts.URL + tt.path)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != tt.code {
			t.Errorf("want %d got %d", tt.code, resp.StatusCode)
		}
	})
}
```

---

## 17.5 本章小结

1. **禁硬编码端口**；优先 **`httptest`** 或 **`Listen(..., ":0")` + `Addr()`**。  
2. **分层选工具**：HTTP 用 **`httptest`**；字节协议用 **`net.Pipe`** 或内存 **`Listener`**。  
3. **注入失败**：超时、半包、EOF、RST——证明错误路径**可观测**而非 panic。  
4. **并行子测**：**捕获循环变量**；**`t` 只在测试 goroutine 用**。

**背诵版**：**动态端口；httptest 隔离 HTTP；Pipe 隔离字节流；并行要 tt:=tt；失败要可注入。**

**前后章节**：[`chapter01` 架构与谬误](../../01-foundation-stage/chapter01-architectural-layers/note.md) · [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md) · [`chapter03` Socket](../../01-foundation-stage/chapter03-socket-programming/note.md) · 附录 A/B 模糊测试与泛型（见五阶段路线）
