package sort

import (
	"fmt"
	"time"
)

type Timer struct {
	notSorted time.Duration
	sorted    time.Duration
}

func (t Timer) Log(title string) {
	fmt.Println(title)

	fmt.Printf("Array não ordenado: %d ms\n", t.notSorted.Milliseconds())
	fmt.Printf("Array já ordenado: %d ms\n", t.sorted.Milliseconds())
}

func MeasureDuration(f func([]int), array []int) Timer {

	var timer Timer

	start := time.Now()
	f(array)
	timer.notSorted = time.Since(start)

	fmt.Print(time.Since(start))

	start = time.Now()
	f(array)
	timer.sorted = time.Since(start)

	return timer

}
