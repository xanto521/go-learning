package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 无形参goroutine
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	//有参数goroutine
	go func(a, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(1, 2)

	for {
		time.Sleep(1 * time.Second)
	}
}
