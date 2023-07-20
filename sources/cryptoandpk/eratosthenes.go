package main

import "fmt"

// Build a sieve of Eratosthenes.
func sieve_of_eratosthenes(max int) []bool {
	sieve := make([]bool, max+1)
	for i := 2; i < len(sieve); i++ { // first 2 are false, mark all others true
		sieve[i] = true
	}

	for i := range sieve {
		if sieve[i] {
			for j := i + i; j < len(sieve); j += i { // if prime, mark all its multiples false
				sieve[j] = false
			}
		}
	}
	return sieve
}

// Print out the primes in the sieve.
func print_sieve(sieve []bool) {
	i := 2
	for i < len(sieve) {
		if sieve[i] {
			fmt.Printf("%d ", i)
		}
		if i > 2 {
			i++
		}
		i++
	}
	fmt.Println()
}

// Convert the sieve into a slice holding prime numbers.
func sieve_to_primes(sieve []bool) []int {
	var primes []int
	i := 2
	for i < len(sieve) {
		if sieve[i] {
			primes = append(primes, i)
		}
		if i > 2 {
			i++
		}
		i++
	}
	return primes
}
