// 示例：Go 语言的一等函数（First-Class Functions）
// 演示函数作为一等公民：赋值给变量、作为参数传递、作为返回值

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ============================================
	// 1. 什么是"一等函数"
	// ============================================
	fmt.Println("=== 什么是\"一等函数\" ===")

	fmt.Println("在 Go 中，函数被视为\"一等公民\"，意味着它可以：")
	fmt.Println("  1. 像整数、字符串一样赋值给变量")
	fmt.Println("  2. 作为参数传递给其他函数")
	fmt.Println("  3. 作为返回值从其他函数返回")
	fmt.Println("  4. 存储在数据结构（如切片、映射）中")
	fmt.Println()

	// ============================================
	// 2. 函数赋值给变量
	// ============================================
	fmt.Println("=== 函数赋值给变量 ===")

	// 方式1：直接赋值函数名（函数值）
	var addFunc func(int, int) int = Add
	result1 := addFunc(10, 20)
	fmt.Printf("addFunc(10, 20) = %d\n", result1)

	// 方式2：使用短变量声明
	multiplyFunc := Multiply
	result2 := multiplyFunc(5, 6)
	fmt.Printf("multiplyFunc(5, 6) = %d\n", result2)

	// 方式3：赋值匿名函数
	squareFunc := func(x int) int {
		return x * x
	}
	result3 := squareFunc(7)
	fmt.Printf("squareFunc(7) = %d\n", result3)

	// 函数可以重新赋值
	operation := Add
	fmt.Printf("operation(3, 4) = %d\n", operation(3, 4))
	operation = Multiply
	fmt.Printf("operation(3, 4) = %d\n", operation(3, 4))

	fmt.Println()

	// ============================================
	// 3. 函数作为参数传递（高阶函数）
	// ============================================
	fmt.Println("=== 函数作为参数传递（高阶函数）===")

	numbers := []int{1, 2, 3, 4, 5}

	// 传递不同的函数实现不同的操作
	squared := ApplyOperation(numbers, Square)
	fmt.Printf("Square([1,2,3,4,5]) = %v\n", squared)

	doubled := ApplyOperation(numbers, Double)
	fmt.Printf("Double([1,2,3,4,5]) = %v\n", doubled)

	// 传递匿名函数
	cubed := ApplyOperation(numbers, func(x int) int {
		return x * x * x
	})
	fmt.Printf("Cube([1,2,3,4,5]) = %v\n", cubed)

	// 火星温度监测站示例
	fmt.Println("\n火星温度监测站示例:")
	tempC := readSensor()
	fmt.Printf("传感器原始读数：%.2f °C\n", tempC)

	kelvin := processTemperature(readSensor, celsiusToKelvin)
	fmt.Printf("转换为开尔文：%.2f K\n", kelvin)

	fahrenheit := processTemperature(readSensor, celsiusToFahrenheit)
	fmt.Printf("转换为华氏度：%.2f °F\n", fahrenheit)

	fmt.Println()

	// ============================================
	// 4. 函数作为返回值（工厂函数）
	// ============================================
	fmt.Println("=== 函数作为返回值（工厂函数）===")

	// 创建带偏移量的温度转换函数
	offsetConverter := createOffsetConverter(5.0)
	adjusted := offsetConverter(-60.0)
	fmt.Printf("温度偏移+5后：%.2f °C\n", adjusted)

	// 创建不同的转换器
	converter1 := createOffsetConverter(10.0)
	converter2 := createOffsetConverter(-5.0)
	fmt.Printf("偏移+10: %.2f °C\n", converter1(-60.0))
	fmt.Printf("偏移-5: %.2f °C\n", converter2(-60.0))

	// 创建乘法器工厂函数
	doubleMultiplier := createMultiplier(2.0)
	tripleMultiplier := createMultiplier(3.0)
	fmt.Printf("2倍: %.2f\n", doubleMultiplier(5.0))
	fmt.Printf("3倍: %.2f\n", tripleMultiplier(5.0))

	fmt.Println()

	// ============================================
	// 5. 闭包（Closure）
	// ============================================
	fmt.Println("=== 闭包（Closure）===")

	// 闭包：函数可以捕获并访问外部函数的变量
	counter1 := createCounter()
	fmt.Printf("counter1() = %d\n", counter1())
	fmt.Printf("counter1() = %d\n", counter1())
	fmt.Printf("counter1() = %d\n", counter1())

	// 每个闭包都有自己独立的变量
	counter2 := createCounter()
	fmt.Printf("counter2() = %d\n", counter2())
	fmt.Printf("counter1() = %d\n", counter1()) // counter1 的状态保持独立

	// 带参数的闭包
	accumulator := createAccumulator(10)
	fmt.Printf("accumulator(5) = %d\n", accumulator(5))
	fmt.Printf("accumulator(3) = %d\n", accumulator(3))
	fmt.Printf("accumulator(7) = %d\n", accumulator(7))

	fmt.Println()

	// ============================================
	// 6. 函数存储在数据结构中
	// ============================================
	fmt.Println("=== 函数存储在数据结构中 ===")

	// 函数切片
	operations := []func(int, int) int{
		Add,
		Multiply,
		Subtract,
	}

	for i, op := range operations {
		result := op(10, 5)
		fmt.Printf("操作 %d: %d\n", i+1, result)
	}

	// 函数映射（策略模式）
	strategies := map[string]func(int, int) int{
		"add":      Add,
		"multiply": Multiply,
		"subtract": Subtract,
	}

	fmt.Println("\n策略模式示例:")
	for name, strategy := range strategies {
		result := strategy(8, 4)
		fmt.Printf("  %s(8, 4) = %d\n", name, result)
	}

	fmt.Println()

	// ============================================
	// 7. 回调函数示例
	// ============================================
	fmt.Println("=== 回调函数示例 ===")

	// 模拟异步操作，完成后调用回调函数
	processAsync("数据处理", func(result string) {
		fmt.Printf("回调函数收到结果: %s\n", result)
	})

	// 多个回调函数
	callbacks := []func(int){
		func(x int) { fmt.Printf("回调1: %d\n", x) },
		func(x int) { fmt.Printf("回调2: %d * 2 = %d\n", x, x*2) },
		func(x int) { fmt.Printf("回调3: %d^2 = %d\n", x, x*x) },
	}

	fmt.Println("\n多个回调函数:")
	for _, callback := range callbacks {
		callback(5)
	}

	fmt.Println()

	// ============================================
	// 8. 函数类型定义
	// ============================================
	fmt.Println("=== 函数类型定义 ===")

	// 定义函数类型，提高代码可读性
	type MathFunc func(int, int) int
	type TransformFunc func(int) int

	var mathOp MathFunc = Add
	fmt.Printf("MathFunc Add(10, 5) = %d\n", mathOp(10, 5))

	mathOp = Multiply
	fmt.Printf("MathFunc Multiply(10, 5) = %d\n", mathOp(10, 5))

	// 使用函数类型作为参数
	transform := ApplyTransform([]int{1, 2, 3}, func(x int) int {
		return x * 2
	})
	fmt.Printf("ApplyTransform([1,2,3], x*2) = %v\n", transform)

	fmt.Println()

	// ============================================
	// 9. 一等函数的实际应用场景
	// ============================================
	fmt.Println("=== 一等函数的实际应用场景 ===")

	// 场景1：数据过滤
	numbers2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := Filter(numbers2, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("偶数: %v\n", evenNumbers)

	// 场景2：数据映射
	strings := []string{"hello", "world", "go"}
	lengths := Map(strings, func(s string) int {
		return len(s)
	})
	fmt.Printf("字符串长度: %v\n", lengths)

	// 场景3：数据归约
	sum := Reduce(numbers2, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Printf("求和: %d\n", sum)

	fmt.Println()

	// ============================================
	// 10. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. 函数可以赋值给变量")
	fmt.Println("2. 函数可以作为参数传递（高阶函数）")
	fmt.Println("3. 函数可以作为返回值（工厂函数）")
	fmt.Println("4. 函数可以存储在数据结构中")
	fmt.Println("5. 闭包可以捕获外部变量")
	fmt.Println("6. 回调函数用于异步操作和事件处理")
	fmt.Println("7. 策略模式通过传递不同函数实现不同行为")
	fmt.Println("8. 一等函数提供了极大的代码灵活性")
}

// ============================================
// 基础函数
// ============================================

// Add 加法函数
func Add(a, b int) int {
	return a + b
}

// Multiply 乘法函数
func Multiply(a, b int) int {
	return a * b
}

// Subtract 减法函数
func Subtract(a, b int) int {
	return a - b
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
// 高阶函数：函数作为参数
// ============================================

// ApplyOperation 将操作函数应用到切片中的每个元素
func ApplyOperation(numbers []int, op func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = op(num)
	}
	return result
}

// ============================================
// 火星温度监测站示例
// ============================================

// readSensor 模拟温度传感器读取函数
func readSensor() float64 {
	rand.Seed(time.Now().UnixNano())
	// 火星表面温度范围：-125°C ~ 20°C
	return rand.Float64()*145 - 125
}

// celsiusToKelvin 摄氏度转开尔文
func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

// celsiusToFahrenheit 摄氏度转华氏度
func celsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

// processTemperature 处理温度数据并应用转换函数（高阶函数）
func processTemperature(read func() float64, convert func(float64) float64) float64 {
	temp := read()
	return convert(temp)
}

// ============================================
// 工厂函数：函数作为返回值
// ============================================

// createOffsetConverter 创建带偏移量的温度转换函数
func createOffsetConverter(offset float64) func(float64) float64 {
	return func(c float64) float64 {
		return c + offset
	}
}

// createMultiplier 创建乘法器函数
func createMultiplier(factor float64) func(float64) float64 {
	return func(x float64) float64 {
		return x * factor
	}
}

// ============================================
// 闭包示例
// ============================================

// createCounter 创建计数器闭包
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// createAccumulator 创建累加器闭包
func createAccumulator(initial int) func(int) int {
	sum := initial
	return func(x int) int {
		sum += x
		return sum
	}
}

// ============================================
// 回调函数示例
// ============================================

// processAsync 模拟异步处理，完成后调用回调函数
func processAsync(data string, callback func(string)) {
	// 模拟处理时间
	time.Sleep(100 * time.Millisecond)
	result := fmt.Sprintf("处理完成: %s", data)
	callback(result)
}

// ============================================
// 函数类型定义
// ============================================

// TransformFunc 变换函数类型
type TransformFunc func(int) int

// ApplyTransform 应用变换函数
func ApplyTransform(numbers []int, transform TransformFunc) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = transform(num)
	}
	return result
}

// ============================================
// 实际应用场景：函数式编程模式
// ============================================

// Filter 过滤函数：根据条件过滤元素
func Filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

// Map 映射函数：将每个元素转换为新值
func Map(strings []string, mapper func(string) int) []int {
	result := make([]int, len(strings))
	for i, s := range strings {
		result[i] = mapper(s)
	}
	return result
}

// Reduce 归约函数：将切片归约为单个值
func Reduce(numbers []int, initial int, reducer func(int, int) int) int {
	acc := initial
	for _, num := range numbers {
		acc = reducer(acc, num)
	}
	return acc
}

