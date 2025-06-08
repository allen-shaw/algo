package zoom

import (
	"fmt"
	"testing"
)

func zigzagTraverse(matrix [][]int) {
	n := len(matrix)

	for i := range n {
		if i%2 == 0 {
			x, y := i, 0
			for x >= 0 && y <= i {
				fmt.Printf("%v ", matrix[x][y])
				x--
				y++
			}
			println()
		} else {
			x, y := 0, i
			for x <= i && y >= 0 {
				fmt.Printf("%v ", matrix[x][y])
				x++
				y--
			}
			println()
		}
	}
}

func zigzagTraverse2(matrix [][]int) {
	row, column := len(matrix), len(matrix[0])
	i, j := 0, 0
	n := row + column - 1

	for k := range n {
		if k%2 == 0 {
			for i >= 0 && j < column {
				fmt.Printf("%v ", matrix[i][j])
				i--
				j++
			}
			i++
			if j >= column {
				j = column - 1
				i++
			}
			fmt.Println()
		} else {
			for i < row && j >= 0 {
				fmt.Printf("%v ", matrix[i][j])
				i++
				j--
			}
			j++
			if i >= row {
				i = row - 1
				j++
			}
			fmt.Println()
		}
	}
}

func Test_zigzagTraverse(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		// {13, 14, 15, 16, 17},
	}
	zigzagTraverse2(matrix)
}
