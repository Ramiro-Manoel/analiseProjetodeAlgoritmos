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
	shelllSortDuration := sort.MeasureDuration(sort.ShellSort, array)

	copy(array, randomArray)
	mergeSortDuration := sort.MeasureDuration(sort.MergeSort, array)

	copy(array, randomArray)
	quickSortDuration := sort.MeasureDuration(sort.QuickSort, array)

	shelllSortDuration.Log("Shell Sort:")
	mergeSortDuration.Log("Merge Sort:")
	quickSortDuration.Log("Quick Sort:")
}
