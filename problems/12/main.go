package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solve())
}

func solve() int {
	n := 1
	for divisors(triangle(n)) < 500 {
		n++
	}
	return triangle(n)
}

func triangle(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func divisors(n int) int {
	res := 0
	for i := 1; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			d1, d2 := i, n/i
			if d1 == d2 {
				res++
			} else {
				res += 2
			}
		}
	}
	return res
}
