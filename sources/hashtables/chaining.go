package main

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
