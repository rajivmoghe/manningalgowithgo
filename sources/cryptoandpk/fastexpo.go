package main

// Perform fast exponentiation.
func fast_exp(num, pow int64) int64 {
	result := int64(1)
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
	}
	return result
}

// Perform fast exponentiation in a modulus.
func fast_exp_mod(num, pow, mod int64) int64 {
	result := int64(1)
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
			result %= mod
		}
		pow /= 2
		num *= num
		num %= mod
	}
	return result
}
