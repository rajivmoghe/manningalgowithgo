package main

import "fmt"

func main() {
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
