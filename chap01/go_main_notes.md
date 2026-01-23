# Go 语言基础：main 函数与大括号规则

## 1. Package Main 的概念

### 什么是 package？
- `package` 是 Go 语言中组织代码的基本方式
- 一个文件夹里的所有 `.go` 文件必须声明同一个 `package`
- 文件夹名通常就是包名

### Package Main 的特殊性
- `package main` 表示这是一个可执行程序
- 只有包含 `main` 函数的 `main` 包才能被编译成可执行文件（如 Windows 下的 `.exe`）
- 其他包（如 `package utils`）是库包，不能直接运行，只能被其他包导入使用

### 与 JavaScript 的对比
- **JS**: `import` 可以导入单个文件（如 `import xxx from './a.js'`）或 npm 包
- **Go**: `package` 是按文件夹划分的，导入和使用都是以文件夹（包名）为单位
- **JS**: 导入时可以精确到单个文件
- **Go**: 同一个包内的文件可以直接互相调用，不用导入

## 2. Main 函数的规则

### Main 函数的作用
- `main` 函数是 Go 程序的**唯一入口**
- 程序运行时一定会从 `main` 包里的 `main` 函数开始执行
- 就像程序的"启动开关"

### Main 函数的规则
1. **只有 `main` 包才能有 `main` 函数**
   - 如果 `.go` 文件属于其他包（如 `package utils`），不能有 `main` 函数
   - 否则编译器会报错

2. **一个 `main` 包必须有且只有一个 `main` 函数**
   - 不能没有（编译时会提示"找不到入口"）
   - 不能有多个（会提示"main 函数重复定义"）

3. **`main` 函数可以放在 `main` 包的任意一个 `.go` 文件里**
   - 比如在 `playground.go`（`package main`）里写了 `func main()`
   - 同一个 `main` 包下的其他 `.go` 文件（如 `helper.go`）就不能再写 `main` 函数了
   - 但可以写其他辅助函数，`main` 函数可以直接调用它们

### 示例代码结构
```go
// playground.go
package main

import "fmt"

func main() {
    sayHi()  // 可以直接调用同包下的函数
}

// helper.go
package main

func sayHi() {
    fmt.Println("Hi!")
}
```

## 3. Go 语言的大括号规则

### 基本规则
Go 语言对大括号的摆放有**严格规定**，这是为了统一代码风格。

#### 规则总结：**"左不换行，右单独"**

1. **左大括号 `{` 必须与关键字在同一行**
   - `func main()` 后面的 `{` 必须和 `func main()` 放在同一行
   - `if`、`for` 等语句的左大括号也必须和关键字在同一行

2. **右大括号 `}` 必须单独占一行**
   - 右大括号不能和其他代码放在同一行
   - 必须单独占一行

3. **右大括号必须与对应的关键字对齐**
   - 右大括号要和它对应的关键字（如 `func`、`if`、`for`）的首字母对齐
   - 或者和左大括号所在行的开头对齐
   - 视觉上要在同一条竖线上

### 正确的写法
```go
func main() {
    fmt.Println("Hello")
}

if a > b {
    fmt.Println("a 大")
}

for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 错误的写法（会报错）
```go
// ❌ 错误：左大括号换行了
func main()
{
    fmt.Println("Hello")
}

// ❌ 错误：右大括号和其他代码在同一行
func main() {
    fmt.Println("Hello") }

// ❌ 错误：if 的左大括号换行了
if a > b
{
    fmt.Println("a 大")
}
```

### 嵌套语句的大括号规则
- **所有嵌套的语句（if、for、while 等）都必须遵守同样的规则**
- 不管是在函数里嵌套，还是在循环里嵌套条件判断，大括号规则都一样

#### 嵌套示例
```go
func main() {
    for i := 0; i < 5; i++ {
        if i > 2 {
            fmt.Println(i)
        } // 这个右括号跟 if 对齐
    } // 这个右括号跟 for 对齐
} // 这个右括号跟 func 对齐
```

### 格式化工具
- 使用 Go 官方的 `gofmt` 工具可以自动格式化代码
- 大多数编辑器（如 GoLand、VS Code）也会自动帮你对齐大括号
- 不用担心手动计算空格数

## 4. 完整的示例程序

```go
// 声明本代码所属的包
package main

// 导入 fmt (是 format 的缩写) 包,使其可用
import (
    "fmt"
)

// 声明一个名为 main 的函数
func main() {
    // 在屏幕上打印出 "Hello, playground"
    fmt.Println("Hello, playground")
}
```

### 代码解释
- `package main`: 声明这是 main 包，可以编译成可执行程序
- `import "fmt"`: 导入 fmt 包，用于输入输出
- `func main()`: 程序的入口函数，程序从这里开始执行
- `fmt.Println()`: 打印函数，输出内容到控制台
- 注意大括号的摆放：左大括号与 `func main()` 同一行，右大括号单独占一行

## 5. 总结

### Package Main
- 一个文件夹 = 一个包
- `main` 包是特殊的可执行包
- 其他包是库包，供导入使用

### Main 函数
- 程序的唯一入口
- 只有 `main` 包才能有 `main` 函数
- 一个 `main` 包必须有且只有一个 `main` 函数

### 大括号规则
- **左不换行**：左大括号必须与关键字同一行
- **右单独**：右大括号必须单独占一行
- **对齐**：右大括号与对应关键字对齐
- 所有语句（函数、if、for 等）都遵守这个规则

