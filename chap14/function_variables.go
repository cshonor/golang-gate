// 示例：函数变量声明
// 演示用 var 声明函数变量、方法变量、带参数和多返回值的函数变量

package main

import "fmt"

// ============================================
// 基础函数定义
// ============================================

// Add 加法函数
func Add(a, b int) int {
	return a + b
}

// Divide 除法函数（带多返回值）
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// ============================================
// 自定义类型和方法
// ============================================

// Kelvin 开尔文温度类型
type Kelvin float64

// ToCelsius 开尔文转摄氏度（方法）
func (k Kelvin) ToCelsius() float64 {
	return float64(k - 273.15)
}

// ToFahrenheit 开尔文转华氏度（方法，带参数）
func (k Kelvin) ToFahrenheit(offset float64) float64 {
	return float64(k-273.15)*1.8 + 32 + offset
}

func main() {
	// ============================================
	// 1. 函数变量 vs 函数声明
	// ============================================
	fmt.Println("=== 函数变量 vs 函数声明 ===")

	fmt.Println("函数声明：用 func 关键字定义新函数（创建工具）")
	fmt.Println("  func add(a, b int) int { return a + b }")
	fmt.Println()
	fmt.Println("函数变量：用 var 或 := 声明变量，类型是函数类型（存放工具）")
	fmt.Println("  var f func(int, int) int")
	fmt.Println("  f = add")
	fmt.Println()

	// ============================================
	// 2. 用 var 声明函数变量
	// ============================================
	fmt.Println("=== 用 var 声明函数变量 ===")

	// 方式1：var 声明，然后赋值
	var f1 func(int, int) int
	f1 = Add
	fmt.Printf("方式1 - var 声明: f1(10, 20) = %d\n", f1(10, 20))

	// 方式2：var 声明并初始化
	var f2 func(int, int) int = Add
	fmt.Printf("方式2 - var 声明并初始化: f2(10, 20) = %d\n", f2(10, 20))

	// 方式3：短变量声明
	f3 := Add
	fmt.Printf("方式3 - 短变量声明: f3(10, 20) = %d\n", f3(10, 20))

	// 方式4：var 声明匿名函数
	var f4 func(int, int) int = func(a, b int) int {
		return a * b
	}
	fmt.Printf("方式4 - var 声明匿名函数: f4(10, 20) = %d\n", f4(10, 20))

	fmt.Println()

	// ============================================
	// 3. 带参数和多返回值的函数变量
	// ============================================
	fmt.Println("=== 带参数和多返回值的函数变量 ===")

	// 声明函数变量，类型要和函数完全匹配
	var divideFunc func(float64, float64) (float64, error)
	divideFunc = Divide

	// 调用函数变量
	result1, err1 := divideFunc(10, 2)
	if err1 != nil {
		fmt.Printf("错误: %v\n", err1)
	} else {
		fmt.Printf("divideFunc(10, 2) = %.2f\n", result1)
	}

	// 错误情况
	result2, err2 := divideFunc(10, 0)
	if err2 != nil {
		fmt.Printf("divideFunc(10, 0) 错误: %v\n", err2)
	} else {
		fmt.Printf("结果: %.2f\n", result2)
	}

	fmt.Println()

	// ============================================
	// 4. 方法变量（特殊的函数变量）
	// ============================================
	fmt.Println("=== 方法变量（特殊的函数变量）===")

	// 将方法赋值给变量，自动绑定接收者
	temp := Kelvin(300)
	fmt.Printf("原始温度: %.2f K\n", temp)

	// 方式1：无参数方法变量
	converter1 := temp.ToCelsius
	fmt.Printf("方法变量 converter1() = %.2f °C\n", converter1())
	fmt.Println("  说明: 调用时不需要再传接收者，已经绑定到 temp")

	// 方式2：带参数方法变量
	converter2 := temp.ToFahrenheit
	fmt.Printf("方法变量 converter2(0) = %.2f °F\n", converter2(0))
	fmt.Printf("方法变量 converter2(5) = %.2f °F (带偏移量)\n", converter2(5))
	fmt.Println("  说明: 方法参数仍然需要传递")

	// 方式3：用 var 声明方法变量
	var converter3 func() float64
	converter3 = temp.ToCelsius
	fmt.Printf("var 声明的方法变量 converter3() = %.2f °C\n", converter3())

	fmt.Println()

	// ============================================
	// 5. 方法变量的绑定时机
	// ============================================
	fmt.Println("=== 方法变量的绑定时机 ===")

	temp1 := Kelvin(300)
	temp2 := Kelvin(400)

	// 方法变量在赋值时就绑定了接收者
	converter4 := temp1.ToCelsius
	converter5 := temp2.ToCelsius

	fmt.Printf("temp1 = %.2f K, converter4() = %.2f °C\n", temp1, converter4())
	fmt.Printf("temp2 = %.2f K, converter5() = %.2f °C\n", temp2, converter5())

	// 即使修改 temp1，converter4 仍然绑定的是原来的值
	temp1 = 500
	fmt.Printf("修改 temp1 后: temp1 = %.2f K\n", temp1)
	fmt.Printf("但 converter4() 仍然是: %.2f °C (绑定的是旧值)\n", converter4())

	fmt.Println()

	// ============================================
	// 6. 函数变量类型定义
	// ============================================
	fmt.Println("=== 函数变量类型定义 ===")

	// 定义函数类型，提高代码可读性
	type MathFunc func(int, int) int
	type DivideFunc func(float64, float64) (float64, error)

	var mathOp MathFunc = Add
	fmt.Printf("MathFunc: mathOp(10, 5) = %d\n", mathOp(10, 5))

	var divOp DivideFunc = Divide
	result, err := divOp(20, 4)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("DivideFunc: divOp(20, 4) = %.2f\n", result)
	}

	fmt.Println()

	// ============================================
	// 7. 函数变量声明速查表
	// ============================================
	fmt.Println("=== 函数变量声明速查表 ===")
	fmt.Println()
	fmt.Println("1. 无参数无返回值函数:")
	fmt.Println("   var f func()")
	fmt.Println("   f = someFunc")
	fmt.Println()
	fmt.Println("2. 有参数有返回值函数:")
	fmt.Println("   var f func(int, int) int")
	fmt.Println("   f = Add")
	fmt.Println()
	fmt.Println("3. 多返回值函数:")
	fmt.Println("   var f func(float64, float64) (float64, error)")
	fmt.Println("   f = Divide")
	fmt.Println()
	fmt.Println("4. 方法变量（无参数）:")
	fmt.Println("   var f func() float64")
	fmt.Println("   f = temp.ToCelsius")
	fmt.Println()
	fmt.Println("5. 方法变量（有参数）:")
	fmt.Println("   var f func(float64) float64")
	fmt.Println("   f = temp.ToFahrenheit")
	fmt.Println()
	fmt.Println("6. 使用类型别名:")
	fmt.Println("   type MathFunc func(int, int) int")
	fmt.Println("   var f MathFunc = Add")
	fmt.Println()

	// ============================================
	// 8. 实际应用示例
	// ============================================
	fmt.Println("=== 实际应用示例 ===")

	// 示例：根据条件选择不同的函数
	var operation func(int, int) int
	useMultiply := true

	if useMultiply {
		operation = func(a, b int) int {
			return a * b
		}
	} else {
		operation = Add
	}

	fmt.Printf("根据条件选择函数: operation(5, 6) = %d\n", operation(5, 6))

	// 示例：函数变量切片
	operations := []func(int, int) int{
		Add,
		func(a, b int) int { return a * b },
		func(a, b int) int { return a - b },
	}

	fmt.Println("函数变量切片:")
	for i, op := range operations {
		fmt.Printf("  操作 %d: op(10, 5) = %d\n", i+1, op(10, 5))
	}

	fmt.Println()

	// ============================================
	// 9. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. 函数变量用 var 或 := 声明，类型是函数类型")
	fmt.Println("2. 函数变量的类型必须和函数签名完全匹配")
	fmt.Println("3. 方法变量会自动绑定接收者，变成普通函数变量")
	fmt.Println("4. 方法变量在赋值时就绑定了接收者的值")
	fmt.Println("5. 使用类型别名可以提高代码可读性")
	fmt.Println("6. 函数变量可以实现策略模式、回调等设计模式")
}

