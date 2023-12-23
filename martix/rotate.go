package martix

import (
	"fmt"
	"sort"
)

func NewMatrix(row, col int) [][]int {
	m := make([][]int, row)
	for i := 0; i < row; i++ {
		m[i] = make([]int, col)
	}
	return m
}

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := NewMatrix(m, n)

	for i := 0; i < m*n; i++ {
		idx := (i + k) % (m * n)
		ans[idx/n][idx%n] = grid[i/n][i%n]
	}

	return ans
}

func diagonalSort(mat [][]int) [][]int {
	m := make(map[int][]int, 0)

	// 同意对角线的数放到同一个k,v
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			m[j-i] = append(m[j-i], mat[i][j])
		}
	}

	for _, list := range m {
		sort.Ints(list)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			list := m[j-i]
			mat[i][j] = list[0]
			m[j-i] = m[j-i][1:]
		}
	}

	return mat
}

func transpose(matrix [][]int) [][]int {
	ans := NewMatrix(len(matrix[0]), len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ans[j][i] = matrix[i][j]
		}
	}
	return ans
}

func generateMatrix(n int) [][]int {
	ans := NewMatrix(n, n)
	num := 0

	up, down, left, right := 0, n-1, 0, n-1
	for up <= down && left <= right {
		if up <= down {
			for i := left; i <= right; i++ {
				num++
				ans[up][i] = num
			}
		}
		up++

		if left <= right {
			for i := up; i <= down; i++ {
				num++
				ans[i][right] = num
			}
		}
		right--

		if up <= down {
			for i := right; i >= left; i-- {
				num++
				ans[down][i] = num
			}
		}
		down--

		if left <= right {
			for i := down; i >= up; i-- {
				num++
				ans[i][left] = num
			}
		}
		left++
	}

	return ans
}

func spiralOrder(matrix [][]int) []int {
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	ans := make([]int, 0, len(matrix)*len(matrix[0]))

	for len(ans) < len(matrix)*len(matrix[0]) {
		if up <= down {
			for i := left; i <= right; i++ {
				ans = append(ans, matrix[up][i])
			}
		}
		up++

		fmt.Println("up:", up, ans)

		if left <= right {
			for i := up; i <= down; i++ {
				ans = append(ans, matrix[i][right])
			}
		}
		right--
		fmt.Println("right:", right, ans)

		if up <= down {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[down][i])
			}
		}
		down--
		fmt.Println("down:", down, ans)

		if left <= right {
			for i := down; i >= up; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		left++
		fmt.Println("left:", left, ans)
	}

	return ans
}

func rotate(matrix [][]int) {
	diagonalSymmetry(matrix)
	axisymmetric(matrix)
}

func diagonalSymmetry(martix [][]int) {
	for i := 0; i < len(martix); i++ {
		for j := i + 1; j < len(martix[0]); j++ {
			martix[i][j], martix[j][i] = martix[j][i], martix[i][j]
		}
	}
}

func axisymmetric(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		left, right := 0, len(matrix[0])-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
