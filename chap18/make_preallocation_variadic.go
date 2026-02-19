我给你一句最实用、最落地、你以后写代码直接照抄判断的规则：
 
你只记这 3 条，永远不会用错
 
1. 看到 slice / map / channel → 只用 make，绝对不用 new
2. 看到 int / float / bool / 结构体 → 只用 new 或直接赋值，不用 make
3. 平时写业务，new 很少用，大部分用 &结构体{} 就够了
 
 
 
1. 什么时候用 make？（必记）
 
只要你用到这三个东西：
 
- 切片 []T
- map
- channel
 
必须用 make，因为它们需要初始化内部结构，new 做不到。
 
go  
s := make([]int, 3)
m := make(map[string]int)
ch := make(chan int)
 
 
👉 不用 make，直接用会 panic / 无法赋值。
 
 
 
2. 什么时候用 new？
 
你想要：
 
- 一个指针
- 类型是：int、float、bool、结构体 这种值类型
 
才用 new。
 
go  
p := new(int)     // *int
c := new(Coordinate) // *Coordinate
 
 
但实际开发中，大家更爱写：
 
go  
c := &Coordinate{} // 效果和 new(Coordinate) 一样，更常用
 
 
👉 new 很少用，能用 & 就不用 new。
 
 
 
3. 超级简单口诀（背这个）
 
- slice、map、chan → 只能 make
- int、float、结构体 → 用不到 make，想用指针就 & 或 new
- new = 拿指针
- make = 初始化引用类型
 
 
 
你现在这句理解已经完全正确了：
 
new 用于基本/值类型，make 用于引用类型。
 
我再问你一句，你马上就能彻底稳了：
 new([]int)  是对还是错？
// 示例：Go 语言 make 函数预分配切片和可变参数函数
// 演示 make 预分配、可变参数函数的使用和性能优化

package main

import "fmt"

