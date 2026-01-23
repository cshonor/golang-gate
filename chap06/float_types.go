// 示例：浮点数类型
// 演示 float32 和 float64 的定义和使用

package main

import "fmt"

func main() {
	// ============================================
	// 1. 浮点数的类型推断
	// ============================================
	fmt.Println("=== 浮点数类型推断 ===")
	
	// 如果直接赋值一个带小数点的数字，Go 会自动推断为 float64
	var pi = 3.14
	fmt.Printf("pi 的类型: %T, 值: %v\n", pi, pi)  // float64
	
	// 使用 := 也是同样的效果
	pi2 := 3.14
	fmt.Printf("pi2 的类型: %T, 值: %v\n", pi2, pi2)  // float64
	
	// 也可以显式声明类型
	var pi3 float32 = 3.14
	fmt.Printf("pi3 的类型: %T, 值: %v\n", pi3, pi3)  // float32
	
	var pi4 float64 = 3.14
	fmt.Printf("pi4 的类型: %T, 值: %v\n", pi4, pi4)  // float64
	
	fmt.Println()
	
	// ============================================
	// 2. float32 和 float64 的区别
	// ============================================
	fmt.Println("=== float32 vs float64 ===")
	
	// float32：单精度，4 字节，约 6-7 位有效数字
	var num1 float32 = 365.2425
	fmt.Printf("float32: %v\n", num1)
	
	// float64：双精度，8 字节，约 15-16 位有效数字
	var num2 float64 = 365.2425
	fmt.Printf("float64: %v\n", num2)
	
	// 注意：Go 默认推断为 float64，因为精度更高
	var num3 = 365.2425  // 自动推断为 float64
	fmt.Printf("自动推断: %T, %v\n", num3, num3)
	
	fmt.Println()
	
	// ============================================
	// 3. 零值问题
	// ============================================
	fmt.Println("=== 零值问题 ===")
	
	// 整数类型的零值是 0
	var num int
	fmt.Printf("int 零值: %v (类型: %T)\n", num, num)
	
	// 浮点数类型的零值是 0.0
	var pi5 float64
	fmt.Printf("float64 零值: %v (类型: %T)\n", pi5, pi5)
	
	// 赋值 0 会被推断为 int
	var num4 = 0
	fmt.Printf("0 的类型: %T\n", num4)  // int
	
	// 赋值 0.0 会被推断为 float64
	var num5 = 0.0
	fmt.Printf("0.0 的类型: %T\n", num5)  // float64
	
	fmt.Println()
	
	// ============================================
	// 4. 浮点数运算
	// ============================================
	fmt.Println("=== 浮点数运算 ===")
	
	// 1.0 / 3 的结果是浮点数
	result := 1.0 / 3
	fmt.Printf("1.0 / 3 = %v\n", result)
	fmt.Printf("1.0 / 3 = %.3f\n", result)  // 保留 3 位小数
	fmt.Printf("1.0 / 3 = %.10f\n", result) // 保留 10 位小数
	
	fmt.Println()
	
	// ============================================
	// 5. 打印浮点数的格式
	// ============================================
	fmt.Println("=== 打印浮点数格式 ===")
	
	value := 1.0 / 3
	
	// %f：默认格式
	fmt.Printf("默认: %f\n", value)
	
	// %.3f：保留 3 位小数
	fmt.Printf("保留3位: %.3f\n", value)
	
	// %4.2f：总宽度 4，保留 2 位小数
	fmt.Printf("宽度4保留2位: %4.2f\n", value)
	
	// %v：通用占位符，自动识别类型
	fmt.Printf("通用: %v\n", value)
	
	// fmt.Println 也会显示完整精度
	fmt.Println("Println:", value)
}

