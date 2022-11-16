package main

import (
	"fmt"

	"github.com/dyusembaevd/project-euler/models"
)

const (
	shift = 4
)

func main() {
	grid := models.ReadGrid()
	fmt.Println(up(grid))
	fmt.Println(right(grid))
	fmt.Println(diag1(grid))
	fmt.Println(diag2(grid))
}

func diag2(grid models.Grid) int {
	res := 0
	for row := shift - 1; row < len(grid.Rows); row++ {
		for col := 0; col < len(grid.Rows[0])-shift; col++ {
			mult := 1
			for i := 0; i < shift; i++ {
				mult = mult * grid.Rows[row-i][col+i]
			}
			res = max(res, mult)
		}
	}
	return res
}

func diag1(grid models.Grid) int {
	res := 0
	for row := 0; row < len(grid.Rows)-shift; row++ {
		for col := 0; col < len(grid.Rows[0])-shift; col++ {
			mult := 1
			for i := 0; i < shift; i++ {
				mult = mult * grid.Rows[row+i][col+i]
			}
			res = max(res, mult)
		}
	}
	return res
}

func right(grid models.Grid) int {
	res := 0
	for row := 0; row < len(grid.Rows); row++ {
		for col := 0; col < len(grid.Rows[0])-shift; col++ {
			mult := 1
			for i := col; i < col+shift; i++ {
				mult = mult * grid.Rows[row][i]
			}
			res = max(res, mult)
		}
	}
	return res
}

func up(grid models.Grid) int {
	res := 0
	for col := 0; col < len(grid.Rows[0]); col++ {
		for row := 0; row < len(grid.Rows)-shift; row++ {
			mult := 1
			for i := row; i < row+shift; i++ {
				mult = mult * grid.Rows[i][col]
			}
			res = max(res, mult)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
