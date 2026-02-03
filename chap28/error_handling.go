// 示例：Go 语言的错误处理（Error Handling）
// 演示错误处理的核心思想、文件IO错误处理、自定义错误、panic/recover

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// ============================================
	// 1. 错误处理的核心思想
	// ============================================
	fmt.Println("=== 1. 错误处理的核心思想 ===")

	fmt.Println("书里用'消防演习'的比喻非常形象：")
	fmt.Println("  - 软件遇到错误就像学校遇到警报")
	fmt.Println("  - 不能侥幸忽略，而应该像演习一样，提前规划好应对步骤")
	fmt.Println("  - 确保在真正的故障发生时能安全处理")
	fmt.Println()

	fmt.Println("Go鼓励显式、主动的错误处理")
	fmt.Println("而不是依赖异常捕获")
	fmt.Println("这让代码更清晰、更可靠")
	fmt.Println()

	// ============================================
	// 2. Go错误处理的特点
	// ============================================
	fmt.Println("=== 2. Go错误处理的特点 ===")

	fmt.Println("特点1：显式错误返回")
	fmt.Println("  函数通常会返回 (result, error)")
	fmt.Println("  调用者必须显式检查错误")
	fmt.Println("  这避免了'隐藏'的异常")
	fmt.Println()

	fmt.Println("特点2：无异常捕获机制")
	fmt.Println("  Go没有 try/catch")
	fmt.Println("  而是通过返回错误和 panic/recover 来处理不同场景")
	fmt.Println()

	fmt.Println("特点3：错误是值")
	fmt.Println("  错误在Go里是普通的值")
	fmt.Println("  可以像其他值一样被传递、存储和处理")
	fmt.Println("  这让错误处理非常灵活")
	fmt.Println()

	// ============================================
	// 3. 写入文件并处理错误
	// ============================================
	fmt.Println("=== 3. 写入文件并处理错误 ===")

	fmt.Println("在IO操作等常见场景中，如何检查和处理错误")
	fmt.Println()

	// 示例1：写入文件
	fmt.Println("示例1：写入文件")
	err := writeFile("test.txt", "Hello, World!")
	if err != nil {
		fmt.Printf("  写入文件失败: %v\n", err)
	} else {
		fmt.Println("  写入文件成功")
	}
	fmt.Println()

	// 示例2：读取文件
	fmt.Println("示例2：读取文件")
	content, err := readFile("test.txt")
	if err != nil {
		fmt.Printf("  读取文件失败: %v\n", err)
	} else {
		fmt.Printf("  文件内容: %s\n", content)
	}
	fmt.Println()

	// 示例3：文件不存在的情况
	fmt.Println("示例3：文件不存在的情况")
	_, err = readFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("  预期的错误: %v\n", err)
		fmt.Printf("  错误类型: %T\n", err)
	}
	fmt.Println()

	// ============================================
	// 4. 创造性地处理错误
	// ============================================
	fmt.Println("=== 4. 创造性地处理错误 ===")

	fmt.Println("错误包装、错误链传递")
	fmt.Println("让错误信息更具可读性和可追溯性")
	fmt.Println()

	// 示例1：错误包装
	fmt.Println("示例1：错误包装")
	err = processData("data.txt")
	if err != nil {
		fmt.Printf("  包装后的错误: %v\n", err)
		fmt.Printf("  原始错误: %v\n", errors.Unwrap(err))
	}
	fmt.Println()

	// 示例2：错误链传递
	fmt.Println("示例2：错误链传递")
	result, err := complexOperation("input")
	if err != nil {
		fmt.Printf("  错误链: %v\n", err)
		// 可以逐层展开错误
		for err != nil {
			fmt.Printf("    - %v\n", err)
			err = errors.Unwrap(err)
		}
	} else {
		fmt.Printf("  操作成功: %s\n", result)
	}
	fmt.Println()

	// ============================================
	// 5. 创建并标识特定错误
	// ============================================
	fmt.Println("=== 5. 创建并标识特定错误 ===")

	fmt.Println("使用 errors.New() 或 fmt.Errorf() 定义自定义错误")
	fmt.Println("让调用者能识别并处理特定错误类型")
	fmt.Println()

	// 示例1：使用 errors.New()
	fmt.Println("示例1：使用 errors.New()")
	var ErrNotFound = errors.New("资源未找到")
	var ErrPermissionDenied = errors.New("权限被拒绝")

	err = checkResource("resource1")
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("  处理：资源未找到")
		} else if errors.Is(err, ErrPermissionDenied) {
			fmt.Println("  处理：权限被拒绝")
		} else {
			fmt.Printf("  其他错误: %v\n", err)
		}
	}
	fmt.Println()

	// 示例2：使用 fmt.Errorf()
	fmt.Println("示例2：使用 fmt.Errorf()")
	err = validateInput("")
	if err != nil {
		fmt.Printf("  格式化错误: %v\n", err)
	}
	fmt.Println()

	// 示例3：自定义错误类型
	fmt.Println("示例3：自定义错误类型")
	_, err = divideNumbers(10, 0)
	if err != nil {
		var divErr *DivisionError
		if errors.As(err, &divErr) {
			fmt.Printf("  特定错误类型: %v\n", divErr)
			fmt.Printf("  被除数: %d, 除数: %d\n", divErr.Dividend, divErr.Divisor)
		}
	}
	fmt.Println()

	// ============================================
	// 6. 处理"惊恐"（Panic）
	// ============================================
	fmt.Println("=== 6. 处理'惊恐'（Panic）===")

	fmt.Println("理解 panic 和 recover 的使用场景")
	fmt.Println("知道何时用它们来处理不可恢复的错误")
	fmt.Println()

	// 示例1：panic 的基本使用
	fmt.Println("示例1：panic 的基本使用")
	fmt.Println("  panic 用于不可恢复的错误")
	fmt.Println("  通常表示程序遇到了无法继续执行的严重问题")
	fmt.Println()

	// 示例2：recover 的使用
	fmt.Println("示例2：recover 的使用")
	fmt.Println("  使用 defer + recover 捕获 panic")
	safeFunction()
	fmt.Println()

	// 示例3：panic 和 recover 的配合
	fmt.Println("示例3：panic 和 recover 的配合")
	result2 := safeDivide(10, 0)
	fmt.Printf("  安全除法结果: %d\n", result2)
	result3 := safeDivide(10, 2)
	fmt.Printf("  安全除法结果: %d\n", result3)
	fmt.Println()

	// ============================================
	// 7. 错误处理的最佳实践
	// ============================================
	fmt.Println("=== 7. 错误处理的最佳实践 ===")

	fmt.Println("1. 总是检查错误:")
	fmt.Println("   ✅ 不要忽略错误返回值")
	fmt.Println("   ✅ 即使你认为不会出错，也要检查")
	fmt.Println()

	fmt.Println("2. 提供有意义的错误信息:")
	fmt.Println("   ✅ 使用 fmt.Errorf() 添加上下文")
	fmt.Println("   ✅ 错误信息应该帮助调试")
	fmt.Println()

	fmt.Println("3. 使用错误包装:")
	fmt.Println("   ✅ 使用 fmt.Errorf() 的 %w 动词")
	fmt.Println("   ✅ 使用 errors.Unwrap() 展开错误")
	fmt.Println()

	fmt.Println("4. 定义特定错误:")
	fmt.Println("   ✅ 使用 errors.New() 定义错误变量")
	fmt.Println("   ✅ 使用 errors.Is() 检查特定错误")
	fmt.Println()

	fmt.Println("5. 谨慎使用 panic:")
	fmt.Println("   ✅ panic 只用于不可恢复的错误")
	fmt.Println("   ✅ 使用 recover 在必要时恢复")
	fmt.Println("   ✅ 大多数情况下应该返回错误")
	fmt.Println()

	// ============================================
	// 8. 实际应用示例
	// ============================================
	fmt.Println("=== 8. 实际应用示例 ===")

	// 示例1：配置文件读取
	fmt.Println("示例1：配置文件读取")
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Printf("  加载配置失败: %v\n", err)
	} else {
		fmt.Printf("  配置加载成功: %+v\n", config)
	}
	fmt.Println()

	// 示例2：网络请求（模拟）
	fmt.Println("示例2：网络请求（模拟）")
	data, err := fetchData("https://api.example.com/data")
	if err != nil {
		fmt.Printf("  获取数据失败: %v\n", err)
	} else {
		fmt.Printf("  获取数据成功: %s\n", data)
	}
	fmt.Println()

	// 示例3：数据验证
	fmt.Println("示例3：数据验证")
	err = validateUser("", -1)
	if err != nil {
		fmt.Printf("  验证失败: %v\n", err)
	}
	err = validateUser("Alice", 25)
	if err != nil {
		fmt.Printf("  验证失败: %v\n", err)
	} else {
		fmt.Println("  验证成功")
	}
	fmt.Println()

	// ============================================
	// 9. 总结
	// ============================================
	fmt.Println("=== 9. 总结 ===")
	fmt.Println()
	fmt.Println("1. 错误处理的核心思想:")
	fmt.Println("   ✅ 像消防演习一样，提前规划应对步骤")
	fmt.Println("   ✅ 显式、主动的错误处理")
	fmt.Println("   ✅ 让代码更清晰、更可靠")
	fmt.Println()
	fmt.Println("2. Go错误处理的特点:")
	fmt.Println("   ✅ 显式错误返回: (result, error)")
	fmt.Println("   ✅ 无异常捕获机制: 没有 try/catch")
	fmt.Println("   ✅ 错误是值: 可以传递、存储和处理")
	fmt.Println()
	fmt.Println("3. 写入文件并处理错误:")
	fmt.Println("   ✅ 总是检查IO操作的错误")
	fmt.Println("   ✅ 提供有意义的错误信息")
	fmt.Println()
	fmt.Println("4. 创造性地处理错误:")
	fmt.Println("   ✅ 错误包装: 使用 fmt.Errorf() 的 %w")
	fmt.Println("   ✅ 错误链传递: 使用 errors.Unwrap()")
	fmt.Println()
	fmt.Println("5. 创建并标识特定错误:")
	fmt.Println("   ✅ errors.New() 定义错误变量")
	fmt.Println("   ✅ errors.Is() 检查特定错误")
	fmt.Println("   ✅ 自定义错误类型")
	fmt.Println()
	fmt.Println("6. 处理 panic:")
	fmt.Println("   ✅ panic 用于不可恢复的错误")
	fmt.Println("   ✅ recover 在必要时恢复")
	fmt.Println("   ✅ 大多数情况下应该返回错误")
	fmt.Println()
}

