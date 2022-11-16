package main

import (
	"fmt"
	"math"
)

var cache = map[int]int{
	1: 0,
}

func main() {
	solve()
}

func solve() {
	N := 10000
	for i := 2; i < N; i++ {
		sum := divisorsSum(i)
		cache[i] = sum
	}

	cnt := 0
	for a1 := 2; a1 < N; a1++ {
		b1 := cache[a1]
		a2 := cache[b1]
		if a1 == a2 && a1 != b1 {
			cnt += a1 + b1
		}
	}
	fmt.Println(cnt / 2)
}

func divisorsSum(n int) int {
	sum := 0
	for i := 1; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			d1 := i
			d2 := n / i
			if d1 == d2 {
				sum += d1
			} else {
				sum += d1 + d2
			}
		}
	}
	sum -= n
	return sum
}
