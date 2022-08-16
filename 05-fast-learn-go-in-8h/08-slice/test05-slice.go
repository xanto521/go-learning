package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(s), cap(s), s)

	s1 := s[0:2] // 左闭右开
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(s1), cap(s1), s1)

	// 截取前几个数据
	s1 = s[:2] // 左闭右开
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(s1), cap(s1), s1)

	// 截取最后几个数据
	s1 = s[len(s)-2:] // 左闭右开
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(s1), cap(s1), s1)

	// 切片只是数据引用 修改数据会影响原来的数据
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s)

	// copy 深拷贝 将原始数据地址的数据复制一份到新的数据地址
	s2 := make([]int, 3)
	// 将s中的值拷贝到s2中
	copy(s2, s)
	s2[0] = 222

	fmt.Println(&s)
	fmt.Println(&s2)

}
