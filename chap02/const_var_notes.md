# 常量和变量

## 1. 常量（const）

### 定义方式
```go
const 常量名 = 值
```

### 特点
- 使用 `const` 关键字定义
- **值一旦定义就不能再改变**
- 如果尝试重新赋值，编译器会报错

### 示例
```go
const lightSpeed = 3e8        // 光速：3×10^8 米/秒
const marsGravity = 0.3783   // 火星重力系数
const pi = 3.14159           // 圆周率

// ❌ 错误：不能给常量重新赋值
// lightSpeed = 4e8  // 编译错误
```

### 多个常量定义
```go
const (
    gravityEarth = 9.8
    gravityMars  = 3.7
    gravityMoon  = 1.6
)
```

## 2. 变量（var）

### 定义方式1：使用 var 关键字
```go
var 变量名 类型 = 值
```

### 定义方式2：简化定义（推荐）
```go
变量名 := 值  // Go 会自动判断类型
```

### 特点
- 使用 `var` 关键字或 `:=` 定义
- **值可以随时改变**
- 可以重新赋值

### 示例
```go
// 方式1：指定类型
var distance float64 = 5.04e10
var earthWeight float64 = 164.0

// 方式2：自动判断类型（推荐）
distance2 := 5.04e10  // 自动判断为 float64
age := 25             // 自动判断为 int
name := "万正鹏"       // 自动判断为 string

// 变量可以重新赋值
earthWeight = 160.0  // ✅ 正确
```

### 多个变量定义
```go
var (
    weight = 70.0
    height = 175.0
    age    = 25
)
```

## 3. const vs var 对比

| 特性 | const（常量） | var（变量） |
|------|--------------|------------|
| 关键字 | `const` | `var` 或 `:=` |
| 值能否改变 | ❌ 不能 | ✅ 能 |
| 适用场景 | 固定不变的值（如 π、重力系数） | 可能变化的值（如体重、年龄） |

## 4. 在 fmt.Println 中使用常量和变量

### 可以直接进行运算
```go
const lightSpeed = 3e8
var distance = 5.04e10

// 在 Println 中直接计算
fmt.Println(distance/lightSpeed, "seconds")
// 输出：168 seconds
```

### 运算规则
- 常量和变量只要类型兼容（比如都是数字类型），就能直接进行加减乘除运算
- Go 会先算出结果，再打印出来

### 示例
```go
const marsGravity = 0.3783
var earthWeight = 164.0

// 直接运算
fmt.Println("火星体重:", earthWeight*marsGravity, "磅")

// 多个运算
fmt.Println("计算结果:", 10+5, "和", 20-3)
```

## 5. fmt.Println 自动加空格

### 重要特性
**`fmt.Println` 会在括号里的每个参数之间自动加一个空格。**

### 示例
```go
fmt.Println(168, "seconds")
// 输出：168 seconds（中间有空格）

fmt.Println("火星", "体重：", 62.04)
// 输出：火星 体重： 62.04（每个参数之间都有空格）
```

### 对比：fmt.Print vs fmt.Println

| 函数 | 是否自动加空格 | 是否自动换行 |
|------|--------------|------------|
| `fmt.Print` | ❌ 否 | ❌ 否 |
| `fmt.Println` | ✅ 是 | ✅ 是 |

### 示例对比
```go
fmt.Print(168, "seconds\n")      // 输出：168seconds（没有空格，需要手动换行）
fmt.Println(168, "seconds")       // 输出：168 seconds（有空格，自动换行）
```

## 6. 实际应用示例

### 光速计算
```go
package main

import "fmt"

func main() {
    // 常量
    const lightSpeed = 3e8  // 光速：3×10^8 米/秒
    
    // 变量
    var distance = 5.04e10  // 地球到太阳的距离：5.04×10^10 米
    
    // 在 Println 中直接计算并打印
    fmt.Println(distance/lightSpeed, "seconds")
    // 输出：168 seconds
}
```

### 火星计算
```go
package main

import "fmt"

func main() {
    // 常量
    const marsGravity = 0.3783
    const earthDaysPerYear = 365
    const marsDaysPerYear = 687
    
    // 变量
    var earthWeight = 164.0
    var earthAge = 41
    
    // 计算并打印
    fmt.Println("火星体重:", earthWeight*marsGravity, "磅")
    fmt.Println("火星年龄:", earthAge*earthDaysPerYear/marsDaysPerYear, "年")
}
```

## 7. 总结

### 常量（const）
- 使用 `const` 关键字定义
- 值不能改变
- 适合固定不变的值

### 变量（var）
- 使用 `var` 或 `:=` 定义
- 值可以改变
- 适合可能变化的值

### fmt.Println 的特性
1. **可以直接进行运算**：常量和变量可以在 Println 中直接运算
2. **自动加空格**：每个参数之间会自动加一个空格
3. **自动换行**：打印完后自动换行

### 关键点
- 在 `fmt.Println` 中，变量和常量的运算结果会自动和字符串拼接
- 拼接时，每个参数之间会自动加一个空格
- 这比 `fmt.Printf` 更简单，适合简单的输出场景

