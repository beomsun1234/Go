package main

import "fmt"

type Product struct {
	id    string
	name  string
	price int
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) buildId(id string) *Product {
	p.id = id
	return p
}

func (p *Product) buildName(name string) *Product {
	p.name = name
	return p
}

func (p *Product) buildPrice(price int) *Product {
	p.price = price
	return p
}

func main() {
	var p = NewProduct().buildId("2").buildName("test").buildPrice(20)

	fmt.Println(p.price)

	fmt.Println("change price")
	p.buildPrice(30)

	fmt.Println(p.price)

}
