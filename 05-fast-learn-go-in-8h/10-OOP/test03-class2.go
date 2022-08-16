package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("human is eating...")
}

func (h *Human) Walk() {
	fmt.Println("human is walking...")
}

func (h *Human) GetName() {
	fmt.Println(h.name)
}

type SuperMan struct {
	Human // 继承human的属性
	level int
	//name  string // 当继承时候有相同属性的时候 无法处理 编译器认为是两个属性  比如这里就是Human.name 和 SuperMan.name两个属性
}

func (sm *SuperMan) Eat() {
	fmt.Println("superman is eating...")
}

func (sm *SuperMan) Fly() {
	fmt.Println("superman is flying...")
}

func main() {

	h := Human{"zhang3", "female"}

	h.Eat()
	h.Walk()

	//s := SuperMan{
	//	Human: h,
	//	level: 0,
	//	name:  "superman",
	//}
	var s SuperMan
	s.name = "superman" // 此处是superman.name 所以下面打印get name方法的时候无法获取 那里获取的是human.name的属性值
	s.sex = "female"
	s.level = 999

	s.Eat()
	s.Walk()
	s.Fly()
	s.GetName()

}
