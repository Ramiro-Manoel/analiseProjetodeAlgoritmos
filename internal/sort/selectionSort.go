package sort

func SelectionSort(array []int) []int {

	for i, num := range array {
		smallerIndex := i
		for j := i + 1; i < len(array); j++ {
			if array[j] < num {
				smallerIndex = j
			}
		}
		smallerNumber := array[i]
		array[i] = array[smallerIndex]
		array[smallerIndex] = smallerNumber
	}
	return array
}
