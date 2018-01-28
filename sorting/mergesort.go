package sorting

func Merge(left, right []int) []int {
	merge := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merge, right...)
		}
		if len(right) == 0 {
			return append(merge, left...)
		}
		if left[0] <= right[0] {
			merge = append(merge, left[0])
			left = left[1:]
		} else {
			merge = append(merge, right[0])
			right = right[1:]
		}
	}
	return merge
}

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	n := len(array) / 2
	left := MergeSort(array[:n])
	right := MergeSort(array[n:])
	return Merge(left, right)
}
