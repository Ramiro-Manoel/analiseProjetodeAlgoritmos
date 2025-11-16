package intervalPartitioning

import (
	"container/heap"
	"sort"
)

type Interval struct {
	Start int
	End   int
}
type Intervals []Interval

func (a Intervals) Len() int           { return len(a) }
func (a Intervals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Intervals) Less(i, j int) bool { return a[i].Start < a[j].End }

func AlocateIntervals(intervals Intervals) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(intervals)

	minHeap := &minHeap{}
	heap.Init(minHeap)

	for _, interval := range intervals {
		if minHeap.Len() > 0 && (*minHeap)[0] <= interval.Start {
			heap.Pop(minHeap)
			heap.Push(minHeap, interval.End)
		} else {
			heap.Push(minHeap, interval.End)
		}
	}
	return minHeap.Len()
}
