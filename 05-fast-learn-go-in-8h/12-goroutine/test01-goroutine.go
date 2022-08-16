package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine : i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go newTask()
	fmt.Println("main goroutine exit...")
	i := 0
	for {
		i++
		fmt.Printf("main goroutine : i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
}
