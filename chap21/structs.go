// 示例：Go 语言的结构（Struct）
// 演示结构的声明、使用、字段访问和JSON编码

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

// ============================================
// 类型定义（包级别）
// ============================================

// Coordinate 坐标结构
type Coordinate struct {
	Lat  float64
	Long float64
}

func main() {
	// ============================================
	// 1. 结构的核心价值
	// ============================================
	fmt.Println("=== 1. 结构的核心价值 ===")

	fmt.Println("解决的问题：")
	fmt.Println("  当需要描述一个包含多种属性的实体时")
	fmt.Println("  用单独的变量会非常零散且难以管理")
	fmt.Println("  结构可以将这些不同类型的属性打包成一个整体")
	fmt.Println()

	fmt.Println("与映射、切片的区别：")
	fmt.Println("  - 切片和映射只能存储同一类型的数据")
	fmt.Println("  - 结构可以存储不同类型的数据")
	fmt.Println("  - 更适合描述现实世界中的复杂对象")
	fmt.Println()

	fmt.Println("生活中的例子：")
	fmt.Println("  - 描述一个人：姓名（string）+ 年龄（int）+ 身高（float64）")
	fmt.Println("  - 描述一个坐标：纬度（float64）+ 经度（float64）")
	fmt.Println()

	// ============================================
	// 2. 声明与使用结构
	// ============================================
	fmt.Println("=== 2. 声明与使用结构 ===")

	// 示例1：火星坐标
	fmt.Println("示例1：火星坐标")
	fmt.Println("  声明: type Coordinate struct { Lat float64; Long float64 }")

	// 创建一个结构实例
	curiosity := Coordinate{-4.5895, 137.4417}
	fmt.Printf("  创建实例: Coordinate{-4.5895, 137.4417}\n")
	fmt.Printf("  纬度: %.4f\n", curiosity.Lat)
	fmt.Printf("  经度: %.4f\n", curiosity.Long)
	fmt.Println()

	// 访问结构的字段
	fmt.Println("访问结构的字段:")
	fmt.Printf("  curiosity.Lat = %.4f\n", curiosity.Lat)
	fmt.Printf("  curiosity.Long = %.4f\n", curiosity.Long)
	fmt.Println()

	// 修改字段
	fmt.Println("修改字段:")
	curiosity.Lat = -4.5896
	fmt.Printf("  修改后: %.4f, %.4f\n", curiosity.Lat, curiosity.Long)
	fmt.Println()

	// 示例2：使用结构简化函数参数
	fmt.Println("示例2：使用结构简化函数参数")
	fmt.Println("  原本需要4个参数的距离计算函数:")
	fmt.Println("    func distance(lat1, long1, lat2, long2 float64) float64")
	fmt.Println("  可以简化为接收两个 Coordinate 结构:")
	fmt.Println("    func distance(p1, p2 Coordinate) float64")
	fmt.Println()

	// 演示距离计算
	curiositySite := Coordinate{-4.5895, 137.4417}
	opportunitySite := Coordinate{-1.9462, 354.4734}
	dist := distance(curiositySite, opportunitySite)
	fmt.Printf("  Curiosity 站点: (%.4f, %.4f)\n", curiositySite.Lat, curiositySite.Long)
	fmt.Printf("  Opportunity 站点: (%.4f, %.4f)\n", opportunitySite.Lat, opportunitySite.Long)
	fmt.Printf("  距离: %.2f 公里\n", dist)
	fmt.Println()

	// 示例3：描述一个人
	fmt.Println("示例3：描述一个人")
	type Person struct {
		Name   string
		Age    int
		Height float64
	}

	person := Person{
		Name:   "Alice",
		Age:    25,
		Height: 165.5,
	}
	fmt.Printf("  姓名: %s\n", person.Name)
	fmt.Printf("  年龄: %d\n", person.Age)
	fmt.Printf("  身高: %.1f cm\n", person.Height)
	fmt.Println()

	// 示例4：嵌套结构
	fmt.Println("示例4：嵌套结构")
	type Address struct {
		Street string
		City   string
		Zip    string
	}

	type Employee struct {
		Name    string
		Age     int
		Address Address
	}

	employee := Employee{
		Name: "Bob",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "Beijing",
			Zip:    "100000",
		},
	}
	fmt.Printf("  员工: %s, %d岁\n", employee.Name, employee.Age)
	fmt.Printf("  地址: %s, %s, %s\n", employee.Address.Street, employee.Address.City, employee.Address.Zip)
	fmt.Println()

	// ============================================
	// 3. 结构的初始化方式
	// ============================================
	fmt.Println("=== 3. 结构的初始化方式 ===")

	// 方式1：按顺序提供值
	fmt.Println("方式1：按顺序提供值")
	coord1 := Coordinate{-4.5895, 137.4417}
	fmt.Printf("  Coordinate{-4.5895, 137.4417}: %+v\n", coord1)
	fmt.Println()

	// 方式2：使用字段名
	fmt.Println("方式2：使用字段名")
	coord2 := Coordinate{
		Lat:  -4.5895,
		Long: 137.4417,
	}
	fmt.Printf("  Coordinate{Lat: -4.5895, Long: 137.4417}: %+v\n", coord2)
	fmt.Println()

	// 方式3：部分初始化（未初始化的字段为零值）
	fmt.Println("方式3：部分初始化（未初始化的字段为零值）")
	coord3 := Coordinate{Lat: -4.5895}
	fmt.Printf("  Coordinate{Lat: -4.5895}: %+v (Long 为零值)\n", coord3)
	fmt.Println()

	// 方式4：使用 new 函数（返回指针）
	fmt.Println("方式4：使用 new 函数（返回指针）")
	coord4 := new(Coordinate)
	coord4.Lat = -4.5895
	coord4.Long = 137.4417
	fmt.Printf("  new(Coordinate): %+v\n", coord4)
	fmt.Printf("  类型: %T\n", coord4)
	fmt.Println()
