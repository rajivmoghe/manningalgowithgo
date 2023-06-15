package main

func quicksort(arr []int) {
	quicksort_0(arr, 0, len(arr)-1)
}

func quicksort_0(arr []int, low int, high int) {
	if low < high {
		pi := partitionLomuto(arr, low, high)
		quicksort_0(arr, low, pi-1)
		quicksort_0(arr, pi+1, high)
	}
}

func partitionHoare(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func partitionLomuto(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
