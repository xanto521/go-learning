package main

import (
	"fmt"
	"time"
)

func main() {
	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine exit")
		fmt.Println("goroutine start")
		time.Sleep(1 * time.Second)
		c <- 1 //将1写入channel
		fmt.Println("goroutine end")

	}()
	// channel双向等待,如果没有写入,那么就等待,如果没有读取，也同样等待
	time.Sleep(1 * time.Second)
	num := <-c //从channel中读取数据
	fmt.Println("num = ", num)
	fmt.Println("main exit")
}
