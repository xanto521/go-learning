package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}
type Book struct{}

func (this *Book) ReadBook() {
	fmt.Println("Read book")
}
func (this *Book) WriteBook() {
	fmt.Println("Write book")
}

func main() {
	// b:pair<type:book,value:Book{}地址>
	b := &Book{}
	// r:pari<type:book,value:>
	var r Reader
	// r:pair<type:book,value:Book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	// w:pair<type:book,value:Book{}地址>
	w = r.(Writer) // 断言成功，因为类型都为book,book对象实现了Writer的WriteBook方法
	w.WriteBook()

}
