package main

import (
	"fmt"
)

type Animal interface {
	sound()
	getAge()
}

type Dog struct {
	age int
}

func (d *Dog) sound() {
	fmt.Println("왈왈")
}

func (d *Dog) getAge() {
	fmt.Println(d.age)
}

type Cat struct {
	age int
}

func (c *Cat) sound() {
	fmt.Println("야옹")
}
func (c *Cat) getAge() {
	fmt.Println(c.age)
	c.age += 2
}

func main() {
	//선언 방법 1(이 방법은 경고표시가뜬다.)
	var my_cat Animal
	my_cat = &Cat{age: 2}
	my_cat.getAge()
	my_cat.sound()

	//선언 방법 2
	var my_dog = &Dog{age: 3}
	my_dog.getAge()
	my_dog.sound()
}
