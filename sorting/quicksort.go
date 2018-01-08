package sorting

import "math/rand"

func QuickSort(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	left, right := 0, len(array)-1

	pivot := rand.Int() % len(array)

	array[pivot], array[right] = array[right], array[pivot]

	for i := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	QuickSort(array[:left])
	QuickSort(array[left+1:])

	return array
}
