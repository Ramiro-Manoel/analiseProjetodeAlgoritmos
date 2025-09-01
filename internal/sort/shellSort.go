package sort

func ShellSort(array []int) {
	lenght := len(array)

	for gap := lenght / 2; gap > 0; gap /= 2 {
		for i := gap; i < lenght; i++ {
			temp := array[i]
			j := i
			for j >= gap && array[j-gap] > temp {
				array[j] = array[j-gap]
				j -= gap
			}
			array[j] = temp
		}
	}
}
