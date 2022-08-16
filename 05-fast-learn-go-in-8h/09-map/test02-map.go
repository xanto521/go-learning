package main

import "fmt"

func printMap(m map[string]string) {
	// 引用传递,修改会影响原始数据
	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}

}

func changeMap(m map[string]string) {
	m["England"] = "英国"

}
func main() {
	cityMap := make(map[string]string)
	// 添加
	cityMap["China"] = "中国"
	cityMap["Japan"] = "日本"
	cityMap["USA"] = "美国"
	fmt.Println(cityMap)

	// 遍历
	for k, v := range cityMap {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}

	// 删除
	delete(cityMap, "USA")
	fmt.Println(cityMap)

	// 修改
	cityMap["Japan"] = "日本日本"
	fmt.Println(cityMap)
	changeMap(cityMap)

	printMap(cityMap)

}
