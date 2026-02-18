// 示例：Go 语言的指针（Pointers）
// 演示指针的声明、使用、值传递vs指针传递和实际应用

package main

import "fmt"

func main() {
	// ============================================
	// 1. 指针的形象化理解
	// ============================================
	fmt.Println("=== 1. 指针的形象化理解 ===")

	fmt.Println("书里用'店铺搬迁告示'来比喻指针非常贴切：")
	fmt.Println("  - 指针本身就像'告示牌'，它不直接存储数据")
	fmt.Println("  - 而是存储了数据在内存中的地址")
	fmt.Println("  - 当你需要访问真正的数据时，就像根据告示上的新地址找到店铺一样")
	fmt.Println("  - 通过指针的地址去内存中取值")
	fmt.Println()

	// 示例：指针就像地址
	fmt.Println("示例：指针就像地址")
	value := 42
	pointer := &value // 获取value的地址（告示牌）
	fmt.Printf("  值: %d (存储在内存中)\n", value)
	fmt.Printf("  指针: %p (指向内存地址，就像告示牌上的地址)\n", pointer)
	fmt.Printf("  通过指针访问值: %d (根据地址找到的值)\n", *pointer)
	fmt.Println()

	// ============================================
	// 2. 声明和使用指针
	// ============================================
	fmt.Println("=== 2. 声明和使用指针 ===")

	fmt.Println("基本语法：")
	fmt.Println("  & 操作符：取地址（获取变量的内存地址）")
	fmt.Println("  * 操作符：解引用（通过指针访问值）")
	fmt.Println()

	// 示例1：基本指针操作
	fmt.Println("示例1：基本指针操作")
	x := 10
	p := &x // p是指向x的指针
	fmt.Printf("  x的值: %d\n", x)
	fmt.Printf("  x的地址: %p\n", &x)
	fmt.Printf("  p的值（地址）: %p\n", p)
	fmt.Printf("  p指向的值: %d\n", *p)
	fmt.Println()

	// 示例2：通过指针修改值
	fmt.Println("示例2：通过指针修改值")
	*p = 20 // 通过指针修改值
	fmt.Printf("  修改*p后，x的值: %d\n", x)
	fmt.Println()

	// 示例3：指针的类型
	fmt.Println("示例3：指针的类型")
	var ptr *int // 声明一个指向int的指针
	fmt.Printf("  未初始化的指针: %v, nil=%t\n", ptr, ptr == nil)
	num := 100
	ptr = &num
	fmt.Printf("  初始化后的指针: %p, 指向的值: %d\n", ptr, *ptr)
	fmt.Println()

	// ============================================
	// 3. 指针与RAM的关系
	// ============================================
	fmt.Println("=== 3. 指针与RAM的关系 ===")

	fmt.Println("指针本质是内存地址")
	fmt.Println("它让程序可以直接操作内存中的数据")
	fmt.Println("这也是Go能高效处理数据的原因之一")
	fmt.Println()

	// 演示内存地址
	fmt.Println("演示内存地址:")
	a := 1
	b := 2
	c := 3
	fmt.Printf("  a的地址: %p, 值: %d\n", &a, a)
	fmt.Printf("  b的地址: %p, 值: %d\n", &b, b)
	fmt.Printf("  c的地址: %p, 值: %d\n", &c, c)
	fmt.Println("  说明：每个变量在内存中都有唯一的地址")
	fmt.Println()

	// ============================================
	// 4. 指针的使用时机
	// ============================================
	fmt.Println("=== 4. 指针的使用时机 ===")

	fmt.Println("通常在以下情况使用指针：")
	fmt.Println("  1. 需要修改原始数据")
	fmt.Println("  2. 避免大值拷贝")
	fmt.Println("  3. 实现数据共享")
	fmt.Println("  4. 构建复杂数据结构（如链表、树）")
	fmt.Println()

	// 示例1：修改原始数据
	fmt.Println("示例1：修改原始数据")
	original := 10
	fmt.Printf("  修改前: %d\n", original)
	modifyValue(original)
	fmt.Printf("  值传递后: %d (未改变)\n", original)
	modifyPointer(&original)
	fmt.Printf("  指针传递后: %d (已改变)\n", original)
	fmt.Println()

	// 示例2：避免大值拷贝 这就是「大值拷贝」：
 
- 拷贝的数据量很大
​
- 非常消耗内存和 CPU 时间
​
- 会显著拖慢程序性能
如果改成指针传递：
- 传递的是指针，只复制 8 字节
​
- 函数里通过指针直接操作原结构体，没有副本
​
- 既高效，又能修改原数据
	fmt.Println("示例2：避免大值拷贝")
	_ = LargeStruct{
		Data: [1000]int{},
	}
	fmt.Printf("  大结构体大小: %d 字节\n", 1000*8) // 假设int是8字节
	fmt.Println("  值传递会复制整个结构体（8000字节）")
	fmt.Println("  指针传递只复制指针（8字节）")
	fmt.Println()
这是一个数组字面量，表示一个长度为 1000、元素类型为  int  的数组。
 
-  [1000]int ：这是数组的类型，表示“长度固定为 1000 的 int 数组”。
​
-  {} ：这是数组的初始化值。在 Go 里，如果大括号  {}  里面是空的，就表示用该类型的零值来初始化所有元素。
​
- 对于  int  类型，零值就是  0 。
​
- 所以  [1000]int{}  等价于一个包含 1000 个  0  的数组。
	// 示例3：数据共享
	fmt.Println("示例3：数据共享")
	sharedData := 100
	ptr1 := &sharedData
	ptr2 := &sharedData
	fmt.Printf("  ptr1指向的值: %d\n", *ptr1)
	fmt.Printf("  ptr2指向的值: %d\n", *ptr2)
	*ptr1 = 200
	fmt.Printf("  修改ptr1后，ptr2指向的值: %d (共享同一块内存)\n", *ptr2)
	fmt.Println()

	// ============================================
	// 5. Go指针的关键特性
	// ============================================
	fmt.Println("=== 5. Go指针的关键特性 ===")

	fmt.Println("Go的指针相比C/C++更安全：")
	fmt.Println("  ✅ 不支持指针算术运算")
	fmt.Println("  ✅ 不能随意转换指针类型")
	fmt.Println("  ✅ 避免了很多内存错误")
	fmt.Println()

	// 示例：Go指针的限制
	fmt.Println("示例：Go指针的限制")
	p1 := &x
	fmt.Printf("  指针p1: %p\n", p1)
	fmt.Println("  // p1++  // ❌ 错误：Go不支持指针算术")
	fmt.Println("  // p1 = p1 + 1  // ❌ 错误：Go不支持指针算术")
	fmt.Println("  说明：Go指针更安全，不能随意操作内存地址")
	fmt.Println()

	// ============================================
	// 6. 值传递 vs 指针传递
	// ============================================
	fmt.Println("=== 6. 值传递 vs 指针传递 ===")

	fmt.Println("Go函数默认是值传递")
	fmt.Println("用指针可以让函数直接修改外部变量")
	fmt.Println("避免拷贝大结构体带来的性能开销")
	fmt.Println()

	// 对比示例
	fmt.Println("对比示例：")
	num1 := 10
	num2 := 10
	fmt.Printf("  原始值: num1=%d, num2=%d\n", num1, num2)

	// 值传递
	passByValue(num1)
	fmt.Printf("  值传递后: num1=%d (未改变)\n", num1)

	// 指针传递
	passByPointer(&num2)
	fmt.Printf("  指针传递后: num2=%d (已改变)\n", num2)
	fmt.Println()

	// 性能对比
	fmt.Println("性能对比：")
	smallStruct := SmallStruct{Value: 1}
	largeStruct2 := LargeStruct{Data: [1000]int{}}
	fmt.Println("  小结构体（值传递）:")
	processSmallStruct(smallStruct)
	fmt.Println("  大结构体（指针传递）:")
	processLargeStruct(&largeStruct2)
	fmt.Println()

	// ============================================
	// 7. nil指针
	// ============================================
	fmt.Println("=== 7. nil指针 ===")

	fmt.Println("未初始化的指针值为nil")
	fmt.Println("解引用nil指针会触发运行时panic")
	fmt.Println("使用时需要注意判空")
	fmt.Println()

	// 示例1：nil指针
	fmt.Println("示例1：nil指针")
	var nilPtr *int
	fmt.Printf("  nil指针: %v, nil=%t\n", nilPtr, nilPtr == nil)
	// *nilPtr = 10 // ❌ 错误：解引用nil指针会panic
	fmt.Println("  // *nilPtr = 10  // ❌ 错误：会panic")
	fmt.Println()

	// 示例2：安全使用指针
	fmt.Println("示例2：安全使用指针")
	safePtr := new(int) // 使用new创建指针
	fmt.Printf("  new(int)创建的指针: %p, 值: %d\n", safePtr, *safePtr)
	*safePtr = 42
	fmt.Printf("  赋值后: %d\n", *safePtr)
	fmt.Println()

	// 示例3：nil检查
	fmt.Println("示例3：nil检查")
	var checkPtr *int
	if checkPtr == nil {
		fmt.Println("  指针是nil，需要初始化")
		checkPtr = new(int)
		*checkPtr = 100
	}
	fmt.Printf("  安全使用: %d\n", *checkPtr)
	fmt.Println()

	// ============================================
	// 8. 指针的实际应用
	// ============================================
	fmt.Println("=== 8. 指针的实际应用 ===")

	// 应用1：链表节点
	fmt.Println("应用1：链表节点")
	node1 := &Node{Value: 1, Next: nil}
	node2 := &Node{Value: 2, Next: nil}
	node3 := &Node{Value: 3, Next: nil}
	node1.Next = node2
	node2.Next = node3
	fmt.Printf("  链表: %d -> %d -> %d\n", node1.Value, node1.Next.Value, node1.Next.Next.Value)
	fmt.Println()

	// 应用2：交换两个值
	fmt.Println("应用2：交换两个值")
	a1, b1 := 10, 20
	fmt.Printf("  交换前: a=%d, b=%d\n", a1, b1)
	swap(&a1, &b1)
	fmt.Printf("  交换后: a=%d, b=%d\n", a1, b1)
	fmt.Println()

	// 应用3：返回多个值（通过指针）
	fmt.Println("应用3：返回多个值（通过指针）")
	result := 0
	success := divide(10, 2, &result)
	if success {
		fmt.Printf("  除法结果: %d\n", result)
	}
	success2 := divide(10, 0, &result)
	if !success2 {
		fmt.Println("  除法失败: 除数不能为0")
	}
	fmt.Println()

	// ============================================
	// 9. 指针与结构体
	// ============================================
	fmt.Println("=== 9. 指针与结构体 ===")

	// 值接收者 vs 指针接收者
	fmt.Println("值接收者 vs 指针接收者:")
	person1 := Person{Name: "Alice", Age: 25}
	fmt.Printf("  原始: %+v\n", person1)
	person1.SetAgeValue(30) // 值接收者，不会修改原值
	fmt.Printf("  值接收者SetAge后: %+v (未改变)\n", person1)
	person1.SetAgePointer(30) // 指针接收者，会修改原值
	fmt.Printf("  指针接收者SetAge后: %+v (已改变)\n", person1)
	fmt.Println()

	// ============================================
	// 10. 指针数组和数组指针
	// ============================================
	fmt.Println("=== 10. 指针数组和数组指针 ===")

	// 指针数组：数组的元素是指针
	fmt.Println("指针数组：数组的元素是指针")
	ptrArray := [3]*int{&a1, &b1, new(int)}
	*ptrArray[2] = 30
	fmt.Printf("  指针数组: [%d, %d, %d]\n", *ptrArray[0], *ptrArray[1], *ptrArray[2])
	fmt.Println()

	// 数组指针：指向数组的指针
	fmt.Println("数组指针：指向数组的指针")
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Printf("  数组: %v\n", arr)
	fmt.Printf("  通过指针访问: %v\n", *arrPtr)
	(*arrPtr)[0] = 10
	fmt.Printf("  修改后: %v\n", arr)
	fmt.Println()

	// ============================================
	// 11. 指针的最佳实践
	// ============================================
	fmt.Println("=== 11. 指针的最佳实践 ===")

	fmt.Println("1. 何时使用指针:")
	fmt.Println("   ✅ 需要修改函数参数时")
	fmt.Println("   ✅ 传递大结构体时（避免拷贝）")
	fmt.Println("   ✅ 实现数据共享时")
	fmt.Println("   ✅ 构建复杂数据结构时")
	fmt.Println()

	fmt.Println("2. 何时不使用指针:")
	fmt.Println("   ✅ 小值类型（int、bool等）")
	fmt.Println("   ✅ 不需要修改的值")
	fmt.Println("   ✅ 简单的值传递场景")
	fmt.Println()

	fmt.Println("3. 安全使用指针:")
	fmt.Println("   ✅ 总是检查nil指针")
	fmt.Println("   ✅ 使用new()创建指针")
	fmt.Println("   ✅ 避免悬空指针")
	fmt.Println()

	// ============================================
	// 12. 总结
	// ============================================
	fmt.Println("=== 12. 总结 ===")
	fmt.Println()
	fmt.Println("1. 指针的形象化理解:")
	fmt.Println("   ✅ 指针就像'告示牌'，存储内存地址")
	fmt.Println("   ✅ 通过地址访问真正的数据")
	fmt.Println()
	fmt.Println("2. 声明和使用指针:")
	fmt.Println("   ✅ & 操作符：取地址")
	fmt.Println("   ✅ * 操作符：解引用")
	fmt.Println()
	fmt.Println("3. 指针的使用时机:")
	fmt.Println("   ✅ 需要修改原始数据")
	fmt.Println("   ✅ 避免大值拷贝")
	fmt.Println("   ✅ 实现数据共享")
	fmt.Println("   ✅ 构建复杂数据结构")
	fmt.Println()
	fmt.Println("4. Go指针的关键特性:")
	fmt.Println("   ✅ 不支持指针算术运算")
	fmt.Println("   ✅ 不能随意转换指针类型")
	fmt.Println("   ✅ 更安全，避免内存错误")
	fmt.Println()
	fmt.Println("5. 值传递 vs 指针传递:")
	fmt.Println("   ✅ 值传递：复制值，不修改原值")
	fmt.Println("   ✅ 指针传递：传递地址，可以修改原值")
	fmt.Println("   ✅ 大结构体使用指针传递提高性能")
	fmt.Println()
	fmt.Println("6. nil指针:")
	fmt.Println("   ✅ 未初始化的指针值为nil")
	fmt.Println("   ✅ 解引用nil指针会panic")
	fmt.Println("   ✅ 使用前需要检查nil")
	fmt.Println()
}

