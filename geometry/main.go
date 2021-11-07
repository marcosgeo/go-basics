package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	length float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	trg := triangle{base: 6.4, height: 9.4}
	sq := square{length: 8}

	printArea(trg)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}
