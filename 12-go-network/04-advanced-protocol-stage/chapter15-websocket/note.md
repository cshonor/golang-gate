# Chapter 15 — WebSockets 编程实战

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 15 章。  
> **HTTP/1.x 请求-响应**适合文档与 REST；**实时推送、协作、行情**等需要**服务端主动下行**时，常见路径是 **WebSocket（RFC 6455）**：在已有 **TCP** 上先走 **`HTTP Upgrade`**，再切换为**带帧边界的双向消息流**。

**性质**：叠读 [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md)（Upgrade、Server、`TLS`）、[`chapter14` REST](../chapter14-rest/note.md)、[`chapter07` TLS](../../02-general-network-stage/chapter07-security/note.md)（**WSS**）。

---

## 15.1 服务端基础：从 HTTP 到 WS

### 15.1.1 握手与 `101 Switching Protocols`

浏览器先发 **HTTP GET**，携带：

- **`Upgrade: websocket`**、**`Connection: Upgrade`**  
- **`Sec-WebSocket-Key` / `Sec-WebSocket-Version`** 等  

服务端校验后返回 **`101 Switching Protocols`** 与 **`Sec-WebSocket-Accept`**（对 key 的约定变换）。此后**同一 TCP 连接**上改为 **WebSocket 成帧**协议，而不再是「一问一答 HTTP」。

### 15.1.2 并发、内存与僵尸连接

**一连接一 goroutine** 在 Go 里很常见，但 **× 万连接** 时栈与调度开销仍要预算；更关键的是：**读阻塞**与**半开连接**——必须用 **`SetReadDeadline` / `SetWriteDeadline` / `SetDeadline`**（见第 3 章）、**心跳**、**Close** 策略回收，避免 goroutine 与 fd 泄漏。

**轮询 vs WS**：轮询每次整 HTTP 往返，**头字段**相对小 payload 往往过重；WS **帧头**很小（量级 **2～14 字节**基线，视扩展），适合高频小消息（具体仍看业务消息大小与压缩）。

---

## 15.2 `golang.org/x/net/websocket`（原书向）

**生态现状（务必读）**：`golang.org/x/net/websocket` **功能面偏旧**，社区与生产新项目**普遍优先 [`gorilla/websocket`](https://github.com/gorilla/websocket)**（RFC 6455 覆盖、Ping/Pong、`Upgrader`、背压等）。本章保留原书 API 心智：**`Conn` 可当流**、**`Message` / `JSON` / `Codec`**——**新代码请默认 Gorilla**。

**并发写**：**`x/net/websocket.Conn` 非并发安全**；多 goroutine 写需 **单写者**（channel 序列化或 `sync.Mutex`）。示例思路：独立 goroutine 消费 **`writeChan`**，主循环只 **`Receive`**；退出时 **`close(writeChan)`** 并结束写 goroutine，避免泄漏。

**`websocket.JSON`**：大消息要做 **读上限**（`SetReadLimit` 或等价）、**超时**，防 **DoS/OOM**。

**WSS**：**TLS 在 Upgrade 之前**建立；`http.ListenAndServeTLS` + 将 **`websocket.Handler`** 挂到 `http` 上即可得到 **`wss://`**（证书与主机名校验同第 7/8 章）。

**浏览器 `Origin`**：公网服务**必须校验 `Origin`**（允许列表），减轻 **CSWSH** 类风险；**禁止**照抄教学里的 **`CheckOrigin: func(...) bool { return true }`**。

---

## 15.3 `github.com/gorilla/websocket`（推荐落地）

**`websocket.Upgrader`**：控制 **读/写缓冲**、**子协议**、**`CheckOrigin`**（生产写白名单）。

```go
// 生产：按 Host/Origin 白名单校验，勿 return true
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { /* TODO */ return false },
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	c.SetReadLimit(65535)
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			break
		}
		if err := c.WriteMessage(mt, msg); err != nil {
			break
		}
	}
}
```

**写路径**：**`WriteMessage` / `WriteJSON` 同样不应并发裸调**；高并发写用 **`sync.Mutex`** 或**单写 goroutine**。

**心跳**：用 **`Ping`/`Pong`** 或应用层心跳保活，避免中间盒**静默丢连接**；结合 **`ReadDeadline` 刷新**。

**客户端**：`websocket.Dialer` + **`interrupt`/`signal`**；退出路径发送 **`CloseMessage`**（**正常关闭码**）比直接断 TCP 更利于对端与代理观测。

---

## 15.4 工程避坑

1. **`SetReadDeadline`**：每次读到消息后**延长**；无期限 `ReadMessage` 易僵尸化。  
2. **写串行化**：Gorilla 同样要求 **写侧串行**（或 `WriteJSON` 互斥）。  
3. **反向代理**：Nginx 等需转发 **`Upgrade`**、**`Connection`**（及 **`Sec-WebSocket-*`**），否则握手失败。  
4. **Masking**：**客户端→服务端** 帧必须 **mask**（库通常自动处理）；抓包分析时注意与服务端方向不同。

---

## 15.5 小结

**HTTP/2、HTTP/3** 改善了 **HTTP 语义下的多路复用与弱网**；**WebSocket** 仍提供「**升级后近似双向字节流 + 轻量帧**」的产品心智，生态成熟。掌握 **Upgrade → 帧协议 → 心跳/限读/Origin/TLS** 即可覆盖大部分生产排障面。

**背诵版**：**101 升级；读限时；写串行；校验 Origin；公网 WSS；生产用 Gorilla。**

**前后章节**：[`chapter14` REST](../chapter14-rest/note.md) · [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md) · [`chapter07` TLS](../../02-general-network-stage/chapter07-security/note.md) · [`chapter16` Gorilla](../../03-web-core-stage/chapter16-gorilla-toolkit/note.md)
