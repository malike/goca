package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort_MergeSort(t *testing.T) {
	actualArray := []int{9, 5, 7, 3, 6, 2, 1}
	sortedArray := MergeSort(actualArray)
	expectedSort := []int{1, 2, 3, 5, 6, 7, 9}

	if !reflect.DeepEqual(sortedArray, expectedSort) {
		t.Fatalf("Expected %v but got %v", expectedSort, sortedArray)
	}
}
