// 示例：Go 语言的映射（Map）
// 演示映射的本质、特点、声明、初始化、增删改查和迭代

package main

import "fmt"

func main() {
	// ============================================
	// 1. 映射的本质与用途
	// ============================================
	fmt.Println("=== 1. 映射的本质与用途 ===")

	fmt.Println("本质：映射是一种键值对（Key-Value）数据结构")
	fmt.Println("     用于快速通过键查找对应的值")
	fmt.Println("     类似于其他语言中的字典、哈希表或对象")
	fmt.Println()

	fmt.Println("核心用途：")
	fmt.Println("  1. 作为非结构化数据的收集器")
	fmt.Println("     存储那些在运行时才能确定键的数据")
	fmt.Println("  2. 实现高效的查找、插入和删除操作")
	fmt.Println("     时间复杂度接近O(1)")
	fmt.Println()

	// ============================================
	// 2. 跨语言对比
	// ============================================
	fmt.Println("=== 2. 跨语言对比 ===")

	fmt.Println("不同语言中映射的称呼:")
	fmt.Println("  语言          映射的称呼")
	fmt.Println("  ──────────────────────────────")
	fmt.Println("  Go            Map")
	fmt.Println("  Python        Dictionary")
	fmt.Println("  Ruby          Hash")
	fmt.Println("  JavaScript    Object/Map")
	fmt.Println("  PHP           关联数组")
	fmt.Println()

	// ============================================
	// 3. Go语言映射的特点
	// ============================================
	fmt.Println("=== 3. Go语言映射的特点 ===")

	fmt.Println("特点1：键的类型限制")
	fmt.Println("  - 键的类型必须是可比较的")
	fmt.Println("  - 例如：string、int、bool 等")
	fmt.Println("  - 不能是切片、映射或函数这类不可比较的类型")
	fmt.Println()

	fmt.Println("特点2：值的类型灵活")
	fmt.Println("  - 值可以是任意类型")
	fmt.Println("  - 包括切片、映射甚至结构体")
	fmt.Println()

	fmt.Println("特点3：引用类型")
	fmt.Println("  - 映射是引用类型")
	fmt.Println("  - 传递映射时传递的是引用")
	fmt.Println("  - 修改会影响原始映射")
	fmt.Println()

	// ============================================
	// 4. 映射的声明和初始化
	// ============================================
	fmt.Println("=== 4. 映射的声明和初始化 ===")

	// 方式1：使用 make 函数
	fmt.Println("方式1：使用 make 函数")
	map1 := make(map[string]int)
	fmt.Printf("  make(map[string]int): %v\n", map1)
	fmt.Printf("  类型: %T\n", map1)
	fmt.Println()

	// 方式2：使用字面量
	fmt.Println("方式2：使用字面量")
	map2 := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}
	fmt.Printf("  字面量初始化: %v\n", map2)
	fmt.Println()

	// 方式3：空映射字面量
	fmt.Println("方式3：空映射字面量")
	map3 := map[string]int{}
	fmt.Printf("  空映射: %v\n", map3)
	fmt.Println()

	// 方式4：声明后初始化
	fmt.Println("方式4：声明后初始化")
	var map4 map[string]int
	fmt.Printf("  声明后（nil）: %v, nil=%t\n", map4, map4 == nil)
	// map4["key"] = 1 // ❌ 错误：不能向 nil 映射写入
	map4 = make(map[string]int)
	map4["key"] = 1
	fmt.Printf("  初始化后: %v\n", map4)
	fmt.Println()

	// ============================================
	// 5. 映射的增删改查
	// ============================================
	fmt.Println("=== 5. 映射的增删改查 ===")

	// 创建映射
	scores := make(map[string]int)

	// 添加/修改元素
	fmt.Println("添加/修改元素:")
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("  添加后: %v\n", scores)

	// 修改元素
	scores["Alice"] = 98
	fmt.Printf("  修改后: %v\n", scores)
	fmt.Println()

	// 查找元素
	fmt.Println("查找元素:")
	// 方式1：直接访问（如果键不存在，返回零值）
	value1 := scores["Alice"]
	fmt.Printf("  scores[\"Alice\"] = %d\n", value1)

	value2 := scores["David"] // 键不存在
	fmt.Printf("  scores[\"David\"] = %d (键不存在，返回零值)\n", value2)
	fmt.Println()

	// 方式2：检查键是否存在
	fmt.Println("检查键是否存在:")
	value3, exists := scores["Bob"]
	if exists {
		fmt.Printf("  scores[\"Bob\"] = %d (存在)\n", value3)
	} else {
		fmt.Printf("  scores[\"Bob\"] 不存在\n")
	}

	value4, exists := scores["David"]
	if exists {
		fmt.Printf("  scores[\"David\"] = %d (存在)\n", value4)
	} else {
		fmt.Printf("  scores[\"David\"] 不存在\n")
	}
	fmt.Println()

	// 删除元素
	fmt.Println("删除元素:")
	fmt.Printf("  删除前: %v\n", scores)
	delete(scores, "Bob")
	fmt.Printf("  删除 \"Bob\" 后: %v\n", scores)
	fmt.Println()

	// 删除不存在的键（不会报错）
	fmt.Println("删除不存在的键（不会报错）:")
	delete(scores, "Nonexistent")
	fmt.Printf("  删除 \"Nonexistent\" 后: %v (无变化)\n", scores)
	fmt.Println()

	// ============================================
	// 6. 映射的迭代
	// ============================================
	fmt.Println("=== 6. 映射的迭代 ===")

	planets := map[string]string{
		"Mercury": "类地行星",
		"Venus":   "类地行星",
		"Earth":   "类地行星",
		"Mars":    "类地行星",
		"Jupiter": "气态巨行星",
		"Saturn":  "气态巨行星",
	}

	fmt.Println("使用 for range 迭代:")
	for key, value := range planets {
		fmt.Printf("  %s: %s\n", key, value)
	}
	fmt.Println()

	fmt.Println("只迭代键:")
	for key := range planets {
		fmt.Printf("  %s\n", key)
	}
	fmt.Println()

	fmt.Println("只迭代值（使用 _ 忽略键）:")
	for _, value := range planets {
		fmt.Printf("  %s\n", value)
	}
	fmt.Println()

	// 注意：迭代顺序是随机的
	fmt.Println("注意：映射的迭代顺序是随机的（Go 1.0+）")
	fmt.Println("     每次运行可能得到不同的顺序")
	fmt.Println()

	// ============================================
	// 7. 映射的零值和 nil 检查
	// ============================================
	fmt.Println("=== 7. 映射的零值和 nil 检查 ===")

	var nilMap map[string]int
	fmt.Printf("  nil 映射: %v, nil=%t\n", nilMap, nilMap == nil)

	// 检查映射是否为 nil
	if nilMap == nil {
		fmt.Println("  映射是 nil")
	}

	// nil 映射可以读取（返回零值），但不能写入
	zeroValue := nilMap["key"]
	fmt.Printf("  从 nil 映射读取: %d (零值)\n", zeroValue)
	// nilMap["key"] = 1 // ❌ 错误：不能向 nil 映射写入
	fmt.Println()

	// ============================================
	// 8. 映射作为值类型
	// ============================================
	fmt.Println("=== 8. 映射作为值类型 ===")

	fmt.Println("映射的值可以是任意类型，包括切片、映射、结构体等")
	fmt.Println()

	// 映射的值是切片
	fmt.Println("示例1：映射的值是切片")
	groups := map[string][]string{
		"类地行星": {"Mercury", "Venus", "Earth", "Mars"},
		"气态巨行星": {"Jupiter", "Saturn"},
		"冰巨行星": {"Uranus", "Neptune"},
	}
	fmt.Printf("  分组: %v\n", groups)
	fmt.Println()

	// 映射的值是映射（嵌套映射）
	fmt.Println("示例2：映射的值是映射（嵌套映射）")
	students := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"English":  88,
			"Science":  92,
		},
		"Bob": {
			"Math":    87,
			"English":  90,
			"Science":  85,
		},
	}
	fmt.Printf("  学生成绩: %v\n", students)
	fmt.Printf("  Alice 的数学成绩: %d\n", students["Alice"]["Math"])
	fmt.Println()

	// ============================================
	// 9. 映射的引用类型特性
	// ============================================
	fmt.Println("=== 9. 映射的引用类型特性 ===")

	fmt.Println("映射是引用类型，传递映射时传递的是引用")
	fmt.Println()

	original := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Printf("  原始映射: %v\n", original)

	// 赋值（共享引用）
	copy := original
	fmt.Printf("  赋值后: copy = %v\n", copy)

	// 修改副本
	copy["d"] = 4
	fmt.Printf("  修改 copy 后:\n")
	fmt.Printf("    original: %v (也被修改了)\n", original)
	fmt.Printf("    copy: %v\n", copy)
	fmt.Println()

	// 函数参数传递
	fmt.Println("函数参数传递（传递引用）:")
	testMap := map[string]int{"x": 10, "y": 20}
	fmt.Printf("  调用函数前: %v\n", testMap)
	modifyMap(testMap)
	fmt.Printf("  调用函数后: %v (被修改了)\n", testMap)
	fmt.Println()

	// ============================================
	// 10. 映射的长度和容量
	// ============================================
	fmt.Println("=== 10. 映射的长度和容量 ===")

	fmt.Println("映射只有长度（len），没有容量（cap）")
	fmt.Println()

	data := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Printf("  映射: %v\n", data)
	fmt.Printf("  长度: %d\n", len(data))

	// 添加元素
	data["four"] = 4
	fmt.Printf("  添加元素后长度: %d\n", len(data))

	// 删除元素
	delete(data, "one")
	fmt.Printf("  删除元素后长度: %d\n", len(data))
	fmt.Println()

	// ============================================
	// 11. 映射的键类型限制
	// ============================================
	fmt.Println("=== 11. 映射的键类型限制 ===")

	fmt.Println("键的类型必须是可比较的（comparable）")
	fmt.Println()

	// 可用的键类型
	fmt.Println("可用的键类型:")
	validKeys := map[string]string{
		"string": "字符串",
		"int":    "整数",
		"bool":   "布尔值",
		"float":  "浮点数",
	}
	fmt.Printf("  %v\n", validKeys)
	fmt.Println()

	// 不可用的键类型（已注释）
	fmt.Println("不可用的键类型（已注释）:")
	fmt.Println("  // map[[]int]string{}  // ❌ 切片不能作为键")
	fmt.Println("  // map[map[string]int]string{}  // ❌ 映射不能作为键")
	fmt.Println("  // map[func()]string{}  // ❌ 函数不能作为键")
	fmt.Println()

	// ============================================
	// 12. 实际应用示例
	// ============================================
	fmt.Println("=== 12. 实际应用示例 ===")

	// 示例1：计数器
	fmt.Println("示例1：计数器")
	wordCount := make(map[string]int)
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	for _, word := range words {
		wordCount[word]++
	}
	fmt.Printf("  单词计数: %v\n", wordCount)
	fmt.Println()

	// 示例2：分组
	fmt.Println("示例2：分组")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenOdd := make(map[string][]int)
	for _, num := range numbers {
		if num%2 == 0 {
			evenOdd["even"] = append(evenOdd["even"], num)
		} else {
			evenOdd["odd"] = append(evenOdd["odd"], num)
		}
	}
	fmt.Printf("  奇偶分组: %v\n", evenOdd)
	fmt.Println()

	// 示例3：缓存
	fmt.Println("示例3：缓存（模拟）")
	cache := make(map[string]string)
	cache["user:1"] = "Alice"
	cache["user:2"] = "Bob"
	fmt.Printf("  缓存: %v\n", cache)
	// 查找
	if value, exists := cache["user:1"]; exists {
		fmt.Printf("  查找 user:1: %s\n", value)
	}
	fmt.Println()

	// ============================================
	// 13. 映射核心操作清单
	// ============================================
	fmt.Println("=== 13. 映射核心操作清单 ===")
	fmt.Println()
	fmt.Println("1. 声明和初始化:")
	fmt.Println("   ✅ make(map[KeyType]ValueType)")
	fmt.Println("   ✅ map[KeyType]ValueType{key: value, ...}")
	fmt.Println("   ✅ map[KeyType]ValueType{}")
	fmt.Println()
	fmt.Println("2. 添加/修改元素:")
	fmt.Println("   ✅ map[key] = value")
	fmt.Println()
	fmt.Println("3. 查找元素:")
	fmt.Println("   ✅ value := map[key] (键不存在返回零值)")
	fmt.Println("   ✅ value, exists := map[key] (检查键是否存在)")
	fmt.Println()
	fmt.Println("4. 删除元素:")
	fmt.Println("   ✅ delete(map, key)")
	fmt.Println()
	fmt.Println("5. 迭代:")
	fmt.Println("   ✅ for key, value := range map")
	fmt.Println("   ✅ for key := range map")
	fmt.Println("   ✅ for _, value := range map")
	fmt.Println()
	fmt.Println("6. 长度:")
	fmt.Println("   ✅ len(map)")
	fmt.Println()
	fmt.Println("7. 检查 nil:")
	fmt.Println("   ✅ if map == nil")
	fmt.Println()
	fmt.Println("8. 注意事项:")
	fmt.Println("   ⚠️  键的类型必须是可比较的")
	fmt.Println("   ⚠️  映射是引用类型")
	fmt.Println("   ⚠️  迭代顺序是随机的")
	fmt.Println("   ⚠️  nil 映射不能写入，但可以读取（返回零值）")
	fmt.Println()

	// ============================================
	// 14. 总结
	// ============================================
	fmt.Println("=== 14. 总结 ===")
	fmt.Println()
	fmt.Println("1. 映射的本质:")
	fmt.Println("   ✅ 键值对数据结构")
	fmt.Println("   ✅ 快速查找、插入、删除（O(1)）")
	fmt.Println("   ✅ 非结构化数据的收集器")
	fmt.Println()
	fmt.Println("2. Go语言映射的特点:")
	fmt.Println("   ✅ 键的类型必须是可比较的")
	fmt.Println("   ✅ 值的类型灵活（任意类型）")
	fmt.Println("   ✅ 引用类型，传递时传递引用")
	fmt.Println()
	fmt.Println("3. 核心操作:")
	fmt.Println("   ✅ 声明、初始化、增删改查")
	fmt.Println("   ✅ 迭代、长度检查、nil 检查")
	fmt.Println()
	fmt.Println("4. 最佳实践:")
	fmt.Println("   ✅ 使用 value, exists := map[key] 检查键是否存在")
	fmt.Println("   ✅ 注意映射是引用类型")
	fmt.Println("   ✅ 注意迭代顺序是随机的")
	fmt.Println("   ✅ 避免向 nil 映射写入")
	fmt.Println()
}

// ============================================
// 辅助函数
// ============================================

// modifyMap 修改映射（演示引用传递）
func modifyMap(m map[string]int) {
	m["z"] = 30
	fmt.Printf("    函数内部修改: %v\n", m)
}

