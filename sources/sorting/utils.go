package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Make a fixed array for testing
func make_fixed_array() []int {
	arr := make([]int, 11)
	arr[0] = 42
	arr[1] = 17
	arr[2] = 23
	arr[3] = 70
	arr[4] = 34
	arr[5] = 85
	arr[6] = 44
	arr[7] = 12
	arr[8] = 65
	arr[9] = 83
	arr[10] = 45
	return arr
}

// Make an array containing pseudorandom numbers in [0, max).
func make_random_array(num_items, max int) []int {
	rand.NewSource(time.Now().UnixNano())
	arr := make([]int, num_items)
	for i := 0; i < num_items; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}

// Make an array containing pseudorandom numbers in [0, max).
func make_random_Customer_array(num_items, max int) []Customer {
	rand.NewSource(time.Now().UnixNano())
	arr := make([]Customer, num_items)
	for i := 0; i < num_items; i++ {
		id := fmt.Sprintf("C%d", i)
		purchase_count := rand.Intn(max)
		arr[i] = Customer{id: id, num_purchases: purchase_count}
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

// Print at most num_items items.
func print_Customer_array(arr []Customer, num_items int) {
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

func check_Customer_sorted(arr []Customer) {
	for i := 1; i < len(arr); i++ {
		if arr[i-1].num_purchases > arr[i].num_purchases {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}
	fmt.Println("The array is sorted")
}
