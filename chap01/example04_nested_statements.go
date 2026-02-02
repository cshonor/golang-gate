// 示例 4：嵌套语句的大括号规则
// 演示在函数中嵌套 if、for 等语句时，大括号的规则依然适用

package main

import "fmt"

func main() {
	// 外层：for 循环
	// 左大括号与 for 在同一行
	for i := 0; i < 3; i++ {
		// 内层：if 语句
		// if 的左大括号也要与 if 在同一行
		if i%2 == 0 {
			fmt.Printf("i=%d 是偶数\n", i)
		} else {
			fmt.Printf("i=%d 是奇数\n", i)
			// if-else 的右大括号单独占一行，与 if 对齐
		}

		// 内层：嵌套的 for 循环
		for j := 0; j < 2; j++ {
			fmt.Printf("  内层循环: i=%d, j=%d\n", i, j)
			// 内层 for 的右大括号单独占一行，与内层 for 对齐
		}
		// 外层 for 的右大括号单独占一行，与外层 for 对齐
	}
	// main 函数的右大括号单独占一行，与 func main() 对齐
}


