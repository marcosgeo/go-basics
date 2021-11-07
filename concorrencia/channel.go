package main

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func process(base int, c chan int) {
	time.Sleep(time.Second)
	two := 2
	c <- two * base
}

func main() {
	// c := make(chan int)
	// go doisTresQuatroVezes(2, c)

	// a, b := <-c, <-c // recebendo dados dos canal
	// fmt.Printf("a: %v\nb: %v\n", a, b)

	// fmt.Println(<-c)

	valores := []int{2, 4, 6}
	canal := make(chan int)
	soma := 0
	for _, v := range valores {
		go process(v, canal)
		soma += <-canal
	}
	fmt.Printf("\nsoma: %d\n", soma) // 24: (2*2) + (2*4) + (2*6)
}
