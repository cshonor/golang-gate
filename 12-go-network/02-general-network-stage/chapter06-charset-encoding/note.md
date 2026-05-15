# Chapter 06 — Managing Character Sets and Encodings（字符集与编码）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 6 章。  
> 在分布式链路里，**编码一致**不只是「显示不乱码」——它决定**语义是否对齐**、日志与审计是否可信；偏差会退化为 **Mojibake（乱码）**，极端时与**解析歧义、注入、身份校验绕过**等安全问题交织。

**性质**：**表示层**心智；与第 **4** 章（序列化后的**字节形态**）、第 **5** 章（**文本协议**）、第 **7** 章（**密与身份**）叠读。

---

## 6.1 核心定义（Definitions）

先统一术语，避免「字符集 / 编码 / 传输编码」混谈。

| 概念 | 定义 | 工程示例 |
|------|------|----------|
| **字符（Character）** | 抽象符号，与具体存储无关 | 汉字「好」、符号「∑」 |
| **字符谱 / 字符集（Repertoire / Set）** | 允许出现的字符**逻辑集合** | Unicode、ASCII |
| **字符代码 / 码点（Code point）** | 字符在集合中的**数值编号** | `'A'` → U+0041（十进制 65） |
| **字符编码（Encoding）** | 码点 ↔ **字节序列**的规则 | UTF-8、UTF-16、ISO 8859-1 |
| **传输编码（Transport encoding）** | 为通过**文本通道**承载二进制而做的二次包装 | Base64、URL Encoding（百分号编码） |

**小结**：大量互联网控制面语法（HTTP 头行、SMTP 命令等）仍以 **US-ASCII** 的可打印子集为底线约定；万国码时代也要记得这条「**底层指令面**」。

---

## 6.2 ASCII

- **范围**：传统 **7 bit**（0～127），最高位常为 0；在 Go 里自然落在 **`byte`（`uint8`）** 上。  
- **地位**：控制面协议里仍占核心；**不要**把「业务体 UTF-8」与「元数据 ASCII」混成同一套假设。

**ASCII 校验**（按**字节**判断「是否全是 7-bit ASCII」；与「是否为合法 UTF-8」是不同问题）：

```go
import "unicode"

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII { // 0x7F
			return false
		}
	}
	return true
}
```

**局限**：7 bit 无法覆盖全球文字 → 历史上出现 **ISO 8859** 等「单字节扩展」方案。

---

## 6.3 ISO 8859 与碎片化

- **思路**：启用第 8 bit，单字节 256 个码位，例如 **ISO 8859-1（Latin-1）** 覆盖西欧常见字符。  
- **代价**：**同形字节、不同解释**——发送方按 8859-1 编码、接收方按 8859-5（西里尔）解码会得到完全错误的文本；跨国集成维护成本极高。  
- **出路**：统一到 **Unicode 码点空间**，再用明确编码（首选 **UTF-8**）落到线路上。

---

## 6.4 Unicode：码点与编码分离

- **Unicode**：全球字符的**码点**集合（U+xxxx…）。  
- **UTF-8 / UTF-16**：把码点变成字节的**不同编码方案**——二者不可混为一谈。  
- **网络默认**：**UTF-8**——向下兼容 ASCII、**无主机字节序问题**（对比 UTF-16）、变长对英文省带宽，适合 **TCP 字节流** 上承载文本。

---

## 6.5 UTF-8、Go 与 `rune`

- **`string`**：只读字节序列；**源码**与 **`string` 字面量**在 Go 中按 **UTF-8** 解释。  
- **`len(s)`**：**字节数**，不是「字符数」——`len("你好") == 6`（两汉字常各 3 字节）。  
- **`rune`**：`int32` 别名，表示一个 **Unicode 码点**。

**迭代方式**：

```go
s := "Go语言"

// 按字节索引：易切断多字节 UTF-8 序列（仅适合 ASCII 协议或已知全 ASCII）
for i := 0; i < len(s); i++ {
	_ = s[i]
}

// 按码点迭代（文本逻辑首选）
for pos, r := range s {
	_, _ = pos, r // pos 是该码点起始字节下标
}
```

**工程习惯**：需要「字符数」、按码点截取、宽度计算等，用 **`unicode/utf8`**（`RuneCountInString`、`DecodeRuneInString` 等），不要手写一半 UTF-8 边界逻辑。

---

## 6.6 UTF-8 客户端 / 服务端与流式边界

TCP 是**字节流**：一个多字节 UTF-8 字符可能被拆在两个 `Read` 返回里。  
**不要**在无法保证边界时写 `string(buf[:n])` 再当完整 UTF-8 文本解析——除非协议层保证 **n 落在码点边界**（例如整段 UTF-8 文本 + 明确长度或分隔符，且你已重组完整帧）。

**用 `bufio.Reader.ReadRune`** 在缓冲上解码：内部会拼满当前码点再返回；仍要注意 **`utf8.RuneError`**、EOF 与半包残留（对端断连时）。

```go
func handleUTF8Connection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			break
		}
		_ = ch // 业务处理；若 ch==utf8.RuneError 需结合 size/错误区分「真·非法」与「读错误」
	}
}
```

```go
func sendUTF8Data(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	defer conn.Close()
	_, _ = conn.Write([]byte("你好，Go网络编程"))
}
```

