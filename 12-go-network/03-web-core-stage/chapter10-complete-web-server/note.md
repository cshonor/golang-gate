# Chapter 10 — A Complete Web Server（完整 Web 服务器）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 10 章。  
> 以 **中文闪卡（Chinese Flashcards）** 串联 **`net/http`**、**`html/template`**、**表单与重定向**、**共享状态并发安全**与**静态资源**——体现的不是「堆框架」，而是对 **请求生命周期**、**边界** 与 **可运维路径** 的控制。

**性质**：小型整站案例；与 [`chapter08` HTTP](../chapter08-http/note.md)、[`chapter09` 模板](../chapter09-templates/note.md)、[`chapter06` 编码](../../02-general-network-stage/chapter06-charset-encoding/note.md) 叠读。

---

## 10.1 站点图（Routing 蓝图）

先画 **URL → 职责**，再写 `ServeMux`，路由表即中间件挂载点（前缀鉴权、限流、trace 等）。

| 路径 | 职责 |
|------|------|
| **`/`** | 首页 / 闪卡列表 |
| **`/show`** | 单张闪卡学习 |
| **`/manage`** | 词条 CRUD（表单） |
| **`/static/`** | CSS / JS / 图片（`http.FileServer`） |

**架构提示**：例如对 **`/manage`** 前缀统一加 **认证中间件**，业务 `Handler` 保持「只关心业务」。

---

## 10.2 浏览器端文件（SoC）

**关注点分离**：模板与 Go 分离；页面生成用 **`html/template`**，利用**上下文转义**降低 XSS（见第 9 章）。

**原书式文件布局（示例）**

- `templates/base.html` — 全站布局（`head`、导航、页脚），内嵌 **`{{ define "base" }}`** 等根块。  
- `templates/list.html` / `show.html` / `manage.html` — 子页面 **`{{ define "…" }}`** 或 `{{ template "…" . }}`。  
- `static/css/style.css` — 样式（闪卡翻转、响应式等）。

**路径陷阱**：避免依赖进程 **CWD** 写死 `./templates`。更稳做法：`embed` 嵌入、`filepath` 相对**可执行文件**或**显式配置根目录**（配置项 / 环境变量），与 [`chapter08`](../chapter08-http/note.md) 中「路径与资源」心智一致。

---

## 10.3 基础服务器（`http.Server` + 封装）

用 **`application` 结构体** 挂接字典等业务依赖，避免散落全局变量，便于测试与演进。

**静态资源惯用法**（注意 **`StripPrefix` 与 `Handle` 路径一致**，且带尾部 **`/`**）：

```go
// 需 import: log, net/http
fileServer := http.FileServer(http.Dir("static"))
mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
```

原书若写 **`StripPrefix("/static", …)`** 在部分路径下易与 **`/static/`** 不匹配；推荐 **`/static/`** 成对出现。

```go
type application struct {
	dict *Dictionary
}

mux := http.NewServeMux()
mux.HandleFunc("/", app.listFlashCards)
mux.HandleFunc("/manage", app.manageFlashCards)
// mux.HandleFunc("/show", app.showFlashCards)

srv := &http.Server{Addr: ":8080", Handler: mux}
log.Fatal(srv.ListenAndServe())
```

**并发模型**：每请求 **一 goroutine**（实现细节可演进），Handler 内访问的**共享可变状态**必须加锁或通道隔离（见 10.7）。

---

## 10.4 `listFlashCards`（列表 + `ExecuteTemplate`）

- **精确路由**：仅 **`/`** 时列全部，否则 **`http.NotFound`**，避免前缀路由误吞子路径。  
- **模板**：`html/template.ParseFiles("templates/base.html", "templates/list.html", …)` → **`ExecuteTemplate(w, "base", data)`**。

**与第 9 章衔接**：原书为教学可能在 Handler 内 **`ParseFiles`**；生产应 **启动时 `ParseGlob` / `ParseFiles` 一次**，请求路径只做 **`ExecuteTemplate`**（[`chapter09` §9.7](../chapter09-templates/note.md)）。若坚持在 Handler 内解析，至少 **`template.Must`** + **缓存 `*template.Template`**。

**错误处理**：`ExecuteTemplate` 失败时响应可能已写入部分字节——严格场景可先 **`Execute` 到 `bytes.Buffer`** 再一次性 **`Write`**，或接受「记录日志 + 监控」策略。

---

## 10.5 `manageFlashCards`（GET 展示 / POST 提交）

- **按 Method 分派**：GET 渲染表单；POST 处理写入；其它返回 **`405`** 并带 **`Allow`** 头。  
- **防大 body**：**`http.MaxBytesReader(w, r.Body, limit)`** 限制读取，减轻 DoS 面。  
- **`ParseForm`** 后读 **`PostForm`**；服务端校验非空、长度、字符集等。  
- 成功后 **`http.Redirect(w, r, "/", http.StatusSeeOther)`**（POST-Redirect-GET 模式，避免刷新重复提交）。

```go
switch r.Method {
case http.MethodGet:
	// 渲染 manage 模板
case http.MethodPost:
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "无效的表单数据", http.StatusBadRequest)
		return
	}
	// 校验 + app.dict.AddEntry(...)
	// http.Redirect(..., http.StatusSeeOther)
default:
	w.Header().Set("Allow", http.MethodGet+", "+http.MethodPost)
	http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
}
```

