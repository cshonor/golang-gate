# 同时声明一组变量和常量

## 1. 使用 var (...) 同时声明一组变量

### 语法
```go
var (
    变量名1 类型1 = 值1
    变量名2 类型2 = 值2
    变量名3 类型3 = 值3
    ...
)
```

### 特点
- 使用 `var` 关键字 + 小括号 `()`
- 可以同时声明多个变量
- 每个变量占一行，更清晰
- 比分开声明更简洁

### 示例1：指定类型
```go
var (
    earthWeight float64 = 164.0  // 地球体重
    earthAge    int     = 41     // 地球年龄
    name        string  = "万正鹏" // 姓名
    isStudent   bool    = true   // 是否是学生
)
```

### 示例2：自动推断类型
```go
var (
    weight = 70.0    // 自动推断为 float64
    height = 175.0   // 自动推断为 float64
    age    = 25      // 自动推断为 int
    city   = "北京"  // 自动推断为 string
)
```

### 示例3：混合使用
```go
var (
    score    float64 = 95.5  // 指定类型
    grade    = "A"           // 自动推断为 string
    passed   = true          // 自动推断为 bool
    attempts int    = 3      // 指定类型
)
```

## 2. 使用 const (...) 同时声明一组常量

### 语法
```go
const (
    常量名1 = 值1
    常量名2 = 值2
    常量名3 = 值3
    ...
)
```

### 特点
- 使用 `const` 关键字 + 小括号 `()`
- 可以同时声明多个常量
- 每个常量占一行，更清晰
- 比分开声明更简洁

### 示例
```go
const (
    lightSpeed      = 3e8        // 光速：3×10^8 米/秒
    marsGravity     = 0.3783     // 火星重力系数
    earthGravity    = 9.8        // 地球重力
    pi              = 3.14159    // 圆周率
    earthDaysPerYear = 365       // 地球一年天数
    marsDaysPerYear  = 687       // 火星一年天数
)
```

## 3. 对比：分开声明 vs 括号声明

### 方式1：分开声明（繁琐）
```go
var a int = 10
var b string = "hello"
var c float64 = 3.14
```

### 方式2：括号声明（简洁）
```go
var (
    a int     = 10
    b string  = "hello"
    c float64 = 3.14
)
```

### 优点
- **更简洁**：不需要重复写 `var` 或 `const`
- **更清晰**：所有变量/常量集中在一起，一目了然
- **更易维护**：修改或添加变量/常量更方便

## 4. 实际应用示例

### 火星计算程序
```go
package main

import "fmt"

func main() {
    // 使用 const (...) 声明所有常量
    const (
        marsGravity     = 0.3783
        earthDaysPerYear = 365
        marsDaysPerYear  = 687
    )
    
    // 使用 var (...) 声明所有变量
    var (
        earthWeight = 164.0
        earthAge    = 41
    )
    
    // 计算
    marsWeight := earthWeight * marsGravity
    marsAge := earthAge * earthDaysPerYear / marsDaysPerYear
    
    fmt.Println("地球体重:", earthWeight, "磅")
    fmt.Println("火星体重:", marsWeight, "磅")
    fmt.Println("地球年龄:", earthAge, "年")
    fmt.Println("火星年龄:", marsAge, "年")
}
```

## 5. 注意事项

### var (...) 的注意事项
1. **可以混合使用**：部分变量指定类型，部分自动推断
2. **可以重新赋值**：变量声明后可以重新赋值
3. **作用域**：括号内的变量在同一个作用域内

### const (...) 的注意事项
1. **值不能改变**：常量一旦声明，值就不能再改变
2. **可以省略类型**：Go 会自动推断类型
3. **可以用于 iota**：在常量组中可以使用 `iota` 计数器（进阶内容）

## 6. 常见用法

### 场景1：相关变量/常量分组
```go
// 物理常量
const (
    lightSpeed = 3e8
    gravity    = 9.8
    pi         = 3.14159
)

// 用户信息
var (
    name  = "万正鹏"
    age   = 25
    city  = "北京"
)
```

### 场景2：配置信息
```go
const (
    maxUsers    = 1000
    timeout     = 30
    defaultPort = 8080
)

var (
    serverName = "localhost"
    isDebug    = true
)
```

## 7. 总结

| 声明方式 | 语法 | 适用场景 |
|---------|------|---------|
| 单个变量 | `var name = "value"` | 只声明一个变量 |
| 一组变量 | `var (...) { ... }` | 同时声明多个相关变量 |
| 单个常量 | `const name = value` | 只声明一个常量 |
| 一组常量 | `const (...) { ... }` | 同时声明多个相关常量 |

### 关键点
1. **`var (...)`** 用于同时声明一组变量
2. **`const (...)`** 用于同时声明一组常量
3. **括号内每个变量/常量占一行**，更清晰
4. **比分开声明更简洁**，特别适合声明多个相关变量/常量

### 使用建议
- 当需要声明**多个相关**的变量或常量时，使用括号方式
- 当只声明**一个**变量或常量时，可以直接用 `var` 或 `const`
- 括号方式让代码更**清晰、易维护**

