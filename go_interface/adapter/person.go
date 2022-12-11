package adapter

import "fmt"

type Person struct {
	age  int
	name string
}

func NewPerson(age int, name string) *Person {

	return &Person{age: age, name: name}
}

func (p *Person) IntroducePerson() {
	fmt.Printf("hi myname is %s \n", p.name)
}