**业务码**：字段校验失败可用 **`422 Unprocessable Entity`**（`http.StatusUnprocessableEntity`）等与前端约定。

---

## 10.6～10.7 中文词典与 `Dictionary`（并发）

- **UTF-8**：HTTP 与 Go 源码默认 UTF-8；展示前对「字节 vs 字符」敏感逻辑用 **`[]rune` / `unicode/utf8`**（第 6 章）。  
- **`Dictionary`**：用 **`struct` + 方法** 封装 `entries`，比裸 `map` 更易挂行为、日志与不变式。  
- **并发**：多 goroutine 同时 `AddEntry` / `ListWords` 时，用 **`sync.RWMutex`**：**写 `Lock`**，**读 `RLock`**。

```go
// 需 import: sync
type Entry struct {
	ID        int // 若用查询参数 ?id= 展示详情，需稳定主键
	Character string
	Pinyin    string
	Meaning   string
}

type Dictionary struct {
	mu      sync.RWMutex
	entries []Entry
	nextID  int
}

func (d *Dictionary) AddEntry(char, pinyin string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.nextID++
	d.entries = append(d.entries, Entry{ID: d.nextID, Character: char, Pinyin: pinyin})
}
```

---

## 10.8 闪卡集合（Sets）

**集合**是按难度（如 HSK）、主题对 `Entry` 的逻辑分组；同一套引擎可挂多套 Set，利于扩展。实现上可以是 **切片引用**、**过滤谓词** 或 **单独表**——本章心智是「分层：词典 vs 学习视图」。

---

## 10.9 拼音数字调号 → Unicode（`convertTones`）

教学数据常用 **`ma1`…`ma4`**；展示可换成 **带调字母**。简单做法是 **查表 + `strings.ReplaceAll`**：

```go
// 需 import: strings
func convertTones(s string) string {
	repl := map[string]string{
		"a1": "ā", "a2": "á", "a3": "ǎ", "a4": "à",
		// … e/i/o/u/v 同理；ü 常写作 v 或 u:
	}
	out := s
	for from, to := range repl {
		out = strings.ReplaceAll(out, from, to)
	}
	return out
}
```

**工程提醒**：`map` **遍历顺序不定**不影响此处「互不重叠 token」替换；若有 **重叠模式** 或 **ü / u:** 多种写法，应做 **规范化** 或 **最长匹配**；也可考虑 **`golang.org/x/text/unicode/norm`**（与第 6 章正规化呼应）。

---

## 10.10 `ListWords`（过滤与读锁）

读多写少场景：**`RLock`** 遍历；**预分配 `make([]Entry, 0, len(d.entries))`** 减少扩容。过滤逻辑注意 **大小写** 与 **中文子串**（`strings.Contains` 对 UTF-8 子串按 Unicode 码点工作，但仍要警惕「用户输入即模式」带来的性能问题——大数据应索引或前缀树）。

---

## 10.11 `showFlashCards`（查询参数 + 详情）

```go
id := r.URL.Query().Get("id")
entry, err := app.dict.GetByID(id) // 内部 strconv.Atoi + 边界检查
if err != nil {
	http.NotFound(w, r)
	return
}
// ParseFiles / ExecuteTemplate("base", entry) — 生产改为预编译 + html/template
```

**安全**：`id` 必须 **校验为整数** 且 **范围合法**；勿把未校验字符串拼进文件路径或 SQL。

---

## 10.12 浏览器呈现（静态 URL）

**根路径引用静态资源**：**`<link href="/static/css/style.css">`**（以 **`/`** 开头），避免在 **`/manage/...`** 等深层 URL 下，相对路径 **`css/style.css`** 被浏览器解析到错误前缀。

---

## 10.13 运行与验证

```bash
go mod init chinese-flashcards
go mod tidy
go run .
go build -o flashcard_server .
```

**手测**：`/` 列表、`/manage` POST、`han4 yu3` → 展示 **`hàn yǔ`**（依你的 `convertTones` 表）、DevTools 确认 **`/static/...` 无 404**。

---

## 10.14 本章小结

1. **站点图 → 路由 → 中间件挂载点**。  
2. **`html/template` + 静态 `StripPrefix`**；模板与二进制部署路径要可配置。  
3. **POST**：**`MaxBytesReader`**、校验、**PRG 重定向**。  
4. **共享词典**：**`RWMutex`** 或更好的存储后端。  
5. **中文**：UTF-8 全链路；调号显示可做表驱动替换并留「规范化」升级口。

**背诵版**：**路由表清晰；模板预编译；POST 限流与校验；共享状态加锁；静态资源用绝对路径。**  

**下一阶**：[`chapter11` HTML 解析与生成](../../04-advanced-protocol-stage/chapter11-html-parse/note.md)、REST/XML 等；本仓库 [`13-projects-optional`](../../legacy-topic-index/13-projects-optional/) 可作扩展项目池。

**前后章节**：[`chapter09` 模板](../chapter09-templates/note.md) · [`chapter08` HTTP](../chapter08-http/note.md) · [`chapter06` 编码](../../02-general-network-stage/chapter06-charset-encoding/note.md) · [`chapter11` HTML](../../04-advanced-protocol-stage/chapter11-html-parse/note.md)
