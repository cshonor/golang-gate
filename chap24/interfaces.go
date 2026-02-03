// 示例：Go 语言的接口（Interface）
// 演示接口的定义、隐式实现、标准库接口和接口与组合的结合

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// ============================================
// 类型定义和接口定义（包级别）
// ============================================

// Speaker 说话接口
type Speaker interface {
	Speak() string
}

// Dog 狗类型
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.Name)
}

// Cat 猫类型
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!", c.Name)
}

// Reader2 读取接口
type Reader2 interface {
	Read() string
}

// Writer2 写入接口
type Writer2 interface {
	Write(data string)
}

// Closer2 关闭接口
type Closer2 interface {
	Close()
}

// ReadWriter2 读写接口（组合接口）
type ReadWriter2 interface {
	Reader2
	Writer2
}

// File 文件结构体（实现ReadWriter2接口）
type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

// Stringer 字符串化接口
type Stringer interface {
	String() string
}

// Person 人员结构体
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// Movable 可移动接口
type Movable interface {
	Move()
}

// Car 汽车结构体
type Car struct {
	Brand string
}

func (c Car) Move() {
	fmt.Printf("    %s car is moving\n", c.Brand)
}

// Bicycle 自行车结构体
type Bicycle struct {
	Brand string
}

func (b Bicycle) Move() {
	fmt.Printf("    %s bicycle is moving\n", b.Brand)
}

// Vehicle 车辆结构体（组合多个Movable）
type Vehicle struct {
	Movables []Movable
}

func (v *Vehicle) AddMovable(m Movable) {
	v.Movables = append(v.Movables, m)
}

func (v Vehicle) MoveAll() {
	for _, m := range v.Movables {
		m.Move()
	}
}

// PaymentMethod 支付方式接口
type PaymentMethod interface {
	Pay(amount float64) error
}

// CreditCard 信用卡结构体
type CreditCard struct {
	Number string
}

func (c CreditCard) Pay(amount float64) error {
	fmt.Printf("    Paying %.2f with credit card %s\n", amount, c.Number)
	return nil
}

// PayPal PayPal结构体
type PayPal struct {
	Email string
}

func (p PayPal) Pay(amount float64) error {
	fmt.Printf("    Paying %.2f with PayPal %s\n", amount, p.Email)
	return nil
}

// Storage 存储接口
type Storage interface {
	Store(key string, value interface{})
	Retrieve(key string) (interface{}, bool)
}

// MemoryStorage 内存存储结构体
type MemoryStorage struct {
	data map[string]interface{}
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]interface{}),
	}
}

func (m *MemoryStorage) Store(key string, value interface{}) {
	m.data[key] = value
}

func (m *MemoryStorage) Retrieve(key string) (interface{}, bool) {
	value, exists := m.data[key]
	return value, exists
}