// ============================================
// 辅助函数和类型
// ============================================

// writeFile 写入文件
func writeFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

// readFile 读取文件
func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("读取文件 %s 失败: %w", filename, err)
	}
	return string(data), nil
}

// processData 处理数据（错误包装示例）
func processData(filename string) error {
	data, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("处理数据时出错: %w", err)
	}
	if len(data) == 0 {
		return fmt.Errorf("数据为空")
	}
	return nil
}

// complexOperation 复杂操作（错误链传递示例）
func complexOperation(input string) (string, error) {
	result, err := step1(input)
	if err != nil {
		return "", fmt.Errorf("步骤1失败: %w", err)
	}
	result, err = step2(result)
	if err != nil {
		return "", fmt.Errorf("步骤2失败: %w", err)
	}
	return result, nil
}

func step1(input string) (string, error) {
	if input == "" {
		return "", errors.New("输入不能为空")
	}
	return input + "_step1", nil
}

func step2(input string) (string, error) {
	if len(input) < 5 {
		return "", errors.New("输入长度不足")
	}
	return input + "_step2", nil
}

// 定义特定错误
var (
	ErrNotFound         = errors.New("资源未找到")
	ErrPermissionDenied = errors.New("权限被拒绝")
)

// checkResource 检查资源
func checkResource(name string) error {
	if name == "" {
		return ErrNotFound
	}
	if name == "forbidden" {
		return ErrPermissionDenied
	}
	return nil
}

