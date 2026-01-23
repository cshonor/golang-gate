// 示例：浮点数比较
// 演示如何正确比较浮点数

package main

import (
	"fmt"
	"math"
)

func main() {
	// ============================================
	// 1. 错误的比较方式：直接用 ==
	// ============================================
	fmt.Println("=== 错误的比较方式：直接用 == ===")
	
	pinkBlack := 0.1
	pinkBlack += 0.2
	
	// ❌ 错误：直接用 == 比较浮点数
	// 因为精度误差，结果会是 false
	fmt.Printf("pinkBlack = %.20f\n", pinkBlack)
	fmt.Printf("pinkBlack == 0.3? %t\n", pinkBlack == 0.3)  // false
	
	fmt.Println()
	
	// ============================================
	// 2. 正确的比较方式：使用 math.Abs
	// ============================================
	fmt.Println("=== 正确的比较方式：使用 math.Abs ===")
	
	// ✅ 正确：判断两个数的差值是否小于一个很小的数（epsilon）
	// epsilon 通常取 1e-9（0.000000001）
	epsilon := 0.0001  // 也可以使用 1e-9
	
	isEqual := math.Abs(pinkBlack-0.3) < epsilon
	fmt.Printf("pinkBlack ≈ 0.3? %t\n", isEqual)  // true
	
	fmt.Println()
	
	// ============================================
	// 3. 不同 epsilon 值的影响
	// ============================================
	fmt.Println("=== 不同 epsilon 值的影响 ===")
	
	diff := math.Abs(pinkBlack - 0.3)
	fmt.Printf("实际差值: %.20f\n", diff)
	
	// 使用不同的 epsilon
	fmt.Printf("epsilon = 1e-9 (0.000000001): %t\n", diff < 1e-9)
	fmt.Printf("epsilon = 1e-6 (0.000001): %t\n", diff < 1e-6)
	fmt.Printf("epsilon = 0.0001: %t\n", diff < 0.0001)
	fmt.Printf("epsilon = 0.001: %t\n", diff < 0.001)
	
	fmt.Println()
	
	// ============================================
	// 4. 实际应用：比较函数
	// ============================================
	fmt.Println("=== 实际应用：比较函数 ===")
	
	// 定义一个浮点数比较函数
	equal := func(a, b, epsilon float64) bool {
		return math.Abs(a-b) < epsilon
	}
	
	// 使用函数比较
	fmt.Printf("0.1 + 0.2 ≈ 0.3? %t\n", equal(pinkBlack, 0.3, 1e-9))
	fmt.Printf("1.0/3 + 1.0/3 + 1.0/3 ≈ 1.0? %t\n", 
		equal(1.0/3+1.0/3+1.0/3, 1.0, 1e-9))
	
	fmt.Println()
	
	// ============================================
	// 5. epsilon 的选择
	// ============================================
	fmt.Println("=== epsilon 的选择 ===")
	
	fmt.Println("epsilon 的取值要根据场景调整：")
	fmt.Println("- 普通业务场景：1e-9 (0.000000001)")
	fmt.Println("- 高精度科学计算：1e-15 (0.000000000000001)")
	fmt.Println("- 一般计算：0.0001 也可以")
	fmt.Println()
	fmt.Println("原则：")
	fmt.Println("- 太小：可能把真正相等的数误判为不等")
	fmt.Println("- 太大：可能把不相等的数误判为相等")
	fmt.Println("- 要根据实际精度需求选择")
	
	fmt.Println()
	
	// ============================================
	// 6. 更多比较示例
	// ============================================
	fmt.Println("=== 更多比较示例 ===")
	
	// 示例1：比较两个计算结果
	result1 := 1.0 / 3
	result2 := 0.3333333333333333
	fmt.Printf("result1 ≈ result2? %t\n", 
		math.Abs(result1-result2) < 1e-9)
	
	// 示例2：比较三个 1/3 相加
	sum := 1.0/3 + 1.0/3 + 1.0/3
	fmt.Printf("sum ≈ 1.0? %t\n", math.Abs(sum-1.0) < 1e-9)
	
	// 示例3：比较不同精度的数
	num1 := 0.123456789
	num2 := 0.123456788
	fmt.Printf("num1 ≈ num2? %t\n", math.Abs(num1-num2) < 1e-9)  // false
	fmt.Printf("num1 ≈ num2 (更宽松)? %t\n", math.Abs(num1-num2) < 1e-6)  // true
}

