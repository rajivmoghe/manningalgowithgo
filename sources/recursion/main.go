package main

import (
	"fmt"
	"strconv"
)

func main() {
	run_fibonacci()
}

func run_factorial() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}

func run_fibonacci() {
	for {
		// Get n as a string.
		var n_string string
		fmt.Printf("N: ")
		fmt.Scanln(&n_string)

		// If the n string is blank, break out of the loop.
		if len(n_string) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(n_string, 10, 64)
		fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
	}
}
