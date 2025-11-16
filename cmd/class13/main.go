package main

import (
	dp "analiseProjetodeAlgoritmos/internal/dynamicProgramming"
	"fmt"
)

func run(intervals dp.WeightedIntervals, expectedResult dp.Scheduling) {
	fmt.Printf("intervals: %v\n", intervals)
	fmt.Printf("expected result: %v\n", expectedResult)

	fmt.Printf("result %v\n", dp.WeightedIntervalScheduling(intervals))
	fmt.Println()
}

func main() {

	// Exemplo 1
	run(dp.WeightedIntervals{
		{Start: 1, End: 5, Value: 20},
		{Start: 6, End: 9, Value: 30},
		{Start: 2, End: 8, Value: 10},
		{Start: 10, End: 12, Value: 40}},

		dp.Scheduling{
			Value: 90,
			Intervals: dp.WeightedIntervals{
				{Start: 1, End: 5, Value: 20},
				{Start: 6, End: 9, Value: 30},
				{Start: 10, End: 12, Value: 40}},
		},
	)

	// Exemplo 2
	run(dp.WeightedIntervals{
		{Start: 3, End: 6, Value: 40},
		{Start: 7, End: 10, Value: 10},
		{Start: 1, End: 4, Value: 20},
		{Start: 5, End: 8, Value: 30}},

		dp.Scheduling{
			Value: 50,
			Intervals: dp.WeightedIntervals{
				{Start: 3, End: 6, Value: 40},
				{Start: 7, End: 10, Value: 10}},
		},
	)

	// Exemplo 3
	run(dp.WeightedIntervals{
		{Start: 1, End: 3, Value: 10},
		{Start: 4, End: 6, Value: 20},
		{Start: 7, End: 9, Value: 10},
		{Start: 2, End: 5, Value: 40}},

		dp.Scheduling{
			Value: 40,
			Intervals: dp.WeightedIntervals{
				{Start: 2, End: 5, Value: 40}},
		},
	)

	// Exemplo 4
	run(dp.WeightedIntervals{
		{Start: 1, End: 2, Value: 10},
		{Start: 2, End: 3, Value: 20},
		{Start: 3, End: 4, Value: 30},
		{Start: 4, End: 5, Value: 40}},

		dp.Scheduling{
			Value: 100,
			Intervals: dp.WeightedIntervals{
				{Start: 1, End: 2, Value: 10},
				{Start: 2, End: 3, Value: 20},
				{Start: 3, End: 4, Value: 30},
				{Start: 4, End: 5, Value: 40}},
		},
	)

	// Exemplo 5
	run(dp.WeightedIntervals{
		{Start: 10, End: 15, Value: 30},
		{Start: 1, End: 4, Value: 40},
		{Start: 3, End: 6, Value: 20},
		{Start: 16, End: 20, Value: 10}},

		dp.Scheduling{
			Value: 80,
			Intervals: dp.WeightedIntervals{
				{Start: 1, End: 4, Value: 40},
				{Start: 10, End: 15, Value: 30},
				{Start: 16, End: 20, Value: 10}},
		},
	)

}
