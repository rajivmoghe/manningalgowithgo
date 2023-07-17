package main

// Return the average probe sequence length for the items in the table.
func (hash_table *DoubleHashTable) ave_probe_sequence_length() float32 {
	total_length := 0
	num_values := 0
	for _, employee := range hash_table.employees {
		if employee != nil {
			_, probe_length := hash_table.find(employee.name)
			total_length += probe_length
			num_values++
		}
	}
	return float32(total_length) / float32(num_values)
}

// Return the key's index or where it would be if present.
// If the key is not present and the table is full, return -1.
func (hash_table *DoubleHashTable) find(name string) (int, int) {
	// Hash the key.
	hash1 := hash(name) % hash_table.capacity
	hash2 := hash2(name) % hash_table.capacity

	// Keep track of a deleted spot if we find one.
	deleted_index := -1

	// Probe up to hash_table.capacity times.
	for i := 0; i < hash_table.capacity; i++ {
		// index := (hash + i) % hash_table.capacity        // Linear Probing.
		// index := (hash + i * i) % hash_table.capacity    // Quadratic probing.
		index := (hash1 + i*hash2) % hash_table.capacity // Double hashing.

		// If this spot is empty, the value isn't in the table.
		if hash_table.employees[index] == nil {
			// If we found a deleted spot, return its index.
			if deleted_index >= 0 {
				return deleted_index, i + 1
			}

			// Return this index, which holds nil.
			return index, i + 1
		}

		// If this spot is deleted, remember where it is.
		if hash_table.employees[index].deleted {
			if deleted_index < 0 {
				deleted_index = index
			}
		} else if hash_table.employees[index].name == name {
			// If this cell holds the key, return its data.
			return index, i + 1
		}

		// Otherwise continue the loop.
	}

	// If we get here, then the key is not
	// in the table and the table is full.

	// If we found a deleted spot, return it.
	if deleted_index >= 0 {
		return deleted_index, hash_table.capacity
	}

	// There's nowhere to put a new entry.
	return -1, hash_table.capacity
}

// Add an item to the hash table.
func (hash_table *DoubleHashTable) set(name string, phone string) {
	// See where the key belongs.
	index, _ := hash_table.find(name)

	// If the key is missing and the table is full, panic.
	if index < 0 {
		panic("The hash table is full.")
	}

	// See if the key is already in the table.
	if hash_table.employees[index] == nil || hash_table.employees[index].deleted {
		// The key is not in the table. Put it here.
		employee := Employee_D{name, phone, false}
		hash_table.employees[index] = &employee
	} else {
		// The employee is already in the table. Update the phone value.
		hash_table.employees[index].phone = phone
	}
}

// Return an item from the hash table.
func (hash_table *DoubleHashTable) get(name string) string {
	// See where the key belongs.
	index, _ := hash_table.find(name)

	// See if we found an Employee.
	if index < 0 {
		// This name is not in the hash table and the hash table is full.
		return ""
	} else if hash_table.employees[index] == nil {
		// This name is not in the hash table.
		return ""
	} else if hash_table.employees[index].deleted {
		// This name is not in the hash table and we found a deleted spot.
		return ""
	} else {
		// Return this Employee's phone number.
		return hash_table.employees[index].phone
	}
}

// Return true if the person is in the hash table.
func (hash_table *DoubleHashTable) contains(name string) bool {
	// See where the key belongs.
	index, _ := hash_table.find(name)

	// If index < 0, then the key is not present
	// and the table is full.
	if index < 0 {
		return false
	}

	// If this entry is nil, then the key is not present.
	if hash_table.employees[index] == nil {
		return false
	}

	// If this entry was deleted, then the key is not present.
	if hash_table.employees[index].deleted {
		return false
	}

	// See if this Employee has the target key.
	return hash_table.employees[index].name == name
}

// Delete this key's entry.
func (hash_table *DoubleHashTable) delete(name string) {
	// See where the key belongs.
	index, _ := hash_table.find(name)

	// If we found the Employee struct, mark it as deleted.
	if index >= 0 &&
		hash_table.employees[index] != nil &&
		!hash_table.employees[index].deleted {
		hash_table.employees[index].deleted = true
	}
}
