// 示例：Go 语言的字符操作和凯撒加密
// 演示 rune/byte 类型、字符字面量、字符串索引访问和凯撒加密法

package main

import "fmt"

func main() {
	// ============================================
	// 1. rune 和 byte 类型
	// ============================================
	fmt.Println("=== rune 和 byte 类型 ===")

	// rune：Go 中 int32 的别名，用于表示 Unicode 码点
	// 支持全球所有字符（包括中文、希腊字母等）
	var pi rune = 960    // 希腊字母 π 的 Unicode 码点
	var alpha rune = 945 // 希腊字母 α 的 Unicode 码点
	var omega rune = 969 // 希腊字母 ω 的 Unicode 码点

	// byte：Go 中 uint8 的别名，用于表示 ASCII 字符（0-255）
	var bang byte = 33 // 感叹号 ! 的 ASCII 码

	// 打印数值本身
	fmt.Printf("数值: %v %v %v %v\n", pi, alpha, omega, bang)
	// 打印对应的字符
	fmt.Printf("字符: %c%c%c%c\n", pi, alpha, omega, bang)

	fmt.Println()

	// ============================================
	// 2. 字符字面量的三种等价写法
	// ============================================
	fmt.Println("=== 字符字面量的三种写法 ===")

	// 方式1：短变量声明，自动推断为 rune
	grade := 'A'
	fmt.Printf("grade 类型: %T, 值: %v, 字符: %c\n", grade, grade, grade)

	// 方式2：var 声明，自动推断为 rune
	var grade2 = 'A'
	fmt.Printf("grade2 类型: %T, 值: %v, 字符: %c\n", grade2, grade2, grade2)

	// 方式3：显式声明为 rune
	var grade3 rune = 'A'
	fmt.Printf("grade3 类型: %T, 值: %v, 字符: %c\n", grade3, grade3, grade3)

	// byte 类型的字符字面量
	var star byte = '*'
	fmt.Printf("star 类型: %T, 值: %v, 字符: %c\n", star, star, star)

	fmt.Println()

	// ============================================
	// 3. 字符串索引访问
	// ============================================
	fmt.Println("=== 字符串索引访问 ===")

	message := "shalom"
	fmt.Printf("字符串: %s\n", message)
	fmt.Printf("字符串长度（字节数）: %d\n", len(message))

	// 通过索引访问字符串中的字符
	fmt.Println("通过索引访问:")
	for i := 0; i < len(message); i++ {
		fmt.Printf("  message[%d] = %c (类型: %T, ASCII码: %d)\n", i, message[i], message[i], message[i])
	}

	// 注意：对于包含非 ASCII 字符的字符串，索引访问得到的是单个字节
	message2 := "shalom世界"
	fmt.Printf("\n字符串（包含中文）: %s\n", message2)
	fmt.Printf("字符串长度（字节数）: %d\n", len(message2))
	fmt.Printf("message2[6] = %c (这是'世'的第一个字节，不是完整字符)\n", message2[6])

	// 正确的方式：使用 for range 遍历（按 rune 遍历）
	fmt.Println("使用 for range 遍历（按字符）:")
	for i, char := range message2 {
		fmt.Printf("  位置 %d: %c (Unicode码点: U+%04X)\n", i, char, char)
	}

	fmt.Println()

	// ============================================
	// 4. 打印字符串每个字符（每行一个）
	// ============================================
	fmt.Println("=== 打印字符串每个字符（每行一个）===")

	// 方式1：通过索引访问（适合 ASCII 字符串）
	fmt.Println("方式1 - 通过索引访问:")
	for i := 0; i < len(message); i++ {
		fmt.Printf("%c\n", message[i])
	}

	// 方式2：通过 for range 遍历（推荐，支持 Unicode 字符）
	fmt.Println("方式2 - for range 遍历:")
	for _, char := range message {
		fmt.Printf("%c\n", char)
	}

	fmt.Println()

	// ============================================
	// 5. 大小写转换
	// ============================================
	fmt.Println("=== 大小写转换 ===")

	// 小写转大写：c - 'a' + 'A'
	c := 'g'
	upper := c - 'a' + 'A'
	fmt.Printf("'%c' 转大写: %c\n", c, upper)

	// 计算过程说明
	fmt.Printf("计算过程: %d - %d + %d = %d (对应字符 '%c')\n",
		c, 'a', 'A', upper, upper)

	// 大写转小写：c - 'A' + 'a'
	C := 'G'
	lower := C - 'A' + 'a'
	fmt.Printf("'%c' 转小写: %c\n", C, lower)

	// 批量转换示例
	fmt.Println("批量转换:")
	lowercase := "hello"
	uppercase := ""
	for _, char := range lowercase {
		if char >= 'a' && char <= 'z' {
			uppercase += string(char - 'a' + 'A')
		} else {
			uppercase += string(char)
		}
	}
	fmt.Printf("%s -> %s\n", lowercase, uppercase)

	fmt.Println()

	// ============================================
	// 6. 字符串的不可变性
	// ============================================
	fmt.Println("=== 字符串的不可变性 ===")

	msg := "hello"
	fmt.Printf("原始字符串: %s\n", msg)

	// 错误：不能直接修改字符串中的字符
	// msg[0] = 'H' // 编译错误：cannot assign to msg[0]

	// 正确：需要修改时，先转成切片，修改后再转回字符串
	// 方式1：转成 []byte（适合 ASCII 字符）
	bytes := []byte(msg)
	bytes[0] = 'H'
	newMsg1 := string(bytes)
	fmt.Printf("修改后（[]byte）: %s\n", newMsg1)
	fmt.Printf("原始字符串未变: %s\n", msg)

	// 方式2：转成 []rune（适合 Unicode 字符）
	runes := []rune(msg)
	runes[0] = 'H'
	newMsg2 := string(runes)
	fmt.Printf("修改后（[]rune）: %s\n", newMsg2)

	fmt.Println()

	// ============================================
	// 7. 凯撒加密法实现
	// ============================================
	fmt.Println("=== 凯撒加密法 ===")

	// 凯撒加密：将字母表中的每个字母移动固定位数
	plaintext := "hello world"
	shift := 3 // 位移 3 位

	fmt.Printf("原文: %s\n", plaintext)
	fmt.Printf("位移: %d 位\n", shift)

	// 加密
	encrypted := caesarEncrypt(plaintext, shift)
	fmt.Printf("密文: %s\n", encrypted)

	// 解密（反向位移）
	decrypted := caesarDecrypt(encrypted, shift)
	fmt.Printf("解密: %s\n", decrypted)

	fmt.Println()

	// ============================================
	// 8. 凯撒加密法（处理边界回绕）
	// ============================================
	fmt.Println("=== 凯撒加密法（处理边界回绕）===")

	plaintext2 := "xyz ABC"
	shift2 := 3

	fmt.Printf("原文: %s\n", plaintext2)
	fmt.Printf("位移: %d 位\n", shift2)

	encrypted2 := caesarEncryptWithWrap(plaintext2, shift2)
	fmt.Printf("密文: %s\n", encrypted2)

	decrypted2 := caesarDecryptWithWrap(encrypted2, shift2)
	fmt.Printf("解密: %s\n", decrypted2)

	fmt.Println()

	// ============================================
	// 9. 格式化符 %v 和 %c
	// ============================================
	fmt.Println("=== 格式化符 %v 和 %c ===")

	char := 'A'
	fmt.Printf("%%v 打印数值: %v\n", char)   // 打印原始值（数字）
	fmt.Printf("%%c 打印字符: %c\n", char)   // 打印字符
	fmt.Printf("%%d 打印十进制: %d\n", char) // 打印十进制数值
	fmt.Printf("%%x 打印十六进制: %x\n", char) // 打印十六进制数值
	fmt.Printf("%%U 打印Unicode: %U\n", char) // 打印 Unicode 格式

	fmt.Println()

	// ============================================
	// 10. 总结
	// ============================================
	fmt.Println("=== 总结 ===")
	fmt.Println("1. rune 是 int32 的别名，用于表示 Unicode 码点")
	fmt.Println("2. byte 是 uint8 的别名，用于表示 ASCII 字符")
	fmt.Println("3. 字符字面量用单引号 ' ' 包裹，自动推断为 rune")
	fmt.Println("4. 字符串是不可变的，不能直接修改单个字符")
	fmt.Println("5. 字符串索引访问得到的是 byte，不是完整字符")
	fmt.Println("6. for range 遍历字符串按 rune（字符）遍历")
	fmt.Println("7. %v 打印数值，%c 打印字符")
	fmt.Println("8. 小写转大写：c - 'a' + 'A'，大写转小写：c - 'A' + 'a'")
	fmt.Println("9. 凯撒加密通过字母位移实现，需要处理边界回绕")
}