func main() {
	// ============================================
	// 1. 使用 make 函数对切片预分配
	// ============================================
	fmt.Println("=== 1. 使用 make 函数对切片预分配 ===")

	fmt.Println("作用：提前为切片分配足够的底层数组容量")
	fmt.Println("     避免后续 append 操作时频繁扩容，从而提升性能")
	fmt.Println()

	fmt.Println("语法: make([]T, length, capacity)")
	fmt.Println("  - length: 切片初始包含的元素数量")
	fmt.Println("  - capacity: 底层数组的容量，可省略，省略时与 length 相等")
	fmt.Println()

	// 示例1：预分配容量
	fmt.Println("示例1：预分配容量")
	dwarfs := make([]string, 0, 10) // 长度为0，容量为10
	fmt.Printf("  初始: len=%d, cap=%d\n", len(dwarfs), cap(dwarfs))
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Printf("  添加5个元素后: %v\n", dwarfs)
	fmt.Printf("  长度: %d, 容量: %d (未扩容)\n", len(dwarfs), cap(dwarfs))
	fmt.Println("  说明：容量足够，不会触发扩容")
	fmt.Println()

	// 示例2：对比预分配 vs 不预分配
	fmt.Println("示例2：对比预分配 vs 不预分配")
	// 不预分配
	noPrealloc := []string{}
	fmt.Printf("  不预分配初始: len=%d, cap=%d\n", len(noPrealloc), cap(noPrealloc))
	for i := 0; i < 5; i++ {
		oldCap := cap(noPrealloc)
		noPrealloc = append(noPrealloc, fmt.Sprintf("Planet%d", i))
		newCap := cap(noPrealloc)
		if newCap != oldCap {
			fmt.Printf("    添加元素后扩容: %d -> %d\n", oldCap, newCap)
		}
	}
	fmt.Printf("  最终: len=%d, cap=%d (可能多次扩容)\n", len(noPrealloc), cap(noPrealloc))
	fmt.Println()

	// 预分配
	prealloc := make([]string, 0, 5)
	fmt.Printf("  预分配初始: len=%d, cap=%d\n", len(prealloc), cap(prealloc))
	for i := 0; i < 5; i++ {
		oldCap := cap(prealloc)
		prealloc = append(prealloc, fmt.Sprintf("Planet%d", i))
		newCap := cap(prealloc)
		if newCap != oldCap {
			fmt.Printf("    添加元素后扩容: %d -> %d\n", oldCap, newCap)
		}
	}
	fmt.Printf("  最终: len=%d, cap=%d (未扩容)\n", len(prealloc), cap(prealloc))
	fmt.Println()

	// 示例3：不同 length 和 capacity 的组合
	fmt.Println("示例3：不同 length 和 capacity 的组合")
	// length = 0, capacity = 5
	slice1 := make([]int, 0, 5)
	fmt.Printf("  make([]int, 0, 5): len=%d, cap=%d, %v\n", len(slice1), cap(slice1), slice1)

	// length = 3, capacity = 5
	slice2 := make([]int, 3, 5)
	fmt.Printf("  make([]int, 3, 5): len=%d, cap=%d, %v\n", len(slice2), cap(slice2), slice2)

	// length = 5, capacity = 5 (省略 capacity)
	slice3 := make([]int, 5)
	fmt.Printf("  make([]int, 5): len=%d, cap=%d, %v\n", len(slice3), cap(slice3), slice3)
	fmt.Println()

	// ============================================
	// 2. 可变参数函数（Variadic Functions）
	// ============================================
	fmt.Println("=== 2. 可变参数函数（Variadic Functions）===")

	fmt.Println("定义：函数的最后一个参数类型前加上 ...")
	fmt.Println("     表示可以接收0个或多个该类型的参数")
	fmt.Println()

	// 示例1：基础可变参数函数
	fmt.Println("示例1：基础可变参数函数")
	fmt.Println("  函数定义: func terraform(prefix string, worlds ...string) []string")
	fmt.Println()

	// 调用方式1：直接传递多个参数
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Printf("  直接传递多个参数: terraform(\"New\", \"Venus\", \"Mars\")\n")
	fmt.Printf("  结果: %v\n", twoWorlds)
	fmt.Println()

	// 调用方式2：展开切片作为参数
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Printf("  展开切片作为参数: terraform(\"New\", planets...)\n")
	fmt.Printf("  原切片: %v\n", planets)
	fmt.Printf("  结果: %v\n", newPlanets)
	fmt.Println()

	// 示例2：可变参数函数内部处理
	fmt.Println("示例2：可变参数函数内部处理")
	fmt.Println("  函数内部的可变参数会被当作切片处理")
	fmt.Println("  但直接修改这个切片不会影响外部传入的原切片")
	fmt.Println()

	// 演示可变参数函数的行为
	original := []string{"Earth", "Mars"}
	result := terraform("New", original...)
	fmt.Printf("  原切片: %v\n", original)
	fmt.Printf("  函数返回: %v (新切片，不影响原切片)\n", result)
	fmt.Println()

	// 示例3：可变参数函数可以接收0个参数
	fmt.Println("示例3：可变参数函数可以接收0个参数")
	emptyResult := terraform("New")
	fmt.Printf("  terraform(\"New\"): %v (空切片)\n", emptyResult)
	fmt.Println()

	// 示例4：多个可变参数函数示例
	fmt.Println("示例4：多个可变参数函数示例")
	sum1 := sum(1, 2, 3, 4, 5)
	fmt.Printf("  sum(1, 2, 3, 4, 5) = %d\n", sum1)

	numbers := []int{10, 20, 30}
	sum2 := sum(numbers...)
	fmt.Printf("  sum([]int{10, 20, 30}...) = %d\n", sum2)
	fmt.Println()

	// ============================================
	// 3. 核心优化思路
	// ============================================
	fmt.Println("=== 3. 核心优化思路 ===")

	fmt.Println("优化1：预分配")
	fmt.Println("  当你知道切片的大致元素数量时")
	fmt.Println("  用 make 预分配容量可以减少内存分配和数据复制，提升性能")
	fmt.Println()

	// 性能对比示例
	fmt.Println("性能对比示例:")
	// 方式1：不预分配（可能多次扩容）
	fmt.Println("  方式1：不预分配（可能多次扩容）")
	slowSlice := []int{}
	for i := 0; i < 1000; i++ {
		slowSlice = append(slowSlice, i)
	}
	fmt.Printf("    最终: len=%d, cap=%d (可能多次扩容)\n", len(slowSlice), cap(slowSlice))
	fmt.Println()

	// 方式2：预分配（无需扩容）
	fmt.Println("  方式2：预分配（无需扩容）")
	fastSlice := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		fastSlice = append(fastSlice, i)
	}
	fmt.Printf("    最终: len=%d, cap=%d (无需扩容)\n", len(fastSlice), cap(fastSlice))
	fmt.Println()

	fmt.Println("优化2：不可变设计")
	fmt.Println("  可变参数函数中，最好创建新切片返回结果")
	fmt.Println("  而不是直接修改传入的切片")
	fmt.Println("  这样可以避免意外的副作用，让代码更安全")
	fmt.Println()

	// 对比：可变设计 vs 不可变设计
	fmt.Println("对比：可变设计 vs 不可变设计")
	data := []string{"a", "b", "c"}
	fmt.Printf("  原始数据: %v\n", data)

	// 不可变设计（推荐）
	result1 := terraform("New", data...)
	fmt.Printf("  不可变设计: 原数据=%v, 结果=%v (原数据未改变)\n", data, result1)

	// 可变设计（不推荐，已注释）
	fmt.Println("  可变设计（不推荐）:")
	fmt.Println("    直接修改传入的切片，可能产生意外的副作用")
	fmt.Println()

	// ============================================
	// 4. 切片性能优化清单
	// ============================================
	fmt.Println("=== 4. 切片性能优化清单 ===")

	fmt.Println("1. make 预分配:")
	fmt.Println("   ✅ 知道大致元素数量时，使用 make([]T, 0, capacity) 预分配")
	fmt.Println("   ✅ 避免频繁扩容带来的性能开销")
	fmt.Println("   ✅ 减少内存分配和数据复制")
	fmt.Println()

	fmt.Println("2. 三索引切片:")
	fmt.Println("   ✅ 使用 a[low:high:max] 限制容量")
	fmt.Println("   ✅ 避免意外占用过多内存")
	fmt.Println("   ✅ 避免内存泄漏")
	fmt.Println()

	fmt.Println("3. 避免不必要的扩容:")
	fmt.Println("   ✅ 批量添加时，一次性添加多个元素")
	fmt.Println("   ✅ 使用 append(slice, elem1, elem2, ...) 而不是多次 append")
	fmt.Println("   ✅ 使用 append(slice, anotherSlice...) 合并切片")
	fmt.Println()

	fmt.Println("4. 可变参数函数:")
	fmt.Println("   ✅ 使用可变参数函数提高灵活性")
	fmt.Println("   ✅ 函数内部创建新切片返回，避免副作用")
	fmt.Println("   ✅ 使用 ... 展开切片作为参数")
	fmt.Println()

	fmt.Println("5. 容量规划:")
	fmt.Println("   ✅ 根据实际需求合理设置容量")
	fmt.Println("   ✅ 避免过度预分配（浪费内存）")
	fmt.Println("   ✅ 避免容量不足（频繁扩容）")
	fmt.Println()

	// ============================================
	// 5. 实际应用示例
	// ============================================
	fmt.Println("=== 5. 实际应用示例 ===")

	// 示例1：批量处理数据
	fmt.Println("示例1：批量处理数据")
	processData("apple", "banana", "cherry")
	fmt.Println()

	// 示例2：合并多个切片
	fmt.Println("示例2：合并多个切片")
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}
	sliceC := []int{7, 8, 9}
	merged := mergeSlices(sliceA, sliceB, sliceC)
	fmt.Printf("  合并: %v + %v + %v = %v\n", sliceA, sliceB, sliceC, merged)
	fmt.Println()

	// 示例3：高效构建字符串切片
	fmt.Println("示例3：高效构建字符串切片")
	names := buildNames("User", 100)
	fmt.Printf("  构建100个名称: len=%d, cap=%d\n", len(names), cap(names))
	fmt.Printf("  前5个: %v\n", names[:5])
	fmt.Println()

	// 示例4：可变参数函数的高级用法
	fmt.Println("示例4：可变参数函数的高级用法")
	filtered := filter(func(x int) bool { return x%2 == 0 }, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("  过滤偶数: %v\n", filtered)
	fmt.Println()

	// ============================================
	// 6. 常见陷阱和注意事项
	// ============================================
	fmt.Println("=== 6. 常见陷阱和注意事项 ===")

	fmt.Println("陷阱1：忘记接收 append 的返回值")
	fmt.Println("  slice := []int{1, 2, 3}")
	fmt.Println("  append(slice, 4)  // ❌ 错误：没有接收返回值")
	fmt.Println("  slice = append(slice, 4)  // ✅ 正确")
	fmt.Println()

	fmt.Println("陷阱2：可变参数函数中直接修改切片")
	fmt.Println("  func badFunc(items ...string) {")
	fmt.Println("    items[0] = \"changed\"  // ❌ 可能影响外部切片")
	fmt.Println("  }")
	fmt.Println("  应该创建新切片返回")
	fmt.Println()

	fmt.Println("陷阱3：过度预分配")
	fmt.Println("  make([]int, 0, 1000000)  // ❌ 如果只用100个元素，浪费内存")
	fmt.Println("  应该根据实际需求合理设置容量")
	fmt.Println()

	fmt.Println("陷阱4：容量不足导致频繁扩容")
	fmt.Println("  slice := []int{}")
	fmt.Println("  for i := 0; i < 1000; i++ {")
	fmt.Println("    slice = append(slice, i)  // ❌ 可能多次扩容")
	fmt.Println("  }")
	fmt.Println("  应该预分配: make([]int, 0, 1000)")
	fmt.Println()

	// ============================================
	// 7. 总结
	// ============================================
	fmt.Println("=== 7. 总结 ===")
	fmt.Println()
	fmt.Println("1. make 预分配:")
	fmt.Println("   ✅ 提前分配容量，避免频繁扩容")
	fmt.Println("   ✅ 提升性能，减少内存分配")
	fmt.Println("   ✅ 语法: make([]T, length, capacity)")
	fmt.Println()
	fmt.Println("2. 可变参数函数:")
	fmt.Println("   ✅ 最后一个参数类型前加 ...")
	fmt.Println("   ✅ 可以接收0个或多个参数")
	fmt.Println("   ✅ 使用 ... 展开切片作为参数")
	fmt.Println("   ✅ 函数内部当作切片处理")
	fmt.Println()
	fmt.Println("3. 性能优化:")
	fmt.Println("   ✅ 预分配容量")
	fmt.Println("   ✅ 使用三索引切片限制容量")
	fmt.Println("   ✅ 批量添加元素")
	fmt.Println("   ✅ 不可变设计，避免副作用")
	fmt.Println()
	fmt.Println("4. 最佳实践:")
	fmt.Println("   ✅ 根据实际需求合理设置容量")
	fmt.Println("   ✅ 可变参数函数创建新切片返回")
	fmt.Println("   ✅ 避免过度预分配和容量不足")
	fmt.Println("   ✅ 注意 append 的返回值")
	fmt.Println()
}

