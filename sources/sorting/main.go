package main

func main() {
	ar1 := make_fixed_array()
	ar1 = make_random_array(15, 500)
	print_array(ar1, 20)
	check_sorted(ar1)
	// bubblesort(ar1)
	quicksort(ar1, 0, len(ar1)-1)
	print_array(ar1, 30)
	check_sorted(ar1)
}