// caesarEncrypt 凯撒加密（简单版本，不处理边界）
func caesarEncrypt(text string, shift int) string {
	result := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			// 小写字母：位移后转回字符
			shifted := char + rune(shift)
			if shifted > 'z' {
				shifted -= 26 // 回绕
			}
			result += string(shifted)
		} else if char >= 'A' && char <= 'Z' {
			// 大写字母：位移后转回字符
			shifted := char + rune(shift)
			if shifted > 'Z' {
				shifted -= 26 // 回绕
			}
			result += string(shifted)
		} else {
			// 非字母字符保持不变
			result += string(char)
		}
	}
	return result
}

// caesarDecrypt 凯撒解密（简单版本）
func caesarDecrypt(text string, shift int) string {
	// 解密就是反向位移
	return caesarEncrypt(text, -shift)
}

// caesarEncryptWithWrap 凯撒加密（完整版本，处理边界回绕）
func caesarEncryptWithWrap(text string, shift int) string {
	result := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			// 小写字母：处理回绕
			shifted := ((char - 'a' + rune(shift)) % 26) + 'a'
			result += string(shifted)
		} else if char >= 'A' && char <= 'Z' {
			// 大写字母：处理回绕
			shifted := ((char - 'A' + rune(shift)) % 26) + 'A'
			result += string(shifted)
		} else {
			// 非字母字符保持不变
			result += string(char)
		}
	}
	return result
}

// caesarDecryptWithWrap 凯撒解密（完整版本）
func caesarDecryptWithWrap(text string, shift int) string {
	// 解密：反向位移，处理负数情况
	result := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			// 处理负数回绕
			shifted := ((char - 'a' - rune(shift) + 26) % 26) + 'a'
			result += string(shifted)
		} else if char >= 'A' && char <= 'Z' {
			// 处理负数回绕
			shifted := ((char - 'A' - rune(shift) + 26) % 26) + 'A'
			result += string(shifted)
		} else {
			// 非字母字符保持不变
			result += string(char)
		}
	}
	return result
}

