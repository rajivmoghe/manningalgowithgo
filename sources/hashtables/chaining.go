package main

import (
	"fmt"
)

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
	buckets := make([][]*Employee, num_buckets)
	hashTable := &ChainingHashTable{num_buckets: num_buckets, buckets: buckets}
	return hashTable
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

// Find the bucket and Employee holding this key.
// Return the bucket number and Employee number in the bucket.
// If the key is not present, return the bucket number and -1.
func (hash_table *ChainingHashTable) find(name string) (int, int) {
	hash_key := hash(name) % hash_table.num_buckets
	for i, value := range hash_table.buckets[hash_key] {
		if value.name == name {
			return hash_key, i
		}
	}
	return hash_key, -1
}

// Add an item to the hash table.
func (hash_table *ChainingHashTable) set(name string, phone string) {
	hash_key, idx := hash_table.find(name)
	if idx >= 0 {
		empl := hash_table.buckets[hash_key][idx] // element with given name found - just update phone number
		empl.phone = phone
	} else {
		employee := &Employee{name, phone}
		slice_to_add := &hash_table.buckets[hash_key]
		*slice_to_add = append(*slice_to_add, employee)
	}
}

// Return an item from the hash table.
func (hash_table *ChainingHashTable) get(name string) string {
	hash_key, idx := hash_table.find(name)
	if idx >= 0 {
		return hash_table.buckets[hash_key][idx].phone
	} else {
		return ""
	}
}

// Return true if the person is in the hash table.
func (hash_table *ChainingHashTable) contains(name string) bool {
	_, idx := hash_table.find(name)
	return idx >= 0
}

// Delete this key's entry.
func (hash_table *ChainingHashTable) delete(name string) {
	hash_key, idx := hash_table.find(name)
	if idx > -1 {
		hash_table.buckets[hash_key] =
			append(hash_table.buckets[hash_key][:idx],
				hash_table.buckets[hash_key][idx+1:]...)
	}
}
