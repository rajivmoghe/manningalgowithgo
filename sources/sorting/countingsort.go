package main

// Actually sort the array with counting sort
func countingsort(arr []Customer) {
	C := fillUpC(arr)
	B := make([]Customer, len(arr))

	for i := len(arr); i > 0; i-- {
		c_chk := arr[i-1].num_purchases
		b_idx := C[c_chk]
		B[b_idx-1] = arr[i-1]
		C[c_chk] -= 1
	}

	copy(arr, B)
}

func fillUpC(arr []Customer) []int {
	max := -1
	for _, cust := range arr {
		if cust.num_purchases > max {
			max = cust.num_purchases
		}
	}

	counter := make([]int, max+1)
	for _, cust := range arr {
		counter[cust.num_purchases]++
	}

	rolling := counter[0]
	for i := 1; i < len(counter); i++ {
		counter[i] += rolling
		rolling = counter[i]
	}
	return counter
}
