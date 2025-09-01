package main

import (
	"analiseProjetodeAlgoritmos/internal/random"
	"analiseProjetodeAlgoritmos/internal/sort"
	"fmt"
	"log"
	"time"
)

type Timer struct {
	sorted    time.Duration
	notSorted time.Duration
}

func (t Timer) log(title string) {
	fmt.Println(title)

	fmt.Printf("Array não ordenado: %d ms\n", t.notSorted.Milliseconds())
	fmt.Printf("Array já ordenado: %d ms\n", t.sorted.Milliseconds())
}

func measureDuration(f func([]int), array []int) Timer {

	var timer Timer

	start := time.Now()
	f(array)
	timer.notSorted = time.Since(start)

	start = time.Now()
	f(array)
	timer.sorted = time.Since(start)

	return timer

}

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

	bubbleSortTimer := measureDuration(sort.BubbleSort, array)
	selectionSortTime := measureDuration(sort.SelectionSort, array)

	bubbleSortTimer.log("Bubble Sort")
	selectionSortTime.log("Selection Sort")
}
