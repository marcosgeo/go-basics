package main

import "fmt"

// map is a reference type
func main() {
	// simple way of creating a mpa
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)

	// an empty map
	var colors2 map[string]string
	fmt.Println(colors2)

	// using make instruction
	colors3 := make(map[string]string)
	fmt.Println(colors3)

	// adding value value to a map
	colors["white"] = "ffffff"
	fmt.Println(colors)

	// removing a key from a map
	delete(colors, "green")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Hex code for given color %s is %s\n", key, value)
	}
}
