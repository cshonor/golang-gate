// 示例 3：for 循环的大括号规则
// 演示循环语句中大括号的正确用法

package main

import "fmt"

func main() {
	// 方式 1：传统的 for 循环（类似 C 语言）
	// 左大括号必须与 for 在同一行
	for i := 0; i < 5; i++ {
		fmt.Println("循环次数:", i)
		// 右大括号必须单独占一行，与 for 对齐
	}

	// 方式 2：类似 while 的循环（只有条件）
	j := 0
	for j < 3 {
		fmt.Println("j 的值:", j)
		j++
	}

	// 方式 3：无限循环
	k := 0
	for {
		if k >= 3 {
			break // 使用 break 跳出循环
		}
		fmt.Println("k 的值:", k)
		k++
	}

	// 方式 4：遍历数组/切片
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("索引 %d: 值 %d\n", index, value)
	}
}


