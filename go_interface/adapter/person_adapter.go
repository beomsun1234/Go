package adapter

import (
	"fmt"
)

type PersonAdapter struct {
	*Person
}

func (p *PersonAdapter) Sound() {
	p.Person.IntroducePerson()
}

func (p *PersonAdapter) GetAge() {
	fmt.Println(p.Person.age)
}
