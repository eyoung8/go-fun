package sorts

import "testing"

func TestQuickSortSorts(t *testing.T) {
	arr := []int{0, 3, 5, 4, 2, 1, 7, 6}
	sorted := []int{0, 1, 2, 3, 4, 5, 6, 7}
	QuickSort(arr)
	for i, _ := range arr {
		if arr[i] != sorted[i] {
			t.Errorf("Expected arr to be sorted but at index=%d arr was %d when should be %d", i, arr[i], sorted[i])
		}
	}
}

func TestQuickSortEmptyCase(t *testing.T) {
	arr := []int{}
	QuickSort(arr)
	if len(arr) != 0 {
		t.Errorf("Quicksort somehow added to empty array")
	}
}
