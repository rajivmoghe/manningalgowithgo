package main

import "fmt"

var fibonacci_values = []int64{0, 1}

func fibonacci_on_the_fly(n int64) int64 {
	if int64(len(fibonacci_values)) > n {
		fmt.Println("From Slice", n)
		return fibonacci_values[n]
	} else {
		fmt.Println("From loop", n)
		fib_1 := fibonacci_on_the_fly(n - 1)
		fib_2 := fibonacci_on_the_fly(n - 2)
		fibonacci_values = append(fibonacci_values, fib_1+fib_2)
		return fibonacci_values[n]
	}
}
