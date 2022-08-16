package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的类别
}

// 具体的类
type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("cat is sleeping")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (c *Dog) Sleep() {
	fmt.Println("dog is sleeping")
}

func (c *Dog) GetColor() string {
	return c.color
}

func (c *Dog) GetType() string {
	return "Dog"
}

func animalSleep(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("animal's color is", animal.GetColor())
	fmt.Println("animal's type is", animal.GetType())
}

func main() {

	/*
		var animal AnimalIF
		animal = &Cat{color: "black"}
		animal.Sleep()

		animal = &Dog{color: "black"}
		animal.Sleep()
	*/
	cat := Cat{color: "black"}
	animalSleep(&cat)
	dog := Dog{color: "black"}
	animalSleep(&dog)
}
