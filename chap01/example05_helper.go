// 示例 5 的辅助文件
// 这个文件也在 main 包下，但没有 main 函数
// 可以定义其他辅助函数供 main 函数调用

package main

import "fmt"

// sayHello 是一个辅助函数
// 注意：左大括号与 func 在同一行
func sayHello() {
	fmt.Println("Hello from helper function!")
	// 右大括号单独占一行
}

// calculateSum 计算两个数的和
func calculateSum(a int, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}

// 注意：这个文件不能有 main 函数
// 因为同一个 main 包下只能有一个 main 函数
// 如果这里也写 func main()，编译器会报错："main 函数重复定义"


