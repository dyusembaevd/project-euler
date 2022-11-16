package models

import (
	"bufio"
	"os"

	"github.com/dyusembaevd/project-euler/pkg"
)

type Grid struct {
	Rows []Row
}

type Row []int

func ReadGrid() Grid {
	in := bufio.NewReader(os.Stdin)
	rows := pkg.ReadInt(in)

	grid := Grid{
		Rows: make([]Row, 0, rows),
	}

	for i := 0; i < rows; i++ {
		arr := pkg.ReadArrInt(in)
		grid.Rows = append(grid.Rows, arr)
	}

	return grid
}
