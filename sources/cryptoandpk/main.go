package main

import (
	"fmt"
	"time"
)

func main() {
	sieve()
}

func sieve() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := sieve_of_eratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		print_sieve(sieve)

		primes := sieve_to_primes(sieve)
		fmt.Println(primes)
	}
}

func fastexp() {
	num := []int64{8, 7, 9, 213}
	pow := []int64{6, 10, 13, 5}
	mod := []int64{10, 101, 283, 1000}
	fmt.Printf("%6s\t%6s\t%6s\t%16s\t%6s\n", "num", "pow", "mod", "n^p", "n^p%m")
	miniter := len(num)
	for i := 0; i < miniter; i++ {
		numpow := fast_exp(num[i], pow[i])
		numpowmod := fast_exp_mod(num[i], pow[i], mod[i])
		fmt.Printf("%6d\t%6d\t%6d\t%16d\t%6d\n", num[i], pow[i], mod[i], numpow, numpowmod)
	}
}

func gcdamdlcm() {
	a := []int64{270, 7469, 55290}
	b := []int64{192, 2464, 115430}

	fmt.Printf("%6s\t%6s\t%6s\t%10s\n", "A", "B", "GCD", "LCM")
	miniter := len(a)
	for i := 0; i < miniter; i++ {
		g, l := gcdlcm(a[i], b[i])
		fmt.Printf("%6d\t%6d\t%6d\t%10d\n", a[i], b[i], g, l)
	}
}
