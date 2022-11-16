package main

import "fmt"

var cache = map[int]int{}

func main() {
	max := 0
	maxn := 0
	for i := 1; i < 1000000; i++ {
		n, cnt := collatzSequence(i)
		if cnt > max {
			max = cnt
			maxn = n
		}
	}
	fmt.Println(maxn, max)
}

func collatzSequence(n int) (int, int) {
	start := n
	count := 1
	for n > 1 {
		if val, ok := cache[n]; ok {
			count += val
			break
		}
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		count++
	}
	cache[n] = count
	return start, count
}
