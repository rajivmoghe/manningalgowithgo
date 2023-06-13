package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Make a fixed array for testing
func make_fixed_array() []int {
	arr := make([]int, 5)
	arr[0] = 4
	arr[1] = 1
	arr[2] = 2
	arr[3] = 0
	arr[4] = 4
	return arr
}

// Make an array containing pseudorandom numbers in [0, max).
func make_random_array(num_items, max int) []int {
	arr := make([]int, num_items)
	for i := 0; i < num_items; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}

// Print at most num_items items.
func print_array(arr []int, num_items int) {
	if len(arr) <= num_items {
		fmt.Println(arr)
	} else {
		fmt.Println(arr[:num_items])
	}
}

// Verify that the array is sorted.
func check_sorted(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}

	fmt.Println("The array is sorted")
}

// Actually sort the array
func bubblesort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			//if ith element is larger than jth, swap them.
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello World")
	ar1 := make_random_array(20000, 500000)
	print_array(ar1, 20)
	check_sorted(ar1)
	bubblesort(ar1)
	print_array(ar1, 30)
	check_sorted(ar1)
}
