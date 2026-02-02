// 示例：布尔类型和比较运算符
// 演示布尔变量、比较运算、字符串比较等

package main

import (
	"fmt"
	"strings"
)

func main() {
	// ============================================
	// 1. 布尔类型变量
	// ============================================
	fmt.Println("=== 布尔类型变量 ===")
	
	// 布尔变量可以直接用 true 或 false 赋值
	var walkOutside = true
	var takeTheBluePill = false
	
	fmt.Println("walkOutside:", walkOutside)
	fmt.Println("takeTheBluePill:", takeTheBluePill)
	
	// 也可以显式声明类型
	var isSunny bool = true
	var hasUmbrella bool = false
	
	fmt.Println("isSunny:", isSunny)
	fmt.Println("hasUmbrella:", hasUmbrella)
	
	fmt.Println()
	
	// ============================================
	// 2. strings.Contains 函数
	// ============================================
	fmt.Println("=== strings.Contains 函数 ===")
	
	// strings.Contains 检查字符串是否包含子字符串
	// 返回值是布尔类型（true 或 false）
	command := "walk outside"
	containsOutside := strings.Contains(command, "outside")
	
	fmt.Printf("command: %s\n", command)
	fmt.Printf("contains 'outside': %t\n", containsOutside)
	
	if containsOutside {
		fmt.Println("You leave the cave exit")
	}
	
	fmt.Println()
	
	// ============================================
	// 3. 比较运算符
	// ============================================
	fmt.Println("=== 比较运算符 ===")
	
	// 比较运算符的结果是布尔类型，可以用变量接收
	// == 相等
	// != 不等
	// <  小于
	// >  大于
	// <= 小于等于
	// >= 大于等于
	
	var command2 = "go east"
	var isEqual = (command2 == "go east")
	var isNotEqual = (command2 != "go west")
	
	fmt.Printf("command2 == 'go east': %t\n", isEqual)
	fmt.Printf("command2 != 'go west': %t\n", isNotEqual)
	
	// 字符串长度比较
	var isLonger = (len(command2) > 5)
	fmt.Printf("len(command2) > 5: %t\n", isLonger)
	
	fmt.Println()
	
	// ============================================
	// 4. 字符串比较（按 Unicode 编码）
	// ============================================
	fmt.Println("=== 字符串比较 ===")
	
	// 字符串比较是按字符的 Unicode 编码值逐个对比
	// "apple" 和 "banana" 比较
	// 'a' (97) < 'b' (98)，所以 "apple" < "banana"
	
	var result1 = ("apple" > "banana")
	var result2 = ("banana" > "apple")
	
	fmt.Printf("'apple' > 'banana': %t\n", result1)  // false
	fmt.Printf("'banana' > 'apple': %t\n", result2)   // true
	
	// 更多例子
	fmt.Printf("'abc' < 'def': %t\n", "abc" < "def")  // true
	fmt.Printf("'xyz' > 'abc': %t\n", "xyz" > "abc") // true
	
	fmt.Println()
	
	// ============================================
	// 5. 类型严格性：字符串和数值不能直接比较
	// ============================================
	fmt.Println("=== 类型严格性 ===")
	
	// Go 不允许直接将字符串和数值进行比较
	// 下面的代码会编译错误：
	// var result = ("123" == 123)  // ❌ 错误：不能比较 string 和 int
	
	// 必须先把类型统一
	// 字符串转整数：strconv.Atoi
	// 整数转字符串：strconv.Itoa
	
	fmt.Println("Go 要求 == 和 != 运算符两边的操作数必须是相同类型")
	fmt.Println("字符串和数值不能直接比较，必须先转换类型")
	
	fmt.Println()
	
	// ============================================
	// 6. 布尔变量的运算
	// ============================================
	fmt.Println("=== 布尔变量的运算 ===")
	
	var isSunny2 = true
	var hasUmbrella2 = false
	
	// 取反
	fmt.Printf("!isSunny2: %t\n", !isSunny2)
	
	// 与运算
	fmt.Printf("isSunny2 && hasUmbrella2: %t\n", isSunny2 && hasUmbrella2)
	
	// 或运算
	fmt.Printf("isSunny2 || hasUmbrella2: %t\n", isSunny2 || hasUmbrella2)
	
	// 组合运算
	var canGoOut = isSunny2 && !hasUmbrella2
	fmt.Printf("canGoOut (isSunny && !hasUmbrella): %t\n", canGoOut)
	
	// 布尔异或：用 != 实现
	var a = true
	var b = false
	fmt.Printf("a != b (布尔异或): %t\n", a != b)  // true
	fmt.Printf("a == b: %t\n", a == b)            // false
}


