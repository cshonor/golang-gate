// 示例：switch 语句
// 演示 switch 的语法、case 的用法、default 的作用

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// ============================================
	// 1. switch 语句的基本语法
	// ============================================
	fmt.Println("=== switch 基本语法 ===")
	
	command := "go east"
	
	// switch 后面跟变量名，然后是大括号
	// case 后面跟值，多个值用逗号隔开
	// 不需要 break，每个 case 执行完后自动 break
	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "go west":
		fmt.Println("You head into a dense forest.")
	case "go north", "go south":
		fmt.Println("You go north or south.")
	default:
		fmt.Println("Unknown command:", command)
	}
	
	fmt.Println()
	
	// ============================================
	// 2. case 的值类型必须和 switch 变量类型一致
	// ============================================
	fmt.Println("=== case 类型必须一致 ===")
	
	// 字符串类型
	var command2 string = "go east"
	switch command2 {
	case "go east", "go west":  // 字符串用双引号
		fmt.Println("Valid direction")
	default:
		fmt.Println("Invalid direction")
	}
	
	// 整数类型
	var num int = 3
	switch num {
	case 1, 2, 3:  // 整数不用引号
		fmt.Println("Small number")
	case 4, 5, 6:
		fmt.Println("Medium number")
	default:
		fmt.Println("Large number")
	}
	
	// 布尔类型
	var isOk bool = true
	switch isOk {
	case true:  // 布尔值不用引号
		fmt.Println("It's OK")
	case false:
		fmt.Println("It's not OK")
	}
	
	fmt.Println()
	
	// ============================================
	// 3. 标准输入读取的是字符串类型
	// ============================================
	fmt.Println("=== 标准输入是字符串类型 ===")
	
	// 通过标准输入流接收到的内容，默认都是字符串类型
	// 即使输入的是 123，读进来的也是 "123" 这个字符串
	// 如果要在 switch 中使用，case 里要用字符串 "123"
	// 如果想当整数用，需要先用 strconv.Atoi 转换
	
	fmt.Println("标准输入读取的是字符串，case 里要用字符串值")
	fmt.Println("例如：用户输入 123，case 里要写 case \"123\":")
	
	fmt.Println()
	
	// ============================================
	// 4. default 的作用
	// ============================================
	fmt.Println("=== default 的作用 ===")
	
	// default 是 switch 的"兜底"选项
	// 当所有 case 都不匹配时，会执行 default
	// default 的位置不一定要放在最后，但习惯上放在最后
	
	var command3 = "go north"
	switch command3 {
	case "go east":
		fmt.Println("Go east")
	case "go west":
		fmt.Println("Go west")
	default:
		fmt.Println("Unknown command:", command3)  // 会执行这里
	}
	
	fmt.Println()
	
	// ============================================
	// 5. 在 switch 开头声明变量
	// ============================================
	fmt.Println("=== 在 switch 开头声明变量 ===")
	
	// 在 switch 开头声明变量，作用域仅限 switch 语句
	switch month := rand.Intn(12) + 1; month {
	case 1:
		fmt.Println("一月")
	case 2:
		fmt.Println("二月")
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12:
		fmt.Println("冬季")
	default:
		fmt.Println("其他月份:", month)
	}
	// 出了 switch，month 就不能用了
	
	fmt.Println()
	
	// ============================================
	// 6. 无表达式 switch
	// ============================================
	fmt.Println("=== 无表达式 switch ===")
	
	// switch 后面直接跟分号和大括号
	// case 里可以写任意条件
	switch num := rand.Intn(10); {
	case num < 3:
		fmt.Println("小数字:", num)
	case num < 7:
		fmt.Println("中等数字:", num)
	default:
		fmt.Println("大数字:", num)
	}
}