// validateInput 验证输入
func validateInput(input string) error {
	if input == "" {
		return fmt.Errorf("输入验证失败: 输入不能为空")
	}
	if len(input) < 3 {
		return fmt.Errorf("输入验证失败: 输入长度至少为3，当前为%d", len(input))
	}
	return nil
}

// DivisionError 除法错误类型
type DivisionError struct {
	Dividend int
	Divisor  int
	Message  string
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("除法错误: %s (被除数: %d, 除数: %d)", e.Message, e.Dividend, e.Divisor)
}

// divideNumbers 除法运算
func divideNumbers(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, &DivisionError{
			Dividend: dividend,
			Divisor:  divisor,
			Message:  "除数不能为0",
		}
	}
	return dividend / divisor, nil
}

// safeFunction 安全函数（使用 recover）
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  捕获到 panic: %v\n", r)
		}
	}()
	fmt.Println("  执行安全函数")
	panic("这是一个测试 panic")
	fmt.Println("  这行不会执行")
}

// safeDivide 安全除法（使用 recover）
func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("    除法出错: %v\n", r)
		}
	}()
	if b == 0 {
		panic("除数不能为0")
	}
	return a / b
}

// Config 配置结构体
type Config struct {
	Host string
	Port int
}

// loadConfig 加载配置
func loadConfig(filename string) (*Config, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("配置文件 %s 不存在: %w", filename, err)
	}
	// 模拟加载配置
	return &Config{Host: "localhost", Port: 8080}, nil
}

// fetchData 获取数据（模拟网络请求）
func fetchData(url string) (string, error) {
	if url == "" {
		return "", errors.New("URL不能为空")
	}
	// 模拟网络错误
	if url == "https://api.example.com/error" {
		return "", fmt.Errorf("网络请求失败: 连接超时")
	}
	return "数据内容", nil
}

// validateUser 验证用户
func validateUser(name string, age int) error {
	if name == "" {
		return fmt.Errorf("用户名不能为空")
	}
	if age < 0 {
		return fmt.Errorf("年龄不能为负数: %d", age)
	}
	if age > 150 {
		return fmt.Errorf("年龄不合理: %d", age)
	}
	return nil
}

