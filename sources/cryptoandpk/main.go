package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	a := []int{270, 7469, 55290}
	b := []int{192, 2464, 115430}

	fmt.Printf("%6s\t%6s\t%6s\t%10s\n", "A", "B", "GCD", "LCM")
	miniter := min(len(a), len(b))
	for i := 0; i < miniter; i++ {
		g, l := gcdlcm(a[i], b[i])
		fmt.Printf("%6d\t%6d\t%6d\t%10d\n", a[i], b[i], g, l)
	}
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
