package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc called...")
	fmt.Println(arg)

	// interface 区分数据类型
	// .(type) 如 .(string)断言是否为字符串类型

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not a string.")
	} else {
		fmt.Printf("arg type is %T. arg value is %v\n", value, value)
	}

}

type Car struct {
	color string
}

func main() {
	car := Car{"Green"}
	myFunc(car)
	myFunc("abc")
	myFunc(123)
	myFunc(1.23)

}
