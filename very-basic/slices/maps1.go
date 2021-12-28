package main

import "fmt"

func main() {
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[111] = "Maria"
	aprovados[222] = "Pedro"
	aprovados[845] = "Ana"
	fmt.Println(aprovados)

	for num, nome := range aprovados {
		fmt.Printf("%s (ID: %d)\n", nome, num)
	}

	fmt.Println(aprovados[845])
	delete(aprovados, 845)
	fmt.Println("Vazio", aprovados[845])
	fmt.Println(aprovados)
}