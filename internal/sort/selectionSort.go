package sort

func SelectionSort(array []int) {

	for i := range array {
		smallerIndex := i
		for j := i; j < len(array); j++ {
			if array[j] < array[smallerIndex] {
				smallerIndex = j
			}
		}
		array[i], array[smallerIndex] = array[smallerIndex], array[i]
	}
}
