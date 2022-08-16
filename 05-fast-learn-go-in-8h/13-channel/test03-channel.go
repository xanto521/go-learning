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
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main finished.")
}
