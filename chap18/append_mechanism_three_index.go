// 示例：Go 语言 append 函数的底层机制和三索引切片操作
// 演示 append 的扩容机制、三索引切片操作和实用场景

package main

import "fmt"

func main() {
	// ============================================
	// 1. append 函数的底层扩容机制
	// ============================================
	fmt.Println("=== 1. append 函数的底层扩容机制 ===")

	fmt.Println("触发条件：当新元素数量超出切片当前容量时，append 会自动扩容")
	fmt.Println()

	// 扩容策略演示
	fmt.Println("扩容策略:")
	fmt.Println("  - 对于小切片（容量 < 1024），通常会将容量翻倍")
	fmt.Println("  - 对于大切片（容量 >= 1024），容量会增加约25%，以避免过度分配内存")
	fmt.Println()

	// 小切片扩容演示
	fmt.Println("小切片扩容演示（容量 < 1024）:")
	smallSlice := []int{1}
	fmt.Printf("  初始: len=%d, cap=%d\n", len(smallSlice), cap(smallSlice))
	for i := 0; i < 5; i++ {
		oldCap := cap(smallSlice)
		smallSlice = append(smallSlice, i+2)
		newCap := cap(smallSlice)
		if newCap != oldCap {
			ratio := float64(newCap) / float64(oldCap)
			fmt.Printf("  扩容: %d -> %d (倍数: %.2f)\n", oldCap, newCap, ratio)
		}
	}
	fmt.Println()

	// 大切片扩容演示（模拟）
	fmt.Println("大切片扩容演示（容量 >= 1024）:")
	fmt.Println("  当容量 >= 1024 时，扩容策略会改变")
	fmt.Println("  例如：1024 -> 1280 (增加约25%)")
	fmt.Println("  例如：1280 -> 1600 (增加约25%)")
	fmt.Println()

	// 底层变化演示
	fmt.Println("底层变化演示:")
	fmt.Println("  扩容时，Go会创建一个新的、容量更大的底层数组")
	fmt.Println("  它会把原切片的元素复制到新数组，再添加新元素")
	fmt.Println("  新切片会指向这个新数组，原数组则会被垃圾回收")
	fmt.Println()

	// 演示底层数组的变化
	fmt.Println("演示底层数组的变化:")
	original := []int{1, 2, 3}
	reference := original // 引用原切片
	fmt.Printf("  原始: %v, len=%d, cap=%d\n", original, len(original), cap(original))
	fmt.Printf("  引用: %v, len=%d, cap=%d\n", reference, len(reference), cap(reference))
	fmt.Println()

	// 扩容前，两个切片共享底层数组
	original[0] = 999
	fmt.Printf("  修改 original[0] = 999:\n")
	fmt.Printf("    original: %v\n", original)
	fmt.Printf("    reference: %v (共享底层数组，也被修改)\n", reference)
	fmt.Println()

	// 扩容后，新切片指向新数组
	original = append(original, 4, 5, 6) // 触发扩容
	fmt.Printf("  original append(4,5,6) 后（触发扩容）:\n")
	fmt.Printf("    original: %v, len=%d, cap=%d (指向新数组)\n", original, len(original), cap(original))
	fmt.Printf("    reference: %v (仍指向原数组)\n", reference)
	fmt.Println()

	// 修改新切片，不影响原引用
	original[0] = 111
	fmt.Printf("  修改扩容后的 original[0] = 111:\n")
	fmt.Printf("    original: %v (指向新数组)\n", original)
	fmt.Printf("    reference: %v (未受影响，仍指向原数组)\n", reference)
	fmt.Println()

	fmt.Println("关键结论：扩容后，新切片与原切片指向不同的底层数组")
	fmt.Println("         修改新切片不会影响原切片")
	fmt.Println()

	// ============================================
	// 2. 三索引切片操作（a[low:high:max]）
	// ============================================
	fmt.Println("=== 2. 三索引切片操作（a[low:high:max]）===")

	fmt.Println("语法: slice = original[low:high:max]")
	fmt.Println("作用: 显式控制新切片的容量，避免意外占用过多内存")
	fmt.Println()

	fmt.Println("规则:")
	fmt.Println("  - low: 起始索引（包含）")
	fmt.Println("  - high: 结束索引（不包含）")
	fmt.Println("  - max: 新切片的最大容量上限，必须大于等于 high")
	fmt.Println("  - 新切片的长度 = high - low")
	fmt.Println("  - 新切片的容量 = max - low")
	fmt.Println()

	// 基础示例
	fmt.Println("基础示例:")
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	fmt.Printf("  原始切片: %v\n", planets)
	fmt.Printf("  长度: %d, 容量: %d\n", len(planets), cap(planets))
	fmt.Println()

	// 普通切片操作
	terrestrial1 := planets[0:4]
	fmt.Printf("  普通切片 planets[0:4]:\n")
	fmt.Printf("    结果: %v\n", terrestrial1)
	fmt.Printf("    长度: %d, 容量: %d (容量延伸到原数组末尾)\n", len(terrestrial1), cap(terrestrial1))
	fmt.Println()

	// 三索引切片操作
	terrestrial2 := planets[0:4:4]
	fmt.Printf("  三索引切片 planets[0:4:4]:\n")
	fmt.Printf("    结果: %v\n", terrestrial2)
	fmt.Printf("    长度: %d, 容量: %d (容量被限制为4)\n", len(terrestrial2), cap(terrestrial2))
	fmt.Println()

	// 对比：普通切片 vs 三索引切片
	fmt.Println("对比：普通切片 vs 三索引切片")
	slice1 := planets[2:5]
	slice2 := planets[2:5:5]
	fmt.Printf("  普通切片 planets[2:5]: len=%d, cap=%d\n", len(slice1), cap(slice1))
	fmt.Printf("  三索引切片 planets[2:5:5]: len=%d, cap=%d\n", len(slice2), cap(slice2))
	fmt.Println()

	// 演示 append 行为差异
	fmt.Println("演示 append 行为差异:")
	fmt.Printf("  原始: %v\n", planets)

	// 普通切片 append（容量足够，不扩容，会影响原数组）
	slice1 = append(slice1, "NewPlanet")
	fmt.Printf("  普通切片 append 后:\n")
	fmt.Printf("    slice1: %v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("    原数组: %v (被修改了)\n", planets)
	fmt.Println()

	// 重置
	planets = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	slice2 = planets[2:5:5]
	fmt.Printf("  重置后，使用三索引切片:\n")

	// 三索引切片 append（容量不足，会扩容，不影响原数组）
	slice2 = append(slice2, "NewPlanet")
	fmt.Printf("  三索引切片 append 后:\n")
	fmt.Printf("    slice2: %v, len=%d, cap=%d (扩容了)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("    原数组: %v (未受影响)\n", planets)
	fmt.Println()

	// ============================================
	// 3. 核心对比
	// ============================================
	fmt.Println("=== 3. 核心对比 ===")

	fmt.Println("操作对比表:")
	fmt.Println("  操作                          底层数组是否共享  扩容后是否影响原切片")
	fmt.Println("  ──────────────────────────────────────────────────────────────")
	fmt.Println("  普通切片 a[low:high]          是                是（在容量足够时）")
	fmt.Println("  三索引切片 a[low:high:max]    是                否（容量被限制，扩容会创建新数组）")
	fmt.Println()

	// 详细对比示例
	fmt.Println("详细对比示例:")
	base := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("  基础数组: %v\n", base)

	// 普通切片
	normalSlice := base[2:6]
	fmt.Printf("  普通切片 base[2:6]: %v, len=%d, cap=%d\n", normalSlice, len(normalSlice), cap(normalSlice))

	// 三索引切片
	threeIndexSlice := base[2:6:6]
	fmt.Printf("  三索引切片 base[2:6:6]: %v, len=%d, cap=%d\n", threeIndexSlice, len(threeIndexSlice), cap(threeIndexSlice))
	fmt.Println()

	// 测试 append 行为
	fmt.Println("测试 append 行为:")
	normalSlice = append(normalSlice, 99)
	fmt.Printf("  普通切片 append(99) 后:\n")
	fmt.Printf("    normalSlice: %v\n", normalSlice)
	fmt.Printf("    基础数组: %v (被修改了)\n", base)
	fmt.Println()

	// 重置并测试三索引切片
	base = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	threeIndexSlice = base[2:6:6]
	threeIndexSlice = append(threeIndexSlice, 99)
	fmt.Printf("  三索引切片 append(99) 后:\n")
	fmt.Printf("    threeIndexSlice: %v (扩容了)\n", threeIndexSlice)
	fmt.Printf("    基础数组: %v (未受影响)\n", base)
	fmt.Println()

	// ============================================
	// 4. 三索引切片的实用场景
	// ============================================
	fmt.Println("=== 4. 三索引切片的实用场景 ===")

	// 场景1：避免意外修改原数组
	fmt.Println("场景1：避免意外修改原数组")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	safeSlice := data[2:5:5] // 容量限制为3，append 时会扩容
	fmt.Printf("  原始数据: %v\n", data)
	fmt.Printf("  安全切片: %v, len=%d, cap=%d\n", safeSlice, len(safeSlice), cap(safeSlice))
	safeSlice = append(safeSlice, 99, 100)
	fmt.Printf("  append 后:\n")
	fmt.Printf("    安全切片: %v (扩容了)\n", safeSlice)
	fmt.Printf("    原始数据: %v (未受影响)\n", data)
	fmt.Println()

	// 场景2：控制内存占用
	fmt.Println("场景2：控制内存占用")
	largeArray := make([]int, 1000)
	for i := range largeArray {
		largeArray[i] = i
	}
	// 只使用前10个元素，但不想占用整个数组的容量
	smallView := largeArray[0:10:10]
	fmt.Printf("  大数组长度: %d, 容量: %d\n", len(largeArray), cap(largeArray))
	fmt.Printf("  小视图: %v, len=%d, cap=%d (容量被限制)\n", smallView, len(smallView), cap(smallView))
	fmt.Println("  说明：小视图的容量被限制为10，不会占用大数组的容量")
	fmt.Println()

	// 场景3：函数返回独立切片
	fmt.Println("场景3：函数返回独立切片")
	originalData := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	independent := getIndependentSlice(originalData, 2, 5)
	fmt.Printf("  原始数据: %v\n", originalData)
	fmt.Printf("  独立切片: %v, len=%d, cap=%d\n", independent, len(independent), cap(independent))
	independent = append(independent, "x", "y", "z")
	fmt.Printf("  append 后:\n")
	fmt.Printf("    独立切片: %v (扩容了)\n", independent)
	fmt.Printf("    原始数据: %v (未受影响)\n", originalData)
	fmt.Println()

	// 场景4：避免内存泄漏
	fmt.Println("场景4：避免内存泄漏")
	bigData := make([]int, 10000)
	for i := range bigData {
		bigData[i] = i
	}
	// 只使用一小部分，但普通切片会保留整个大数组的引用
	normalView := bigData[0:10]
	limitedView := bigData[0:10:10]
	fmt.Printf("  大数组: len=%d, cap=%d\n", len(bigData), cap(bigData))
	fmt.Printf("  普通视图: len=%d, cap=%d (保留了整个大数组的引用)\n", len(normalView), cap(normalView))
	fmt.Printf("  限制视图: len=%d, cap=%d (只保留10个元素的引用)\n", len(limitedView), cap(limitedView))
	fmt.Println("  说明：限制视图可以避免大数组无法被垃圾回收")
	fmt.Println()

	// ============================================
	// 5. 三索引切片的边界检查
	// ============================================
	fmt.Println("=== 5. 三索引切片的边界检查 ===")

	fmt.Println("三索引切片的约束条件:")
	fmt.Println("  0 <= low <= high <= max <= cap(original)")
	fmt.Println("  如果违反约束，会在运行时 panic")
	fmt.Println()

	// 正确示例
	fmt.Println("正确示例:")
	valid := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	validSlice := valid[2:5:7]
	fmt.Printf("  valid[2:5:7]: %v, len=%d, cap=%d\n", validSlice, len(validSlice), cap(validSlice))
	fmt.Println("  说明：low=2, high=5, max=7，满足 2 <= 5 <= 7 <= 10")
	fmt.Println()

	// 错误示例（已注释，避免 panic）
	fmt.Println("错误示例（已注释，避免 panic）:")
	fmt.Println("  invalid := []int{1, 2, 3, 4, 5}")
	fmt.Println("  invalid[2:5:3]  // ❌ panic: max(3) < high(5)")
	fmt.Println("  invalid[2:5:10] // ❌ panic: max(10) > cap(5)")
	fmt.Println()

	// ============================================
	// 6. 实际应用示例
	// ============================================
	fmt.Println("=== 6. 实际应用示例 ===")

	// 示例1：处理大文件的部分数据
	fmt.Println("示例1：处理大文件的部分数据")
	fileData := make([]byte, 10000)
	for i := range fileData {
		fileData[i] = byte(i % 256)
	}
	// 只处理前100个字节，但不想保留整个文件的引用
	chunk := fileData[0:100:100]
	fmt.Printf("  文件大小: %d 字节\n", len(fileData))
	fmt.Printf("  处理块: %d 字节, cap=%d (容量被限制)\n", len(chunk), cap(chunk))
	fmt.Println("  说明：处理完后，大文件可以被垃圾回收")
	fmt.Println()

	// 示例2：API 返回数据的子集
	fmt.Println("示例2：API 返回数据的子集")
	apiResponse := []string{"user1", "user2", "user3", "user4", "user5", "user6", "user7", "user8"}
	// 只返回前3个用户，后续修改不会影响原数据
	userSubset := apiResponse[0:3:3]
	fmt.Printf("  API 响应: %v\n", apiResponse)
	fmt.Printf("  用户子集: %v, len=%d, cap=%d\n", userSubset, len(userSubset), cap(userSubset))
	userSubset = append(userSubset, "newUser")
	fmt.Printf("  append 后:\n")
	fmt.Printf("    用户子集: %v (扩容了)\n", userSubset)
	fmt.Printf("    API 响应: %v (未受影响)\n", apiResponse)
	fmt.Println()

	// ============================================
	// 7. 总结
	// ============================================
	fmt.Println("=== 7. 总结 ===")
	fmt.Println()
	fmt.Println("1. append 函数的底层扩容机制:")
	fmt.Println("   ✅ 触发条件：新元素数量超出切片当前容量")
	fmt.Println("   ✅ 小切片（<1024）：容量翻倍")
	fmt.Println("   ✅ 大切片（>=1024）：容量增加约25%")
	fmt.Println("   ✅ 扩容后指向新底层数组")
	fmt.Println()
	fmt.Println("2. 三索引切片操作:")
	fmt.Println("   ✅ 语法：a[low:high:max]")
	fmt.Println("   ✅ 显式控制新切片的容量")
	fmt.Println("   ✅ 长度 = high - low")
	fmt.Println("   ✅ 容量 = max - low")
	fmt.Println()
	fmt.Println("3. 核心对比:")
	fmt.Println("   ✅ 普通切片：容量足够时 append 会影响原数组")
	fmt.Println("   ✅ 三索引切片：容量被限制，append 会扩容，不影响原数组")
	fmt.Println()
	fmt.Println("4. 实用场景:")
	fmt.Println("   ✅ 避免意外修改原数组")
	fmt.Println("   ✅ 控制内存占用")
	fmt.Println("   ✅ 函数返回独立切片")
	fmt.Println("   ✅ 避免内存泄漏")
	fmt.Println()
	fmt.Println("5. 最佳实践:")
	fmt.Println("   ✅ 需要独立切片时使用三索引切片")
	fmt.Println("   ✅ 处理大数据时使用三索引切片避免内存泄漏")
	fmt.Println("   ✅ 注意边界检查：0 <= low <= high <= max <= cap(original)")
	fmt.Println()
}

// ============================================
// 辅助函数
// ============================================

// getIndependentSlice 返回一个独立的切片，不会影响原数组
func getIndependentSlice(data []string, start, end int) []string {
	if end > len(data) {
		end = len(data)
	}
	// 使用三索引切片，限制容量
	return data[start:end:end]
}

