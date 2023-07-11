package main

// Return the key's index or where it would be if present and
// the probe sequence length.
// If the key is not present and the table is full, return -1 for the index.
func (hash_table *LinearProbingHashTable) find(name string) (int, int) {
	hash_key := hash(name) % hash_table.capacity
	for i := 0; i < hash_table.capacity; i++ {
		idx := (hash_key + i) % hash_table.capacity
		if hash_table.employees[idx] == nil || hash_table.employees[idx].name == name {
			return idx, i + 1
		}
	}

	return -1, hash_table.capacity
}

// Add an item to the hash table.
func (hash_table *LinearProbingHashTable) set(name string, phone string) {
	key_location, _ := hash_table.find(name)
	if key_location == -1 {
		panic("Table is full")
	} else {
		if hash_table.employees[key_location] == nil { //empty slot
			new_employee := &Employee{name, phone}
			hash_table.employees[key_location] = new_employee
		} else {
			hash_table.employees[key_location].phone = phone
		}
	}
}

// Return an item from the hash table.
func (hash_table *LinearProbingHashTable) get(name string) string {
	key_location, _ := hash_table.find(name)
	if key_location < 0 {
		return ""
	} else {
		if hash_table.employees[key_location] == nil {
			return ""
		} else {
			return hash_table.employees[key_location].phone
		}
	}
}

// Return true if the person is in the hash table.
func (hash_table *LinearProbingHashTable) contains(name string) bool {
	key_location, _ := hash_table.find(name)
	if key_location < 0 || hash_table.employees[key_location] == nil {
		return false
	}
	return true
}

// Return the average probe sequence length for the items in the table.
func (hash_table *LinearProbingHashTable) ave_probe_sequence_length() float32 {
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