// ============================================
// 辅助函数
// ============================================

// terraform 可变参数函数示例
// 给所有行星名称加上前缀，返回新切片
func terraform(prefix string, worlds ...string) []string {
	// 预分配容量
	newWorlds := make([]string, 0, len(worlds))
	for i := range worlds {
		newWorlds = append(newWorlds, prefix+" "+worlds[i])
	}
	return newWorlds
}

// sum 可变参数函数：计算整数和
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// processData 处理数据（可变参数）
func processData(items ...string) {
	fmt.Printf("  处理 %d 个元素: %v\n", len(items), items)
	// 处理逻辑...
}

// mergeSlices 合并多个切片
func mergeSlices(slices ...[]int) []int {
	// 计算总长度
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}
	// 预分配容量
	result := make([]int, 0, totalLen)
	// 合并所有切片
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// buildNames 高效构建名称切片
func buildNames(prefix string, count int) []string {
	// 预分配容量
	names := make([]string, 0, count)
	for i := 1; i <= count; i++ {
		names = append(names, fmt.Sprintf("%s%d", prefix, i))
	}
	return names
}

// filter 过滤函数（可变参数）
func filter(fn func(int) bool, numbers ...int) []int {
	result := make([]int, 0, len(numbers))
	for _, num := range numbers {
		if fn(num) {
			result = append(result, num)
		}
	}
	return result
}

