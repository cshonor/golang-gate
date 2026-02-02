// 示例：Go 语言的函数
// 演示函数声明、参数、返回值、多返回值、高阶函数和递归

package main

import "fmt"

func main() {
	// ============================================
	// 1. 函数的基本结构
	// ============================================
	fmt.Println("=== 函数的基本结构 ===")

	// Go 语言中函数的声明格式：
	// func 函数名(参数列表) (返回值列表) {
	//     // 函数体
	// }

	// 示例：简单的加法函数
	result := Add(10, 20)
	fmt.Printf("Add(10, 20) = %d\n", result)

	fmt.Println()

	// ============================================
	// 2. 温度转换函数（REMS 项目）
	// ============================================
	fmt.Println("=== 温度转换函数（REMS 项目）===")

	// 火星表面的典型温度（摄氏度）
	tempC := -60.0
	tempF := CToF(tempC)
	fmt.Printf("火星表面温度：%.1f°C = %.1f°F\n", tempC, tempF)

	// 反向转换
	tempC2 := FToC(tempF)
	fmt.Printf("反向转换：%.1f°F = %.1f°C\n", tempF, tempC2)

	// 更多温度转换示例
	temperatures := []float64{-60.0, -40.0, 0.0, 20.0, 100.0}
	fmt.Println("\n温度转换表（摄氏度 -> 华氏度）:")
	for _, tc := range temperatures {
		tf := CToF(tc)
		fmt.Printf("  %.1f°C = %.1f°F\n", tc, tf)
	}

	fmt.Println()

	// ============================================
	// 3. 多返回值
	// ============================================
	fmt.Println("=== 多返回值 ===")

	// Go 支持函数返回多个值，通常用于同时返回结果和错误信息
	result1, err1 := Divide(10, 2)
	if err1 != nil {
		fmt.Printf("错误: %v\n", err1)
	} else {
		fmt.Printf("Divide(10, 2) = %.2f\n", result1)
	}

	result2, err2 := Divide(10, 0)
	if err2 != nil {
		fmt.Printf("Divide(10, 0) 错误: %v\n", err2)
	} else {
		fmt.Printf("Divide(10, 0) = %.2f\n", result2)
	}

	// 多返回值示例：计算商和余数
	quotient, remainder := DivideWithRemainder(17, 5)
	fmt.Printf("DivideWithRemainder(17, 5) = 商: %d, 余数: %d\n", quotient, remainder)

	// 忽略某个返回值（使用 _）
	_, remainder2 := DivideWithRemainder(20, 3)
	fmt.Printf("只获取余数: %d\n", remainder2)

	fmt.Println()

	// ============================================
	// 4. 命名返回值
	// ============================================
	fmt.Println("=== 命名返回值 ===")

	// Go 支持命名返回值，可以在函数体中直接使用
	sum, product := Calculate(3, 4)
	fmt.Printf("Calculate(3, 4) = 和: %d, 积: %d\n", sum, product)

	// 命名返回值可以简化 return 语句
	result3 := NamedReturn(5)
	fmt.Printf("NamedReturn(5) = %d\n", result3)

	fmt.Println()

	// ============================================
	// 5. 可变参数函数
	// ============================================
	fmt.Println("=== 可变参数函数 ===")

	// 使用 ... 表示可变参数
	sum1 := Sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum(1, 2, 3, 4, 5) = %d\n", sum1)

	sum2 := Sum(10, 20)
	fmt.Printf("Sum(10, 20) = %d\n", sum2)

	// 可变参数也可以传入切片
	numbers := []int{1, 2, 3, 4, 5}
	sum3 := Sum(numbers...)
	fmt.Printf("Sum([]int{1,2,3,4,5}...) = %d\n", sum3)

	fmt.Println()

	// ============================================
	// 6. 函数作为参数（高阶函数）
	// ============================================
	fmt.Println("=== 函数作为参数（高阶函数）===")

	// Go 支持将函数作为参数传递给其他函数
	numbers2 := []int{1, 2, 3, 4, 5}

	// 使用不同的操作函数
	squared := ApplyOperation(numbers2, Square)
	fmt.Printf("Square([1,2,3,4,5]) = %v\n", squared)

	doubled := ApplyOperation(numbers2, Double)
	fmt.Printf("Double([1,2,3,4,5]) = %v\n", doubled)

	// 使用匿名函数
	cubed := ApplyOperation(numbers2, func(x int) int {
		return x * x * x
	})
	fmt.Printf("Cube([1,2,3,4,5]) = %v\n", cubed)

	fmt.Println()

	// ============================================
	// 7. 递归函数
	// ============================================
	fmt.Println("=== 递归函数 ===")

	// 计算阶乘
	factorial5 := Factorial(5)
	fmt.Printf("Factorial(5) = %d\n", factorial5)

	// 计算斐波那契数列
	fmt.Println("斐波那契数列（前10项）:")
	for i := 0; i < 10; i++ {
		fmt.Printf("  F(%d) = %d\n", i, Fibonacci(i))
	}

	// 计算最大公约数（欧几里得算法）
	gcd := GCD(48, 18)
	fmt.Printf("GCD(48, 18) = %d\n", gcd)

	fmt.Println()

	// ============================================
	// 8. 函数类型
	// ============================================
	fmt.Println("=== 函数类型 ===")

	// 定义函数类型
	type MathOperation func(int, int) int

	// 使用函数类型
	var op MathOperation = Add
	result4 := op(10, 20)
	fmt.Printf("使用函数类型: op(10, 20) = %d\n", result4)

	// 函数类型可以作为变量
	op = Multiply
	result5 := op(10, 20)
	fmt.Printf("切换函数: op(10, 20) = %d\n", result5)

	fmt.Println()

	// ============================================
	// 9. 闭包
	// ============================================
	fmt.Println("=== 闭包 ===")

	// 闭包：函数可以访问外部作用域的变量
	counter := CreateCounter()
	fmt.Printf("counter() = %d\n", counter())
	fmt.Printf("counter() = %d\n", counter())
	fmt.Printf("counter() = %d\n", counter())

	// 创建多个独立的计数器
	counter1 := CreateCounter()
	counter2 := CreateCounter()
	fmt.Printf("counter1() = %d\n", counter1())
	fmt.Printf("counter2() = %d\n", counter2())
	fmt.Printf("counter1() = %d\n", counter1())
	fmt.Printf("counter2() = %d\n", counter2())

	fmt.Println()

	// ============================================
	// 10. 延迟执行（defer）
	// ============================================
	fmt.Println("=== 延迟执行（defer）===")

	// defer 语句会在函数返回前执行
	fmt.Println("调用 DeferExample:")
	DeferExample()

	fmt.Println()

	// ============================================
	// 11. 函数总结
	// ============================================
	fmt.Println("=== 函数总结 ===")
	fmt.Println("1. 函数基本结构: func 函数名(参数) (返回值) { 函数体 }")
	fmt.Println("2. 支持多返回值，常用于返回结果和错误")
	fmt.Println("3. 支持命名返回值，简化 return 语句")
	fmt.Println("4. 支持可变参数，使用 ... 表示")
	fmt.Println("5. 函数可以作为参数传递（高阶函数）")
	fmt.Println("6. 支持递归函数")
	fmt.Println("7. 函数可以作为类型使用")
	fmt.Println("8. 支持闭包，函数可以访问外部变量")
	fmt.Println("9. defer 语句用于延迟执行")
}

