# Go 语言 range 循环完整指南

## 1. 什么是 range？

`range` 是 Go 语言中用于遍历集合类型的关键字，可以遍历：
- **数组**（Array）
- **切片**（Slice）
- **映射**（Map）
- **字符串**（String）
- **通道**（Channel）

## 2. 空白标识符 `_` 的作用

在 Go 语言里，`_` 是一个**空白标识符**，它的作用是"接收但不使用"某个值，避免编译器报"变量未使用"的错误。

### 为什么需要 `_`？
- `range` 会返回**两个值**：索引/键 和 对应的值
- 如果你只需要其中一个，必须用 `_` 占位另一个位置
- 这样既符合 Go 语言的语法要求，也能让代码更清晰

```go
// ✅ 正确：用 _ 忽略索引
for _, value := range numbers {
    fmt.Println(value)
}

// ❌ 错误：不能只写一个变量
// for value := range numbers { ... }
```

## 3. 遍历数组和切片

### 3.1 只获取值（忽略索引）
```go
numbers := []int{1, 2, 3, 4, 5}
for _, value := range numbers {
    fmt.Print(value, " ")
}
// 输出: 1 2 3 4 5
```

### 3.2 获取索引和值
```go
numbers := []int{10, 20, 30, 40, 50}
for index, value := range numbers {
    fmt.Printf("索引=%d, 值=%d\n", index, value)
}
// 输出:
// 索引=0, 值=10
// 索引=1, 值=20
// 索引=2, 值=30
// 索引=3, 值=40
// 索引=4, 值=50
```

### 3.3 只获取索引（忽略值）
```go
numbers := []int{1, 2, 3, 4, 5}
for index := range numbers {
    fmt.Print(index, " ")
}
// 输出: 0 1 2 3 4
```

### 3.4 数组和切片的区别
- **数组**：长度固定，如 `[5]int{1, 2, 3, 4, 5}`
- **切片**：长度可变，如 `[]int{1, 2, 3, 4, 5}`
- 遍历方式完全相同

```go
// 数组
arr := [5]int{1, 2, 3, 4, 5}
for i, v := range arr {
    fmt.Printf("arr[%d]=%d\n", i, v)
}

// 切片
slice := []int{1, 2, 3, 4, 5}
for i, v := range slice {
    fmt.Printf("slice[%d]=%d\n", i, v)
}
```

## 4. 遍历 Map（映射）

### 4.1 获取键和值
```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Charlie": 92,
}

for name, score := range scores {
    fmt.Printf("%s 的分数是 %d\n", name, score)
}
// 输出顺序是随机的（map 是无序的）
```

### 4.2 只获取键（忽略值）
```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
}

for name := range scores {
    fmt.Println("姓名:", name)
}
```

### 4.3 只获取值（忽略键）
```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
}

for _, score := range scores {
    fmt.Println("分数:", score)
}
```

### 4.4 Map 遍历的特点
- **顺序随机**：每次遍历的顺序可能不同
- **并发安全**：在遍历过程中修改 map 会导致运行时错误
- **键值对**：返回的是键和值，不是索引

## 5. 遍历字符串

### 5.1 获取索引和字符（字节）
```go
text := "Hello"
for index, char := range text {
    fmt.Printf("索引=%d, 字符=%c, Unicode=%d\n", index, char, char)
}
// 输出:
// 索引=0, 字符=H, Unicode=72
// 索引=1, 字符=e, Unicode=101
// 索引=2, 字符=l, Unicode=108
// 索引=3, 字符=l, Unicode=108
// 索引=4, 字符=o, Unicode=111
```

### 5.2 只获取字符（忽略索引）
```go
text := "Go语言"
for _, char := range text {
    fmt.Printf("%c ", char)
}
// 输出: G o 语 言
```

### 5.3 字符串遍历的特点
- **按 Unicode 字符遍历**：不是按字节
- **中文支持**：可以正确处理中文字符
- **索引是字节位置**：第一个返回值是字节索引，不是字符索引

