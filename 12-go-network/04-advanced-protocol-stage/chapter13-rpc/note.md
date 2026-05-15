# Chapter 13 — 远程过程调用（RPC）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 13 章。  
> **RPC** 追求 **访问透明性**：像调本地函数一样调远程过程，但**故障模型、延迟、版本与幂等**仍与本地调用不同（对照第 1 章分布式谬误）。Go 标准库 **`net/rpc`** 以 **gob** 为默认线路编码，适合 **Go 同构** 内网；跨语言场景转向 **JSON-RPC（`net/rpc/jsonrpc`）** 或生产主流的 **gRPC**。

**性质**：与 [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md)、[`chapter04` 序列化](../../03-web-core-stage/chapter04-data-serialization/note.md)、[`chapter05` 协议](../../01-foundation-stage/chapter05-application-protocols/note.md) 叠读。

---

## 13.1 Go 的 `net/rpc` 机制

**服务方法五条硬约束**（违反则 **注册失败** 或运行期不符合预期）：

1. **服务类型导出**（`Register` 时用的 receiver 类型名会成为服务名的一部分）。  
2. **方法导出**。  
3. **两个参数**：均为 **导出类型或内建类型**；第一个是 **请求**（常指针），第二个是 **响应（必须指针）**。  
4. **第二个参数必须是指针**（用于写回结果）。  
5. **单一返回值 `error`**。

**默认编码 `gob`**：**同构 Go** 时紧凑、快；**跨语言**不可用，应换 **JSON-RPC** 或 **gRPC/Protobuf**。

```go
// 需 import: net/rpc
type Args struct{ A, B int }
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
```

---

## 13.2 HTTP RPC 服务端（`rpc.HandleHTTP`）

**工程动机**：走 **HTTP/1.x** 与现有 **防火墙、LB、Ingress** 路径一致，运维友好。

**底层事实（勿与 HTTP CONNECT 混淆）**：`rpc.HandleHTTP()` 在 **`http.DefaultServeMux`** 上注册处理器（默认路径 **`/debug/rpc`**），用 **HTTP POST** 承载 **gob** 编解码的 RPC 帧——**不是**「CONNECT 劫持后脱离 HTTP」那种语义。排障时要查：**路径是否被代理转发**、**Body 是否被缓冲/改写**、**是否仅允许 POST**。

```go
// 需 import: log, net/http, net/rpc
arith := new(Arith)
if err := rpc.Register(arith); err != nil {
	log.Fatal(err)
}
rpc.HandleHTTP()
log.Fatal(http.ListenAndServe(":1234", nil))
```

---

## 13.3 HTTP RPC 客户端（`DialHTTP`）

```go
// 需 import: log, net/rpc
client, err := rpc.DialHTTP("tcp", "localhost:1234")
if err != nil {
	log.Fatal(err)
}
defer client.Close()
var reply int
if err := client.Call("Arith.Multiply", &Args{10, 20}, &reply); err != nil {
	log.Fatal(err)
}
```

**超时**：标准库 **`rpc.Client` 无 `Context`** 集成；生产需在 **连接建立**（自定义 `DialHTTP` 路径较少见）或 **更上层** 做期限控制，或迁移 **gRPC** 等现成方案。至少避免无限阻塞的裸 **`Call`**。

---

## 13.4 TCP RPC 服务端（`rpc.ServeConn`）

去掉 HTTP 头开销，适合 **内网低延迟**（仍要处理 TLS、鉴权、背压等——本章仅传输层心智）。

```go
// 需 import: log, net, net/rpc
if err != nil {
	log.Fatal(err)
}
for {
	conn, err := ln.Accept()
	if err != nil {
		continue
	}
	go rpc.ServeConn(conn)
}
```

---

## 13.5 TCP RPC 客户端（`Dial` 与连接复用）

```go
client, err := rpc.Dial("tcp", "localhost:1234")
if err != nil {
	log.Fatal(err)
}
defer client.Close()
// 同一 client 上多次 Call，复用一条连接
```

**端口与 `TIME_WAIT`**：高频 **Dial → Close** 会堆积 **`TIME_WAIT`**，可能 **耗尽临时端口**；应 **复用 `*rpc.Client`** 或做 **连接池**（或上移协议栈）。

---

## 13.6 值匹配与类型兼容（`gob`）

`gob` 带 **类型描述**，对结构体常见行为包括：**缺字段填零值**、**多余字段可忽略**（利于一定范围内的滚动升级）。**改类型**（如 `int` → `string`）或 **重命名** 仍可能 **不兼容**——把 **DTO 与版本策略**写进评审清单。

---

## 13.7 JSON-RPC 线缆（互操作）

Go 的 **`net/rpc/jsonrpc`** 实现的是较老的 **JSON-RPC 1.0 风格**（`method` / `params` / `id` 等），与 **JSON-RPC 2.0** 生态不完全同一套校验规则；跨语言对接前务必 **对齐规范版本**。

示意（概念，非唯一字段名实现）：

```json
{ "method": "Arith.Multiply", "params": [{ "A": 7, "B": 8 }], "id": 0 }
```

---

## 13.8 JSON-RPC 服务端（`jsonrpc.NewServerCodec`）

```go
// 需 import: log, net, net/rpc, net/rpc/jsonrpc
ln, err := net.Listen("tcp", ":1234")
if err != nil {
	log.Fatal(err)
}
for {
	conn, err := ln.Accept()
	if err != nil {
		continue
	}
	go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
}
```

**权衡**：JSON **可读、跨语言**；通常 **CPU 与体积**劣于 **gob/Protobuf**——数字需 **profile**，勿抄固定百分比教条。

---

## 13.9 JSON-RPC 客户端与 JSON 数值坑

```go
// 需 import: log, net/rpc/jsonrpc
client, err := jsonrpc.Dial("tcp", "localhost:1234")
if err != nil {
	log.Fatal(err)
}
defer client.Close()
```

**大整数 / ID**：JSON **不区分**整型与浮点；若中间层用 **`float64`** 承载，**`int64` 精度**可能受损。对策：**字符串化 ID**、**`json.Number`**、或 **避免 JSON 作为强类型整数通道**。

---

## 13.10 小结与进阶

| 维度 | Gob RPC | JSON-RPC（`jsonrpc`） |
|------|---------|------------------------|
| 互操作 | 基本限 **Go** | **跨语言** |
| 典型代价 | 低（二进制） | 更高 CPU / 更大报文（视负载） |
| 运维 | TCP 或 HTTP 路径 | 同左 |

**避坑**：**超时**、**方法签名五规则**、**Client 复用**、**JSON 数值与版本契约**。

**背诵版**：**同构用 gob；异构用 JSON-RPC 或 gRPC；Call 要有期限；连接要复用。**

**进阶**：工业界大规模系统多选 **gRPC + Protobuf + HTTP/2**（流、元数据、跨语言、生态工具链）。标准库 **`net/rpc`** 更适合 **教材级最小示例** 与历史代码维护。

**前后章节**：[`chapter12` XML](../chapter12-xml-parse/note.md) · [`chapter08` HTTP](../../03-web-core-stage/chapter08-http/note.md) · 第 14 章 REST（[`chapter14-rest`](../chapter14-rest/README.md)）
