package sorting

import "testing"

func TestQuickSort_QuickSort(t *testing.T) {
	actualArray := []int{9,5,7,3,6,2,1}
	expectedArray := QuickSort(actualArray)

	if expectedArray[len(actualArray) - 1] != 9  {
		t.Fatalf("Expected %v but got %v", expectedArray, actualArray)
	}
}