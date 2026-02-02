# Go 语言变量作用域

## 1. 作用域从大到小

### 1. 跨包全局变量（最大作用域）
- **定义**：首字母大写的全局变量
- **作用域**：可以被其他包通过 `import` 引入后访问
- **示例**：
```go
var PublicVar int = 100  // 首字母大写，可以被其他包访问
```

### 2. 包内全局变量
- **定义**：首字母小写的全局变量，定义在函数外面
- **作用域**：整个包内的所有 `.go` 文件都可以使用
- **特点**：
  - 同一个 `package` 下的所有 `.go` 文件都能访问
  - 不需要加文件名前缀（如 `scope_rules.year`）
  - 直接使用变量名即可
- **示例**：
```go
package main

var year = 2024  // 包内全局变量

func main() {
    fmt.Println(year)  // 可以直接使用
}

func otherFunc() {
    fmt.Println(year)  // 也可以使用
}
```

### 3. 函数作用域
- **定义**：函数内部定义的变量和函数参数
- **作用域**：仅限函数内部
- **包括**：
  - 函数内用 `var` 或 `:=` 声明的变量
  - 函数参数
- **示例**：
```go
func test(value int) {  // value 是函数参数，作用域在函数内
    var local = "局部变量"  // local 作用域在函数内
    fmt.Println(value, local)
}
// 出了函数，value 和 local 都不能用了
```

### 4. 块作用域（最小作用域）
- **定义**：大括号 `{}` 内定义的变量
- **作用域**：仅限当前块内部
- **包括**：
  - `if` 语句块
  - `for` 循环块
  - `switch` 语句块
  - 直接使用 `{}` 创建的块
- **示例**：
```go
if age := 25; age >= 18 {
    // age 只能在这个 if 块内使用
}
// 出了 if，age 就不能用了
```

## 2. 在控制语句开头声明变量

### 核心思想
**在 `for`、`if`、`switch` 的开头声明变量，可以把变量的作用域限制在当前控制语句内，避免污染外部作用域。**

### 2.1 for 循环
```go
// 在 for 开头声明变量
for count := 10; count > 0; count-- {
    fmt.Println(count)
    // count 在循环条件、迭代表达式、循环体内都可以使用
}
// 出了 for 循环，count 就不能用了
```

**作用域范围**：
- 初始化表达式（`count := 10`）
- 条件判断（`count > 0`）
- 迭代表达式（`count--`）
- 循环体 `{}` 内的代码

### 2.2 if 语句
```go
// 在 if 开头声明变量
if age := 25; age >= 18 {
    fmt.Println("成年人", age)
} else {
    fmt.Println("未成年人", age)
}
// 出了 if-else，age 就不能用了
```

**作用域范围**：
- `if` 条件判断（`age >= 18`）
- `if` 块内的代码
- `else` 块内的代码

### 2.3 switch 语句
```go
// 在 switch 开头声明变量
switch month := rand.Intn(12) + 1; month {
case 1:
    fmt.Println("一月", month)
case 2:
    fmt.Println("二月", month)
default:
    fmt.Println("其他月份", month)
}
// 出了 switch，month 就不能用了
```

**语法说明**：
- `month := rand.Intn(12) + 1`：在 switch 开头声明并初始化变量
- 分号 `;`：分隔变量声明和 switch 表达式
- `month`：根据这个变量的值进行匹配
- `case` 分支内可以使用 `month`

**无表达式 switch**：
```go
// switch 后面直接跟分号和大括号
switch num := rand.Intn(10); {
case num < 3:
    fmt.Println("小数字", num)
case num < 7:
    fmt.Println("中等数字", num)
default:
    fmt.Println("大数字", num)
}
```

## 3. 包内全局变量的使用

### 3.1 什么是"整个包"
- **整个包** = 同一个 `package` 下的所有 `.go` 文件
- 比如 `package main` 下的 `main.go` 和 `utils.go` 都属于同一个包
- 包内全局变量可以在这些文件之间共享

### 3.2 如何使用
```go
// main.go
package main

var year = 2024  // 包内全局变量

func main() {
    fmt.Println(year)  // 直接使用，不需要前缀
}

// utils.go
package main

func test() {
    fmt.Println(year)  // 也可以直接使用，不需要 main.year
}
```

### 3.3 避免重名
- **编译期检查**：Go 编译器会在编译时检查重名问题
- **同一作用域不能重名**：同一个包内不能有两个同名的全局变量
- **不同作用域可以重名**：
  ```go
  var year = 2024  // 全局变量
  
  func test() {
      year := 2025  // 局部变量，遮蔽全局的 year
      fmt.Println(year)  // 输出：2025（使用局部的）
  }
  
  fmt.Println(year)  // 输出：2024（使用全局的）
  ```

## 4. 变量遮蔽（Shadowing）

### 定义
内层作用域的变量会"遮蔽"外层作用域的同名变量。

### 示例
```go
year := 2023  // 外层变量

if true {
    year := 2024  // 内层变量，遮蔽外层的 year
    fmt.Println(year)  // 输出：2024（内层的）
}

fmt.Println(year)  // 输出：2023（外层的，没被改变）
```

## 5. 垃圾回收（GC）

### 自动回收
- Go 语言有**自动垃圾回收机制**
- 当变量离开作用域，并且没有任何引用时，GC 会自动清理内存
- 不需要手动管理内存

### 示例
```go
func test() {
    var x = 10  // 局部变量
    fmt.Println(x)
}  // 函数结束，x 离开作用域，GC 会自动清理
```

## 6. 作用域的好处

### 6.1 避免变量污染
```go
// 在不同的控制语句中使用同名变量，不会冲突
for i := 0; i < 3; i++ {
    // 第一个 i
}

for i := 0; i < 3; i++ {
    // 第二个 i（和第一个不冲突）
}

if i := 5; i > 0 {
    // 第三个 i（和前两个不冲突）
}
```

### 6.2 代码更清晰
- 变量作用域明确，不会意外访问到不应该访问的变量
- 临时变量用完就"消失"，不会污染外部作用域

### 6.3 减少内存占用
- 变量离开作用域后可以被 GC 回收
- 不需要等到整个函数或程序结束

## 7. 总结

### 作用域从大到小
1. **跨包全局变量**（首字母大写）
2. **包内全局变量**（首字母小写，整个包可用）
3. **函数作用域**（函数内和函数参数）
4. **块作用域**（`if`、`for`、`switch`、`{}`）

### 在控制语句开头声明变量
- **for**：`for count := 10; count > 0; count-- { ... }`
- **if**：`if age := 25; age >= 18 { ... }`
- **switch**：`switch month := rand.Intn(12)+1; month { ... }`

### 关键点
1. **大括号对应一个作用域**
2. **内层作用域可以访问外层作用域的变量**
3. **内层作用域的同名变量会遮蔽外层变量**
4. **变量离开作用域后会被 GC 自动回收**
5. **编译器会在编译期检查重名问题**


