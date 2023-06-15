package main

type Customer struct {
	id            string
	num_purchases int
}

func main() {
	// ar1 := make_fixed_array()
	// ar1 := make_random_array(15, 500)
	ar1 := make_random_Customer_array(10, 7)
	// print_array(ar1, 20)
	print_Customer_array(ar1, 30)
	// check_sorted(ar1)
	check_Customer_sorted(ar1)
	// bubblesort(ar1)
	// quicksort(ar1)
	countingsort(ar1)
	// print_array(ar1, 30)
	print_Customer_array(ar1, 30)
	// check_sorted(ar1)
	check_Customer_sorted(ar1)
}
