// 示例：Go 语言带有方法的切片
// 演示如何为自定义切片类型绑定方法，以及本章小结

package main

import (
	"fmt"
	"sort"
	"strings"
)

// ============================================
// 类型定义（包级别）
// ============================================

// Planets 类型：用于 terraform 实验题
type Planets []string

// StringList 类型：字符串列表
type StringList []string

// IntSlice 类型：整数切片
type IntSlice []int

// IntList 类型：整数列表
type IntList []int

// StringSet 类型：字符串集合
type StringSet []string

// NumberList 类型：数字列表
type NumberList []int

// ============================================
// 方法实现
// ============================================

// terraform 方法：给所有行星名称加上"New"前缀，并对火星、天王星、海王星进行特殊处理
func (p Planets) terraform() {
	for i := range p {
		planet := p[i]
		// 特殊处理：火星、天王星、海王星
		if planet == "Mars" {
			p[i] = "New " + planet + " (红色星球)"
		} else if planet == "Uranus" {
			p[i] = "New " + planet + " (冰巨行星)"
		} else if planet == "Neptune" {
			p[i] = "New " + planet + " (蓝色星球)"
		} else {
			p[i] = "New " + planet
		}
	}
}

// Uppercase 方法：将所有字符串转为大写
func (sl StringList) Uppercase() {
	for i := range sl {
		sl[i] = strings.ToUpper(sl[i])
	}
}

// Trim 方法：去除所有字符串的前后空格
func (sl StringList) Trim() {
	for i := range sl {
		sl[i] = strings.TrimSpace(sl[i])
	}
}

// Sort 方法：对整数切片进行排序（简单冒泡排序）
func (is IntSlice) Sort() {
	for i := 0; i < len(is)-1; i++ {
		for j := 0; j < len(is)-1-i; j++ {
			if is[j] > is[j+1] {
				is[j], is[j+1] = is[j+1], is[j]
			}
		}
	}
}

// Reverse 方法：反转整数列表
func (il IntList) Reverse() {
	for i, j := 0, len(il)-1; i < j; i, j = i+1, j-1 {
		il[i], il[j] = il[j], il[i]
	}
}

