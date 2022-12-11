package main

import (
	"fmt"
	"go/Go/go_interface/adapter"
)

type Animal interface {
	Sound()
	GetAge()
}

type Dog struct {
	age int
}

func (d *Dog) Sound() {
	fmt.Println("왈왈")
}

func (d *Dog) GetAge() {
	fmt.Println(d.age)
}

type Cat struct {
	age int
}

func (c *Cat) Sound() {
	fmt.Println("야옹")
}
func (c *Cat) GetAge() {
	fmt.Println(c.age)
	c.age += 2
}

func main() {
	//선언 방법 1(이 방법은 경고표시가뜬다.)
	var my_cat Animal
	my_cat = &Cat{age: 2}
	my_cat.GetAge()
	my_cat.Sound()

	//선언 방법 2
	var my_dog = &Dog{age: 3}
	my_dog.GetAge()
	my_dog.Sound()

	//adapter--------------------/
	fmt.Println("------------------------------------")
	var p = adapter.NewPerson(20, "park")
	var pa = &adapter.PersonAdapter{Person: p}
	pa.Sound()
	pa.GetAge()

}
