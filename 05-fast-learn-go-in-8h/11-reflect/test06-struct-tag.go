package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"姓名"`
	Sex  string `info:"sex"`
}

func findTag(input interface{}) {
	// 传入的是对象就可以用上一节的方法，传入的是指针就可以用elem取出指针的对象结构体
	t := reflect.TypeOf(input).Elem()
 
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("taginfo:", taginfo, "tagdoc:", tagdoc)
	}
}

func main() {
	var r resume
	findTag(&r)
}
