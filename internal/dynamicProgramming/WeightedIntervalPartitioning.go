package WeightedIntervalPartitioning

import (
	"sort"
)

type WeightedInterval struct {
	Start int
	End   int
	Value int
}

type WeightedIntervals []WeightedInterval

type Scheduling struct {
	Value     int
	Intervals WeightedIntervals
}

type schedulings []Scheduling

func (a WeightedIntervals) Len() int           { return len(a) }
func (a WeightedIntervals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a WeightedIntervals) Less(i, j int) bool { return a[i].End < a[j].End }

func getIntervals(DP []int, precedings []int, intervals WeightedIntervals, i int) {

}

func WeightedIntervalScheduling(intervals WeightedIntervals) Scheduling {
	if len(intervals) == 0 {
		return Scheduling{}
	}

	sort.Sort(intervals)
	intervalsSize := len(intervals)
	precedings := findPrecedings(intervals)

	DP := make(schedulings, intervalsSize)
	DP[0] = Scheduling{
		Value:     intervals[0].Value,
		Intervals: WeightedIntervals{intervals[0]}}

	for i := 1; i < intervalsSize; i++ {
		currentValue := intervals[i].Value

		if precedings[i] != -1 {
			currentValue += DP[precedings[i]].Value
		}

		previousValue := DP[i-1].Value

		if currentValue > previousValue {
			var newIntervals WeightedIntervals

			if precedings[i] != -1 {
				newIntervals = append(DP[precedings[i]].Intervals, intervals[i])
			} else {
				newIntervals = WeightedIntervals{intervals[i]}
			}
			DP[i] = Scheduling{
				Value:     currentValue,
				Intervals: newIntervals}
		} else {
			DP[i] = Scheduling{
				Value:     previousValue,
				Intervals: DP[i-1].Intervals}
		}
	}

	return DP[intervalsSize-1]
}
