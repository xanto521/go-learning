package main

import "fmt"

const (
	BEIJING = iota
	SHANGHAI
	GUANGZHOU
	SHENZHEN
)

func main() {
	// 常量 只读 不可修改
	const length int = 10
	fmt.Println("length: ", length)

	// 不可修改
	// length = 2
	// fmt.Println("length: ", length)
	fmt.Println("BEIJING:", BEIJING)
	fmt.Println("SHANGHAI:", SHANGHAI)
	fmt.Println("GUANGZHOU:", GUANGZHOU)
	fmt.Println("SHENZHEN:", SHENZHEN)

	// iota 只能在const中使用
	// var i = iota
}
