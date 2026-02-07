package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigNum := new(big.Int)

	// 十进制
	fmt.Println("=== 十进制 ===")
	bigNum.SetString("86400", 10)
	fmt.Printf("SetString(\"86400\", 10) = %s\n", bigNum.String())
	fmt.Println("解释：字符串 \"86400\" 按十进制解析，就是 86400")
	fmt.Println()

	// 二进制
	fmt.Println("=== 二进制 ===")
	bigNum.SetString("1010", 2)
	fmt.Printf("SetString(\"1010\", 2) = %s\n", bigNum.String())
	fmt.Println("解释：字符串 \"1010\" 按二进制解析")
	fmt.Println("  1×2³ + 0×2² + 1×2¹ + 0×2⁰")
	fmt.Println("  = 1×8 + 0×4 + 1×2 + 0×1")
	fmt.Println("  = 8 + 0 + 2 + 0")
	fmt.Println("  = 10（十进制）")
	fmt.Println()

	// 十六进制
	fmt.Println("=== 十六进制 ===")
	bigNum.SetString("FF", 16)
	fmt.Printf("SetString(\"FF\", 16) = %s\n", bigNum.String())
	fmt.Println("解释：字符串 \"FF\" 按十六进制解析")
	fmt.Println("  F = 15, F = 15")
	fmt.Println("  15×16¹ + 15×16⁰")
	fmt.Println("  = 15×16 + 15×1")
	fmt.Println("  = 240 + 15")
	fmt.Println("  = 255（十进制）")
	fmt.Println()

	// 对比：同样的字符串，不同进制得到不同结果
	fmt.Println("=== 对比：同样的字符串，不同进制 ===")
	
	// "1010" 在不同进制下的结果
	fmt.Println("\n字符串 \"1010\" 在不同进制下：")
	bigNum.SetString("1010", 2)
	fmt.Printf("  二进制(2): %s\n", bigNum.String())
	bigNum.SetString("1010", 10)
	fmt.Printf("  十进制(10): %s\n", bigNum.String())
	bigNum.SetString("1010", 16)
	fmt.Printf("  十六进制(16): %s\n", bigNum.String())
	
	fmt.Println("\n为什么结果不同？")
	fmt.Println("  因为进制决定了每个位置的权重！")
	fmt.Println("  二进制：每个位置是 2 的幂次（2⁰, 2¹, 2², 2³...）")
	fmt.Println("  十进制：每个位置是 10 的幂次（10⁰, 10¹, 10², 10³...）")
	fmt.Println("  十六进制：每个位置是 16 的幂次（16⁰, 16¹, 16², 16³...）")
}