```go
text := "Go语言"
for i, char := range text {
    fmt.Printf("字节索引=%d, 字符=%c\n", i, char)
}
// 输出:
// 字节索引=0, 字符=G
// 字节索引=1, 字符=o
// 字节索引=2, 字符=语  (注意：这是第2个字节，不是第2个字符)
// 字节索引=5, 字符=言
```

## 6. 常见用法总结

### 6.1 快速参考表

| 数据类型 | 返回值 | 只获取值 | 只获取索引/键 | 获取索引/键和值 |
|---------|--------|---------|--------------|----------------|
| **数组/切片** | `(index, value)` | `for _, v := range arr` | `for i := range arr` | `for i, v := range arr` |
| **Map** | `(key, value)` | `for _, v := range m` | `for k := range m` | `for k, v := range m` |
| **字符串** | `(index, rune)` | `for _, char := range s` | `for i := range s` | `for i, char := range s` |

### 6.2 实际应用示例

#### 示例1：统计字符出现次数
```go
text := "hello"
count := make(map[rune]int)
for _, char := range text {
    count[char]++
}
fmt.Println(count)  // map[101:1 104:1 108:2 111:1]
```

#### 示例2：查找最大值
```go
numbers := []int{3, 7, 2, 9, 1}
max := numbers[0]
for _, value := range numbers[1:] {
    if value > max {
        max = value
    }
}
fmt.Println("最大值:", max)  // 最大值: 9
```

#### 示例3：过滤元素
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var evens []int
for _, num := range numbers {
    if num%2 == 0 {
        evens = append(evens, num)
    }
}
fmt.Println("偶数:", evens)  // 偶数: [2 4 6 8 10]
```

#### 示例4：遍历 map 并排序键
```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Charlie": 92,
}

// 先收集所有键
var names []string
for name := range scores {
    names = append(names, name)
}

// 排序键
sort.Strings(names)

// 按排序后的键遍历
for _, name := range names {
    fmt.Printf("%s: %d\n", name, scores[name])
}
```

## 7. 注意事项

### 7.1 不能省略两个返回值
```go
// ❌ 错误：range 必须返回两个值
// for value := range numbers { ... }

// ✅ 正确：用 _ 忽略不需要的值
for _, value := range numbers { ... }
```

### 7.2 遍历时修改集合
- **数组/切片**：可以修改元素值，但不能修改长度
- **Map**：遍历时不能修改 map（会导致运行时错误）
- **字符串**：字符串是不可变的，不能修改

```go
// ✅ 可以修改切片元素的值
numbers := []int{1, 2, 3}
for i := range numbers {
    numbers[i] *= 2  // 可以
}
fmt.Println(numbers)  // [2 4 6]

// ❌ 不能在遍历时修改 map
// scores := map[string]int{"Alice": 95}
// for k := range scores {
//     delete(scores, k)  // 运行时错误！
// }
```

### 7.3 性能考虑
- `range` 遍历是**值拷贝**，对于大结构体可能影响性能
- 如果只需要索引，用 `for i := range arr` 比 `for i, _ := range arr` 更高效
- 对于大切片，考虑使用传统 for 循环通过索引访问

## 8. 与其他语言的对比

| 语言 | 语法 | 说明 |
|-----|------|------|
| **Go** | `for i, v := range arr` | 必须接收两个返回值 |
| **Python** | `for i, v in enumerate(arr)` | 可选接收索引 |
| **JavaScript** | `arr.forEach((v, i) => ...)` | 回调函数形式 |
| **Java** | `for (int v : arr)` | 只能获取值 |

## 9. 总结

### 核心要点
1. **`range` 总是返回两个值**：索引/键 和 值
2. **用 `_` 忽略不需要的值**：避免"变量未使用"错误
3. **支持多种集合类型**：数组、切片、map、字符串、通道
4. **遍历是值拷贝**：修改不会影响原集合（除非是指针）

### 常用模式
- 只关心值：`for _, v := range collection`
- 只关心索引/键：`for i := range collection`
- 两者都需要：`for i, v := range collection`

### 最佳实践
- 明确表达意图：用 `_` 明确表示"我不需要这个值"
- 注意性能：大结构体考虑用索引访问
- 注意并发：遍历 map 时不要修改 map

