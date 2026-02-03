// 示例：Go 语言没有类，通过结构+方法+组合实现面向对象
// 演示Go的面向对象思路、结构体绑定方法、嵌套组合和接口实现

package main

import (
	"fmt"
	"math"
)

// ============================================
// 类型定义和方法实现（包级别）
// ============================================

// Celsius 摄氏度类型
type Celsius float64

func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c + 273.15)
}

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Kelvin 开尔文类型
type Kelvin float64

// Fahrenheit 华氏度类型
type Fahrenheit float64

// Point 点结构体
type Point struct {
	X, Y float64
}

func (p Point) DistanceTo(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Point) DistanceFromOrigin() float64 {
	return p.DistanceTo(Point{0, 0})
}

// Account 账户结构体（封装示例）
type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string, initialBalance float64) *Account {
	return &Account{
		owner:   owner,
		balance: initialBalance,
	}
}

func (a *Account) GetOwner() string {
	return a.owner
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
	}
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return fmt.Errorf("余额不足")
	}
	if amount > 0 {
		a.balance -= amount
	}
	return nil
}

// Shape 形状接口（多态示例）
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle 圆形
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle 矩形
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Person 人员结构体
type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() string {
	return fmt.Sprintf("我是%s，%d岁", p.Name, p.Age)
}

// Employee 员工结构体（嵌入Person）
type Employee struct {
	Person        // 嵌入，可以直接访问Person的字段和方法
	EmployeeID    int
	Department    string
	Salary        float64
}

// Address 地址结构体
type Address struct {
	Street string
	City   string
	Zip    string
}

// Contact 联系方式结构体
type Contact struct {
	Email   string
	Phone   string
	Address Address
}

// Manager 经理结构体（多层嵌套）
type Manager struct {
	Employee
	Contact
	TeamSize int
}

// Reader 读取接口
type Reader interface {
	Read() string
}

// Writer 写入接口
type Writer interface {
	Write(data string)
}

// Closer 关闭接口
type Closer interface {
	Close()
}

// File 文件结构体（实现多个接口）
type File struct {
	name    string
	content string
	closed  bool
}

func (f *File) Read() string {
	if f.closed {
		return ""
	}
	return f.content
}

func (f *File) Write(data string) {
	if !f.closed {
		f.content = data
	}
}

func (f *File) Close() {
	f.closed = true
}

// TemperatureSensor 温度传感器（组合示例）
type TemperatureSensor struct {
	Location string
	Unit     Celsius
}

func (ts *TemperatureSensor) UpdateTemperature(temp Celsius) {
	ts.Unit = temp
}

func (ts TemperatureSensor) Describe() string {
	return fmt.Sprintf("%s的温度是%.1f°C", ts.Location, ts.Unit)
}

// Handler 处理器接口
type Handler interface {
	Handle() string
}

// MyHandler 自定义处理器
type MyHandler struct {
	name string
}

func (h *MyHandler) Handle() string {
	return fmt.Sprintf("Handler %s processed request", h.name)
}

func processRequest(handler Handler) {
	fmt.Printf("  处理请求: %s\n", handler.Handle())
}

// Connection 连接结构体
type Connection struct {
	host string
	port int
}

func (c *Connection) Connect() {
	fmt.Printf("    连接到 %s:%d\n", c.host, c.port)
}

func (c *Connection) Disconnect() {
	fmt.Printf("    断开连接\n")
}

// Logger 日志结构体
type Logger struct {
	name string
}

func (l *Logger) Log(message string) {
	fmt.Printf("    [%s] %s\n", l.name, message)
}

// Database 数据库（组合Connection和Logger）
type Database struct {
	*Connection
	*Logger
}

