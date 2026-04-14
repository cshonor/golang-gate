# Go 字符串 string 与 rune

## 1. string 本质

- 是**只读字节切片**
- 底层结构：指针 + 长度
- 内容不可修改，要改必须转 []byte

## 2. UTF-8 编码

- Go 字符串默认 UTF-8
- 一个中文占 3 字节
- 按 index 遍历是**字节遍历**，不是字符遍历

## 3. rune = int32

- 代表一个 Unicode 字符
- range 遍历 string 得到的是 rune

## 4. 示例

```go
s := "Go 语言"
fmt.Println(s[0]) // 字节 'G'

for _, c := range s {
    fmt.Printf("%c\n", c) // 正确遍历字符
}
```

## 5. 常见坑

- 直接 s[i] 修改会编译报错
- 中文长度 len(s) 是字节数，不是字符数
