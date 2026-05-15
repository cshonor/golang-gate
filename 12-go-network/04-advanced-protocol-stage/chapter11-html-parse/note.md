# Chapter 11 — HTML 编程实战

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 11 章。  
> **HTML 的生成与解析**不是「前端琐事」：生成侧关系 **XSS 与模板安全**；解析侧关系 **爬虫、监控、协议适配** 的吞吐与内存。工业上应避免用**裸字符串拼接**拼 HTML——无结构感知、易引入注入面。

**性质**：衔接 [`chapter09` 模板](../../03-web-core-stage/chapter09-templates/note.md) 与 [`chapter10` 整站](../../03-web-core-stage/chapter10-complete-web-server/note.md)；解析侧与 **第 12 章 XML** 对读。

---

## 11.1 `html/template`：安全模板引擎

**与 `text/template` 的本质差异**：同一套语法，**`html/template` 在 `Parse` 阶段**结合 HTML/CSS/JS **语境**做分析，对动态插入做**上下文敏感转义**（可理解为针对不同插入点的不同规则），降低 **XSS**。

**同一段数据在不同语境**（示意）：

- 在 **HTML 文本节点**（如 `<div>{{.}}</div>`）中：`<` `>` 等走 **HTML 实体**转义。  
- 在 **`<script>` 内的 JS 字面量**语境中：转义规则会偏向 **不破坏 JS 词法**（例如避免提前闭合字符串）。

**预编译（ParseOnce）**：`Parse` / `ParseFiles` 属于重操作；生产应在 **启动阶段**完成（`template.Must` 做强校验），请求路径只做 **`Execute`**（与第 9、10 章一致）。

**`template.HTML` / `template.JS` / `template.JSStr` 等**：表示「**我已保证这段内容安全**」，引擎**跳过**对应语境的自动转义。**误用 = 主动开洞**；仅在对数据源做过 **严格洗刷（sanitize）** 或 **结构可证明安全**（如你自己 `json.Marshal` 的配置对象）时使用。

**铁律**：Web 页面输出用 **`html/template`**；**禁止**用 **`text/template`** 处理不可信 HTML 片段。

```go
// 需 import: html/template, os
// 同一 .Content 在 <div> 与 <script> 中会得到不同语境的转义（自行 Execute 到 stdout 观察）
tmpl := template.Must(template.New("demo").Parse(
	`<div>{{.Content}}</div><script>var x = {{.Content}};</script>`))
_ = tmpl.Execute(os.Stdout, struct{ Content string }{
	Content: "<b>hi</b>",
})
```

---

## 11.2 HTML 词法分析（`golang.org/x/net/html`）

依赖：`go get golang.org/x/net`

**Tokenizer（流式）** vs **`html.Parse`（DOM 树）**：

| 模式 | 优点 | 代价 |
|------|------|------|
| **`html.NewTokenizer`** | 不建整棵树，**内存更平**；适合高吞吐爬虫/扫描 | 手写状态机处理标签/属性 |
| **`html.Parse`** | 得 `*html.Node`，便于 DFS/BFS 查询 | 大文档 **分配与指针追踪** 更重 |

**`TagAttr()`**：返回的 **`[]byte` 往往指向 tokenizer 内部缓冲**——若要在 **`Next()` 之后**保留字符串，请 **`append([]byte(nil), val...)`** 或 **`string(append([]byte(nil), val...))`** 拷贝。

**ErrorToken 与 EOF**：`ErrorToken` 时 **`z.Err()`** 常为 **`io.EOF`**（正常结束）；**不要**用 **`z.Err().Error() == "EOF"`** 字符串比较，应使用 **`errors.Is(z.Err(), io.EOF)`**。

```go
// 需 import: errors, fmt, io, strings, golang.org/x/net/html
func printHrefsFragment(htmlSrc string) {
	z := html.NewTokenizer(strings.NewReader(htmlSrc))
	for {
		switch z.Next() {
		case html.ErrorToken:
			if errors.Is(z.Err(), io.EOF) {
				return
			}
			panic(z.Err())
		case html.StartTagToken, html.SelfClosingTagToken:
			name, hasAttr := z.TagName()
			if string(name) == "a" && hasAttr {
				for {
					key, val, more := z.TagAttr()
					if string(key) == "href" {
						fmt.Println(string(val))
					}
					if !more {
						break
					}
				}
			}
		}
	}
}
```

---

## 11.3 XHTML 与 HTML5

- **`x/net/html`**：遵循 **HTML5** 容错规则，能**修复**缺失闭合标签等——灵活，但可能**掩盖上游脏数据**。  
- **严格 XHTML / 已知 Well-formed XML**：可评估 **`encoding/xml`**（或专用解析器）走 **早失败（fail-fast）**；在**结构强约束**场景下往往更易断言数据质量。  
- **「XML 一定快 20～30%」**之类数字仅作**不可盲信**的粗量级直觉：以 **profile + 真实负载** 为准。

---

## 11.4 HTML 中的 JSON（Data seeding）

常见模式：

1. **`data-*` 属性**：把 JSON 放进 **`data-conf="..."`**；由 **`html/template`** 做 **属性语境**转义。  
2. **`<script>` 内嵌全局配置**：需要 **`json.Marshal` 成功**且内容在 **JS 语境**下安全。

**`<script>` 内嵌对象字面量**：把 **`json.RawMessage` 或 `[]byte` 转成 `template.JS`**，表示「这是一段 **JS 表达式**（此处为 JSON 文本）」——避免被当成普通字符串再包一层引号。仍需防范 JSON 中出现 **`</script>`** 子序列破坏 HTML 封装；高要求场景可对 JSON 做额外转义或改用 **`data-*` + `JSON.parse`**。

```go
// 需 import: encoding/json, html/template
type AppConfig struct {
	Region  string `json:"region"`
	Version string `json:"version"`
}
```

```go
// 示意：Marshal 后交给 template.JS（仅当内容为受控 JSON 时）
raw, _ := json.Marshal(AppConfig{Region: "ap-1", Version: "v2.1.0"})
data := map[string]any{
	"JSONData": template.JS(raw), // 非用户任意字符串
}
```

**反模式**：把 **`string(jsonBytes)`** 当普通字符串塞进 **`const cfg = {{.}}`** 却不匹配引擎期望类型，易出现**错误转义**或**语法断裂**——统一用 **`template.JS`** / 官方推荐模式。

---

## 11.5 本章小结

1. **生成 HTML**：**`html/template` + 预编译**；慎用 **`template.HTML`/`template.JS`** 绕过转义。  
2. **解析 HTML**：大流量优先 **`Tokenizer`**；注意 **EOF** 判断与 **`[]byte` 生命周期**。  
3. **策略**：HTML5 容错 vs XML **早失败**，按数据源与 SLO 选型。  
4. **JSON 进页**：属性 vs 脚本两条路径；脚本路径用 **`template.JS`** 并处理 **`</script>`** 风险。

**背诵版**：**生成用 html/template 预编译；解析用 Tokenizer 控内存；JSON 进 script 用 template.JS 且防闭合脚本序列。**

**前后章节**：[`chapter09` 模板](../../03-web-core-stage/chapter09-templates/note.md) · [`chapter10` 整站](../../03-web-core-stage/chapter10-complete-web-server/note.md) · [`chapter12` XML](../chapter12-xml-parse/note.md)
