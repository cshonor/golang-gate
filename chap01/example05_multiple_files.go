// 示例 5：同一个 main 包下的多个文件
// 这个文件包含 main 函数

package main

import "fmt"

// main 函数是程序的入口
// 注意：整个 main 包只能有一个 main 函数
func main() {
	fmt.Println("这是主函数")
	
	// 可以直接调用同包下其他文件的函数，不需要导入
	sayHello()
	calculateSum(10, 20)
}


