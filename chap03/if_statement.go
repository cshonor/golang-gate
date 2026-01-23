// 示例：if 语句
// 演示 if、else if、else 的语法和逻辑运算符

package main

import (
	"fmt"
	"strings"
)

func main() {
	// ============================================
	// 1. if 语句的基本语法
	// ============================================
	fmt.Println("=== if 语句基本语法 ===")
	
	// Go 的 if 语句特点：
	// 1. 条件不需要小括号
	// 2. 左大括号必须和条件在同一行
	// 3. 右大括号必须单独占一行
	
	command := "go east"
	
	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Unknown command:", command)
	}
	
	fmt.Println()
	
	// ============================================
	// 2. 逻辑运算符：&& (与)、|| (或)、! (非)
	// ============================================
	fmt.Println("=== 逻辑运算符 ===")
	
	var isSunny = true
	var hasUmbrella = false
	var command2 = "go east"
	
	// 逻辑与 &&
	if command2 == "go east" && isSunny {
		fmt.Println("Go east and it's sunny")
	}
	
	// 逻辑或 ||
	if command2 == "go east" || command2 == "go west" {
		fmt.Println("Go east or go west")
	}
	
	// 逻辑非 !
	if !hasUmbrella {
		fmt.Println("No umbrella")
	}
	
	// 组合使用（注意优先级，不确定时用括号）
	if (command2 == "go east" || command2 == "go west") && !hasUmbrella {
		fmt.Println("Go east or west, but no umbrella")
	}
	
	fmt.Println()
	
	// ============================================
	// 3. 在 if 开头声明变量
	// ============================================
	fmt.Println("=== 在 if 开头声明变量 ===")
	
	// 在 if 开头声明变量，作用域仅限 if-else 块
	if age := 25; age >= 18 {
		fmt.Println("成年人，年龄:", age)
	} else {
		fmt.Println("未成年人，年龄:", age)
	}
	// 出了 if-else，age 就不能用了
	
	// 多个条件
	if score := 95; score >= 90 {
		fmt.Println("优秀，分数:", score)
	} else if score >= 60 {
		fmt.Println("及格，分数:", score)
	} else {
		fmt.Println("不及格，分数:", score)
	}
	
	fmt.Println()
	
	// ============================================
	// 4. 实际应用示例
	// ============================================
	fmt.Println("=== 实际应用示例 ===")
	
	var walkOutside = true
	var takeTheBluePill = false
	
	if walkOutside {
		fmt.Println("You leave the cave exit")
	}
	
	if !takeTheBluePill {
		fmt.Println("You didn't take the blue pill")
	}
	
	// 组合条件
	var command3 = "walk outside"
	if strings.Contains(command3, "outside") && walkOutside {
		fmt.Println("You leave the cave exit")
	}
}

