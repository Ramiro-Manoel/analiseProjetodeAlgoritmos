package main

import (
	"fmt"
	"log"
)

func main() {

	var size, max int

	fmt.Print("Array size: ")
	if _, err := fmt.Scan(&size); err != nil {
		log.Fatal("The array size must be a number")
	}
	fmt.Print("Max random number: ")
	if _, err := fmt.Scan(&max); err != nil {
		log.Fatal("The max random number must be a number")
	}

	array, err := randomArrayGenerator(size, max)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("random array: ", array)

	array = bubbleSort(array)
	fmt.Println("sorted array: ", array)

	fmt.Println("prime numbers: ", getPrimeNumbers(array))
}
