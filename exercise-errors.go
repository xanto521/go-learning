package main

//import "fmt"
//
//type ErrNegativeSqrt float64
//
//func (e ErrNegativeSqrt) Error() string {
//	if e < 0 {
//		return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
//	}
//	return ""
//}
//
//func Sqrt(x float64) (float64, error) {
//	i := ErrNegativeSqrt(x)
//	if i.Error() == "" {
//		z := 1.0
//		for i := 0; i < 10; i++ {
//			old := z
//			z -= (z*z - x) / (2 * z)
//			if z-old == 0 {
//				return z, nil
//			}
//			fmt.Printf("[%v], x is %v, z is %v\n", i, x, z)
//		}
//		return z, nil
//
//	}
//	return 0, i
//}
//
//func main() {
//	fmt.Println(Sqrt(2))
//	fmt.Println(Sqrt(-2))
//}
