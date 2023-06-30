package main

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