func main() {
	// ============================================
	// 1. "让类型'说话'"的含义
	// ============================================
	fmt.Println("=== 1. '让类型说话'的含义 ===")

	fmt.Println("在Go里，接口是一种'行为契约'")
	fmt.Println("只要一个类型实现了接口定义的所有方法，它就自动'属于'这个接口类型")
	fmt.Println("无需显式声明")
	fmt.Println()

	// 示例：Speaker接口
	fmt.Println("示例：Speaker接口")
	fmt.Println("  定义Speaker接口和Dog、Cat类型")

	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// 都可以作为Speaker接口使用
	var speaker1 Speaker = dog
	var speaker2 Speaker = cat

	fmt.Printf("  %s\n", speaker1.Speak())
	fmt.Printf("  %s\n", speaker2.Speak())
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - Dog和Cat都没有显式声明实现Speaker接口")
	fmt.Println("  - 但它们都实现了Speak()方法")
	fmt.Println("  - 所以它们自动满足了Speaker接口")
	fmt.Println("  - 这就是让类型'说话'的含义")
	fmt.Println()

	// ============================================
	// 2. 按需使用接口
	// ============================================
	fmt.Println("=== 2. 按需使用接口 ===")

	fmt.Println("Go的接口设计非常'小而专'")
	fmt.Println("通常只定义1到2个方法")
	fmt.Println("可以根据场景需要，定义最精简的接口")
	fmt.Println("避免设计臃肿的'万能接口'")
	fmt.Println()

	// 示例：小而专的接口
	fmt.Println("示例：小而专的接口")
	fmt.Println("  定义Reader2、Writer2、Closer2和ReadWriter2接口")

	fmt.Println("  小接口的优势：")
	fmt.Println("    ✅ Reader: 只定义读取行为")
	fmt.Println("    ✅ Writer: 只定义写入行为")
	fmt.Println("    ✅ Closer: 只定义关闭行为")
	fmt.Println("    ✅ ReadWriter: 组合多个小接口")
	fmt.Println()

	// ============================================
	// 3. 标准库中的接口
	// ============================================
	fmt.Println("=== 3. 标准库中的接口 ===")

	fmt.Println("Writer接口是最经典的例子之一")
	fmt.Println("只有一个Write方法，却被几十种类型实现")
	fmt.Println()

	// 演示Writer接口的使用
	fmt.Println("演示Writer接口的使用:")

	// 方式1：写入文件
	fmt.Println("  方式1：写入文件")
	osFile, err := os.Create("test.txt")
	if err == nil {
		writeToWriter(osFile, "Hello, File!")
		osFile.Close()
		fmt.Println("    已写入文件")
	}

	// 方式2：写入内存缓冲区
	fmt.Println("  方式2：写入内存缓冲区")
	var buf bytes.Buffer
	writeToWriter(&buf, "Hello, Buffer!")
	fmt.Printf("    缓冲区内容: %s\n", buf.String())

	// 方式3：写入标准输出
	fmt.Println("  方式3：写入标准输出")
	writeToWriter(os.Stdout, "Hello, Stdout!\n")
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - os.File、bytes.Buffer、os.Stdout都实现了Writer接口")
	fmt.Println("  - 可以用统一的方式写入数据")
	fmt.Println("  - 完全不用关心底层实现")
	fmt.Println()

	// ============================================
	// 4. 接口的隐式实现
	// ============================================
	fmt.Println("=== 4. 接口的隐式实现 ===")

	fmt.Println("Go的接口实现是隐式的")
	fmt.Println("不需要显式声明'implements'关键字")
	fmt.Println()

	// 定义接口
	fmt.Println("  定义Stringer接口，Person类型自动实现")

	person := Person{Name: "Alice", Age: 25}
	var stringer Stringer = person
	fmt.Printf("  Person实现了Stringer接口: %s\n", stringer.String())
	fmt.Println()

	// ============================================
	// 5. 接口与组合的结合
	// ============================================
	fmt.Println("=== 5. 接口与组合的结合 ===")

	fmt.Println("接口 + 组合 = Go面向对象的黄金搭档")
	fmt.Println("组合让你复用代码，接口让你定义行为契约")
	fmt.Println("两者结合就能实现高度解耦的设计")
	fmt.Println()

	// 示例：可移动的对象
	fmt.Println("示例：可移动的对象")
	fmt.Println("  定义Movable接口和Car、Bicycle、Vehicle类型")

	vehicle := Vehicle{}
	vehicle.AddMovable(Car{Brand: "Toyota"})
	vehicle.AddMovable(Bicycle{Brand: "Giant"})
	fmt.Println("  移动所有对象:")
	vehicle.MoveAll()
	fmt.Println()

	// ============================================
	// 6. 替代继承
	// ============================================
	fmt.Println("=== 6. 替代继承 ===")

	fmt.Println("接口的隐式实现机制，让你摆脱了继承链的束缚")
	fmt.Println("通过'鸭子类型'实现多态，比传统继承更灵活")
	fmt.Println()

	// 示例：支付系统
	fmt.Println("示例：支付系统")
	fmt.Println("  定义PaymentMethod接口和CreditCard、PayPal类型")

	// 统一的支付处理函数
	processPayment := func(method PaymentMethod, amount float64) {
		method.Pay(amount)
	}

	processPayment(CreditCard{Number: "1234-5678"}, 100.0)
	processPayment(PayPal{Email: "user@example.com"}, 200.0)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - CreditCard和PayPal都实现了PaymentMethod接口")
	fmt.Println("  - 不需要继承关系")
	fmt.Println("  - 只要实现了Pay方法，就可以作为PaymentMethod使用")
	fmt.Println()

	// ============================================
	// 7. 接口的嵌套（组合接口）
	// ============================================
	fmt.Println("=== 7. 接口的嵌套（组合接口）===")

	fmt.Println("接口可以嵌套其他接口，形成更大的接口")
	fmt.Println()

	// 定义小接口
	fmt.Println("  定义Reader2、Writer2和ReadWriter2接口")
	fmt.Println("  File结构体实现ReadWriter2接口")

	fileRW := &File{}
	var rw ReadWriter2 = fileRW
	rw.Write("Hello, World!")
	fmt.Printf("  写入后读取: %s\n", rw.Read())
	fmt.Println()

	// ============================================
	// 8. 空接口（interface{}）
	// ============================================
	fmt.Println("=== 8. 空接口（interface{}）===")

	fmt.Println("空接口可以表示任何类型")
	fmt.Println("类似于其他语言中的'Object'或'any'类型")
	fmt.Println()

	var any interface{}
	any = 42
	fmt.Printf("  any = %v (类型: %T)\n", any, any)
	any = "hello"
	fmt.Printf("  any = %v (类型: %T)\n", any, any)
	any = true
	fmt.Printf("  any = %v (类型: %T)\n", any, any)
	fmt.Println()

	// ============================================
	// 9. 类型断言
	// ============================================
	fmt.Println("=== 9. 类型断言 ===")

	fmt.Println("类型断言用于从接口中提取具体类型")
	fmt.Println()

	var value interface{} = "hello"

	// 方式1：类型断言
	str, ok := value.(string)
	if ok {
		fmt.Printf("  值是一个字符串: %s\n", str)
	}

	// 方式2：类型switch
	switch v := value.(type) {
	case string:
		fmt.Printf("  类型switch: 字符串 %s\n", v)
	case int:
		fmt.Printf("  类型switch: 整数 %d\n", v)
	default:
		fmt.Printf("  类型switch: 其他类型\n")
	}
	fmt.Println()

	// ============================================
	// 10. 实际应用示例
	// ============================================
	fmt.Println("=== 10. 实际应用示例 ===")

	// 示例1：排序接口
	fmt.Println("示例1：排序接口")
	fmt.Println("  定义Sortable接口和IntSlice类型")

	numbers := IntSlice{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("  排序前: %v\n", numbers)
	bubbleSort(numbers)
	fmt.Printf("  排序后: %v\n", numbers)
	fmt.Println()

	// 示例2：存储接口
	fmt.Println("示例2：存储接口")
	fmt.Println("  定义Storage接口和MemoryStorage类型")

	storage := NewMemoryStorage()
	storage.Store("name", "Alice")
	storage.Store("age", 25)
	if name, ok := storage.Retrieve("name"); ok {
		fmt.Printf("  存储和检索: name = %v\n", name)
	}
	fmt.Println()

	// ============================================
	// 11. 总结
	// ============================================
	fmt.Println("=== 11. 总结 ===")
	fmt.Println()
	fmt.Println("1. '让类型说话'的含义:")
	fmt.Println("   ✅ 接口是行为契约")
	fmt.Println("   ✅ 实现接口的方法就自动满足接口")
	fmt.Println("   ✅ 无需显式声明")
	fmt.Println()
	fmt.Println("2. 按需使用接口:")
	fmt.Println("   ✅ 小而专的接口设计")
	fmt.Println("   ✅ 通常只定义1到2个方法")
	fmt.Println("   ✅ 避免臃肿的'万能接口'")
	fmt.Println()
	fmt.Println("3. 标准库中的接口:")
	fmt.Println("   ✅ Writer接口：统一写入方式")
	fmt.Println("   ✅ 被几十种类型实现")
	fmt.Println("   ✅ 完全不用关心底层实现")
	fmt.Println()
	fmt.Println("4. 接口的隐式实现:")
	fmt.Println("   ✅ 不需要显式声明'implements'")
	fmt.Println("   ✅ 实现方法就自动满足接口")
	fmt.Println()
	fmt.Println("5. 接口与组合:")
	fmt.Println("   ✅ 接口 + 组合 = 黄金搭档")
	fmt.Println("   ✅ 组合复用代码，接口定义行为契约")
	fmt.Println("   ✅ 实现高度解耦的设计")
	fmt.Println()
	fmt.Println("6. 替代继承:")
	fmt.Println("   ✅ 通过'鸭子类型'实现多态")
	fmt.Println("   ✅ 比传统继承更灵活")
	fmt.Println("   ✅ 摆脱继承链的束缚")
	fmt.Println()
}

// ============================================
// 辅助函数
// ============================================

// writeToWriter 写入到Writer接口
func writeToWriter(w io.Writer, data string) {
	w.Write([]byte(data))
}

// Sortable 排序接口
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// IntSlice 整数切片（实现Sortable接口）
type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// bubbleSort 冒泡排序（使用Sortable接口）
func bubbleSort(data Sortable) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

