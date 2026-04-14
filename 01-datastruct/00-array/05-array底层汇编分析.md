# Go 数组底层汇编分析（可选）

本篇目标：用汇编验证两件事就够了：

1. **数组按值传参会拷贝**（大数组很贵）。  
2. **通过 `*[N]T` 访问元素**就是简单的地址计算 + 读写指令。  

> 不同 GOOS/GOARCH、Go 版本的指令不同；结论以当前工具链为准。

---

## 1. 怎么自己看

1. 写最小示例（如 `[3]int` 赋值、传参、取地址）。  
2. 运行：`go build -gcflags="-S -l" 2>&1`。  
3. 在输出中搜索函数名、或搜索 `MOV` / `LEA` 相关片段。  

---

## 2. 最小示例（建议）

```go
package main

func access(a *[3]int) int {
	return a[1]
}

func passByValue(a [1024]byte) int {
	return int(a[0])
}

func main() {}
```

你通常能在汇编里观察到：

- `access`：计算 `a + 1*elemSize` 后读取。  
- `passByValue`：函数入口附近有一段把参数拷到栈上的动作（大数组更明显）。  

---

## 3. 延伸阅读

- [00-array/04-array常见坑与最佳实践.md](./04-array常见坑与最佳实践.md)  
- [01-slice/02-slice扩容.md](../01-slice/02-slice扩容.md)
