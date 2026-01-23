// 示例 6：错误的大括号格式（这些写法会导致编译错误）
// ⚠️ 注意：这个文件中的代码都是错误的示例，不要直接运行

package main

import "fmt"

// ❌ 错误示例 1：左大括号换行了
// func main()
// {
//     fmt.Println("错误：左大括号不能换行")
// }

// ❌ 错误示例 2：右大括号和其他代码在同一行
// func main() {
//     fmt.Println("错误：右大括号不能和其他代码在同一行") }

// ❌ 错误示例 3：if 语句的左大括号换行了
// func main() {
//     if true
//     {
//         fmt.Println("错误：if 的左大括号不能换行")
//     }
// }

// ❌ 错误示例 4：for 循环的右大括号和其他代码在同一行
// func main() {
//     for i := 0; i < 3; i++ {
//         fmt.Println(i) }
// }

// ✅ 正确的写法（作为对比）
func main() {
	// 正确的格式
	if true {
		fmt.Println("正确：左大括号与 if 在同一行")
	}

	for i := 0; i < 3; i++ {
		fmt.Println("正确：右大括号单独占一行")
	}
}

