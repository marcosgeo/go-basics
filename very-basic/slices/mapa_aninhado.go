package main

import "fmt"

func main() {
	empregadosPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereia": 8456.78,
		},
		"J": {
			"José João": 11324.98,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(empregadosPorLetra, "P")
	for letra, empreg := range empregadosPorLetra {
		fmt.Println(letra, empreg)
	}
}