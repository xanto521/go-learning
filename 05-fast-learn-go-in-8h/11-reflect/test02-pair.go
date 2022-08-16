package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// tty: pair<type:*os.File, value:"/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// 此时r并未赋值 r: pair<type: ,value:>
	var r io.Reader
	// 赋值完毕,此时 r:  pair<type:*os.File, value:"/dev/tty"文件描述符>
	r = tty

	// 此时r并未赋值 w: pair<type: ,value:>
	var w io.Writer
	// 赋值完毕,此时 w:  pair<type:*os.File, value:"/dev/tty"文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("hello golang.\n"))
}
