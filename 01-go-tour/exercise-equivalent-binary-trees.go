package main

//import (
//	"fmt"
//	"golang.org/x/tour/tree"
//)
//
//func Walk(t *tree.Tree, ch chan int) {
//	if t == nil {
//		return
//	}
//	// 先遍历左子树，取当前节点的值存入通道，再遍历右子树。这样的结果就是从小到大的顺序。
//	Walk(t.Left, ch)
//	ch <- t.Value
//	Walk(t.Right, ch)
//}
//
//func Same(t1, t2 *tree.Tree) bool {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go Walk(t1, ch1)
//	go Walk(t2, ch2)
//	for i := 0; i < 10; i++ {
//		if <-ch1 != <-ch2 {
//			return false
//		}
//	}
//	return true
//}
//
//func main() {
//	a := tree.New(1)
//	fmt.Println(a.String())
//
//	b := tree.New(1)
//	fmt.Println(b)
//
//	// 带&的是指针，不带&的是值。 一个变量:=指针后，变量对象会变成**类型，即指针指针。
//	//c := &a
//	c := a
//	fmt.Println(c)
//
//	//fmt.Println(unsafe.Pointer(a))
//	//fmt.Println(unsafe.Pointer(&a))
//	//
//	//fmt.Println(unsafe.Pointer(c))
//	//fmt.Println(unsafe.Pointer(&c))
//
//	fmt.Println(Same(a, b))
//	fmt.Println(Same(a, c))
//
//}
