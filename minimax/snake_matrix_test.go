package minimax

import (
	"fmt"
	"testing"
)

//m n [2,10]
// 3 4   4 3

// 1 2 6  7
// 3 5 8  11
// 4 9 10 12

func SnakeMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	f := true
	matrix[0][0] = 1

	val := 1
	i, j := 0, 0
	for val < m*n {
		if f {
			x, y := i, j+1
			if y == n {
				y = n - 1
				x = x + 1
			}
			for x >= 0 && x < m && y >= 0 && y < n {
				val++
				matrix[x][y] = val
				x++
				y--
			}
			x--
			y++
			i, j = x, y
			f = !f
		}
		if !f {
			x, y := i+1, j
			if x == m {
				x = m - 1
				y = y + 1
			}
			for x >= 0 && x < m && y >= 0 && y < n {
				val++
				matrix[x][y] = val
				x--
				y++
			}
			x++
			y--
			i, j = x, y
			f = !f
		}
	}

	return matrix
}

func TestSnakeMatrix(t *testing.T) {
	ans := SnakeMatrix(3, 4)
	fmt.Println(ans)
	ans2 := SnakeMatrix(4, 3)
	fmt.Println(ans2)
}
