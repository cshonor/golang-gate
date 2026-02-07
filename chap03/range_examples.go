// 示例：range 循环的完整用法
// 演示 Go 语言中 range 遍历数组、切片、map 和字符串的各种方式

package main

import (
	"fmt"
	"sort"
)

func main() {
	// ============================================
	// 1. 遍历数组和切片
	// ============================================
	fmt.Println("=== 1. 遍历数组和切片 ===")
	
	numbers := []int{10, 20, 30, 40, 50}
	
	// 1.1 只获取值（忽略索引）
	fmt.Println("\n1.1 只获取值（忽略索引）:")
	for _, value := range numbers {
		fmt.Print(value, " ")
	}
	fmt.Println()
	
	// 1.2 获取索引和值
	fmt.Println("\n1.2 获取索引和值:")
	for index, value := range numbers {
		fmt.Printf("索引=%d, 值=%d\n", index, value)
	}
	
	// 1.3 只获取索引（忽略值）
	fmt.Println("\n1.3 只获取索引（忽略值）:")
	for index := range numbers {
		fmt.Print(index, " ")
	}
	fmt.Println()
	
	// 1.4 数组和切片遍历方式相同
	fmt.Println("\n1.4 数组遍历:")
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}
	
	// ============================================
	// 2. 遍历 Map（映射）
	// ============================================
	fmt.Println("\n=== 2. 遍历 Map（映射） ===")
	
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
		"David":   88,
	}
	
	// 2.1 获取键和值
	fmt.Println("\n2.1 获取键和值（顺序随机）:")
	for name, score := range scores {
		fmt.Printf("%s 的分数是 %d\n", name, score)
	}
	
	// 2.2 只获取键（忽略值）
	fmt.Println("\n2.2 只获取键（忽略值）:")
	for name := range scores {
		fmt.Println("姓名:", name)
	}
	
	// 2.3 只获取值（忽略键）
	fmt.Println("\n2.3 只获取值（忽略键）:")
	for _, score := range scores {
		fmt.Println("分数:", score)
	}
	
	// 2.4 按排序后的键遍历（需要先收集键并排序）
	fmt.Println("\n2.4 按排序后的键遍历:")
	var names []string
	for name := range scores {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, scores[name])
	}
	
	// ============================================
	// 3. 遍历字符串
	// ============================================
	fmt.Println("\n=== 3. 遍历字符串 ===")
	
	// 3.1 获取索引和字符（Unicode 字符）
	fmt.Println("\n3.1 获取索引和字符（英文）:")
	text := "Hello"
	for index, char := range text {
		fmt.Printf("索引=%d, 字符=%c, Unicode=%d\n", index, char, char)
	}
	
	// 3.2 只获取字符（忽略索引）
	fmt.Println("\n3.2 只获取字符（忽略索引）:")
	for _, char := range text {
		fmt.Printf("%c ", char)
	}
	fmt.Println()
	
	// 3.3 中文字符串遍历
	fmt.Println("\n3.3 中文字符串遍历:")
	chinese := "Go语言"
	for i, char := range chinese {
		fmt.Printf("字节索引=%d, 字符=%c\n", i, char)
	}
	
	// ============================================
	// 4. 实际应用示例
	// ============================================
	fmt.Println("\n=== 4. 实际应用示例 ===")
	
	// 4.1 统计字符出现次数
	fmt.Println("\n4.1 统计字符出现次数:")
	text2 := "hello"
	charCount := make(map[rune]int)
	for _, char := range text2 {
		charCount[char]++
	}
	fmt.Println("字符统计:", charCount)
	
	// 4.2 查找最大值
	fmt.Println("\n4.2 查找最大值:")
	numbers2 := []int{3, 7, 2, 9, 1, 5}
	max := numbers2[0]
	for _, value := range numbers2[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Printf("数组: %v\n", numbers2)
	fmt.Printf("最大值: %d\n", max)
	
	// 4.3 过滤元素（筛选偶数）
	fmt.Println("\n4.3 过滤元素（筛选偶数）:")
	numbers3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evens []int
	for _, num := range numbers3 {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Printf("原数组: %v\n", numbers3)
	fmt.Printf("偶数: %v\n", evens)
	
	// 4.4 修改切片元素
	fmt.Println("\n4.4 修改切片元素（每个元素乘以2）:")
	numbers4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前: %v\n", numbers4)
	for i := range numbers4 {
		numbers4[i] *= 2
	}
	fmt.Printf("修改后: %v\n", numbers4)
	
	// ============================================
	// 5. 常见错误示例（注释掉，避免编译错误）
	// ============================================
	fmt.Println("\n=== 5. 常见错误示例 ===")
	fmt.Println("以下代码会导致编译错误，已注释：")
	
	// ❌ 错误1：不能只写一个变量接收 range 的返回值
	// numbers := []int{1, 2, 3}
	// for value := range numbers {  // 编译错误！
	//     fmt.Println(value)
	// }
	fmt.Println("❌ for value := range numbers { ... }")
	fmt.Println("   错误：range 必须返回两个值，不能只接收一个")
	fmt.Println("✅ 正确：for _, value := range numbers { ... }")
	
	// ❌ 错误2：遍历 map 时不能修改 map
	// scores := map[string]int{"Alice": 95}
	// for k := range scores {
	//     delete(scores, k)  // 运行时错误！
	// }
	fmt.Println("\n❌ 在遍历 map 时删除元素会导致运行时错误")
	fmt.Println("✅ 正确：先收集要删除的键，遍历后再删除")
	
	// ============================================
	// 6. 性能提示
	// ============================================
	fmt.Println("\n=== 6. 性能提示 ===")
	fmt.Println("• range 遍历是值拷贝，大结构体可能影响性能")
	fmt.Println("• 如果只需要索引，用 for i := range arr 更高效")
	fmt.Println("• 对于大切片，考虑使用传统 for 循环通过索引访问")
}

