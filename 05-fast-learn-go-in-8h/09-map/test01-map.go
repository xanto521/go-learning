package main

import "fmt"

func main() {
	fmt.Println("main runing...")
	// 第一种使用方式
	var myMap map[string]string
	// 判断map是否为空

	if myMap == nil {
		fmt.Println("myMap 是一个空map")
	}

	//在使用map前要先分配
	myMap = make(map[string]string, 10)
	myMap["one"] = "java"
	myMap["two"] = "c++"
	myMap["three"] = "python"
	fmt.Println(myMap)

	// 第二种
	myMap2 := make(map[int]string, 3)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"
	fmt.Println(myMap2)

	// 第三种
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println(myMap3)

}
