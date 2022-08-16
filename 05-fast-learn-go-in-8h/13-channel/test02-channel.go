package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5) //带有缓冲的channel
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))
	go func() {
		defer fmt.Println("sub goroutine exit")
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("sub goroutine is running, value=", i, ", len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		num := <-c
		fmt.Println("num = ", num, ", len(c) = ", len(c), ", cap(c) = ", cap(c))
	}
	fmt.Println("main goroutine exit")
}
