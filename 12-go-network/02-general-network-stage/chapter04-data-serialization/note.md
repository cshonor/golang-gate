# Chapter 04 — Data Serialization（数据序列化）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 4 章。  
> 网络传输的底层是**无业务语义的原始字节**；序列化/反序列化把「线路格式」与 **Go 里的结构化对象**对齐，是从传输层走进应用逻辑的关键一步。

**Legacy 深度笔记**（粘包、定长帧、Protobuf 等）：[`../../legacy-topic-index/08-framing-protocols/`](../../legacy-topic-index/08-framing-protocols/)

---

## 4.1 结构化数据（Structured Data）

**战略意义**：字节流必须经过**结构化**才能变成程序可操作的模型；Go 中最常见载体是 **`struct`**。

- **内存布局**：运行时的对象布局、对齐、指针与切片头，与「线上一串字节」不是同一回事。  
- **线路格式（Wire format）**：跨网络必须变成**连续、双方可解释**的字节序列（或文本流）。

**`net.IP` 心智**：`net.ParseIP("127.0.0.1")` 把文本变成 **4 或 16 字节**的语义化 IP 表示——正是「模糊输入 → 精确结构」的一类工程实践。

**类比**：序列化像**打包快递**——字段顺序、类型、版本若未约定，接收方就无法可靠拆包（反序列化）。

### 新手易错：字段可见性（Exported fields）

`encoding/json`、`encoding/asn1` 等依赖反射访问**导出字段**（首字母大写）。首字母小写的字段**不会**出现在 JSON/ASN.1 输出里，初学者常因此得到「几乎是空对象」的序列化结果。

---

## 4.2 共同协议（Mutual Agreement）

无自描述元数据时，双方必须**预先约定**布局。

- **固定格式协议**：字段偏移与长度写死 → 解析极快，但**演进痛苦**；务必考虑 **版本号** 与兼容策略。  
- **字节序（Endianness）**：多字节整数在大端/小端机器上不同；跨平台二进制应使用 **`encoding/binary`** 的 `BigEndian` / `LittleEndian` **显式**读写。

**实践**：在协议 **Header** 预留 **Version** 字段，让接收端按版本分支解析——往往是后期救命的扩展点。

---

## 4.3 自描述数据（Self-Describing Data）

如 **JSON**、**XML**：载荷中带字段名等元数据，接收端即使结构体滞后，也常能「多读少错」地解析。

- **弱契约 / 演进**：较易做**增字段不破坏旧客户端**（旧端忽略未知字段）。  
- **代价**：体积与解析 CPU；JSON 重复传输键名（如 `"price":100`），大流量下明显。

---

## 4.4 编码包概览（`encoding/*` 与 `io`）

Go 标准库在 **`encoding/`** 下提供多种线路格式的 **Marshal/Unmarshal**（或 Encoder/Decoder）模型，并与 **`io.Reader` / `io.Writer`** 深度集成。

- **流式 JSON**：`json.NewEncoder(conn).Encode(v)` 把对象直接写入连接，**减少**「先整块 `[]byte` 再 Write」的中间缓冲（不等于全链路零拷贝；仍可能有内部缓冲与分配）。  
- **常见子包**：`json`（Web/API）、`gob`（Go↔Go）、`asn1`（证书、LDAP 等）、`binary`（定长二进制）、`xml` 等。

---

## 4.5 ASN.1 与 DER

**ASN.1** 在电信、安全（**X.509**、LDAP）中地位高。Go 的 **`encoding/asn1`** 常配合 **DER**（Distinguished Encoding Rules）——**同值同 DER 字节唯一**，对**数字签名**友好。

- 依赖反射较多，超大吞吐场景要实测。  
- **可选字段**需在结构体标签中标注 **`asn1:"optional"`** 等；字符串等类型支持与 Protobuf/JSON 不同，需查文档。

完整可运行示例见下文 **附录 A**（使用 `127.0.0.1:0` 动态端口，避免硬编码端口冲突）。

---

## 4.6 JSON

事实标准：**`encoding/json`** + **struct tag**（`json:"name,omitempty"` 等）。

- **`omitempty`**：零值省略，减小体积。  
- **极致性能**：反射可能成为热点；生态里有 **jsoniter**、**easyjson**（代码生成）等第三方方案——选型需评估依赖、安全更新与团队规范。

**流式示例要点**：`json.NewDecoder(conn).Decode(&v)` / `NewEncoder(conn).Encode(&v)`；注意 **Accept 与 Dial 的时序**（测试中常用 `net.Listener.Addr()` 或 `sync.WaitGroup`）。

---

## 4.7 Gob（Go 原生二进制）

**`encoding/gob`**：面向 **Go 进程之间**的高效流；流首会传输**类型描述**，支持一定形式的类型演进（新字段可被旧接收端忽略等语义需在协议层再确认）。

