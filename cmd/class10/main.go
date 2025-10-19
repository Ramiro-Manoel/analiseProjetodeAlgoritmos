package main

import (
	"analiseProjetodeAlgoritmos/internal/intervalPartitioning"
	"fmt"
)

func runIntervalPartitioning(intervals intervalPartitioning.Intervals, expectedResult int) {
	fmt.Printf("intervals: %v\n", intervals)
	fmt.Printf("expected result: %v\n", expectedResult)

	fmt.Printf("result %v\n", intervalPartitioning.AlocateIntervals(intervals))
	fmt.Println()
}

func main() {
	runIntervalPartitioning(intervalPartitioning.Intervals{
		{Start: 1, End: 3},
		{Start: 3, End: 5},
		{Start: 5, End: 7},
	}, 1)

	runIntervalPartitioning(intervalPartitioning.Intervals{
		{Start: 1, End: 5},
		{Start: 2, End: 6},
		{Start: 3, End: 7},
	}, 3)

	runIntervalPartitioning(intervalPartitioning.Intervals{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	}, 2)

	runIntervalPartitioning(intervalPartitioning.Intervals{
		{Start: 9, End: 10},
		{Start: 9, End: 12},
		{Start: 9, End: 10},
		{Start: 11, End: 13},
		{Start: 11, End: 12},
		{Start: 12, End: 13},
	}, 4)

	runIntervalPartitioning(intervalPartitioning.Intervals{
		{Start: 0, End: 5},
		{Start: 1, End: 3},
		{Start: 1, End: 8},
		{Start: 3, End: 5},
		{Start: 4, End: 7},
		{Start: 6, End: 9},
		{Start: 9, End: 10},
	}, 4)

}
