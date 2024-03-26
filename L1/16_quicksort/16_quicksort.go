package quicksort



func partitionQuick(arr []int, low, high int) int {
	i := low
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		p := partitionQuick(arr, low, high)
		QuickSort(arr, low, p-1)
		QuickSort(arr, p+1, high)
	}
}
