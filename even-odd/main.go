package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range numbers {
		if n%2 == 0 {
			fmt.Printf("%d is Even", n)
		} else {
			fmt.Printf("%d is Odd", n)
		}
		fmt.Println()
	}
}
