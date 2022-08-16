package main

import "fmt"

/*
			  变量
			/     \
        +----------------+
		| type     value |   ==> pair
        +----------------+
         /	  \
static type   concrete type

*/
func main() {
	var a string
	//pair<type:string,value:"abc">
	a = "abc"

	var allType interface{}
	// pair<type:interface{},value:>
	allType = a
	// pair<type:string,value:"abc">
	str, _ := allType.(string)
	fmt.Println(str)

}
