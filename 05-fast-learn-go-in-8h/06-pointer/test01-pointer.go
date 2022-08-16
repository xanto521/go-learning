package main

import "fmt"

/*
func swap(a, b int) {
	var tmp int
	tmp = a
	a = b
	b = tmp
}
*/

func swap(pa *int, pb *int) {
	var tmp int
	tmp = *pa // tmp = main::a
	*pa = *pb // main::a = main::b
	*pb = tmp // main::b = tmp

}

// 用指针可以修改函数体外部数据
// * 是取指针地址中变量数据的值， &是取变量的指针

func main() {
	var a int = 10
	var b int = 20

	// swap交换
	// swap(a, b)
	swap(&a, &b)

	fmt.Println("a = ", a, ", b = ", b)

	var p *int // p为一级指针
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int // p为二级指针
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
	fmt.Println(**pp)

}
