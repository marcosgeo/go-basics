package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria 1", "Pq voce não fala comigo", 3)
	// fale("João 1", "Só posso falar depois de você.", 1)

	// go fale("Maria 2", "Ei", 5)
	// go fale("João 2", "Opa", 5)

	//go fale("Maria 3", "Entendi!!!", 5)
	fale("João 3", "Parabéns", 5)
}
