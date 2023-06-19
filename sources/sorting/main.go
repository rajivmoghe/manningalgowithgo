package main

import (
	"fmt"
	"strconv"
)

type Customer struct {
	id            string
	num_purchases int
}

func main() {
	binarysearch()
}

func binarysearch() {
	// Make and sort an array.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)
	ar1 := make_random_array(num_items, max)
	bubblesort(ar1)
	print_array(ar1, 20)
	check_sorted(ar1)

	for {
		// Get the target as a string.
		var target_string string
		fmt.Printf("Target <Enter> to quit:")
		fmt.Scanln(&target_string)

		// If the target string is blank, break out of the loop.
		if len(target_string) == 0 {
			break
		}

		// Convert to int and find it.
		target, _ := strconv.Atoi(target_string)
		index, num_tests := binary_search(ar1, target)
		if index < 0 || index >= len(ar1) {
			fmt.Printf("Target %d not found, %d tests\n", target, num_tests)
		} else {
			fmt.Printf("arr[%d] = %d, %d tests\n", index, ar1[index], num_tests)
		}
	}
}

func linearsearch() {
	// Make and sort an array.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)
	ar1 := make_random_array(num_items, max)
	print_array(ar1, 20)
	// check_sorted(ar1)

	for {
		// Get the target as a string.
		var target_string string
		fmt.Printf("Target <Enter> to quit:")
		fmt.Scanln(&target_string)

		// If the target string is blank, break out of the loop.
		if len(target_string) == 0 {
			break
		}

		// Convert to int and find it.
		target, _ := strconv.Atoi(target_string)
		index, num_tests := linear_search(ar1, target)
		if index < 0 || index >= len(ar1) {
			fmt.Printf("Target %d not found, %d tests\n", target, num_tests)
		} else {
			fmt.Printf("arr[%d] = %d, %d tests\n", index, ar1[index], num_tests)
		}
	}

}

func sortstuff() {
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
