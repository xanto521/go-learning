package main

//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	tick := time.Tick(2 * time.Second)
//	boom := time.After(5 * time.Second)
//	for {
//		select {
//		case <-tick:
//			fmt.Println("tick.")
//		case <-boom:
//			fmt.Println("BOOM!")
//		default:
//			fmt.Println("    .")
//			time.Sleep(1 * time.Second)
//		}
//	}
//}
