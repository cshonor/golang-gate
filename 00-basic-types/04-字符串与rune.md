# Go 字符串 `string` 与 `rune`（强化版）

## 1. `string` 本质（先把模型摆正）

- `string` **不是** `[]byte` 这种切片类型；它的语义是：**只读的 UTF-8 字节序列**（独立类型）。
- 常见实现里可以把它想成一个小头部：**数据指针 + 长度**（只读指向一段字节）。
- **内容不可原地修改**；想“改字符串”，正确路径是：**解码 → 在新缓冲区里改 → 再构造新的 `string`**。

### 最关键一句：`string` ↔ `[]byte` 往往会**复制字节**

```go
s := "abc"
b := []byte(s) // 通常会分配 slice 并把 UTF-8 字节复制一份
b[0] = 'x'
s2 := string(b) // 产生新的 string；原来的 s 仍是 "abc"
```

所以口语总结很准确：**`string` 本身不可变，谁也改不了“旧的那一份”的语义内容**（只能得到新串）。

补充：`string([]byte{...})` / `[]byte(s)` 在少数场景下编译器可能做优化，但**心智模型**仍建议按「**可能要复制**」理解，尤其在 `s` 很长、或你要依赖“是否共内存”时别赌优化。

## 2. UTF-8 与 `len`

- Go 源码与 `string` 默认按 **UTF-8** 存字节。
- 常见汉字在 UTF-8 里多是 **3 字节**（也有例外，别死记成全世界都 3）。

**千万记住：`len(s)` 统计的是「字节数」，不是「字符数 / Unicode 码点数」。**

```go
s := "Go语言"
// G(1) + o(1) + 语(3) + 言(3) = 8 字节（不是 6）
_ = len(s) // 8

s2 := "中国"
_ = len(s2) // 6：两个汉字常见各 3 字节
```

- 用**下标** `s[i]`、`s[i:j]` 走的是**字节下标**，不是“第 i 个字符”。

## 3. `rune` = `int32`（Unicode 码点）

- `rune` 是 `int32` 的别名，表示 **Unicode 码点（code point）**。
- 用 `for i, c := range s` 遍历 `string` 时，`c` 的类型是 `rune`；`i` 是该 `rune` 在 UTF-8 里的**起始字节偏移**（不是“第几个字符”的整数下标）。

工程上常说：**要做“按字符（码点）”的处理，用 `rune` / `range`，别用裸下标。**

一句进阶真话：**用户眼里“一个字符”（字素簇 grapheme cluster）有时由多个码点组成**（组合符号、部分 emoji 序列等）。Go 的 `rune` 粒度是**码点**，不是排版引擎里的“视觉字符”。

## 4. 最容易混淆的两行（字节 vs 码点）

```go
s := "Go 语言"

fmt.Println(s[0])        // 字节值（uint8）：71（'G' 的 ASCII）
fmt.Printf("%c\n", s[0]) // 按“单字节字符”打印：G

// range 按 UTF-8 解码成 rune（码点）遍历
for _, c := range s {
	fmt.Printf("%c\n", c) // G、o、空格、语、言
}
```

## 5. 常见坑（面试高频）

### 坑 1：`string` 不能原地改

```go
s := "hello"
// s[0] = 'x' // 编译报错：string 元素不可赋值
```

### 坑 2：`len(string)` 是字节长度

```go
s := "中国"
len(s) // 6（不是 2）
```

### 坑 3：想要「码点数量」别盲信 `len`

```go
s := "中国"
len([]rune(s))              // 2：分配一个 rune slice
utf8.RuneCountInString(s)   // 2：通常更省（不建 []rune）
```

## 6. 一张速记对照（背这个就够用）

```
string   → UTF-8 字节序列（只读语义）
[]byte   → 可写字节序列（自己管容量/底层数组）
rune     → Unicode 码点（int32）

len(s)                    → 字节数
utf8.RuneCountInString(s) → 码点数（常用）
[]rune(s)                 → 码点切片（可能要分配）
range s                   → 按码点遍历（同时给你字节偏移 i）
```

## 7. 一句话终极总结

**`string` 是只读 UTF-8 字节流：下标/`len` 都在跟字节打交道；`rune`（以及 `range`）才是按 Unicode 码点“像字符一样”处理的主路径。**

## 8. 可运行小抄：遍历 / 长度 / 修改路径 / 中文

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Go语言"
	fmt.Println("bytes:", len(s), "runes:", utf8.RuneCountInString(s), "runes(slice):", len([]rune(s)))

	// 字节视角
	fmt.Printf("s[0]=%d %c\n", s[0], s[0])

	// 码点视角
	for i, c := range s {
		fmt.Printf("i=%d c=%U %c\n", i, c, c)
	}

	// “修改”：复制到 []rune，改完再 string
	rs := []rune(s)
	rs[2] = '文'
	s2 := string(rs)
	fmt.Println("after:", s2)
}
```
