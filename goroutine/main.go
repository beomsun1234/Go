package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch3 := make(chan bool)
	ch2 := make(chan string)
	sg := &sync.WaitGroup{}
	sg.Add(2)

	go goRT1("1", ch1, ch3)
	go goRT2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case a := <-ch1:
			fmt.Println("ch1 = ", a, ", ch3 = ", <-ch3)
			sg.Done()
		case b := <-ch2:
			fmt.Println("ch2 =", b)
			sg.Done()
		}
	}
	sg.Wait()
	fmt.Println("-------------------------------------------")
	fmt.Println("end")
	close(ch1)
	close(ch2)
	close(ch3)

}

func goRT1(gubun string, ch chan int, ch2 chan bool) {
	fmt.Println("------start goRT1------------")
	if gubun == "1" {
		ch <- 1
		ch2 <- true
	} else {
		ch <- 0
		ch2 <- false
	}

}
func goRT2(ch chan string) {
	fmt.Println("------start goRT2------------")
	ch <- "s"
}
