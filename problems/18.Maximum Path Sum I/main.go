package main

import (
	"fmt"

	"github.com/dyusembaevd/project-euler/models"
)

func main() {
	solve()
}

func solve() {
	grid := models.ReadGrid()
	for row := 1; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			left := col - 1
			right := col
			if col == 0 {
				grid[row][col] += grid[row-1][col]
			} else if col == len(grid[row])-1 {
				grid[row][col] += grid[row-1][col-1]
			} else {
				grid[row][col] += max(grid[row-1][left], grid[row-1][right])
			}
		}
	}
	for i := 0; i < len(grid[len(grid)-1]); i++ {
		fmt.Println(grid[len(grid)-1][i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
