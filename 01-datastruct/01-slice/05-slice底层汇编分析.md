# 05 - slice 底层汇编分析（可选）

本篇为**拓展**：通过反汇编观察 `slice` 头、`append` 等与机器码的对应关系，**非必读**。

---

## 1. 怎么自己看

1. 写最小示例（如 `s := make([]int, 0, 4); s = append(s, 1)`）。  
2. `go build -gcflags="-S -l" 2>&1`，在输出中筛选与 `slice`/`growslice` 相关片段。  
3. 或用 `go tool compile -S` 对单文件生成汇编。

不同 **GOOS/GOARCH**、**Go 版本** 指令不同，**结论以当前工具链为准**。

---

## 2. 通常能验证什么

- 切片头（指针、len、cap）在寄存器/栈上的传递方式。  
- `append` 可能走 **growslice** 路径（与 [02-slice扩容](./02-slice扩容.md) 对应）。

---

## 延伸阅读

- [02-slice扩容](./02-slice扩容.md)  
- [README.md](../README.md)
