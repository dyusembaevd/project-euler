package models

import (
	"bufio"
	"os"

	"github.com/dyusembaevd/project-euler/pkg"
)

type Grid []Row

type Row []int

func ReadGrid() Grid {
	in := bufio.NewReader(os.Stdin)
	rows := pkg.ReadInt(in)

	grid := make([]Row, 0, rows)

	for i := 0; i < rows; i++ {
		arr := pkg.ReadArrInt(in)
		grid = append(grid, arr)
	}

	return grid
}
