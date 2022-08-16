package main

import "fmt"

func main() {
	// 直接用var声明变量,自带默认值
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a: %T\n", a)

	// 初始化一个变量并设置默认值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b: %T\n", b)

	var bb string = "abcd"
	fmt.Println("bb = ", bb)
	fmt.Printf("type of bb: %T\n", bb)

	// 用var定义变量的时候直接省略变量类型, 自动推断变量类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c: %T\n", c)

	var cc = "qwer"
	fmt.Println("cc = ", cc)
	fmt.Printf("type of cc: %T\n", cc)

	// 省去var关键字(常用)  自动匹配变量类型
	// := 不能用作全局变量的声明
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("type of d: %T\n", d)

	dd := 1.23
	fmt.Println("dd = ", dd)
	fmt.Printf("type of dd: %T\n", dd)

	ddd := "111"
	fmt.Println("ddd = ", ddd)
	fmt.Printf("type of ddd: %T\n", ddd)

	// 声明多个变量
	var aaaa, bbbb string = "aaaa", "bbbb"
	fmt.Println("aaaa = ", aaaa, "bbbb = ", bbbb)

	var jj, kk = 1, "2"
	fmt.Println("jj = ", jj, "kk = ", kk)

	// 多行变量的声明 放在括号内
	var (
		xx        = 111
		yy string = "222"
		zz bool   = true
	)
	fmt.Println("xx = ", xx, "yy = ", yy, "zz = ", zz)

}
