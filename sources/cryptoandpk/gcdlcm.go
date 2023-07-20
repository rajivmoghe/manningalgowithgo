package main

func gcd(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a < b {
		if a == 0 {
			return b
		} else {
			return gcd(b%a, a)
		}
	} else {
		return gcd(b, a)
	}
}

func lcm(a, b int64) (int64, int64) {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	gcd := gcd(a, b)
	lcm := b / gcd
	lcm *= a

	return gcd, lcm
}

func gcdlcm(a, b int64) (int64, int64) {
	return lcm(a, b)
}
