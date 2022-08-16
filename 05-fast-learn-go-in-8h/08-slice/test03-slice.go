package main

import "fmt"

func main() {
	// 声明一个切片,并且初始化,默认值是1,2,3 长度是3
	// slice1 := []int{1, 2, 3}

	// 声明一个切片,但是没有分配空间
	// var slice1 []int
	// slice1 = make([]int, 3) //开辟3个空间
	// slice1[0] = 100

	// 声明一个切片
	// var slice1 []int = make([]int, 3)

	// 声明一个切片,自动推导是切片类型
	slice1 := make([]int, 3)

	fmt.Printf("len=%d, slice=%v\n", len(slice1), slice1)
	// 判断slice是否为空
	if slice1 == nil {
		fmt.Println("slice是一个空切片")
	} else {
		fmt.Println("slice有空间")
	}

}
