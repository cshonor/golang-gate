// 示例：变量作用域规则
// 演示 Go 语言中变量的作用域从大到小

package main

import (
	"fmt"
	"math/rand"
)

// ============================================
// 1. 包内全局变量（作用域：整个包）
// ============================================
// year 是包内全局变量，整个包内的所有函数都可以使用
// 注意：首字母小写，所以只能在本包内使用
var year = 2024

// 包内全局变量可以在任何函数中使用
func testGlobal() {
	fmt.Println("在 testGlobal 函数中访问 year:", year)
	year = 2025  // 可以修改全局变量
	fmt.Println("修改后的 year:", year)
}

// ============================================
// 2. 函数作用域
// ============================================
func testFunctionScope() {
	// 函数内定义的局部变量，作用域仅限这个函数
	var localVar = "我是局部变量"
	fmt.Println("函数内的局部变量:", localVar)
	
	// 函数参数也是函数作用域
	testParam(100)
}

// 函数参数的作用域仅限函数内部
func testParam(value int) {
	fmt.Println("函数参数 value:", value)
	// value 只能在这个函数内使用
}

// ============================================
// 3. 块作用域：if 语句
// ============================================
func testIfScope() {
	// 在 if 开头声明变量，作用域仅限 if-else 块
	if age := 25; age >= 18 {
		fmt.Println("成年人，年龄:", age)
		// age 在这个 if 块内可以使用
	} else {
		fmt.Println("未成年人，年龄:", age)
		// age 在 else 块内也可以使用
	}
	// 出了 if-else，age 就不能用了
	// fmt.Println(age)  // ❌ 错误：undefined: age
}

// ============================================
// 4. 块作用域：for 循环
// ============================================
func testForScope() {
	// 在 for 开头声明变量，作用域仅限整个 for 循环
	for count := 10; count > 0; count-- {
		fmt.Print(count, " ")
		// count 在循环条件、迭代表达式、循环体内都可以使用
	}
	fmt.Println()
	// 出了 for 循环，count 就不能用了
	// fmt.Println(count)  // ❌ 错误：undefined: count
}

// ============================================
// 5. 块作用域：switch 语句
// ============================================
func testSwitchScope() {
	// 在 switch 开头声明变量，作用域仅限 switch 语句
	// month := rand.Intn(12) + 1 生成 1-12 的随机数
	// 分号后面的 month 表示根据 month 的值进行匹配
	switch month := rand.Intn(12) + 1; month {
	case 2:
		fmt.Println("2月有28或29天")
	case 4, 5, 6:
		fmt.Println("4、5、6月有30天")
	case 4, 6, 9, 11:
		fmt.Println("4、6、9、11月有30天")
	default:
		fmt.Println("其他月份有31天，当前月份:", month)
	}
	// 出了 switch，month 就不能用了
	// fmt.Println(month)  // ❌ 错误：undefined: month
}

// ============================================
// 6. 作用域嵌套和变量遮蔽
// ============================================
func testShadowing() {
	// 外层作用域：函数内的变量
	year := 2023
	fmt.Println("外层 year:", year)
	
	// 内层作用域：if 块内的变量
	if true {
		year := 2024  // 内层的 year 遮蔽了外层的 year
		fmt.Println("内层 year:", year)  // 输出：2024
	}
	
	// 出了 if 块，又回到外层作用域
	fmt.Println("外层 year:", year)  // 输出：2023（外层的值没变）
}

// ============================================
// 7. 直接使用大括号创建作用域
// ============================================
func testBlockScope() {
	// 直接使用大括号创建作用域
	{
		var blockVar = "块内变量"
		fmt.Println("块内变量:", blockVar)
	}
	// 出了大括号，blockVar 就不能用了
	// fmt.Println(blockVar)  // ❌ 错误：undefined: blockVar
}

// ============================================
// 8. 实际应用示例
// ============================================
func main() {
	fmt.Println("=== 包内全局变量 ===")
	fmt.Println("全局变量 year:", year)
	testGlobal()
	fmt.Println("main 函数中的 year:", year)
	
	fmt.Println("\n=== 函数作用域 ===")
	testFunctionScope()
	
	fmt.Println("\n=== if 语句作用域 ===")
	testIfScope()
	
	fmt.Println("\n=== for 循环作用域 ===")
	testForScope()
	
	fmt.Println("\n=== switch 语句作用域 ===")
	testSwitchScope()
	
	fmt.Println("\n=== 变量遮蔽 ===")
	testShadowing()
	
	fmt.Println("\n=== 块作用域 ===")
	testBlockScope()
}


