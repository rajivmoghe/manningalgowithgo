package main

import "fmt"

// djb2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}
	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}

type Employee struct {
	name  string
	phone string
}
type ChainingHashTable struct {
	num_buckets int
	buckets     [][]*Employee
}

// Initialize a ChainingHashTable and return a pointer to it.
func NewChainingHashTable(num_buckets int) *ChainingHashTable {
	return &ChainingHashTable{num_buckets: num_buckets, buckets: make([][]*Employee, num_buckets)}
}

// Display the hash table's contents.
func (hash_table *ChainingHashTable) dump() {
	for i, elements := range hash_table.buckets {
		fmt.Printf("bucket: %d\n", i)
		for _, value := range elements {
			fmt.Printf("\t%s: %s\n", value.name, value.phone)
		}
	}
}

type LinearProbingHashTable struct {
	capacity  int
	employees []*Employee
}

// Initialize a LinearProbingHashTable and return a pointer to it.
func NewLinearProbingHashTable(capacity int) *LinearProbingHashTable {
	return &LinearProbingHashTable{capacity: capacity, employees: make([]*Employee, capacity)}
}

// Display the hash table's contents.
func (hash_table *LinearProbingHashTable) dump() {
	for i, element := range hash_table.employees {
		if element == nil {
			fmt.Printf("%2d: ---\n", i)
		} else {
			fmt.Printf("%2d: %15s\t%s\n", i, element.name, element.phone)
		}
	}
}

// Make a display showing whether each slice entry is nil.
func (hash_table *LinearProbingHashTable) dump_concise() {
	// Loop through the slice.
	for i, employee := range hash_table.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else {
			// Display this entry.
			fmt.Printf("O")
		}
		if i%50 == 49 {
			fmt.Println()
		}
	}
	fmt.Println()
}
