package sort

func QuickSort(array []int, low int, high int) {
	if low < high {
		pivot := partition(array, low, high)

		QuickSort(array, low, pivot-1)
		QuickSort(array, pivot+1, high)
	}
}

func partition(array []int, low int, high int) int {

	pivot := array[high]
	i := low - 1

	for j := low; j < high; j++ {
		if array[j] < pivot {
			i++
			swap(array, i, j)
		}
	}
	swap(array, i+1, high)
	return i + 1
}

func swap(array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}
