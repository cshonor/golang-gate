// 示例：美化的计算器
// 演示基本的数学运算和打印

package main

import "fmt"

func main() {
	// 基本运算示例
	fmt.Println("=== 基本运算 ===")
	
	// 加法
	fmt.Println("10 + 5 =", 10+5)
	
	// 减法
	fmt.Println("10 - 5 =", 10-5)
	
	// 乘法
	fmt.Println("10 * 5 =", 10*5)
	
	// 除法
	fmt.Println("10 / 5 =", 10/5)
	
	// 浮点数除法（注意：整数除法会截断小数部分）
	fmt.Println("10 / 3 =", 10/3)           // 整数除法，结果是 3
	fmt.Println("10.0 / 3.0 =", 10.0/3.0)   // 浮点数除法，结果是 3.333...
	
	fmt.Println("\n=== 火星计算示例 ===")
	
	// 地球体重（磅）
	earthWeight := 164.0
	
	// 火星重力系数（火星重力是地球的 37.83%）
	marsGravity := 0.3783
	
	// 计算火星体重
	marsWeight := earthWeight * marsGravity
	fmt.Printf("地球体重: %.1f 磅\n", earthWeight)
	fmt.Printf("火星体重: %.2f 磅\n", marsWeight)
	fmt.Printf("在火星上，你会轻 %.2f 磅\n", earthWeight-marsWeight)
	
	fmt.Println("\n=== 年龄计算示例 ===")
	
	// 地球年龄（年）
	earthAge := 41
	
	// 地球一年天数
	earthDaysPerYear := 365
	
	// 火星一年天数（火星一年约 687 个地球日）
	marsDaysPerYear := 687
	
	// 计算火星年龄
	// 先算出总天数，再除以火星一年的天数
	marsAge := earthAge * earthDaysPerYear / marsDaysPerYear
	fmt.Printf("地球年龄: %d 年\n", earthAge)
	fmt.Printf("火星年龄: %d 年\n", marsAge)
	fmt.Printf("在火星上，你会年轻 %d 年\n", earthAge-marsAge)
}

