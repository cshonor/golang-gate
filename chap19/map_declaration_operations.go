// 示例：Go 语言映射的声明与基本操作
// 演示映射的声明语法、初始化、赋值、取值和典型适用场景

package main

import (
	"fmt"
	"strings"
)

func main() {
	// ============================================
	// 1. 映射的声明语法
	// ============================================
	fmt.Println("=== 1. 映射的声明语法 ===")

	fmt.Println("映射的声明格式: map[KeyType]ValueType")
	fmt.Println("  - KeyType: 键的类型，必须是可比较的（如 string、int、bool 等）")
	fmt.Println("  - ValueType: 值的类型，可以是任意类型")
	fmt.Println()

	// 示例：不同键值类型的映射声明
	fmt.Println("示例：不同键值类型的映射声明")
	var map1 map[string]int        // 键为字符串，值为整数
	var map2 map[int]string        // 键为整数，值为字符串
	var map3 map[string]bool       // 键为字符串，值为布尔值
	var map4 map[int][]string      // 键为整数，值为字符串切片
	var map5 map[string]map[string]int // 键为字符串，值为映射

	fmt.Printf("  map[string]int: %T\n", map1)
	fmt.Printf("  map[int]string: %T\n", map2)
	fmt.Printf("  map[string]bool: %T\n", map3)
	fmt.Printf("  map[int][]string: %T\n", map4)
	fmt.Printf("  map[string]map[string]int: %T\n", map5)
	fmt.Println()

	// ============================================
	// 2. 映射的初始化与赋值
	// ============================================
	fmt.Println("=== 2. 映射的初始化与赋值 ===")

	// 方式1：复合字面量初始化
	fmt.Println("方式1：复合字面量初始化")
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Printf("  初始化: %v\n", temperature)
	fmt.Println()

	// 方式2：先声明后赋值
	fmt.Println("方式2：先声明后赋值")
	var scores map[string]int
	scores = make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	fmt.Printf("  赋值后: %v\n", scores)
	fmt.Println()

	// 方式3：使用 make 函数
	fmt.Println("方式3：使用 make 函数")
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	fmt.Printf("  使用 make: %v\n", ages)
	fmt.Println()

	// 修改值
	fmt.Println("修改值:")
	fmt.Printf("  修改前: %v\n", temperature)
	temperature["Earth"] = 16 // 修改已有键的值
	fmt.Printf("  修改 Earth 后: %v\n", temperature)
	fmt.Println()

	// 添加新的键值对
	fmt.Println("添加新的键值对:")
	temperature["Venus"] = 464 // 添加新的键值对
	fmt.Printf("  添加 Venus 后: %v\n", temperature)
	fmt.Println()

	// ============================================
	// 3. 映射的取值
	// ============================================
	fmt.Println("=== 3. 映射的取值 ===")

	// 直接取值
	fmt.Println("直接取值:")
	temp := temperature["Earth"]
	fmt.Printf("  temperature[\"Earth\"] = %d\n", temp)
	fmt.Printf("  On average the Earth is %v C.\n", temp)
	fmt.Println()

	// 键不存在的情况
	fmt.Println("键不存在的情况:")
	marsTemp := temperature["Mars"]
	fmt.Printf("  temperature[\"Mars\"] = %d\n", marsTemp)
	fmt.Printf("  On average Mars is %v C.\n", marsTemp)

	// 不存在的键返回零值
	unknownTemp := temperature["Jupiter"]
	fmt.Printf("  temperature[\"Jupiter\"] = %d (键不存在，返回零值)\n", unknownTemp)
	fmt.Println()

	// 检查键是否存在
	fmt.Println("检查键是否存在:")
	earthTemp, exists := temperature["Earth"]
	if exists {
		fmt.Printf("  Earth 温度: %d C (存在)\n", earthTemp)
	} else {
		fmt.Printf("  Earth 不存在\n")
	}

	jupiterTemp, exists := temperature["Jupiter"]
	if exists {
		fmt.Printf("  Jupiter 温度: %d C (存在)\n", jupiterTemp)
	} else {
		fmt.Printf("  Jupiter 不存在 (返回零值: %d)\n", jupiterTemp)
	}
	fmt.Println()

	// ============================================
	// 4. 映射的典型适用场景
	// ============================================
	fmt.Println("=== 4. 映射的典型适用场景 ===")

	// 场景1：统计计数
	fmt.Println("场景1：统计计数（统计单词出现次数）")
	text := "apple banana apple cherry banana apple"
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	// 统计每个单词出现的次数
	for _, word := range words {
		wordCount[word]++ // 如果键不存在，会先初始化为0，然后++
	}
	fmt.Printf("  文本: %q\n", text)
	fmt.Printf("  单词计数: %v\n", wordCount)
	fmt.Println()

	// 场景2：配置存储
	fmt.Println("场景2：配置存储（存储键值对形式的配置信息）")
	config := map[string]string{
		"host":     "localhost",
		"port":     "8080",
		"database": "mydb",
		"username": "admin",
	}
	fmt.Printf("  配置信息: %v\n", config)
	fmt.Printf("  数据库主机: %s\n", config["host"])
	fmt.Printf("  端口: %s\n", config["port"])
	fmt.Println()

	// 场景3：快速查找（用户ID到用户信息的映射）
	fmt.Println("场景3：快速查找（用户ID到用户信息的映射）")
	type User struct {
		Name  string
		Email string
		Age   int
	}

	users := map[int]User{
		1: {Name: "Alice", Email: "alice@example.com", Age: 25},
		2: {Name: "Bob", Email: "bob@example.com", Age: 30},
		3: {Name: "Charlie", Email: "charlie@example.com", Age: 28},
	}

	// 通过用户ID快速查找用户信息
	userID := 2
	if user, exists := users[userID]; exists {
		fmt.Printf("  用户ID %d: %s (%s, %d岁)\n", userID, user.Name, user.Email, user.Age)
	} else {
		fmt.Printf("  用户ID %d 不存在\n", userID)
	}
	fmt.Println()

	// 场景4：分组
	fmt.Println("场景4：分组（按类别分组）")
	items := []struct {
		Name  string
		Category string
	}{
		{"Apple", "Fruit"},
		{"Banana", "Fruit"},
		{"Carrot", "Vegetable"},
		{"Tomato", "Vegetable"},
		{"Orange", "Fruit"},
	}

	categoryGroups := make(map[string][]string)
	for _, item := range items {
		categoryGroups[item.Category] = append(categoryGroups[item.Category], item.Name)
	}
	fmt.Printf("  分组结果: %v\n", categoryGroups)
	fmt.Println()

	// 场景5：缓存
	fmt.Println("场景5：缓存（存储计算结果）")
	fibCache := make(map[int]int)
	fibCache[0] = 0
	fibCache[1] = 1
	fibCache[2] = 1
	fibCache[3] = 2
	fmt.Printf("  斐波那契缓存: %v\n", fibCache)
	fmt.Printf("  fib(3) = %d (从缓存获取)\n", fibCache[3])
	fmt.Println()

	// ============================================
	// 5. 映射操作速查表
	// ============================================
	fmt.Println("=== 5. 映射操作速查表 ===")
	fmt.Println()
	fmt.Println("1. 声明:")
	fmt.Println("   var m map[KeyType]ValueType")
	fmt.Println("   m := make(map[KeyType]ValueType)")
	fmt.Println("   m := map[KeyType]ValueType{key: value, ...}")
	fmt.Println()
	fmt.Println("2. 初始化:")
	fmt.Println("   m := map[string]int{\"key\": value}")
	fmt.Println("   m := make(map[string]int)")
	fmt.Println("   m := map[string]int{}")
	fmt.Println()
	fmt.Println("3. 添加/修改元素:")
	fmt.Println("   m[\"key\"] = value")
	fmt.Println()
	fmt.Println("4. 获取元素:")
	fmt.Println("   value := m[\"key\"]")
	fmt.Println("   value, exists := m[\"key\"]  // 检查键是否存在")
	fmt.Println()
	fmt.Println("5. 删除元素:")
	fmt.Println("   delete(m, \"key\")")
	fmt.Println()
	fmt.Println("6. 迭代:")
	fmt.Println("   for key, value := range m { ... }")
	fmt.Println("   for key := range m { ... }")
	fmt.Println("   for _, value := range m { ... }")
	fmt.Println()
	fmt.Println("7. 长度:")
	fmt.Println("   len(m)")
	fmt.Println()
	fmt.Println("8. 检查 nil:")
	fmt.Println("   if m == nil { ... }")
	fmt.Println()
	fmt.Println("9. 判断键是否存在:")
	fmt.Println("   value, exists := m[\"key\"]")
	fmt.Println("   if exists { ... }")
	fmt.Println()

	// ============================================
	// 6. 实际应用示例
	// ============================================
	fmt.Println("=== 6. 实际应用示例 ===")

	// 示例1：温度转换表
	fmt.Println("示例1：温度转换表")
	celsiusToFahrenheit := map[string]float64{
		"Freezing": 0.0,
		"Boiling":  100.0,
		"Room":     20.0,
	}
	for name, celsius := range celsiusToFahrenheit {
		fahrenheit := celsius*9/5 + 32
		fmt.Printf("  %s: %.1f°C = %.1f°F\n", name, celsius, fahrenheit)
	}
	fmt.Println()

	// 示例2：字符频率统计
	fmt.Println("示例2：字符频率统计")
	text2 := "hello world"
	charFreq := make(map[rune]int)
	for _, char := range text2 {
		if char != ' ' { // 忽略空格
			charFreq[char]++
		}
	}
	fmt.Printf("  文本: %q\n", text2)
	fmt.Printf("  字符频率: %v\n", charFreq)
	fmt.Println()

	// 示例3：索引映射
	fmt.Println("示例3：索引映射（值到索引的映射）")
	items2 := []string{"apple", "banana", "cherry", "apple", "banana"}
	indexMap := make(map[string][]int)
	for i, item := range items2 {
		indexMap[item] = append(indexMap[item], i)
	}
	fmt.Printf("  数组: %v\n", items2)
	fmt.Printf("  索引映射: %v\n", indexMap)
	fmt.Println()

	// ============================================
	// 7. 常见操作模式
	// ============================================
	fmt.Println("=== 7. 常见操作模式 ===")

	// 模式1：默认值处理
	fmt.Println("模式1：默认值处理")
	settings := map[string]int{
		"timeout": 30,
	}
	timeout := settings["timeout"]
	if timeout == 0 {
		timeout = 60 // 默认值
	}
	fmt.Printf("  超时设置: %d 秒\n", timeout)
	fmt.Println()

	// 模式2：存在性检查
	fmt.Println("模式2：存在性检查")
	userMap := map[string]string{
		"admin": "Alice",
		"user":  "Bob",
	}
	if name, exists := userMap["admin"]; exists {
		fmt.Printf("  管理员: %s\n", name)
	}
	if name, exists := userMap["guest"]; !exists {
		fmt.Printf("  访客不存在，返回零值: %q\n", name)
	}
	fmt.Println()

	// 模式3：初始化嵌套映射
	fmt.Println("模式3：初始化嵌套映射")
	matrix := make(map[string]map[string]int)
	matrix["row1"] = make(map[string]int)
	matrix["row1"]["col1"] = 1
	matrix["row1"]["col2"] = 2
	matrix["row2"] = make(map[string]int)
	matrix["row2"]["col1"] = 3
	matrix["row2"]["col2"] = 4
	fmt.Printf("  矩阵: %v\n", matrix)
	fmt.Println()

	// ============================================
	// 8. 总结
	// ============================================
	fmt.Println("=== 8. 总结 ===")
	fmt.Println()
	fmt.Println("1. 映射的声明语法:")
	fmt.Println("   ✅ map[KeyType]ValueType")
	fmt.Println("   ✅ 键的类型必须是可比较的")
	fmt.Println("   ✅ 值的类型可以是任意类型")
	fmt.Println()
	fmt.Println("2. 映射的初始化与赋值:")
	fmt.Println("   ✅ 复合字面量初始化")
	fmt.Println("   ✅ 使用 make 函数")
	fmt.Println("   ✅ 通过键直接赋值")
	fmt.Println()
	fmt.Println("3. 映射的取值:")
	fmt.Println("   ✅ 直接通过键获取值")
	fmt.Println("   ✅ 键不存在返回零值")
	fmt.Println("   ✅ 使用 value, exists := map[key] 检查键是否存在")
	fmt.Println()
	fmt.Println("4. 典型适用场景:")
	fmt.Println("   ✅ 统计计数")
	fmt.Println("   ✅ 配置存储")
	fmt.Println("   ✅ 快速查找")
	fmt.Println("   ✅ 分组")
	fmt.Println("   ✅ 缓存")
	fmt.Println()
	fmt.Println("5. 最佳实践:")
	fmt.Println("   ✅ 使用 value, exists := map[key] 检查键是否存在")
	fmt.Println("   ✅ 注意键不存在时返回零值")
	fmt.Println("   ✅ 合理使用映射进行快速查找")
	fmt.Println()
}

