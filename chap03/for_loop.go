// 示例：for 循环
// 演示 Go 语言中 for 循环的各种形式

package main

import (
	"fmt"
	"time"
)

func main() {
	// ============================================
	// 1. for 循环像 while 一样（只写条件）
	// ============================================
	fmt.Println("=== for 循环（类似 while）===")
	
	// Go 没有 while 关键字，用 for + 条件 代替 while
	// 条件判断、循环体、变量修改都在相应位置
	
	count := 10
	for count > 0 {
		fmt.Print(count, " ")
		count--  // 在循环体内修改条件
		time.Sleep(100 * time.Millisecond)  // 延迟，方便观察
	}
	fmt.Println()
	fmt.Println("Liftoff!")
	
	fmt.Println()
	
	// ============================================
	// 2. 传统的 for 循环（三个部分）
	// ============================================
	fmt.Println("=== 传统 for 循环（三个部分）===")
	
	// 类似 Java 的 for，但不用小括号
	// 格式：for 初始化; 条件; 后续操作 { ... }
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// 多个变量
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}
	
	fmt.Println()
	
	// ============================================
	// 3. 无限循环
	// ============================================
	fmt.Println("=== 无限循环 ===")
	
	// 只写 for，不写条件，就是无限循环
	// 可以用 break 退出
	counter := 0
	for {
		counter++
		if counter >= 5 {
			break  // 退出循环
		}
		fmt.Print(counter, " ")
	}
	fmt.Println()
	
	fmt.Println()
	
	// ============================================
	// 4. 遍历数组/切片
	// ============================================
	fmt.Println("=== 遍历数组/切片 ===")
	
	numbers := []int{1, 2, 3, 4, 5}
	
	// 方式1：只获取值
	for _, value := range numbers {
		fmt.Print(value, " ")
	}
	fmt.Println()
	
	// 方式2：获取索引和值
	for index, value := range numbers {
		fmt.Printf("index=%d, value=%d\n", index, value)
	}
	
	fmt.Println()
	
	// ============================================
	// 5. continue 跳过本次循环
	// ============================================
	fmt.Println("=== continue 跳过本次循环 ===")
	
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue  // 跳过偶数
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	fmt.Println()
	
	// ============================================
	// 6. 实际应用：倒计时
	// ============================================
	fmt.Println("=== 实际应用：倒计时 ===")
	
	fmt.Println("倒计时开始：")
	for count := 10; count > 0; count-- {
		fmt.Printf("%d... ", count)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("\n发射！")
}


