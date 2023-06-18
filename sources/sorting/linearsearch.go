package main

func linear_search(arr []int, target int) (index, num_tests int) {
	index = -1
	num_tests = 0

	for jidx, val := range arr {
		num_tests++
		if val == target {
			index = jidx
			break
		}
	}

	return index, num_tests
}
