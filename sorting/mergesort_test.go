package sorting

import "testing"

func TestMergeSort_MergeSort(t *testing.T) {
	actualArray := []int{9,5,7,3,6,2,1}
	expectedArray := MergeSort(actualArray)

	if expectedArray[len(actualArray) - 1] != 9  {
		t.Fatalf("Expected %v but got %v", expectedArray, actualArray)
	}
}