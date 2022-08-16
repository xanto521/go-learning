package main

import "fmt"

func pirntArray(arr []int) {
	// 引用传递 不定长数组

	for idx, value := range arr {
		fmt.Printf("idx=%d, value=%v\n", idx, value)
	}

	// 修改后这个数组的原始数据就会更改
	arr[0] = 100

}
func main() {
	myArr1 := []int{1, 2, 3, 4} // 动态数组 切片
	fmt.Printf("type of myArr1: %T\n", myArr1)
	pirntArray(myArr1)

	for idx, value := range myArr1 {
		fmt.Printf("idx=%d, value=%v\n", idx, value)
	}

}
