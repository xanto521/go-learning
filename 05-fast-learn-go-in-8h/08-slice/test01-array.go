package main

import "fmt"

func printArray(arr [4]int) {
	// 值拷贝 只是将参数的值从外部传进来
	for idx, value := range arr {
		fmt.Printf("idx=%d, value=%v\n", idx, value)
	}

	arr[0] = 111

}
func main() {
	fmt.Println("main running...")
	// 固定长度的数组
	var myArr1 [10]int

	myArr2 := [10]int{1, 2, 3, 4}
	myArr3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArr1); i++ {
		fmt.Println(myArr1[i])
	}

	for index, value := range myArr2 {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}

	// 查看数组类型
	fmt.Printf("type of myArr1: %T\n", myArr1)
	fmt.Printf("type of myArr2: %T\n", myArr2)
	fmt.Printf("type of myArr3: %T\n", myArr3)

	printArray(myArr3)
	// 不能传入长度不同的数组
	// printArray(myArr2)

	// ---
	fmt.Println("-----")
	for index, value := range myArr3 {
		fmt.Printf("index=%d,value=%d\n", index, value)
	}

}
