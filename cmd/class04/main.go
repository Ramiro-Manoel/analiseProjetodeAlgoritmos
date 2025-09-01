package main

import (
	"analiseProjetodeAlgoritmos/internal/random"
	"fmt"
	"log"
)

func main() {

	var size int

	fmt.Print("Array size: ")
	if _, err := fmt.Scan(&size); err != nil {
		log.Fatal("The array size must be a number")
	}

	array, err := random.RandomArrayGenerator(size, 1500)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(array)
}
