package main

//
//import (
//	"fmt"
//	"math"
//)
//
//func Sqrt(x float64) float64 {
//	z := 1.0
//	for i := 0; i < 10; i++ {
//		old := z
//		z -= (z*z - x) / (2 * z)
//		if z-old == 0 {
//			return z
//		}
//
//		fmt.Printf("[%v], x is %v, z is %v\n", i, x, z)
//
//	}
//	return z
//}
//
//func main() {
//	n := 99.
//	fmt.Println(Sqrt(n))
//	fmt.Println(math.Sqrt(n))
//}
