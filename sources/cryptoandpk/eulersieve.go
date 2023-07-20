package main

// Build an Euler's sieve.
func eulers_sieve(max int) []bool {
	sieve := make([]bool, max+1)
	sieve[2] = true

	for i := 3; i <= max; i += 2 {
		sieve[i] = true
	}

	for i := range sieve {
		if sieve[i] {
			maxI := (max / i)
			if maxI%2 == 0 {
				maxI--
			}

			for j := maxI; j >= i; j -= 2 {
				if sieve[i] {
					sieve[i*j] = false
				}
			}
		}
	}
	return sieve
}
