# 第三章：条件判断和循环

## 1. 布尔类型和比较运算符

### 1.1 布尔类型变量
- 布尔变量可以直接用 `true` 或 `false` 赋值
- Go 会自动推断类型为 `bool`
- 也可以显式声明：`var isSunny bool = true`

```go
var walkOutside = true
var takeTheBluePill = false
```

### 1.2 比较运算符
- `==` 相等
- `!=` 不等
- `<` 小于
- `>` 大于
- `<=` 小于等于
- `>=` 大于等于

**特点**：
- 比较运算符的结果是布尔类型（`true` 或 `false`）
- 可以用变量接收比较结果
- 例如：`var isEqual = (command == "go east")`

### 1.3 字符串比较
- 字符串比较是按字符的 Unicode 编码值逐个对比
- 例如：`"apple" < "banana"` 因为 `'a' (97) < 'b' (98)`

### 1.4 类型严格性
- **Go 要求 `==` 和 `!=` 运算符两边的操作数必须是相同类型**
- 字符串和数值不能直接比较
- 必须先转换类型：
  - 字符串转整数：`strconv.Atoi("123")`
  - 整数转字符串：`strconv.Itoa(123)`

```go
// ❌ 错误：不能比较 string 和 int
// var result = ("123" == 123)

// ✅ 正确：先转换类型
num, _ := strconv.Atoi("123")
var result = (num == 123)
```

### 1.5 布尔变量的运算
- **逻辑与**：`&&`（两个 `&`）
- **逻辑或**：`||`（两个 `|`）
- **逻辑非**：`!`
- **布尔异或**：用 `!=` 实现（`a != b`）

```go
var isSunny = true
var hasUmbrella = false

!isSunny                    // 取反
isSunny && hasUmbrella      // 与运算
isSunny || hasUmbrella      // 或运算
isSunny && !hasUmbrella     // 组合运算
```

## 2. if 语句

### 2.1 基本语法
```go
if 条件 {
    // 代码块
} else if 条件 {
    // 代码块
} else {
    // 代码块
}
```

**特点**：
- 条件**不需要小括号**
- **左大括号必须和条件在同一行**
- 右大括号必须单独占一行

### 2.2 逻辑运算符
- `&&` 逻辑与
- `||` 逻辑或
- `!` 逻辑非

**优先级**：`&&` 比 `||` 优先级高，不确定时用括号

```go
if (command == "go east" || command == "go west") && !hasUmbrella {
    // 代码
}
```

### 2.3 在 if 开头声明变量
```go
if age := 25; age >= 18 {
    fmt.Println("成年人", age)
} else {
    fmt.Println("未成年人", age)
}
// 出了 if-else，age 就不能用了
```

## 3. switch 语句

### 3.1 基本语法
```go
switch 变量 {
case 值1, 值2:
    // 代码块
case 值3:
    // 代码块
default:
    // 代码块
}
```

### 3.2 重要特点

#### 1. case 的值类型必须和 switch 变量类型一致
- **字符串**：case 里用双引号，如 `case "go east", "go west":`
- **整数**：case 里直接写数字，如 `case 1, 2, 3:`
- **布尔**：case 里写 `true` 或 `false`，如 `case true:`

```go
// 字符串类型
switch command {
case "go east", "go west":  // 字符串用双引号
    // ...
}

// 整数类型
switch num {
case 1, 2, 3:  // 整数不用引号
    // ...
}
```

#### 2. 多个值用逗号隔开
```go
case "go east", "go west", "go north":
    // 只要匹配其中任意一个值，就执行这个 case
```

#### 3. 不需要 break
- Go 的 switch **默认每个 case 执行完后自动 break**
- 不需要像 C 或 Java 那样手动写 `break`
- 除非使用 `fallthrough`（很少用）

#### 4. default 的作用
- 当所有 case 都不匹配时，会执行 `default`
- `default` 的位置不一定要放在最后，但习惯上放在最后

### 3.3 标准输入是字符串类型
- 通过标准输入流接收到的内容，**默认都是字符串类型**
- 即使输入的是 `123`，读进来的也是 `"123"` 这个字符串
- 如果要在 switch 中使用，case 里要用字符串 `"123"`
- 如果想当整数用，需要先用 `strconv.Atoi` 转换

### 3.4 在 switch 开头声明变量
```go
switch month := rand.Intn(12) + 1; month {
case 1:
    fmt.Println("一月")
default:
    fmt.Println("其他月份", month)
}
// 出了 switch，month 就不能用了
```

### 3.5 无表达式 switch
```go
switch num := rand.Intn(10); {
case num < 3:
    fmt.Println("小数字")
case num < 7:
    fmt.Println("中等数字")
default:
    fmt.Println("大数字")
}
```

## 4. for 循环

### 4.1 Go 没有 while 关键字
- Go 里**没有专门的 while 关键字**
- 用 `for` + 条件 代替 `while` 的功能
- 例如：`for count > 0 { ... }` 相当于 `while (count > 0) { ... }`

### 4.2 for 循环的三种形式

#### 1. 类似 while（只写条件）
```go
count := 10
for count > 0 {
    fmt.Println(count)
    count--  // 在循环体内修改条件
}
```

#### 2. 传统 for 循环（三个部分）
```go
// 格式：for 初始化; 条件; 后续操作 { ... }
// 注意：不用小括号
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

#### 3. 无限循环
```go
for {
    // 代码
    if 条件 {
        break  // 退出循环
    }
}
```

### 4.3 遍历数组/切片
```go
numbers := []int{1, 2, 3, 4, 5}

// 只获取值
for _, value := range numbers {
    fmt.Println(value)
}

// 获取索引和值
for index, value := range numbers {
    fmt.Printf("index=%d, value=%d\n", index, value)
}
```

### 4.4 continue 和 break
- `continue`：跳过本次循环，继续下一次
- `break`：退出循环

## 5. 总结

### 布尔类型
- 布尔变量：`var isSunny = true`
- 比较运算符：`==`, `!=`, `<`, `>`, `<=`, `>=`
- 逻辑运算符：`&&`, `||`, `!`
- 类型严格：字符串和数值不能直接比较

### if 语句
- 条件不需要小括号
- 左大括号必须和条件同一行
- 逻辑运算符：`&&`, `||`, `!`

### switch 语句
- case 值类型必须和 switch 变量类型一致
- 多个值用逗号隔开
- 不需要 break（自动 break）
- default 是兜底选项

### for 循环
- 没有 while 关键字，用 `for` + 条件代替
- 三种形式：只写条件、三个部分、无限循环
- 可以遍历数组/切片

