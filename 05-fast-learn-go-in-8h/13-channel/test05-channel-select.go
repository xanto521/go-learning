package main

import "fmt"

func fibonacci(c, q chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// 一行实现交换
			//x, y = y, x+y

			//多行实现交换,需要间接参数
			tmp := x
			x = y
			y = tmp + y
		case <-q:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
