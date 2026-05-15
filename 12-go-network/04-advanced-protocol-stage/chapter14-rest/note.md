# Chapter 14 — REST 架构风格与 Go 实现

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 14 章。  
> **REST** 不是「又一个 CRUD 框架」，而是一套**资源—表述—统一接口**的约束集合；与 **RPC 的动作动词心智**（见 [`chapter13`](../chapter13-rpc/note.md)）对照，能更清楚何时用 **HTTP 语义** 表达业务。

**性质**：叠读 [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md)、[`chapter05` 协议](../../01-foundation-stage/chapter05-application-protocols/note.md)、[`chapter10` 闪卡整站](../../03-web-core-stage/chapter10-complete-web-server/note.md)。

---

## 14.1 URI 与资源

**资源**：领域实体（用户、闪卡、任务…）；**URI** 是其稳定标识，**不是**过程名。

- **标识与存储解耦**：`/cards/101` 不必泄露表名、主库或微服务拓扑 → **位置透明**。  
- **命名习惯**：**名词、复数集合**（`/cards`）、**层级从属**（`/users/123/cards`）；避免 **`/getCards`**、避免 **`.php` / `.json` 技术后缀**（表述由 **`Accept` / `Content-Type`** 协商，而非路径扩展名）。

**非 CRUD 操作**：优先**资源化**——例如 **`POST /cards/123/resets`**（创建一次「重置」动作记录），优于 **`?action=reset`** 这类 RPC 风格 query。

---

## 14.2 资源表述与内容协商

同一资源可有 **JSON / XML / HTML** 等多种 **Representation**；Go 里常用 **`encoding/json`**、**`encoding/xml`** 与结构体标签。

```go
type Card struct {
	ID       int    `json:"id" xml:"id"`
	Question string `json:"question" xml:"question"`
	Answer   string `json:"answer" xml:"answer"`
}
```

**`Accept` / `Content-Type`**：客户端用 **`Accept`** 表达偏好（可带 **`q`** 权重）；服务端选择**最佳表述**并 **`Content-Type`** 回应。生产级解析 **`Accept`** 宜用 **`mime.ParseMediaType`** 等对 **`q`** 排序的实现，**避免**仅靠 **`strings.Contains`** 误判优先级。

```go
accept := r.Header.Get("Accept")
if strings.Contains(accept, "application/xml") {
	w.Header().Set("Content-Type", "application/xml")
	_ = xml.NewEncoder(w).Encode(card)
	return
}
w.Header().Set("Content-Type", "application/json")
_ = json.NewEncoder(w).Encode(card)
```

---

## 14.3 REST 谓词（HTTP 方法语义）

| 方法 | 典型语义 | 幂等 / 安全 |
|------|-----------|-------------|
| **GET** | 读 | **安全**、**幂等**；可缓存（**ETag** / **Last-Modified**） |
| **PUT** | **整资源替换**（不存在可 **201 创建**） | **幂等**；勿当 **PATCH** 用 |
| **PATCH** | 部分更新 | 是否幂等依语义设计 |
| **DELETE** | 删除 | 常设计为 **幂等**；工程上常见 **软删**（`deleted_at`） |
| **POST** | **集合下创建**或非幂等动作 | **非幂等**；成功创建常 **`201` + `Location`** |

**POST 创建示例**：

```go
defer r.Body.Close()
var c Card
if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
	http.Error(w, "Invalid payload", http.StatusBadRequest)
	return
}
id := h.Store.Save(c)
w.Header().Set("Location", fmt.Sprintf("/cards/%d", id))
w.WriteHeader(http.StatusCreated)
_ = json.NewEncoder(w).Encode(map[string]int{"id": id})
```

---

## 14.4 无状态（Statelessness）

每个请求应携带**理解自身所需的上下文**（认证、trace、分页游标等），服务端**不依赖特定节点上的会话内存** → 横向扩展与故障切换更简单。

**JWT 等**：把**可验证的声明**放在令牌里，服务端无共享会话亦可鉴权（注意**吊销、轮换、密钥与泄露面**）。

---

## 14.5 HATEOAS（Richardson Level 3）

响应中嵌入**链接**（如 **HAL** 风格 **`_links.self`**），客户端**跟链**而非硬编码路径 → API 演进时更稳。

**成本**：序列化体积、设计复杂度上升；多数内部 API **Level 2**（正确动词 + 状态码）已是高性价比。

---

## 14.6 REST 中的「事务」

HTTP 无跨请求事务。常见模式：**把事务做成资源**——`POST /transactions` 得 **`tx_id`**，多步 **`PATCH`/`PUT`…?tx_id=…`**，最后 **`POST /transactions/{id}/commit`**（或 **`DELETE` 取消**）——把一致性边界显式化。

---

## 14.7 理查德森成熟度模型（简要）

| Level | 要点 |
|-------|------|
| 0 | 把 HTTP 当隧道（如 POX/SOAP 风格） |
| 1 | **多 URI** 区分资源 |
| 2 | **动词 + 状态码** 用对 |
| 3 | **HATEOAS** 超媒体驱动 |

---

## 14.8 实战：闪卡 REST 化（相对第 10 章）

**目标路由（示例）**

- **`GET /cards`**：列表  
- **`POST /cards`**：创建（**`201` + `Location`**）  
- **`GET /cards/{id}`**：单条  

**教材级 `ServeMux` 注意**：单函数里用 **`strings.Split`** 时，要区分 **`GET /cards` 与 `GET /cards/1`**，并对 **`strconv.Atoi`** **检查 `error`**；**`405`** 时补 **`Allow`**。原书片段若只处理「带 ID 的 GET」、忽略列表，应在笔记中**显式补全**或改用 **Go 1.22+ `ServeMux` 模式路由** / **`gorilla/mux`**。

**客户端**：**`http.Client` 设超时**、检查 **`NewRequest` / `Do` / `ReadAll`** 的错误，并 **`io.Copy(io.Discard, resp.Body)`** 利于连接复用（见第 8 章）。

---

## 14.9 REST 与 RPC 的权衡

| 维度 | REST（HTTP + 资源） | RPC（gRPC / JSON-RPC 等） |
|------|----------------------|---------------------------|
| 心智 | **名词 + 统一接口** | **过程 + 契约（IDL）** |
| 缓存 / 生态 | **HTTP 缓存、CDN、调试工具**成熟 | 二进制 RPC **默认不经 CDN 缓存** |
| 契约 | OpenAPI 等 | **Protobuf / IDL** 强类型 |
| 典型场景 | **公网 API**、浏览器友好 | **同构/高性能内部调用** |

---

## 14.10 本章小结

1. **URI 表资源**，表述靠 **`Accept`/`Content-Type`**。  
2. **动词与状态码**表达语义；**POST 创建**用 **`201` + `Location`**。  
3. **无状态**支撑扩展；**HATEOAS** 按需。  
4. **分布式事务**用**资源化事务**等显式模型。  
5. **与 RPC 分工**：对外 HTTP 语义 vs 对内高性能过程调用。

**背诵版**：**资源名词化；动词表语义；201+Location；无状态；Level 2 为默认基线。**  

**下一章**：实时与全双工 → [`chapter15` WebSocket](../chapter15-websocket/note.md)

**前后章节**：[`chapter13` RPC](../chapter13-rpc/note.md) · [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md) · [`chapter10` 整站](../../03-web-core-stage/chapter10-complete-web-server/note.md) · [`chapter15` WebSocket](../chapter15-websocket/note.md)
