package main

import "fmt"

// 声明自己的数据类型
type myint int

// 声明一个新的数据结构体
type Book struct {
	title  string
	auther string
}

func changeBook(book Book) {
	// 值传递 只修改方法体内的值 源数据对象不变
	book.title = "666"
}

func changeBook2(book *Book) {
	// 指针传递 修改后源对象的数据随之改变
	book.title = "123"
}

func main() {
	fmt.Println("main...")
	var book1 Book
	book1.title = "Golang"
	book1.auther = "zhangsan"
	fmt.Println(book1)

	changeBook(book1)
	fmt.Println(book1)

	changeBook2(&book1)
	fmt.Println(book1)
}