func main() {
	// ============================================
	// 1. Go的面向对象思路
	// ============================================
	fmt.Println("=== 1. Go的面向对象思路 ===")

	fmt.Println("Go不支持类和继承，但通过以下方式实现面向对象：")
	fmt.Println("  - 为结构体绑定方法")
	fmt.Println("  - 利用结构体的嵌套组合")
	fmt.Println("  - 实现其他语言中类的大部分功能")
	fmt.Println("  - 这是'组合优于继承'的Go式实践")
	fmt.Println()

	// ============================================
	// 2. 为结构化数据提供行为（结构体绑定方法）
	// ============================================
	fmt.Println("=== 2. 为结构化数据提供行为（结构体绑定方法）===")

	fmt.Println("给结构体定义方法，让数据（结构体）和行为（方法）绑定在一起")
	fmt.Println("模拟'对象'的特性")
	fmt.Println()

	// 示例1：温度转换
	fmt.Println("示例1：温度转换")
	celsius := Celsius(25.0)
	fmt.Printf("  摄氏度: %.1f°C\n", celsius)
	fmt.Printf("  转开尔文: %.1f K\n", celsius.ToKelvin())
	fmt.Printf("  转华氏度: %.1f°F\n", celsius.ToFahrenheit())
	fmt.Println()

	// 示例2：坐标计算
	fmt.Println("示例2：坐标计算")
	point1 := Point{X: 0, Y: 0}
	point2 := Point{X: 3, Y: 4}
	fmt.Printf("  点1: %v\n", point1)
	fmt.Printf("  点2: %v\n", point2)
	fmt.Printf("  距离: %.2f\n", point1.DistanceTo(point2))
	fmt.Printf("  点1到原点的距离: %.2f\n", point1.DistanceFromOrigin())
	fmt.Println()

	// ============================================
	// 3. 应用面向对象设计原则
	// ============================================
	fmt.Println("=== 3. 应用面向对象设计原则 ===")

	fmt.Println("封装：通过结构体和方法实现数据封装")
	fmt.Println("多态：通过接口实现多态")
	fmt.Println()

	// 封装示例
	fmt.Println("封装示例：")
	account := NewAccount("Alice", 1000.0)
	fmt.Printf("  账户: %s\n", account.GetOwner())
	fmt.Printf("  余额: %.2f\n", account.GetBalance())
	account.Deposit(500.0)
	fmt.Printf("  存款500后余额: %.2f\n", account.GetBalance())
	err := account.Withdraw(2000.0)
	if err != nil {
		fmt.Printf("  取款失败: %v\n", err)
	}
	fmt.Println()

	// 多态示例（通过接口）
	fmt.Println("多态示例（通过接口）:")
	var shapes []Shape
	shapes = append(shapes, Circle{Radius: 5.0})
	shapes = append(shapes, Rectangle{Width: 4.0, Height: 6.0})
	for i, shape := range shapes {
		fmt.Printf("  形状 %d: 面积=%.2f, 周长=%.2f\n", i+1, shape.Area(), shape.Perimeter())
	}
	fmt.Println()

	// ============================================
	// 4. 结构体的嵌套组合（组合优于继承）
	// ============================================
	fmt.Println("=== 4. 结构体的嵌套组合（组合优于继承）===")

	fmt.Println("通过在一个结构体中嵌入另一个结构体")
	fmt.Println("可以直接复用其字段和方法，模拟'继承'的效果")
	fmt.Println("但比继承更灵活")
	fmt.Println()

	// 示例1：基础结构体
	fmt.Println("示例1：基础结构体")
	fmt.Println("  定义Person结构体和Introduce方法")

	// 嵌入Person结构体
	fmt.Println("  定义Employee结构体，嵌入Person")

	employee := Employee{
		Person: Person{
			Name: "Bob",
			Age:  30,
		},
		EmployeeID: 1001,
		Department: "技术部",
		Salary:     15000.0,
	}

	fmt.Printf("  员工: %s\n", employee.Introduce()) // 直接使用嵌入结构体的方法
	fmt.Printf("  员工ID: %d\n", employee.EmployeeID)
	fmt.Printf("  部门: %s\n", employee.Department)
	fmt.Printf("  姓名: %s (直接访问嵌入字段)\n", employee.Name)
	fmt.Println()

	// 示例2：多层嵌套
	fmt.Println("示例2：多层嵌套")
	fmt.Println("  定义Address、Contact和Manager结构体")

	manager := Manager{
		Employee: Employee{
			Person: Person{
				Name: "Charlie",
				Age:  35,
			},
			EmployeeID: 2001,
			Department: "产品部",
			Salary:     25000.0,
		},
		Contact: Contact{
			Email: "charlie@example.com",
			Phone: "13800138000",
			Address: Address{
				Street: "123 Main St",
				City:   "Beijing",
				Zip:    "100000",
			},
		},
		TeamSize: 10,
	}

	fmt.Printf("  经理: %s\n", manager.Introduce())
	fmt.Printf("  邮箱: %s\n", manager.Email)
	fmt.Printf("  城市: %s\n", manager.Address.City)
	fmt.Printf("  团队规模: %d\n", manager.TeamSize)
	fmt.Println()

	// ============================================
	// 5. 接口与结构体的组合
	// ============================================
	fmt.Println("=== 5. 接口与结构体的组合 ===")

	fmt.Println("一个结构体可以实现多个接口")
	fmt.Println("不同接口的方法组合在一起，让结构体具备多维度的行为能力")
	fmt.Println()

	// 定义多个接口
	fmt.Println("  定义Reader、Writer、Closer接口")
	fmt.Println("  File结构体实现多个接口")

	file := &File{name: "test.txt", content: "Hello", closed: false}

	// 作为Reader使用
	var reader Reader = file
	fmt.Printf("  作为Reader读取: %s\n", reader.Read())

	// 作为Writer使用
	var writer Writer = file
	writer.Write("World")
	fmt.Printf("  作为Writer写入后: %s\n", file.Read())

	// 作为Closer使用
	var closer Closer = file
	closer.Close()
	fmt.Printf("  关闭后读取: %s (空)\n", reader.Read())
	fmt.Println()

	// ============================================
	// 6. 组合的协同效应
	// ============================================
	fmt.Println("=== 6. 组合的协同效应 ===")

	fmt.Println("'协同效应'：类型+基于类型的方法+结构体的组合")
	fmt.Println("产生了'1+1>2'的效果")
	fmt.Println()

	// 示例：温度传感器系统
	fmt.Println("示例：温度传感器系统")
	sensor := TemperatureSensor{
		Location: "Mars",
		Unit:     Celsius(0),
	}

	// 组合了多个能力
	fmt.Printf("  传感器位置: %s\n", sensor.Location)
	fmt.Printf("  当前温度: %.1f°C\n", sensor.Unit)
	fmt.Printf("  温度描述: %s\n", sensor.Describe())
	sensor.UpdateTemperature(Celsius(-65))
	fmt.Printf("  更新后温度: %.1f°C\n", sensor.Unit)
	fmt.Println()

	// ============================================
	// 7. 实际应用示例
	// ============================================
	fmt.Println("=== 7. 实际应用示例 ===")

	// 示例1：HTTP处理器（模拟标准库的设计）
	fmt.Println("示例1：HTTP处理器（模拟标准库的设计）")
	handler := &MyHandler{name: "API Handler"}
	processRequest(handler)
	fmt.Println()

	// 示例2：数据库连接（组合模式）
	fmt.Println("示例2：数据库连接（组合模式）")
	db := &Database{
		Connection: &Connection{host: "localhost", port: 5432},
		Logger:     &Logger{name: "DB"},
	}
	db.Connect()
	db.Log("Query executed")
	db.Disconnect()
	fmt.Println()

	// ============================================
	// 8. Go vs 传统面向对象语言
	// ============================================
	fmt.Println("=== 8. Go vs 传统面向对象语言 ===")

	fmt.Println("传统面向对象语言（如Java、C++）:")
	fmt.Println("  - 使用类和继承")
	fmt.Println("  - 类包含数据和方法")
	fmt.Println("  - 通过继承实现代码复用")
	fmt.Println()

	fmt.Println("Go语言:")
	fmt.Println("  - 没有类，使用结构体")
	fmt.Println("  - 结构体包含数据，方法绑定到类型")
	fmt.Println("  - 通过组合实现代码复用")
	fmt.Println("  - 更灵活，避免继承的复杂性")
	fmt.Println()

	// ============================================
	// 9. 总结
	// ============================================
	fmt.Println("=== 9. 总结 ===")
	fmt.Println()
	fmt.Println("1. Go的面向对象思路:")
	fmt.Println("   ✅ 通过结构体+方法+组合实现面向对象")
	fmt.Println("   ✅ 组合优于继承")
	fmt.Println("   ✅ 更灵活，避免继承的复杂性")
	fmt.Println()
	fmt.Println("2. 为结构化数据提供行为:")
	fmt.Println("   ✅ 给结构体定义方法")
	fmt.Println("   ✅ 数据和行为绑定在一起")
	fmt.Println("   ✅ 模拟'对象'的特性")
	fmt.Println()
	fmt.Println("3. 应用面向对象设计原则:")
	fmt.Println("   ✅ 封装：通过结构体和方法实现")
	fmt.Println("   ✅ 多态：通过接口实现")
	fmt.Println()
	fmt.Println("4. 结构体的嵌套组合:")
	fmt.Println("   ✅ 嵌入结构体，复用字段和方法")
	fmt.Println("   ✅ 模拟'继承'的效果")
	fmt.Println("   ✅ 比继承更灵活")
	fmt.Println()
	fmt.Println("5. 接口与结构体的组合:")
	fmt.Println("   ✅ 一个结构体可以实现多个接口")
	fmt.Println("   ✅ 不同接口的方法组合，多维度行为能力")
	fmt.Println()
	fmt.Println("6. 组合的协同效应:")
	fmt.Println("   ✅ 类型+方法+结构体的组合")
	fmt.Println("   ✅ 产生'1+1>2'的效果")
	fmt.Println("   ✅ 实现强大的功能")
	fmt.Println()
}


