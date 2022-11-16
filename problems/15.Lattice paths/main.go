package main

import (
	"fmt"

	"github.com/dyusembaevd/project-euler/models"
)

func main() {
	solve()
}

func solve() {
	N := 21
	grid := make([]models.Row, N)
	for i := range grid {
		grid[i] = make([]int, N)
	}

	grid[0][0] = 1

	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < len(grid); row++ {
			if col == 0 && row == 0 {
				continue
			}
			sum := 0
			if col > 0 {
				sum += grid[row][col-1]
			}
			if row > 0 {
				sum += grid[row-1][col]
			}
			grid[row][col] = sum
		}
	}

	fmt.Println(grid[N-1][N-1])
}
