package main

//import "fmt"
//
//func fibonacci() func() int {
//	f1 := 0
//	f2 := 1
//	n := 0
//
//	return func() int {
//		n += 1
//		if n == 1 {
//			return 0
//		}
//		f2, f1 = f1+f2, f2
//		return f1
//	}
//
//}
//
//func main() {
//	f := fibonacci()
//	for i := 0; i < 10; i++ {
//		fmt.Println(f())
//	}
//}