**硬限制**：**不能**作为与 Python/Java 互通的格式。

**`interface{}` / `any` 字段**：对端 `Decode` 前需 **`gob.Register(具体类型)`**，且**收发双方**对注册类型集合要一致，否则解码失败。

---

## 4.8 二进制 → 文本：Base64

在 **HTTP 头**、**SMTP** 等偏文本的载体里传密钥/图片等二进制时，常用 **Base64**。

- 体积约 **+33%**。  
- **URL 场景**：使用 **`base64.URLEncoding`**（`+` `/` 替换为 `-` `_`），避免与 URL 保留字符冲突。  
- **大对象**：用 **`base64.NewEncoder(enc *Encoding, w io.Writer)`** 对流写入（第一个参数是 `base64.StdEncoding` 或 `base64.URLEncoding`，第二个是 `io.Writer`），避免一次性分配巨大 `[]byte`。

---

## 4.9 Protocol Buffers（Protobuf）

跨语言、高性能线路格式的工业界主流之一。

**典型流程**

1. 安装 **`protoc`**（协议编译器）。  
2. 安装 Go 插件，例如：  
   `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`  
3. 编写 **`.proto`**（IDL），执行：  
   `protoc --go_out=. path/to/*.proto`  
4. 在业务代码中使用生成的 `*.pb.go`（**字段号 field number 不可随意改**，关乎兼容性与 wire 格式）。

体积常显著小于 JSON（量级上常见 **数倍** 差异，依消息形状而定），但引入 **IDL 与生成物** 的维护成本。

---

## 4.10 本章小结：选型矩阵

| 场景 | 推荐 | 优势 | 劣势 | 演进/契约 |
|------|------|------|------|-----------|
| 通用 Web / 开放 API | **JSON** | 可读、生态极大 | 体积与解析成本 | 较灵活（增字段常可兼容） |
| 纯 Go 高性能进程间 / RPC | **Gob** | 原生、少配置 | **不跨语言** | 类型流，演进需小心约定 |
| 大规模微服务、多语言 | **Protobuf** | 紧凑、跨语言、工具链成熟 | 需 `protoc` 与生成代码 | **字段号** 纪律是关键 |
| 证书、LDAP、电信系 | **ASN.1 / DER** | 标准强、签名友好 | 复杂、学习曲线陡 | 演进成本高 |
| 文本协议里夹二进制 | **Base64** | 纯文本兼容 | +33% 体积 | 非结构化协议本身 |

**一句话**：在**性能压力**、**跨语言**、**可读与调试**、**Schema 治理成本**之间做权衡；没有银弹。

---

## 附录 A：ASN.1 Daytime 可运行最小例（动态端口）

```go
package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"time"
)

type Daytime struct {
	Time time.Time
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	addr := ln.Addr().String()
	done := make(chan struct{})

	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		b, err := asn1.Marshal(Daytime{Time: time.Now()})
		if err != nil {
			return
		}
		_, _ = conn.Write(b)
	}()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	var d Daytime
	if _, err := asn1.Unmarshal(buf[:n], &d); err != nil {
		panic(err)
	}
	fmt.Println("ASN.1:", d.Time.UTC().Format(time.RFC3339))
	<-done
}
```

---

## 附录 B：JSON 流式收发（动态端口）

```go
package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Config struct {
	Host string   `json:"host"`
	Port int      `json:"port"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	addr := ln.Addr().String()
	done := make(chan struct{})

	go func() {
		defer close(done)
		c, err := ln.Accept()
		if err != nil {
			return
		}
		defer c.Close()
		var cfg Config
		if err := json.NewDecoder(c).Decode(&cfg); err != nil {
			return
		}
		fmt.Printf("JSON: %+v\n", cfg)
	}()

	client, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	_ = json.NewEncoder(client).Encode(Config{Host: "127.0.0.1", Port: 8080})
	<-done
}
```

---

## 附录 C：Gob + `interface{}`（双方 Register）

```go
package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Payload struct {
	Data any
}

func main() {
	gob.Register(map[string]int{})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	addr := ln.Addr().String()
	done := make(chan struct{})

	go func() {
		defer close(done)
		c, err := ln.Accept()
		if err != nil {
			return
		}
		defer c.Close()
		var p Payload
		if err := gob.NewDecoder(c).Decode(&p); err != nil {
			return
		}
		fmt.Printf("Gob: %v\n", p.Data)
	}()

	client, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	if err := gob.NewEncoder(client).Encode(Payload{Data: map[string]int{"status": 200}}); err != nil {
		panic(err)
	}
	<-done
}
```

（若你的 Go 版本低于 1.18，将 `any` 改为 `interface{}`。）
