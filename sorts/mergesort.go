package sorts

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, high, mid)
}

func merge(arr []int, low int, high int, mid int) {
	// copy halves into new arrays
	lenA := mid - low + 1
	lenB := high - mid
	a := make([]int, lenA)
	b := make([]int, lenB)
	copy(a, arr[low:mid+1])
	copy(b, arr[mid+1:high+1])
	// rewrite old array from in order halves
	i, j, k := 0, 0, low
	for i < lenA && j < lenB {
		if a[i] <= b[j] {
			arr[k] = a[i]
			i++
		} else {
			arr[k] = b[j]
			j++
		}
		k++
	}
	// fill in the leftover
	for ; i < lenA; i++ {
		arr[k] = a[i]
		k++
	}
	for ; j < lenB; j++ {
		arr[k] = b[j]
		k++
	}
}
