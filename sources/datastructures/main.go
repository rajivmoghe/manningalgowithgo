package main

import "fmt"

func main() {
	list_loops()
}

func list_main() {
	// small_list_test()

	// Make a list from an array of values.
	greek_letters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := make_linked_list()       // make an empty holder list
	list.add_range(greek_letters)    // add nodes to the list
	fmt.Println(list.to_string(" ")) // print it out
	fmt.Println("linked list done")

	// Demonstrate a stack.
	stack := make_linked_list()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.is_empty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.to_string(" "))
	}
	fmt.Println("linked list done")
}

func list_loops() {
	// Make a list from a slice of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := make_linked_list()
	list.add_range(values)

	fmt.Println(list.to_string(" "))
	if list.has_loop() {
		fmt.Println("Has loop x")
	} else {
		fmt.Println("No loop x")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.to_string_max(" ", 10))
	if list.has_loop() {
		fmt.Println("Has loop y")
	} else {
		fmt.Println("No loop y")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.to_string_max(" ", 10))
	if list.has_loop() {
		fmt.Println("Has loop z")
	} else {
		fmt.Println("No loop z")
	}
}
