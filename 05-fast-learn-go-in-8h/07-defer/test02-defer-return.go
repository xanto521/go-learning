package main

import "fmt"

func DeferFunc() int {
	fmt.Println("defer func runing...")
	return 0
}

func ReturnFunc() int {
	fmt.Println("return func runing...")
	return 0
}

func DoFuc() int {
	defer DeferFunc()
	return ReturnFunc()
}

/*
defer 和 return的顺序
*/
func main() {
	DoFuc()
}
