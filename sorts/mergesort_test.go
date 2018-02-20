package sorts

import "testing"

func TestMergeSortSorts(t *testing.T) {
	arr := []int{0, 3, 5, 4, 2, 1, 7, 6}
	sorted := []int{0, 1, 2, 3, 4, 5, 6, 7}
	MergeSort(arr)
	for i, _ := range arr {
		if arr[i] != sorted[i] {
			t.Errorf("Expected arr to be sorted but at index=%d arr was %d when should be %d", i, arr[i], sorted[i])
		}
	}
}

func TestMergeSortEmptyCase(t *testing.T) {
	arr := []int{}
	MergeSort(arr)
	if len(arr) != 0 {
		t.Errorf("Mergesort somehow added to empty array")
	}
}
