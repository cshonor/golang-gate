// 示例：Go 语言的方法（Method）
// 演示方法声明、接收者类型、值接收者与指针接收者的区别

package main

import "fmt"

// ============================================
// 温度类型定义
//Go 的 type 有两种完全不同的用法：
 
//1. 类型别名：只是换个名字，本质还是同一个类型
​
//2. 定义新类型：全新的类型，和原来不是一个东西

//- 有  =  → 别名（同一个类型）
​
//- 没  =  → 新类型（不同类型） ============================================

// Kelvin 开尔文温度类型
type Kelvin float64

// Celsius 摄氏度类型
type Celsius float64

// Fahrenheit 华氏度类型
type Fahrenheit float64

// ============================================
// 温度类型的方法
// ============================================

// ToCelsius 开尔文转摄氏度（值接收者）
func (k Kelvin) ToCelsius() Celsius {
	return Celsius(k - 273.15)
//Celsius(xxx)  不是真的“构造函数”，Go 没有构造函数这个语法  显式类型转换（type conversion）
}

// ToKelvin 摄氏度转开尔文（值接收者）
func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c + 273.15)
}

// ToFahrenheit 摄氏度转华氏度（值接收者）
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*1.8 + 32)
}

// ToCelsius 华氏度转摄氏度（值接收者）
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) / 1.8)
}

func main() {
	// ============================================
	// 1. 方法与函数的区别
	// ============================================
	fmt.Println("=== 方法与函数的区别 ===")

	fmt.Println("特性对比:")
	fmt.Println("  特性          函数（Function）          方法（Method）")
	fmt.Println("  定义方式      func 名(参数)            func (接收者) 名(参数)")
	fmt.Println("  绑定关系      独立，不绑定类型         必须绑定到接收者类型")
	fmt.Println("  调用方式      函数名(参数)             接收者.方法名(参数)")
	fmt.Println("  核心作用      通用代码复用             为特定类型增加行为")

	fmt.Println()

	// ============================================
	// 2. 声明新类型
	// ============================================
	fmt.Println("=== 声明新类型 ===")

	// 基于已有类型创建自定义类型
	var k Kelvin = 300.0
	var c Celsius = 26.85
	var f Fahrenheit = 80.0

	fmt.Printf("Kelvin: %.2f K\n", k)
	fmt.Printf("Celsius: %.2f °C\n", c)
	fmt.Printf("Fahrenheit: %.2f °F\n", f)

	fmt.Println()

	// ============================================
	// 3. 温度转换：函数转方法
	// ============================================
	fmt.Println("=== 温度转换：函数转方法 ===")

	// 使用方法的温度转换
	tempK := Kelvin(300.0)
	tempC := tempK.ToCelsius()
	fmt.Printf("%.2f K = %.2f °C\n", tempK, tempC)

	tempC2 := Celsius(26.85)
	tempK2 := tempC2.ToKelvin()
	fmt.Printf("%.2f °C = %.2f K\n", tempC2, tempK2)

	// 摄氏度转华氏度
	tempC3 := Celsius(0.0)
	tempF := tempC3.ToFahrenheit()
	fmt.Printf("%.2f °C = %.2f °F\n", tempC3, tempF)

	// 华氏度转摄氏度
	tempF2 := Fahrenheit(32.0)
	tempC4 := tempF2.ToCelsius()
	fmt.Printf("%.2f °F = %.2f °C\n", tempF2, tempC4)

	fmt.Println()

	// ============================================
	// 4. 值接收者 vs 指针接收者
	// ============================================
	fmt.Println("=== 值接收者 vs 指针接收者 ===")

	// 值接收者：会复制数据，不能修改原始值
	counter1 := Counter{value: 10}
	fmt.Printf("原始值: %d\n", counter1.value)
	counter1.IncrementByValue() // 值接收者，不会修改原始值
	fmt.Printf("调用值接收者方法后: %d (未改变)\n", counter1.value)

	// 指针接收者：可以修改原始值
	counter2 := Counter{value: 10}
	fmt.Printf("原始值: %d\n", counter2.value)
	counter2.IncrementByPointer() // 指针接收者，会修改原始值
	fmt.Printf("调用指针接收者方法后: %d (已改变)\n", counter2.value)

	fmt.Println()

	// ============================================
	// 5. 值接收者示例：温度转换
	// ============================================
	fmt.Println("=== 值接收者示例：温度转换 ===")

	// 值接收者适合不修改原始值的操作（如转换）
	tempK3 := Kelvin(273.15)
	tempC5 := tempK3.ToCelsius()
	fmt.Printf("值接收者转换: %.2f K -> %.2f °C (原始值未改变)\n", tempK3, tempC5)
	fmt.Printf("原始值仍然: %.2f K\n", tempK3)

	fmt.Println()

	// ============================================
	// 6. 指针接收者示例：修改状态
	// ============================================
	fmt.Println("=== 指针接收者示例：修改状态 ===")

	// 指针接收者适合需要修改原始值的操作
	bankAccount := BankAccount{balance: 1000.0}
	fmt.Printf("初始余额: $%.2f\n", bankAccount.balance)

	bankAccount.Deposit(500.0)
	fmt.Printf("存款 $500 后: $%.2f\n", bankAccount.balance)

	bankAccount.Withdraw(200.0)
	fmt.Printf("取款 $200 后: $%.2f\n", bankAccount.balance)

	fmt.Println()

	// ============================================
	// 7. 方法链式调用
	// ============================================
	fmt.Println("=== 方法链式调用 ===")

	// 如果方法返回接收者类型，可以实现链式调用
	builder := NewStringBuilder()
	result := builder.Append("Hello").Append(" ").Append("World").String()
	fmt.Printf("链式调用结果: %s\n", result)

	fmt.Println()

	// ============================================
	// 8. 为基本类型添加方法
	// ============================================
	fmt.Println("=== 为基本类型添加方法 ===")

	// 为 int 类型添加方法
	var num MyInt = 5
	fmt.Printf("原始值: %d\n", num)
	squared := num.Square()
	fmt.Printf("平方: %d\n", squared)
	doubled := num.Double()
	fmt.Printf("翻倍: %d\n", doubled)

	fmt.Println()

	// ============================================
	// 9. 方法集规则
	// ============================================
	fmt.Println("=== 方法集规则 ===")

	// 值类型可以调用值接收者和指针接收者的方法
	point1 := Point{X: 1, Y: 2}
	point1.MoveByValue(3, 4) // 值接收者方法
	fmt.Printf("Point1: (%d, %d)\n", point1.X, point1.Y)

	point1.MoveByPointer(5, 6) // 指针接收者方法（Go 自动转换）
	fmt.Printf("Point1: (%d, %d)\n", point1.X, point1.Y)

	// 指针类型也可以调用值接收者和指针接收者的方法
	point2 := &Point{X: 10, Y: 20}
	point2.MoveByValue(1, 2) // 值接收者方法（Go 自动解引用）
	fmt.Printf("Point2: (%d, %d)\n", point2.X, point2.Y)

	point2.MoveByPointer(3, 4) // 指针接收者方法
	fmt.Printf("Point2: (%d, %d)\n", point2.X, point2.Y)

	fmt.Println()

	// ============================================
	// 10. 方法与面向对象
	// ============================================
	fmt.Println("=== 方法与面向对象 ===")

	// Go 没有类，但通过"自定义类型+方法"可以实现类似类的封装
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("原始信息: %s, %d 岁\n", person.Name, person.Age)

	person.SetAge(31) // 使用指针接收者修改
	fmt.Printf("修改后: %s, %d 岁\n", person.Name, person.Age)

	info := person.GetInfo() // 使用值接收者获取信息
	fmt.Printf("信息: %s\n", info)

	fmt.Println()

	// ============================================
	// 11. 方法总结
	// ============================================
	fmt.Println("=== 方法总结 ===")
	fmt.Println("1. 方法是绑定到特定类型上的函数")
	fmt.Println("2. 声明方式: func (接收者 类型) 方法名(参数) 返回值")
	fmt.Println("3. 值接收者: 复制数据，不能修改原始值")
	fmt.Println("4. 指针接收者: 可以修改原始值")
	fmt.Println("5. 值类型可以调用值接收者和指针接收者的方法")
	fmt.Println("6. 指针类型也可以调用值接收者和指针接收者的方法")
	fmt.Println("7. Go 通过'自定义类型+方法'实现面向对象特性")
	fmt.Println("8. 方法是学习接口的基础")
}

