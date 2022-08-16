package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}
		// 关闭channel
		close(c)
	}()
	// range从channel一直读取数据，直到channel被关闭
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("main finished.")
}
