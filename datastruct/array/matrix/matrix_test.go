package matrix

import "sort"

// 00 01 02 03 ... 0n
// 10 11 12 13 ... 1n
// 20 21 22 23 ... 2n
// ...
// m0 m1 m2 m3 ... mn
func diagonalSort(mat [][]int) [][]int {
	dict := make(map[int][]int)
	m, n := len(mat), len(mat[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := j - i
			dict[idx] = append(dict[idx], mat[i][j])
		}
	}

	for i := range dict {
		sort.Ints(dict[i])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := j - i
			mat[i][j] = dict[idx][0]
			dict[idx] = dict[idx][1:]
		}
	}

	return mat
}

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	var shiftGrid = func(grid [][]int) [][]int {
		for i := 0; i < m; i++ {
			temp := grid[i][n-1]
			for j := n - 1; j > 0; j-- {
				grid[i][j] = grid[i][j-1]
			}
			grid[i][0] = temp
		}

		temp := grid[n-1][0]
		for i := n - 1; i > 0; i-- {
			grid[i][0] = grid[i-1][0]
		}
		grid[0][0] = temp
		return grid
	}

	for i := 0; i < k; i++ {
		shiftGrid(grid)
	}
	return grid
}

