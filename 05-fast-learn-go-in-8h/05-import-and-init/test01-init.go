package main

import (
	"fmt"
	"learning/05-fast-learn-go-in-8h/05-import-and-init/lib1"

	// . "learning/05-fast-learn-go-in-8h/05-import-and-init/lib2"
	_ "learning/05-fast-learn-go-in-8h/05-import-and-init/lib2"
)

/*

libxxx 	"xxx.github.com/xx/xxx" 导入并重命名
_ 		匿名导入 只导入包中的init方法
. 		包内的方法全部导入 不建议 有多个包方法名重复风险

*/

func main() {
	lib1.Lib1Test()
	// lib2.Lib2Test()
	fmt.Println("main func running ")
}
