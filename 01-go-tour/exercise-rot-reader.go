package main

//import (
//	"io"
//	"os"
//	"strings"
//)
//
//type rot13router struct {
//	r io.Reader
//}
//
//func rot13(b byte) byte {
//	if b >= 'A' && b <= 'Z' {
//		b = 'A' + (b-'A'+13)%26
//	} else if b >= 'a' && b <= 'z' {
//		b = 'a' + (b-'a'+13)%26
//	}
//	return b
//}
//
//func (r rot13router) Read(b []byte) (n int, err error) {
//
//	n, err = r.r.Read(b)
//	for i := 0; i < n; i++ {
//		//fmt.Printf("old: %v\n", b[i])
//		b[i] = rot13(b[i])
//		//fmt.Printf("new: %v\n", b[i])
//	}
//	//fmt.Printf("A")
//
//	return n, err
//}
//
//func main() {
//	s := strings.NewReader("Lbh penpxrq gur pbqr!")
//	r := rot13router{s}
//	io.Copy(os.Stdout, &r)
//}
