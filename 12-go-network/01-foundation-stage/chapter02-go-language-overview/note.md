# Chapter 02 — Overview of the Go Language（Go 语言概览）

> 对应原书 *Network Programming with Go Language, 2nd Ed.* 第 2 章。  
> 对网络/分布式场景而言，Go 不仅是语法：**Goroutine、快速编译、显式错误**，是把第 1 章的架构假设落到**可维护代码**上的常用工具链。

更系统的语言笔记见仓库：[00-basic-types](../../../00-basic-types/README.md)、[07-GMP and channel](../../../07-GMP%20and%20channel/README.md)、[03-error_handling](../../../03-error_handling/README.md)。

---

## 2.1 类型（Types）

- **`byte`**（`uint8`）：处理**原始字节流**、协议缓冲区的基本单位。  
- **`rune`**（`int32`）：Unicode 码点；国际化与多字节文本需注意与 `string`/`[]byte` 的关系（见 [04-字符串与rune](../../../00-basic-types/04-字符串与rune.md)）。

---

## 2.2 切片与数组（Slices and Arrays）

- **数组 `[N]T`**：定长值类型，内存连续。  
- **切片 `[]T`**：对底层数组的**视图**，语义上可理解为包含 **指针、len、cap** 的结构；`append` 可能触发扩容与底层数组迁移。

```go
// 长度为 50、容量 100，减少频繁扩容（按负载调参）
x := make([]int, 50, 100)

a := [5]int{-1, -2, -3, -4, -5}
s := a[1:4] // 索引 1,2,3 → 值 -2,-3,-4
// len(s)==3, cap(s)==4（从索引 1 到数组末尾可容纳的元素数）
```

**工程直觉**：网络 I/O 中常用 `[]byte` 缓冲与切片视图减少不必要的整段拷贝（仍要理解**何时会复制**，见基础类型篇 `string`/`[]byte`）。

---

## 2.3 映射（Maps）

```go
dataMap := map[string]int{"status": 200}
delete(dataMap, "status")
```

- **`nil` map 不能写入**：须 `make` 或字面量初始化后再赋值。  
- **遍历顺序**：`for range map` **不可依赖稳定顺序**（语言有意不保证）；与「Go 1.12 才随机」的坊间说法不同，**不要**写依赖遍历顺序的业务逻辑。

---

## 2.4 指针（Pointers）

无指针算术；`*T` 与 `nil` 常用于表达「可选 / 未初始化」——例如配置指针、尚未 `Dial` 的连接句柄等。**解引用前**必须保证非 `nil`。

---

## 2.5 函数（Functions）

多返回值 + **`error`** 是网络 I/O 的惯用法：每个可能失败点显式分支。

```go
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
```

**工程提示**：`os.Exit` 适合**小工具/教学**；**库代码**应 **`return err`** 交由调用方决定策略（重试、降级、打日志），避免在库内直接退出进程。

---

## 2.6 结构体与方法（Structures & Methods）

无类继承，靠**组合**与**方法集**表达行为。

```go
type Person struct {
	Family, Personal string
}

func (p Person) String() string {
	return p.Personal + " " + p.Family
}
```

**函数为一等值**、高阶函数常用于中间件、装饰 `http.Handler` 等（见后续 HTTP 章）。

---

## 2.7 多线程与并发（Multithreading）

`go` 启动 **Goroutine**：调度由 runtime 完成，相比「一线程一连接」更易在高并发下伸缩（具体仍受文件描述符、内存与调度模型约束，见本仓库 **netpoll / GMP** 笔记）。

---

## 2.8 包与模块（Packages & Modules）

- **导出规则**：首字母大写 = 包外可见。  
- **`go mod`**：`go mod init`、`go mod tidy` 维护依赖图，保证**可重复构建**（生产与 CI 的默认路径）。

---

## 2.9 类型转换与适配器（Type Conversion）

网络读写中常见 **`[]byte` ↔ `string`**（注意拷贝与 UTF-8 语义，见基础类型篇）。

**适配器模式**：`http.HandlerFunc` 将「满足签名的普通函数」适配为 `http.Handler`，是标准库里小而典型的设计（第 8 章展开）。

---

## 2.10 GOPATH 与构建环境

现代项目以 **Go Modules** 为主；了解 **`GOPATH` 下 `src` / `pkg` / `bin`** 仍有助于阅读旧仓库与理解工具链缓存行为。

---

## 2.11 标准库（与网络编程强相关）

| 包 | 用途 |
|----|------|
| **`net`** | TCP/UDP、拨号、监听、解析等 |
| **`crypto` / `crypto/tls`** | 安全传输 |
| **`encoding/json` 等** | 与对端协定的数据表示 |

---

## 2.12 错误处理哲学

相对「异常栈一路弹出」，Go 选择在**调用点**处理 `error`：代码行数可能变多，但**失败路径显式、可组合**（`errors.Is` / `errors.As`、包装错误），更适合分布式里「错误是常态」的现实。

---

## 2.13 本章总结

切片与缓冲、`map`、指针、结构体 + 方法、**Goroutine**、**模块**与**显式错误**，是把第 1 章的分布式假设写进工程的主工具。  
下一章进入 **Socket**：**IP 与端口**成为端到端通信的钥匙。

**Legacy 深度笔记**：仍以 [chapter03 README](../chapter03-socket-programming/README.md) 中 `legacy-topic-index` 链接为主干。
