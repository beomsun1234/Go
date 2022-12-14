package main

import (
	"fmt"
)

type account_no_pointer struct {
	balance int
}

func (a *account_no_pointer) changeBalance(b int) {
	fmt.Printf("a의 메모리주소 %p\n", a)
	a.balance = b
}

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가르키는 메모리값: %d\n", *p)

	fmt.Println("-------------------------------------")
	account := account_no_pointer{balance: 100}
	account.changeBalance(10)

	fmt.Printf("account의 메모리주소 %p\n", &account)
	fmt.Printf("balnce = %d\n", account.balance)
}