// ============================================
// 基础函数示例
// ============================================

// Add 简单的加法函数
func Add(a, b int) int {
	return a + b
}

// Multiply 乘法函数
func Multiply(a, b int) int {
	return a * b
}

// ============================================
// 温度转换函数
// ============================================

// CToF 摄氏度转华氏度
// 公式：F = C × 1.8 + 32
func CToF(celsius float64) float64 {
	return celsius*1.8 + 32
}

// FToC 华氏度转摄氏度
// 公式：C = (F - 32) / 1.8
func FToC(fahrenheit float64) float64 {
	return (fahrenheit - 32) / 1.8
}

// ============================================
// 多返回值函数
// ============================================

// Divide 除法函数，返回结果和错误
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// DivideWithRemainder 整数除法，返回商和余数
func DivideWithRemainder(a, b int) (int, int) {
	return a / b, a % b
}

// ============================================
// 命名返回值
// ============================================

// Calculate 计算两个数的和与积（使用命名返回值）
func Calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // 可以省略返回值，直接 return
}

// NamedReturn 命名返回值示例
func NamedReturn(n int) (result int) {
	result = n * 2
	return // 直接 return，会自动返回 result
}

// ============================================
// 可变参数函数
// ============================================

// Sum 计算多个整数的和（可变参数）
func Sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// ============================================
// 高阶函数
// ============================================

// ApplyOperation 将操作函数应用到切片中的每个元素
func ApplyOperation(numbers []int, op func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = op(num)
	}
	return result
}

// Square 平方函数
func Square(x int) int {
	return x * x
}

// Double 翻倍函数
func Double(x int) int {
	return x * 2
}

// ============================================
// 递归函数
// ============================================

// Factorial 计算阶乘（递归实现）
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci 计算斐波那契数列（递归实现）
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// GCD 计算最大公约数（欧几里得算法，递归实现）
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// ============================================
// 闭包
// ============================================

// CreateCounter 创建一个计数器闭包
func CreateCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ============================================
// defer 示例
// ============================================

// DeferExample 演示 defer 的使用
func DeferExample() {
	fmt.Println("  1. 函数开始")
	defer fmt.Println("  3. defer 语句（函数返回前执行）")
	fmt.Println("  2. 函数执行中")
	// 函数返回时，defer 语句会执行
}

