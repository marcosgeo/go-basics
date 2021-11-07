package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	content := make([]byte, 9999)
	file.Read(content)
	fmt.Println(string(content))
}
