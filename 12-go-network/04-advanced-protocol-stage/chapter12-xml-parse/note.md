# Chapter 12 — XML 处理实战

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 12 章。  
> 在 **JSON** 主导的 REST 生态之外，**XML** 仍在 **SOAP、NETCONF、企业配置交换、遗留集成** 中承担「强 schema + 命名空间」角色。Go 的 **`encoding/xml`** 提供 **`Marshal` / `Unmarshal`** 与 **`Decoder.Token` 流式**两条路径；选型取决于**报文规模**、**是否需早失败**与**安全面**。

**性质**：与 [`chapter11` HTML](../../04-advanced-protocol-stage/chapter11-html-parse/note.md) 对读（**HTML 容错 ≠ XML 严格**）；与 [`chapter04` 序列化](../../03-web-core-stage/chapter04-data-serialization/note.md) 的「线路格式」心智一致。

---

## 12.1 反序列化（`xml.Unmarshal`）

**入口语义**：把**不可信网络上的字节**变成 **DTO**（传输对象），再映射到领域模型——避免把外部 schema **直接绑死**在内核模型上。

**反射与标签**：`Unmarshal` 依赖反射；常用标签：

| 标签 | 含义 |
|------|------|
| **`xml:"name"`** | 匹配子元素 **Local 名**（可配合命名空间，见官方文档） |
| **`xml:"name,attr"`** | 映射为属性 |
| **`xml:",chardata"`** | 该元素下的字符数据（不含子元素树） |
| **`xml:",innerxml"`** | 子树原始 XML 文本（常用于透传、审计；与**其它字段切分同一子树**时需谨慎设计） |

**嵌套路径**：**`xml:"endpoint>host"`** 表示 **`endpoint` 下的 `host`**（与结构体嵌套二选一，按可读性选型）。

```go
// 需 import: encoding/xml, fmt, log
type ServiceConfig struct {
	XMLName  xml.Name `xml:"service"`
	Type     string   `xml:"type,attr"`
	Priority int      `xml:"priority"`
	Host     string   `xml:"endpoint>host"`
	Port     int      `xml:"endpoint>port"`
}

const blob = `<service type="BGP"><priority>10</priority><endpoint><host>192.168.1.1</host><port>179</port></endpoint></service>`

var cfg ServiceConfig
if err := xml.Unmarshal([]byte(blob), &cfg); err != nil {
	log.Fatal(err)
}
fmt.Printf("%s %s:%d\n", cfg.Type, cfg.Host, cfg.Port)
```

**易错点**：仅**导出字段**可解码；**大小写**与 **命名空间**敏感；schema 演进时 DTO 要版本化。

---

## 12.2 序列化（`xml.Marshal`）

- **`xml.Marshal`**：紧凑输出，适合**热路径**。  
- **`xml.MarshalIndent`**：人类可读，**额外分配与体积**——默认避免在 QPS 核心链路上用。  
- **`xml.Header`**：标准 XML 声明前缀（常 **`Write([]byte(xml.Header))` 再 `Marshal`**）。

**命名空间**：根上可用 **`xml.Name{Space: "http://...", Local: "response"}`** 配合结构体 **`xml:"http://... response"`** 形式（**空格**分隔 **URI 与 local**）。

```go
// 需 import: encoding/xml, os
type NetworkResponse struct {
	XMLName   xml.Name `xml:"http://standards.example/ns response"`
	Status    string   `xml:"status,attr"`
	SessionID string   `xml:"session_id"`
}

out, err := xml.Marshal(NetworkResponse{Status: "SUCCESS", SessionID: "X-9921"})
if err != nil {
	return
}
_, _ = os.Stdout.Write([]byte(xml.Header))
_, _ = os.Stdout.Write(out)
```

**扩展**：实现 **`xml.Marshaler` / `xml.Unmarshaler`** 可接管局部序列化（时间格式、非标准封装等）。**转义**：`Marshal` 会对文本/属性做 **XML 转义**，降低注入面（仍需整体安全设计）。

---

## 12.3 流式解析（`xml.Decoder` + `Token`）

**动机**：**整篇 `Unmarshal` 到大结构**在 GB 级日志 / 巨型配置上易导致 **OOM**；**`Token` 循环**以近似**常数内存**推进。

**Token 类型**（类型断言识别）：**`xml.StartElement`**、**`xml.EndElement`**、**`xml.CharData`**、**`xml.Comment`**、**`xml.ProcInst`**、**`xml.Directive`**。

**混合模式**：在 **`StartElement`** 上 **`DecodeElement(ptr, &el)`**，把**当前子树**解到局部 DTO，兼顾**定位**与**强类型**。

```go
// 需 import: encoding/xml, fmt, io, strings
func parseNodes(src string) error {
	dec := xml.NewDecoder(strings.NewReader(src))
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local != "node" {
				continue
			}
			var n struct {
				ID      string `xml:"id,attr"`
				Content string `xml:",chardata"`
			}
			if err := dec.DecodeElement(&n, &t); err != nil {
				return err
			}
			fmt.Printf("id=%s content=%q\n", n.ID, n.Content)
		}
	}
}
```

**安全**：关注 **实体扩展 / XXE** 类风险；对**不可信输入**限制体积、启用 **`Strict`**（若适用）、避免盲目解析 **DOCTYPE 外部实体**（具体以 Go 版本与配置为准，生产需威胁建模）。

---

## 12.4 XHTML、HTML 与 `encoding/xml`

**XML 解析器「不容错」**：未闭合的 **`img` / `br`**、隐式标签等 **HTML** 常使 **`encoding/xml`** **立即报错**——**不要**把它当通用 **HTML 爬虫**引擎。

**Web 脏 HTML**：用 **`golang.org/x/net/html`**（见第 11 章）；**良构 XHTML** 或 **配置/协议 XML** 再用 **`encoding/xml`**。

---

## 12.5 小结：XML 在 Go 网络栈中的位置

| 维度 | XML（`encoding/xml`） | JSON（`encoding/json`） |
|------|-------------------------|-------------------------|
| 典型强项 | 命名空间、企业/电信协议、复杂配置 | 轻量 API、微服务默认可读格式 |
| 解析成本 | 标签与命名空间更重 | 通常更低（视体积与路径） |
| 大文件 | **`Decoder` 流式**优先 | **`Decoder` / 分块`** 同理 |

**黄金法则**

1. **大报文 / 未知大小**：优先 **`Token` + `DecodeElement`**，避免一次性 DOM 式 `Unmarshal`。  
2. **不可信 XML**：防 **炸弹与注入**；限长、限深、关危险特性（按需求配置）。  
3. **解耦**：**DTO + struct tag** 隔离外部协议与内部模型。  
4. **HTML vs XML**：**容错 HTML ≠ 严格 XML**，选型错误会直接体现在线上故障率上。

**背诵版**：**小用 Unmarshal/Marshal；大用 Decoder.Token；HTML 用 x/net/html；不可信 XML 当安全面做。**

**前后章节**：[`chapter11` HTML](../chapter11-html-parse/note.md) · [`chapter04` 序列化](../../03-web-core-stage/chapter04-data-serialization/note.md) · [`chapter13` RPC](../chapter13-rpc/note.md)
