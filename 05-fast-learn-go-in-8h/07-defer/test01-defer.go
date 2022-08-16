package main

import "fmt"

func main() {
	// defer 可以理解为栈 先入后出 后入先出
	defer fmt.Println("main end 1")
	defer fmt.Println("main end 2")
	defer fmt.Println("main end 3")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 1")
}
