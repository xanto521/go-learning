package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {

}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	for {
		select {}
	}
	Walk(t1, ch1)
	Walk(t2, ch2)
	return false
}

func HelpFunc() {

}
func main() {
	a := tree.New(1)
	b := tree.New(1)

	fmt.Println(a.Right == b.Right)
	fmt.Println(a.Right)
	fmt.Println(b.Right)
	fmt.Println(a == b)

}