// ============================================
// 值接收者 vs 指针接收者示例
// ============================================

// Counter 计数器类型
type Counter struct {
	value int
}

// IncrementByValue 值接收者方法（不会修改原始值）
func (c Counter) IncrementByValue() {
	c.value++ // 只修改副本
}

// IncrementByPointer 指针接收者方法（会修改原始值）
func (c *Counter) IncrementByPointer() {
	c.value++ // 修改原始值
}

// ============================================
// 指针接收者示例：银行账户
// ============================================

// BankAccount 银行账户类型
type BankAccount struct {
	balance float64
}

// Deposit 存款（指针接收者，修改余额）
func (ba *BankAccount) Deposit(amount float64) {
	ba.balance += amount
}

// Withdraw 取款（指针接收者，修改余额）
func (ba *BankAccount) Withdraw(amount float64) {
	if ba.balance >= amount {
		ba.balance -= amount
	}
}

// GetBalance 获取余额（值接收者，只读）
func (ba BankAccount) GetBalance() float64 {
	return ba.balance
}

// ============================================
// 方法链式调用示例
// ============================================

// StringBuilder 字符串构建器
type StringBuilder struct {
	data []byte
}

// NewStringBuilder 创建新的字符串构建器
func NewStringBuilder() *StringBuilder {
	return &StringBuilder{data: make([]byte, 0)}
}

// Append 追加字符串（返回指针，支持链式调用）
func (sb *StringBuilder) Append(s string) *StringBuilder {
	sb.data = append(sb.data, []byte(s)...)
	return sb
}

// String 转换为字符串
func (sb *StringBuilder) String() string {
	return string(sb.data)
}

// ============================================
// 为基本类型添加方法
// ============================================

// MyInt 自定义 int 类型
type MyInt int

// Square 计算平方（值接收者）
func (m MyInt) Square() MyInt {
	return m * m
}

// Double 翻倍（值接收者）
func (m MyInt) Double() MyInt {
	return m * 2
}

// ============================================
// 方法集规则示例
// ============================================

// Point 点类型
type Point struct {
	X, Y int
}

// MoveByValue 值接收者方法
func (p Point) MoveByValue(dx, dy int) {
	p.X += dx
	p.Y += dy
	// 注意：这里不会修改原始值，因为 p 是副本
}

// MoveByPointer 指针接收者方法
func (p *Point) MoveByPointer(dx, dy int) {
	p.X += dx
	p.Y += dy
	// 这里会修改原始值
}

// ============================================
// 面向对象示例
// ============================================

// Person 人类类型
type Person struct {
	Name string
	Age  int
}

// SetAge 设置年龄（指针接收者，修改值）
func (p *Person) SetAge(age int) {
	p.Age = age
}

// GetInfo 获取信息（值接收者，只读）
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s, %d 岁", p.Name, p.Age)
}

