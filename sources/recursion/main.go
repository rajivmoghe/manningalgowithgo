package main

import "fmt"

func main() {
	run_factorial()
}

func run_factorial() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}