// Unique 方法：去重（返回新切片）
func (il IntList) Unique() IntList {
	seen := make(map[int]bool)
	result := IntList{}
	for _, v := range il {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Filter 方法：过滤元素（返回新切片）
func (il IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}
	for _, v := range il {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Intersection 方法：求交集（返回新切片）
func (il IntList) Intersection(other IntList) IntList {
	seen := make(map[int]bool)
	for _, v := range il {
		seen[v] = true
	}
	result := IntList{}
	for _, v := range other {
		if seen[v] {
			result = append(result, v)
		}
	}
	return result
}

// Unique 方法：去重（返回新切片）
func (ss StringSet) Unique() StringSet {
	seen := make(map[string]bool)
	result := StringSet{}
	for _, v := range ss {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// ============================================
	// 1. 带有方法的切片
	// ============================================
	fmt.Println("=== 1. 带有方法的切片 ===")

	fmt.Println("核心原理:")
	fmt.Println("  - Go语言允许为切片类型绑定方法")
	fmt.Println("  - 为自定义的切片类型声明方法，而不是直接为原生切片类型添加方法")
	fmt.Println("  - 这让切片的功能更丰富，也更接近面向对象的风格")
	fmt.Println()

	// 示例1：标准库的 sort.StringSlice
	fmt.Println("示例1：标准库的 sort.StringSlice")
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	fmt.Printf("  原始顺序: %v\n", planets)
	sort.StringSlice(planets).Sort() // 按字母排序
	fmt.Printf("  排序后: %v\n", planets)
	fmt.Println()

	// 说明 sort.StringSlice 的本质
	fmt.Println("sort.StringSlice 的本质:")
	fmt.Println("  type StringSlice []string")
	fmt.Println("  func (p StringSlice) Sort() { /* 排序逻辑 */ }")
	fmt.Println("  使用时，只需将原生切片类型转换为该自定义类型，即可调用方法")
	fmt.Println()

	// 示例2：自定义切片类型和方法
	fmt.Println("示例2：自定义切片类型和方法")
	numbers := IntSlice{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("  原始: %v\n", numbers)
	numbers.Sort()
	fmt.Printf("  排序后: %v\n", numbers)
	fmt.Println()

	// ============================================
	// 2. 自定义切片类型和方法实现
	// ============================================
	fmt.Println("=== 2. 自定义切片类型和方法实现 ===")

	// 使用包级别定义的 StringList 类型
	var list StringList = StringList{"apple", "banana", "cherry"}
	fmt.Printf("  原始列表: %v\n", list)
	list.Uppercase()
	fmt.Printf("  转大写后: %v\n", list)
	fmt.Println()

	// ============================================
	// 3. 实验题：terraform.go 实现
	// ============================================
	fmt.Println("=== 3. 实验题：terraform.go 实现 ===")

	fmt.Println("题目要求:")
	fmt.Println("  给所有行星名称加上\"New\"前缀")
	fmt.Println("  对火星、天王星、海王星进行特殊处理")
	fmt.Println()

	// 实现思路
	fmt.Println("实现思路:")
	fmt.Println("  1. 定义一个 Planets 类型，作为 []string 的别名")
	fmt.Println("  2. 为 Planets 类型实现 terraform() 方法")
	fmt.Println("  3. 在方法中完成前缀添加和特殊处理")
	fmt.Println("  4. 在主函数中，将原生切片转换为 Planets 类型并调用方法")
	fmt.Println()

	// 实现代码
	planets2 := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	fmt.Printf("原始行星: %v\n", planets2)
	Planets(planets2).terraform()
	fmt.Printf("terraform 后: %v\n", planets2)
	fmt.Println()

	// ============================================
	// 4. 更多自定义切片方法示例
	// ============================================
	fmt.Println("=== 4. 更多自定义切片方法示例 ===")

	// 示例：反转切片
	intList := IntList{1, 2, 3, 4, 5}
	fmt.Printf("  原始: %v\n", intList)
	intList.Reverse()
	fmt.Printf("  反转后: %v\n", intList)
	fmt.Println()

	// 示例：去重
	stringSet := StringSet{"apple", "banana", "apple", "cherry", "banana"}
	fmt.Printf("  原始: %v\n", stringSet)
	unique := stringSet.Unique()
	fmt.Printf("  去重后: %v\n", unique)
	fmt.Println()

	// 示例：过滤
	numbers2 := NumberList{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("  原始: %v\n", numbers2)
	even := IntList(numbers2).Filter(func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("  偶数: %v\n", even)
	fmt.Println()

	// ============================================
	// 5. 本章核心知识点小结
	// ============================================
	fmt.Println("=== 5. 本章核心知识点小结 ===")

	fmt.Println("1. 切片本质:")
	fmt.Println("   ✅ 切片是指向数组的'视图'或'窗口'")
	fmt.Println("   ✅ 并非独立的数据结构")
	fmt.Println("   ✅ 底层指向原始数组")
	fmt.Println()

	fmt.Println("2. 共享底层数组:")
	fmt.Println("   ✅ 切片在赋值或作为函数参数传递时")
	fmt.Println("   ✅ 会与新变量共享底层数组")
	fmt.Println("   ✅ 修改一个切片会影响其他共享的切片")
	fmt.Println()

	fmt.Println("3. 迭代方式:")
	fmt.Println("   ✅ 使用 for range 可以方便地迭代切片中的元素")
	fmt.Println("   ✅ for i, v := range slice")
	fmt.Println("   ✅ for i := range slice")
	fmt.Println()

	fmt.Println("4. 字面量初始化:")
	fmt.Println("   ✅ 支持用复合字面量直接创建切片")
	fmt.Println("   ✅ []Type{value1, value2, ...}")
	fmt.Println("   ✅ 无需先定义数组")
	fmt.Println()

	fmt.Println("5. 方法绑定:")
	fmt.Println("   ✅ 可以为自定义切片类型绑定方法")
	fmt.Println("   ✅ 实现更强大的功能")
	fmt.Println("   ✅ 更接近面向对象的风格")
	fmt.Println()

	fmt.Println("6. 类型区别:")
	fmt.Println("   ✅ 数组类型包含长度: [5]string")
	fmt.Println("   ✅ 切片类型不包含长度: []string")
	fmt.Println("   ✅ 切片更适合作为函数参数")
	fmt.Println()

	// ============================================
	// 6. 实际应用场景
	// ============================================
	fmt.Println("=== 6. 实际应用场景 ===")

	// 场景1：数据处理管道
	fmt.Println("场景1：数据处理管道")
	data := []string{"  apple  ", "  banana  ", "  cherry  "}
	processed := StringList(data)
	processed.Trim()
	processed.Uppercase()
	fmt.Printf("  处理结果: %v\n", processed)
	fmt.Println()

	// 场景2：集合操作
	fmt.Println("场景2：集合操作")
	set1 := IntList{1, 2, 3, 4, 5}
	set2 := IntList{4, 5, 6, 7, 8}
	intersection := set1.Intersection(set2)
	fmt.Printf("  集合1: %v\n", set1)
	fmt.Printf("  集合2: %v\n", set2)
	fmt.Printf("  交集: %v\n", intersection)
	fmt.Println()

	// ============================================
	// 7. 最佳实践
	// ============================================
	fmt.Println("=== 7. 最佳实践 ===")

	fmt.Println("1. 自定义切片类型:")
	fmt.Println("   ✅ 为特定用途创建语义化的类型别名")
	fmt.Println("   ✅ 提高代码可读性和可维护性")
	fmt.Println()
	fmt.Println("2. 方法设计:")
	fmt.Println("   ✅ 方法应该操作接收者本身")
	fmt.Println("   ✅ 考虑是否需要修改原切片或返回新切片")
	fmt.Println()
	fmt.Println("3. 与标准库配合:")
	fmt.Println("   ✅ 参考 sort.StringSlice 的设计模式")
	fmt.Println("   ✅ 保持与标准库一致的风格")
	fmt.Println()
	fmt.Println("4. 性能考虑:")
	fmt.Println("   ✅ 注意切片共享底层数组的特性")
	fmt.Println("   ✅ 需要独立副本时使用 copy() 函数")
	fmt.Println()
}

