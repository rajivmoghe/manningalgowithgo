package main

import (
	"fmt"
	"math/rand"
	"time"
	//"math/rand"
)

func main() {
	linear()
}

func chaining() {
	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101"},
		Employee{"Bob Baker", "202-555-0102"},
		Employee{"Cindy Cant", "202-555-0103"},
		Employee{"Dan Deever", "202-555-0104"},
		Employee{"Edwina Eager", "202-555-0105"},
		Employee{"Fred Franklin", "202-555-0106"},
		Employee{"Gina Gable", "202-555-0107"},
		Employee{"Herb Henshaw", "202-555-0108"},
		Employee{"Ida Iverson", "202-555-0109"},
		Employee{"Jeb Jacobs", "202-555-0110"},
		Employee{"Hello World", "202-555-0101"},
	}

	hash_table := NewChainingHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hash_table.delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
}

func linear() {
	// Make some names.
	employees := []Employee{
		// Employee{"Hello World", "202-555-0000"},
		Employee{"Ann Archer", "202-555-0101"},
		Employee{"Bob Baker", "202-555-0102"},
		Employee{"Cindy Cant", "202-555-0103"},
		Employee{"Dan Deever", "202-555-0104"},
		Employee{"Edwina Eager", "202-555-0105"},
		Employee{"Fred Franklin", "202-555-0106"},
		Employee{"Gina Gable", "202-555-0107"},
	}

	hash_table := NewLinearProbingHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	// // fmt.Println("Deleting Dan Deever")
	// // hash_table.delete("Dan Deever")
	// fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Println("Changing Gina Gable")
	hash_table.set("Gina Gable", "202-555-0111")
	fmt.Println("Changing Hello World")
	hash_table.set("Hello World", "202-555-1111")
	hash_table.dump_concise()
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Printf("Gina Gable: %s\n", hash_table.get("Gina Gable"))
	fmt.Printf("Hello World: %s\n", hash_table.get("Hello World"))

	// Look at clustering.
	random := rand.New(rand.NewSource(time.Now().UnixMilli())) // Initialize with a fixed seed
	big_capacity := 1009
	big_hash_table := NewLinearProbingHashTable(big_capacity)
	num_items := int(float32(big_capacity) * 0.9)
	for i := 0; i < num_items; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		big_hash_table.set(str, str)
	}
	big_hash_table.dump_concise()
	fmt.Printf("Average probe sequence length: %f\n",
		big_hash_table.ave_probe_sequence_length())
}
