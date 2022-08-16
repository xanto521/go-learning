package main

import "fmt"

func main() {
	// 3个长度有数据，5个容量
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 3)
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 4)
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(numbers), cap(numbers), numbers)

	// 自动扩容 当数据长度大于容量时 进行扩充 扩充的空间为之前的2倍
	numbers = append(numbers, 5)
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(numbers), cap(numbers), numbers)

	// 继续填充则继续扩容 为之前的2倍
	numbers = append(numbers, 6, 7, 8, 9, 10)
	fmt.Printf("len=%d, cap=%d, sclice=%v\n", len(numbers), cap(numbers), numbers)

}
