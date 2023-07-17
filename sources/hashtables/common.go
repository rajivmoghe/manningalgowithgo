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

type Employee_D struct {
	name    string
	phone   string
	deleted bool
}

type LinearProbingHashTable_D struct {
	capacity  int
	employees []*Employee_D
}

// Initialize a LinearProbingHashTable and return a pointer to it.
func NewLinearProbingHashTable_D(capacity int) *LinearProbingHashTable_D {
	return &LinearProbingHashTable_D{capacity: capacity, employees: make([]*Employee_D, capacity)}
}

// Display the hash table's contents.
func (hash_table *LinearProbingHashTable_D) dump() {
	for i, element := range hash_table.employees {
		if element == nil {
			fmt.Printf("%2d: ---\n", i)
		} else if element.deleted {
			fmt.Printf("%2d: xxx\n", i)
		} else {
			fmt.Printf("%2d: %15s\t%s\n", i, element.name, element.phone)
		}

	}
}

// Make a display showing whether each slice entry is nil.
func (hash_table *LinearProbingHashTable_D) dump_concise() {
	// Loop through the slice.
	for i, employee := range hash_table.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else if employee.deleted {
			fmt.Printf("x")
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

// Show this key's probe sequence.
func (hash_table *LinearProbingHashTable_D) probe(name string) (int) {
    // Hash the key.
    hash := hash(name) % hash_table.capacity
    fmt.Printf("Probing %s (%d)\n", name, hash)

    // Keep track of a deleted spot if we find one.
    deleted_index := -1

    // Probe up to hash_table.capacity times.
    for i := 0; i < hash_table.capacity; i++ {
        index := (hash + i) % hash_table.capacity

        fmt.Printf("    %d: ", index)
        if hash_table.employees[index] == nil {
            fmt.Printf("---\n")
        } else if hash_table.employees[index].deleted {
            fmt.Printf("xxx\n")
        } else {
            fmt.Printf("%s\n", hash_table.employees[index].name)
        }

        // If this spot is empty, the value isn't in the table.
        if hash_table.employees[index] == nil {
            // If we found a deleted spot, return its index.
            if deleted_index >= 0 {
                fmt.Printf("    Returning deleted index %d\n", deleted_index)
                return deleted_index
            }

            // Return this index, which holds nil. 
            fmt.Printf("    Returning nil index %d\n", index)
            return index
        }

        // If this spot is deleted, remember where it is.
        if hash_table.employees[index].deleted {
            if deleted_index < 0 {
                deleted_index = index
            }
        } else if hash_table.employees[index].name == name {
            // If this cell holds the key, return its data.
            fmt.Printf("    Returning found index %d\n", index)
            return index
        }

        // Otherwise continue the loop.
    }

    // If we get here, then the key is not
    // in the table and the table is full.

    // If we found a deleted spot, return it.
    if deleted_index >= 0 {
        fmt.Printf("    Returning deleted index %d\n", deleted_index)
        return deleted_index
    }

    // There's nowhere to put a new entry.    
    fmt.Printf("    Table is full\n")
    return -1
}

// Quadratic Probing Hashtable element
type QuadraticProbingHashTable struct {
    capacity    int
    num_entries int
    employees   []*Employee_D
}

// Display the hash table's contents.
func (hash_table *QuadraticProbingHashTable) dump() {
	// Loop through the array.
	for i, employee := range hash_table.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf("%2d: ---\n", i)
		} else if employee.deleted {
			// This value has been deleted.
			fmt.Printf("%2d: xxx\n", i)
		} else {
			// Display this entry.
			fmt.Printf("%2d: %-15s %s\n", i, employee.name, employee.phone)
		}
	}
}

func (hash_table *QuadraticProbingHashTable) dump_concise() {
	// Loop through the slice.
	for i, employee := range hash_table.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else if employee.deleted {
			fmt.Printf("x")
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

// Show this key's probe sequence.
func (hash_table *QuadraticProbingHashTable) probe(name string) (int) {
    // Hash the key.
    hash := hash(name) % hash_table.capacity
    fmt.Printf("Probing %s (%d)\n", name, hash)

    // Keep track of a deleted spot if we find one.
    deleted_index := -1

    // Probe up to hash_table.capacity times.
    for i := 0; i < hash_table.capacity; i++ {
        // index := (hash + i) % hash_table.capacity    // Linear Probing.
        index := (hash + i * i) % hash_table.capacity   // Quadratic probing.

        fmt.Printf("    %d: ", index)
        if hash_table.employees[index] == nil {
            fmt.Printf("---\n")
        } else if hash_table.employees[index].deleted {
            fmt.Printf("xxx\n")
        } else {
            fmt.Printf("%s\n", hash_table.employees[index].name)
        }

        // If this spot is empty, the value isn't in the table.
        if hash_table.employees[index] == nil {
            // If we found a deleted spot, return its index.
            if deleted_index >= 0 {
                fmt.Printf("    Returning deleted index %d\n", deleted_index)
                return deleted_index
            }

            // Return this index, which holds nil. 
            fmt.Printf("    Returning nil index %d\n", index)
            return index
        }

        // If this spot is deleted, remember where it is.
        if hash_table.employees[index].deleted {
            if deleted_index < 0 {
                deleted_index = index
            }
        } else if hash_table.employees[index].name == name {
            // If this cell holds the key, return its data.
            fmt.Printf("    Returning found index %d\n", index)
            return index
        }

        // Otherwise continue the loop.
    }

    // If we get here, then the key is not
    // in the table and the table is full.

    // If we found a deleted spot, return it.
    if deleted_index >= 0 {
        fmt.Printf("    Returning deleted index %d\n", deleted_index)
        return deleted_index
    }

    // There's nowhere to put a new entry.    
    fmt.Printf("    Table is full\n")
    return -1
}