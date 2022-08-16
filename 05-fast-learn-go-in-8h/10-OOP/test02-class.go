package main

import "fmt"

// 类属性首字母大写标识公开属性 小写为私有属性
type Hero struct {
	Name  string
	Age   int
	Level int
}

func (h *Hero) Show() {
	fmt.Printf("<%T> Name=%v,Age=%v,Level=%v\n", h, h.Name, h.Age, h.Level)
}
func (h *Hero) GetName() {
	fmt.Println("this hero's name is", h.Name)
}

func (h *Hero) SetName(newname string) {
	h.Name = newname
}

func main() {

	hero := Hero{Name: "zhang3", Age: 18, Level: 1}
	hero.Show()

	hero.GetName()

	hero.SetName("wang5")
	hero.Show()

}
