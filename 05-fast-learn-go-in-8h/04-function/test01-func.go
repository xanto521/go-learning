package main

import "fmt"

// 无参数
func f1() {
	fmt.Println("f1 running..")
}

// 有形参 无返回值
func f2(a string, b int) {
	fmt.Println("f2 running...")
	fmt.Println("a:", a, "b:", b)
}

// 有形参 有返回值 匿名返回形参
func f3(a string, b int) bool {
	fmt.Println("f3 running...")
	fmt.Println("a:", a, "b:", b)
	return true
}

// 多个返回值 匿名返回形参 return后面必须和形参严格一一对应
func f4(a string, b int) (bool, int) {
	fmt.Println("f4 running...")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return true, 0

}

// 返回值不为匿名参数
func f5(a string, b int) (res1 bool, res2 int) {
	fmt.Println("f5 running...")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// fmt.Println("res1 = ", res1)
	// fmt.Println("res2 = ", res2)

	res1 = true
	res2 = 333

	// return res1, res2
	// 函数体return后面不写返回的形参也可以
	return

}
func main() {

	f1()

	f2("f2 f2", 222)

	fmt.Println(f3("111", 222))

	fmt.Println(f4("abc", 123))

	fmt.Println(f5("abc", 123))

}
