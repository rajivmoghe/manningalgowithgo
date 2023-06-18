package main

func binary_search(arr []int, target int) (index, num_tests int) {
	num_tests = 0
	index, num_tests = binarySearchNonRecursive(arr, target)
	return index, num_tests
}

// Recursive binary search
func binarySearchRecursive(ar2 []int, low int, high int, x int, tests int) (index, num_tests int) {
	tests += 1
	if low <= high {
		mid := low + (high-low)/2
		if ar2[mid] == x {
			return mid, tests
		} else if ar2[mid] < x {
			return binarySearchRecursive(ar2, mid+1, high, x, tests)
		} else {
			return binarySearchRecursive(ar2, low, mid-1, x, tests)
		}
	}
	return -1, tests
}

// Non-recursive binary search
func binarySearchNonRecursive(arr []int, x int) (index, num_tests int) {
	low := 0
	high := len(arr) - 1
	num_tests = 0
	for low <= high {
		mid := low + (high-low)/2
		num_tests++
		if arr[mid] == x {
			return mid, num_tests
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, num_tests
}
