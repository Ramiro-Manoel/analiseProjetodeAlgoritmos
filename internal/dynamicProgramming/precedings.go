package WeightedIntervalPartitioning

func findPreceding(intervals []WeightedInterval, i int) int {
	for j := i - 1; j >= 0; j-- {
		if intervals[j].End <= intervals[i].Start {
			return j
		}
	}
	return -1
}

func findPrecedings(intervals []WeightedInterval) []int {
	precedings := make([]int, len(intervals))
	for i := range intervals {
		precedings[i] = findPreceding(intervals, i)
	}
	return precedings
}
