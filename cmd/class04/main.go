package main

import (
	"analiseProjetodeAlgoritmos/internal/random"
	"analiseProjetodeAlgoritmos/internal/sort"
	"fmt"
	"log"
)

func main() {

	var size int

	fmt.Print("Array size: ")
	if _, err := fmt.Scan(&size); err != nil {
		log.Fatal("The array size must be a number")
	}

	array, err := random.RandomArrayGenerator(size)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(sort.QuickSort(array))

}