// ============================================
// 辅助类型和函数
// ============================================

// LargeStruct 大结构体
type LargeStruct struct {
	Data [1000]int
}

// SmallStruct 小结构体
type SmallStruct struct {
	Value int
}

// Node 链表节点
type Node struct {
	Value int
	Next  *Node
}

// Person 人员结构体
type Person struct {
	Name string
	Age  int
}

// SetAgeValue 值接收者（不会修改原值）
func (p Person) SetAgeValue(age int) {
	p.Age = age
}

// SetAgePointer 指针接收者（会修改原值）
func (p *Person) SetAgePointer(age int) {
	p.Age = age
}

// modifyValue 值传递（不会修改原值）
func modifyValue(x int) {
	x = 100
}

// modifyPointer 指针传递（会修改原值）
func modifyPointer(x *int) {
	*x = 100
}

// passByValue 值传递
func passByValue(n int) {
	n = 999
}

// passByPointer 指针传递
func passByPointer(n *int) {
	*n = 999
}

// processSmallStruct 处理小结构体（值传递）
func processSmallStruct(s SmallStruct) {
	s.Value = 999
	fmt.Printf("    函数内部: %+v\n", s)
}

// processLargeStruct 处理大结构体（指针传递）
func processLargeStruct(s *LargeStruct) {
	s.Data[0] = 999
	fmt.Printf("    函数内部: Data[0]=%d\n", s.Data[0])
}

// swap 交换两个值
func swap(a, b *int) {
	*a, *b = *b, *a
}

// divide 除法（通过指针返回结果）
func divide(a, b int, result *int) bool {
	if b == 0 {
		return false
	}
	*result = a / b
	return true
}

