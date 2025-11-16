package main

import (
	"analiseProjetodeAlgoritmos/internal/random"
	"analiseProjetodeAlgoritmos/internal/sort"
	"fmt"
	"log"
)

func main() {

	var size int
	var array []int

	fmt.Print("Array size: ")
	if _, err := fmt.Scan(&size); err != nil {
		log.Fatal("The array size must be a number")
	}

	randomArray, err := random.RandomArrayGenerator(size, 1500)
	if err != nil {
		log.Fatal(err)
	}

	copy(array, randomArray)
	bubbleSortTimer := sort.MeasureDuration(sort.BubbleSort, array)

	copy(array, randomArray)
	selectionSortTime := sort.MeasureDuration(sort.SelectionSort, array)

	bubbleSortTimer.Log("Bubble Sort")
	selectionSortTime.Log("Selection Sort")
}
