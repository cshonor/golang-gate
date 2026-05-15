# Chapter 09 — Templates（模板）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 9 章。  
> 模板引擎不是「字符串拼接」：它是**表现层**与**数据模型**之间的契约面。Go 用同一套语法分裂为 **`text/template`**（通用文本）与 **`html/template`**（Web 安全默认），网络服务里**生成 HTML 必须用后者**。

**性质**：与 [`chapter08` HTTP](../chapter08-http/note.md) 的 `Handler` 输出、[`chapter10`](../chapter10-complete-web-server/README.md) 整站案例衔接；导出规则见 [`chapter02`](../../01-foundation-stage/chapter02-go-language-overview/note.md)。

---

## 9.1 插入对象值（Dot 与 `Execute`）

**调用链**：**`template.New(name)`** → **`Parse` / `ParseFiles`**（构建 AST）→ **`Execute(w, data)`** 把数据灌进模板。

**`{{ . }}` 的「点」**：表示**当前作用域的根值**；在根模板里通常就是 `Execute` 传入的 `data`。字段访问 **`{{ .Field}}`** 依赖**反射**，且仅 **`Exported`（首字母大写）** 字段可见——小写字段在模板侧**不可达**，常见后果是**执行期报错**或呈现**零值/空**，应在代码审查与测试里兜住。

**性能**：`Parse` 含 I/O 与编译成本，**禁止在每个 HTTP 请求里 Parse**。应在 **`init` / 启动阶段**解析一次，请求路径只做 **`Execute`**（必要时配合 **`Clone`** 做并发安全的分支模板）。

```go
// 需 import: os, text/template
type Name struct {
	Family, Personal string
}

tmpl, err := template.New("welcome").Parse("尊敬的 {{.Personal}} · {{.Family}}，您好！\n")
if err != nil {
	panic(err)
}
_ = tmpl.Execute(os.Stdout, Name{Family: "Newmarch", Personal: "Jan"})
```

---

## 9.2 管道（Pipelines）

管道语义接近 UNIX：**`{{ a | f | g }}`** 本质是 **`g(f(a))`**（前一个结果为**下一个函数的最后一个参数**；多参函数时其余参数由模板内置规则填充）。

用于展示层轻量格式化（`printf`、内置字符串函数等），避免把「仅为排版服务」的逻辑塞进业务层。

```go
tmpl, _ := template.New("pipe").Parse(`用户标签: {{ .Personal | printf "[%s]" }} {{ .Family }}` + "\n")
```

---

## 9.3 自定义函数（`FuncMap`）

**`template.FuncMap`**：本质是 **`map[string]any`**（历史写法 `interface{}`）。把 Go 函数挂进模板命名空间。

**硬规则**：**必须先 `Funcs(funcMap)`，再 `Parse`**。否则 AST 里出现未知标识符会直接 **Parse 失败**。

**架构边界**：复杂分支、聚合、权限决策放在 **Go**；模板侧只做展示需要的**纯格式化 / 轻量映射**（日期格式、货币单位等）。

```go
// 需 import: os, strings, text/template, time
funcMap := template.FuncMap{
	"upper": strings.ToUpper,
	"formatDate": func(t time.Time) string { return t.Format("2006-01-02") },
}
tmpl, err := template.New("f").Funcs(funcMap).Parse("生成日期: {{ .Time | formatDate }}\n操作员: {{ .User | upper }}\n")
if err != nil {
	panic(err)
}
_ = tmpl.Execute(os.Stdout, map[string]any{"Time": time.Now(), "User": "admin_jan"})
```

---

## 9.4 变量（`$`）

**`{{ $v := . }}`**：把当前「点」快照到 **`$v`**，在 **`range` / `with`** 等改变「点」的作用域里仍可读。

**`range` 双变量**：**`{{ range $idx, $email := . }}`** 在根数据为切片时同时拿到**下标与元素**（与 Go `for range` 心智一致；具体语法以当前 Go 版本文档为准）。

```go
tmplStr := `邮件列表:
{{ range $idx, $email := . }}
  第 {{ $idx }} 项: [{{ $email.Kind }}] -> {{ $email.Address }}
{{ end }}`
```

---

## 9.5 条件与循环（`if` / `range` / `with`）

- **`if`**：布尔分支；**`with`**：若值「非空」则进入块并把 **`.`** 切到该值上，便于深层字段展示。  
- **`range`**：`nil` 或长度为 0 的切片：**零次迭代**，不报错。

**作用域与 `$`**：在 **`range` / `with` 内**，**`.`** 指向当前项；要访问 **`Execute` 根上下文**用 **`{{ $.Field }}`**（`$` 绑定根数据）。

综合示例（`with` 可换成 `if` + 深层字段，按可读性选择）：

```go
package main

import (
	"os"
	"text/template"
)

const tmplStr = `
用户: {{ .Name.Personal }}
{{ if .Emails }}
联系方式:
{{ range $e := .Emails }}
  - [{{ $e.Kind }}] {{ $e.Address }}
{{ end }}
{{ else }}
  (未记录联系方式)
{{ end }}`

type Email struct{ Kind, Address string }
type Person struct {
	Name   struct{ Personal string }
	Emails []Email
}

func main() {
	tmpl, err := template.New("flow").Parse(tmplStr)
	if err != nil {
		panic(err)
	}
	var p Person
	p.Name.Personal = "Jan"
	p.Emails = []Email{{Kind: "Work", Address: "jan@work.com"}}
	_ = tmpl.Execute(os.Stdout, p)
}
```

---

## 9.6 `html/template` 与 XSS

**上下文敏感转义**：引擎理解占位符落在 **文本节点、属性、URL、`script` 块**等不同语境，并套用不同转义策略（实体、URL 分段、JS 字符串规则等），显著降低「拼接 HTML」导致的 **XSS**。

**规范**：向浏览器输出 HTML 页面时，**默认使用 `html/template`**，不要用 **`text/template`** 拼 HTML。

**已知安全 HTML 的「逃生舱」**：**`template.HTML`**、**`template.CSS`** 等类型声明「我已保证安全」——**绕过转义**，误用即漏洞。富文本场景应先用 **`bluemonday`** 等做 HTML 清理再决定是否标记为 `template.HTML`。

| 场景 | `text/template` | `html/template` |
|------|-----------------|-------------------|
| 用户输入 `<script>…` | 常原样写出（危险） | 按语境转义，默认不执行 |

```go
// 需 import: html/template, os
tmpl, _ := template.New("safe").Parse("用户评论: {{ . }}\n")
_ = tmpl.Execute(os.Stdout, "<script>alert('XSS')</script>")
```

---

## 9.7 小结与工程清单

**知识链**：反射字段映射 → 管道 → **`Funcs` 扩展** → **`$` / `$.` 作用域** → **`html/template` 安全默认**。

**清单**

1. **ParseOnce**：启动期 `Parse`；热路径只 `Execute`（或 `ExecuteTemplate`）。  
2. **Web 用 `html/template`**。  
3. **逻辑归 Go、展现归模板**；`FuncMap` 只做格式化与薄封装。  
4. **导出字段**、**`Funcs` 先于 `Parse`**、**内层用 `$.` 指根**。

**背诵版**：**预编译模板；Web 用 html/template；FuncMap 前置；块内用 $. 找回根；慎用 template.HTML。**

**前后章节**：[`chapter08` HTTP](../chapter08-http/note.md) · [`chapter10` 完整 Web 服务器](../chapter10-complete-web-server/note.md) · [`chapter02` Go 概览](../../01-foundation-stage/chapter02-go-language-overview/note.md)
