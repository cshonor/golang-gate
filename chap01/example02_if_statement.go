// 示例 2：if 语句的大括号规则
// 演示条件判断语句中大括号的正确用法

package main

import "fmt"

func main() {
	// 定义两个变量
	a := 10
	b := 5

	// if 语句：左大括号必须与 if 在同一行
	// 右大括号必须单独占一行，与 if 对齐
	if a > b {
		fmt.Println("a 大于 b")
	}

	// if-else 语句：else 的左大括号也要与 else 在同一行
	if a < b {
		fmt.Println("a 小于 b")
	} else {
		fmt.Println("a 不小于 b")
	}

	// if-else if-else 语句
	if a > b {
		fmt.Println("a 大于 b")
	} else if a == b {
		fmt.Println("a 等于 b")
	} else {
		fmt.Println("a 小于 b")
	}
}


