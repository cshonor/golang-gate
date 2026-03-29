// 独立运行：go run ./chap18/slice_expand_verify
// 演示：扩容前后底层数组地址变化；共享切片在扩容后与兄弟切片分离。
package main

import (
	"fmt"
	"unsafe"
)

// backingAddr 返回切片底层数组第一个元素的地址；空切片返回 0。
func backingAddr(s []int) uintptr {
	if len(s) == 0 {
		return 0
	}
	return uintptr(unsafe.Pointer(&s[0]))
}

func main() {
	fmt.Println("=== 1. append 触发扩容：指针、len、cap ===")
	s := []int{1, 2, 3}
	fmt.Printf("初始: data=%#x len=%d cap=%d val=%v\n", backingAddr(s), len(s), cap(s), s)
	old := backingAddr(s)
	s = append(s, 4)
	fmt.Printf("追加 4: data=%#x len=%d cap=%d val=%v (首元素地址变化=%v)\n",
		backingAddr(s), len(s), cap(s), s, backingAddr(s) != old)

	fmt.Println("\n=== 2. 共享底层数组：未扩容 vs 扩容后分离 ===")
	s1 := []int{1, 2, 3}
	s2 := s1[:2]
	fmt.Printf("s1=%v s2=%v | addr(s1)=%#x addr(s2)=%#x (共享=%v)\n",
		s1, s2, backingAddr(s1), backingAddr(s2), backingAddr(s1) == backingAddr(s2))
	s2[0] = 100
	fmt.Printf("s2[0]=100 后 s1=%v s2=%v (未扩容，改一处两边都可见)\n", s1, s2)

	s1 = []int{1, 2, 3}
	s2 = s1[:2]
	// 未接收返回值，但若 cap 仍够，append 可能已在共享数组 len 位置写入 —— s1 也会被改
	_ = append(s2, 999)
	fmt.Printf("\n易错: _=append(s2,999) 后 s1=%v s2=%v（s1[2] 可能被写成 999）\n", s1, s2)

	s2 = append(s2, 4, 5, 6) // len 将超 cap，触发扩容
	fmt.Printf("s2=append(s2,4,5,6) 后: s1=%v s2=%v\n", s1, s2)
	fmt.Printf("addr(s1)=%#x addr(s2)=%#x (已分离=%v)\n",
		backingAddr(s1), backingAddr(s2), backingAddr(s1) != backingAddr(s2))
	s2[0] = 200
	fmt.Printf("s2[0]=200 后: s1=%v s2=%v (只改 s2，s1 不变)\n", s1, s2)
}