new(结构体名)  是 Go 自带的函数
​
- 作用：创建一个结构体实例，并返回它的「指针」
	// ============================================
	// 4. 结构与JSON编码
	// ============================================
	fmt.Println("=== 4. 结构与JSON编码 ===")

	fmt.Println("Go语言的 encoding/json 包可以很方便地将结构序列化为JSON")
	fmt.Println("这一特性在开发API或处理配置文件时非常实用")
	fmt.Println()

	// 示例1：基本JSON编码
	fmt.Println("示例1：基本JSON编码")
	curiosityJSON := Coordinate{-4.5895, 137.4417}
	data, err := json.Marshal(curiosityJSON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  结构: %+v\n", curiosityJSON)
	fmt.Printf("  JSON: %s\n", string(data))
	fmt.Println()

	// 示例2：使用JSON标签自定义字段名
	fmt.Println("示例2：使用JSON标签自定义字段名")
	type Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	location := Location{-4.5895, 137.4417}
	locationData, err := json.Marshal(location)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  结构: %+v\n", location)
	fmt.Printf("  JSON: %s (字段名已改变)\n", string(locationData))
	fmt.Println()

	// 示例3：JSON解码
	fmt.Println("示例3：JSON解码")
	jsonStr := `{"latitude":-4.5895,"longitude":137.4417}`
	var decodedLocation Location
	err = json.Unmarshal([]byte(jsonStr), &decodedLocation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  JSON字符串: %s\n", jsonStr)
	fmt.Printf("  解码后: %+v\n", decodedLocation)
	fmt.Println()

	// 示例4：复杂结构的JSON编码
	fmt.Println("示例4：复杂结构的JSON编码")
	type User struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Location Location `json:"location"`
	}

	user := User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
		Location: Location{
			Latitude:  -4.5895,
			Longitude: 137.4417,
		},
	}
	userData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  用户结构: %+v\n", user)
	fmt.Printf("  JSON: %s\n", string(userData))
	fmt.Println()

	// 示例5：格式化JSON输出
	fmt.Println("示例5：格式化JSON输出（缩进）")
	userDataPretty, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("  格式化JSON:")
	fmt.Println(string(userDataPretty))
	fmt.Println()

	// ============================================
	// 5. 结构的零值
	// ============================================
	fmt.Println("=== 5. 结构的零值 ===")

	var zeroCoord Coordinate
	fmt.Printf("  零值结构: %+v\n", zeroCoord)
	fmt.Printf("  Lat: %.1f (零值)\n", zeroCoord.Lat)
	fmt.Printf("  Long: %.1f (零值)\n", zeroCoord.Long)
	fmt.Println()

	// ============================================
	// 6. 结构的方法
	// ============================================
	fmt.Println("=== 6. 结构的方法 ===")

	fmt.Println("可以为结构类型定义方法")
	fmt.Println()

	// 为 Coordinate 添加方法
	coord := Coordinate{-4.5895, 137.4417}
	fmt.Printf("  坐标: %+v\n", coord)
	fmt.Printf("  字符串表示: %s\n", coord.String())
	fmt.Printf("  距离原点的距离: %.2f\n", coord.DistanceFromOrigin())
	fmt.Println()

	// ============================================
	// 7. 结构指针
	// ============================================
	fmt.Println("=== 7. 结构指针 ===")

	fmt.Println("使用指针可以避免复制大型结构")
	fmt.Println()

	// 值传递
	coord5 := Coordinate{-4.5895, 137.4417}
	fmt.Printf("  原始: %+v\n", coord5)
	modifyCoordinate(coord5)
	fmt.Printf("  值传递后: %+v (未改变)\n", coord5)
	fmt.Println()

	// 指针传递
	coord6 := Coordinate{-4.5895, 137.4417}
	fmt.Printf("  原始: %+v\n", coord6)
	modifyCoordinatePtr(&coord6)
	fmt.Printf("  指针传递后: %+v (已改变)\n", coord6)
	fmt.Println()

	// ============================================
	// 8. 实际应用示例
	// ============================================
	fmt.Println("=== 8. 实际应用示例 ===")

	// 示例1：API响应结构
	fmt.Println("示例1：API响应结构")
	type APIResponse struct {
		Status  string      `json:"status"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}

	response := APIResponse{
		Status:  "success",
		Code:    200,
		Data:    user,
		Message: "User retrieved successfully",
	}
	responseData, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println("  API响应:")
	fmt.Println(string(responseData))
	fmt.Println()

	// 示例2：配置结构
	fmt.Println("示例2：配置结构")
	type Config struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
		Debug    bool   `json:"debug"`
	}

	config := Config{
		Host:     "localhost",
		Port:     8080,
		Database: "mydb",
		Debug:    true,
	}
	configData, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println("  配置:")
	fmt.Println(string(configData))
	fmt.Println()

	// ============================================
	// 9. 总结
	// ============================================
	fmt.Println("=== 9. 总结 ===")
	fmt.Println()
	fmt.Println("1. 结构的核心价值:")
	fmt.Println("   ✅ 将不同类型的属性打包成一个整体")
	fmt.Println("   ✅ 更适合描述现实世界中的复杂对象")
	fmt.Println("   ✅ 提高代码的可读性和可维护性")
	fmt.Println()
	fmt.Println("2. 声明与使用结构:")
	fmt.Println("   ✅ type StructName struct { Field1 Type1; Field2 Type2 }")
	fmt.Println("   ✅ 创建实例：StructName{value1, value2}")
	fmt.Println("   ✅ 访问字段：instance.Field")
	fmt.Println()
	fmt.Println("3. 结构与JSON编码:")
	fmt.Println("   ✅ json.Marshal() 将结构编码为JSON")
	fmt.Println("   ✅ json.Unmarshal() 将JSON解码为结构")
	fmt.Println("   ✅ 使用 json 标签自定义字段名")
	fmt.Println("   ✅ json.MarshalIndent() 格式化输出")
	fmt.Println()
	fmt.Println("4. 最佳实践:")
	fmt.Println("   ✅ 使用结构简化函数参数")
	fmt.Println("   ✅ 大型结构使用指针传递")
	fmt.Println("   ✅ 使用JSON标签提高API兼容性")
	fmt.Println("   ✅ 合理使用嵌套结构")
	fmt.Println()
}

// ============================================
// 辅助函数和方法
// ============================================

// distance 计算两个坐标之间的距离（使用结构简化参数）
func distance(p1, p2 Coordinate) float64 {
	// 简化的距离计算（实际应该使用更精确的公式）
	latDiff := p1.Lat - p2.Lat
	longDiff := p1.Long - p2.Long
	return math.Sqrt(latDiff*latDiff + longDiff*longDiff) * 111.0 // 粗略转换为公里
}

// String 为 Coordinate 实现 Stringer 接口
func (c Coordinate) String() string {
	return fmt.Sprintf("(%.4f, %.4f)", c.Lat, c.Long)
}

// DistanceFromOrigin 计算距离原点的距离
func (c Coordinate) DistanceFromOrigin() float64 {
	return math.Sqrt(c.Lat*c.Lat + c.Long*c.Long)
}

// modifyCoordinate 修改坐标（值传递，不会改变原值）
func modifyCoordinate(c Coordinate) {
	c.Lat = 0
	c.Long = 0
}

// modifyCoordinatePtr 修改坐标（指针传递，会改变原值）
func modifyCoordinatePtr(c *Coordinate) {
	c.Lat = 0
	c.Long = 0
}

