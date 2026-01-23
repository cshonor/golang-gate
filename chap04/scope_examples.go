// 示例：各种作用域的具体示例
// 详细演示 for、if、switch 中在开头声明变量的写法

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// ============================================
	// 1. for 循环中在开头声明变量
	// ============================================
	fmt.Println("=== for 循环作用域 ===")
	
	// 在 for 开头声明 count，作用域仅限整个 for 循环
	// count := 10 是初始化
	// count > 0 是条件判断（可以使用 count）
	// count-- 是迭代表达式（可以使用 count）
	// 循环体内也可以使用 count
	for count := 10; count > 0; count-- {
		fmt.Print(count, " ")
	}
	fmt.Println()
	// 出了 for 循环，count 就不能用了
	
	// 另一个例子：多个变量
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}
	
	fmt.Println()
	
	// ============================================
	// 2. if 语句中在开头声明变量
	// ============================================
	fmt.Println("=== if 语句作用域 ===")
	
	// 在 if 开头声明 age，作用域仅限 if-else 块
	if age := 25; age >= 18 {
		fmt.Println("成年人，年龄:", age)
		// age 在 if 块内可以使用
	} else {
		fmt.Println("未成年人，年龄:", age)
		// age 在 else 块内也可以使用
	}
	// 出了 if-else，age 就不能用了
	
	// 另一个例子：多个条件
	if score := 95; score >= 90 {
		fmt.Println("优秀，分数:", score)
	} else if score >= 60 {
		fmt.Println("及格，分数:", score)
	} else {
		fmt.Println("不及格，分数:", score)
	}
	
	fmt.Println()
	
	// ============================================
	// 3. switch 语句中在开头声明变量
	// ============================================
	fmt.Println("=== switch 语句作用域 ===")
	
	// 在 switch 开头声明 month
	// month := rand.Intn(12) + 1 生成 1-12 的随机数
	// 分号后面的 month 表示根据 month 的值进行匹配
	switch month := rand.Intn(12) + 1; month {
	case 1:
		fmt.Println("一月，当前月份:", month)
	case 2:
		fmt.Println("二月，当前月份:", month)
	case 3, 4, 5:
		fmt.Println("春季，当前月份:", month)
	case 6, 7, 8:
		fmt.Println("夏季，当前月份:", month)
	case 9, 10, 11:
		fmt.Println("秋季，当前月份:", month)
	case 12:
		fmt.Println("冬季，当前月份:", month)
	default:
		fmt.Println("其他月份:", month)
	}
	// 出了 switch，month 就不能用了
	
	fmt.Println()
	
	// ============================================
	// 4. 无表达式 switch（在开头声明变量）
	// ============================================
	fmt.Println("=== 无表达式 switch ===")
	
	// 无表达式 switch：switch 后面直接跟分号和大括号
	// 这种写法中，case 里可以写任意条件
	switch num := rand.Intn(10); {
	case num < 3:
		fmt.Println("小数字:", num)
	case num < 7:
		fmt.Println("中等数字:", num)
	default:
		fmt.Println("大数字:", num)
	}
	
	fmt.Println()
	
	// ============================================
	// 5. 作用域的好处：避免变量污染
	// ============================================
	fmt.Println("=== 作用域的好处 ===")
	
	// 在不同的控制语句中使用同名变量，不会冲突
	// 因为每个变量的作用域都限制在各自的块内
	
	// 第一个 i：作用域在第一个 for 循环
	for i := 0; i < 3; i++ {
		fmt.Print("第一个循环 i=", i, " ")
	}
	fmt.Println()
	
	// 第二个 i：作用域在第二个 for 循环（和第一个 i 不冲突）
	for i := 0; i < 3; i++ {
		fmt.Print("第二个循环 i=", i, " ")
	}
	fmt.Println()
	
	// 第三个 i：作用域在 if 块内（和前两个 i 不冲突）
	if i := 5; i > 0 {
		fmt.Println("if 块内的 i=", i)
	}
	
	// 所有这些 i 都不会互相影响，因为它们的作用域是分开的
}