---

## 6.7 「全 ASCII」假设的现代风险

ASCII 协议里常隐含 **1 字节 = 1 码点** 的直觉，实现简单。  
若通道实际混入 **UTF-8 多字节** 或 **二进制**，而解析器仍按「逐字节控制字符」去解释，可能出现**帧错乱、状态机跑偏**甚至崩溃——根因是**协议声明的字符集 / 二进制边界**与实现不一致。

---

## 6.8 UTF-16 与 Go 标准库

互联网与 Unix 系以 **UTF-8** 为主；**Java、Windows、部分遗留中间件**常见 **UTF-16**（内存或线路）。  
每个 UTF-16 **码元**固定 16 bit，但**增补平面**码点要用**代理对**，因此对「完整 Unicode 码点」而言仍是变长序列。  

Go 标准库：

- **`unicode/utf16`**：UTF-16 **码元（uint16）** 与 **rune** 的编解码（含代理对）。  
- 线路上的 **UTF-16 + BOM / 字节序** 通常用 **`golang.org/x/text/encoding/unicode`** 包成 `transform.Transformer` 更省事。

---

## 6.9 字节序与 BOM

- **大端（BE）**：高位字节在前；**网络字节序**传统上指大端。  
- **小端（LE）**：低位在前；典型 **x86** 内存布局。  
- **BOM**：如 UTF-16 开头的 **U+FEFF** 字节序列，用于声明端序 / UTF-8 签名（是否写 BOM 取决于协议与生态，别默认「总有」）。

**显式整数网络序**（与 UTF-16 无关，但同属「多字节怎么摆」）：

```go
import "encoding/binary"

func encodeUint16BE(val uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, val)
	return b
}
```

---

## 6.10 UTF-16 流：用 `x/text` 归一到 UTF-8

依赖：`go get golang.org/x/text`

```go
import (
	"io"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func fetchUTF16AsUTF8(conn net.Conn) ([]byte, error) {
	// 大端 + 若存在则消费 BOM；与对端约定一致时再改 LittleEndian / IgnoreBOM
	dec := unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewDecoder()
	return io.ReadAll(transform.NewReader(conn, dec))
}
```

读完应 **`conn.Close()`** 或由上层管理生命周期；示例省略错误分支以突出转换链。

---

## 6.11 Unicode 工程陷阱

1. **组合字符（Combining characters）**：视觉上「一个字符」可能对应多个码点（基字符 + 组合记号）。  
2. **正规化（Normalization）**：同一视觉结果可有 **NFC / NFD** 等不同码点序列 → **直接 `==` 比较**可能误判。

```go
import "golang.org/x/text/unicode/norm"

func equalNormalizedNFC(a, b string) bool {
	return norm.NFC.String(a) == norm.NFC.String(b)
}
```

3. **同形异义字（Homograph）**：用易混淆字符伪造域名/账号等——属于**产品与安全**层防御（字表、Punycode、显示规范），不单靠编码 API。

---

## 6.12 ISO 8859 与 `charmap`

对接欧洲遗留二进制文本时，用 **`golang.org/x/text/encoding/charmap`**：

```go
import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func iso88591ToUTF8(legacy []byte) ([]byte, error) {
	return io.ReadAll(transform.NewReader(bytes.NewReader(legacy), charmap.ISO8859_1.NewDecoder()))
}
```

---

## 6.13 GBK、Big5 等中文遗留编码

政企、金融老系统常见 **GBK / GB18030 / Big5**。原则：**边界立刻转**，**内部全 UTF-8**。

```go
import (
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func gbkReader(conn net.Conn) io.Reader {
	return transform.NewReader(conn, simplifiedchinese.GBK.NewDecoder())
}

func gbkWriter(conn net.Conn) io.Writer {
	return transform.NewWriter(conn, simplifiedchinese.GBK.NewEncoder())
}
```

实际服务里再包 **`bufio`**、**超时**、**长度上限**，避免无限读。

---

## 6.14 本章小结与检查清单

编码管理是**防御式工程**：在协议与 I/O 边界把假设写清楚。

**检查清单**

- [ ] **内部统一**：业务与持久化默认 **UTF-8**；需要「字符」语义处用 **`rune` / `unicode/utf8`**。  
- [ ] **边界隔离**：外部字节流是否在**入口**就转为 UTF-8（或明确保留为 `[]byte` 二进制通道）？  
- [ ] **流式安全**：TCP 文本是否在重组帧后解码，或用 **`bufio.Reader`** 等避免切断 UTF-8？  
- [ ] **比较与索引**：敏感比较、唯一约束是否约定 **NFC/NFD**（与数据库 collation 一致）？  
- [ ] **多字节协议**：UTF-16 或二进制整型是否写清 **Endian / BOM**？HTTP/JSON 等是否写明 **charset**？

**背诵版**：**码点（Unicode）≠ 线路字节（UTF-8/16）；入口转码；流上缓冲；比较要正规化；协议写清字符集与字节序。**

**前后章节**：[`chapter04` 序列化](../chapter04-data-serialization/note.md) · [`chapter05` 应用层协议](../../01-foundation-stage/chapter05-application-protocols/note.md) · [`chapter07` 安全](../chapter07-security/note.md)（TLS 不改变「编码要先做对」的事实）
