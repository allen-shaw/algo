package matrix

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

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

// 位于 grid[i][j]（j < n - 1）的元素将会移动到 grid[i][j + 1]
// 位于 grid[i][n - 1] 的元素将会移动到 grid[i + 1][0]
// 位于 grid[m - 1][n - 1] 的元素将会移动到 grid[0][0]
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	var shift = func(grid [][]int) {
		temp := make([]int, m)
		for i := 0; i < len(grid); i++ {
			temp[i] = grid[i][n-1]
			for j := n - 2; j >= 0; j-- {
				grid[i][j+1] = grid[i][j]
			}
		}
		grid[0][0] = temp[m-1]
		for i := 0; i < m-1; i++ {
			grid[i+1][0] = temp[i]
		}
	}

	for i := 0; i < k; i++ {
		shift(grid)
	}
	return grid
}

func Test_shiftGrid(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	grid = shiftGrid(grid, 9)
	fmt.Println(grid)
}

func shiftGrid2(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k = k % (m * n)

	var index = func(i, j, k int) (int, int) {
		a, b := k/n, k%n
		return (i - a + m) % m, (j - b + n) % n
	}

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x, y := index(i, j, k)
			fmt.Printf("(%v,%v) => (%v,%v)\n", i, j, x, y)
			ans[i][j] = grid[x][y]
		}
	}
	return ans
}

func Test_shiftGrid2(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	grid = shiftGrid2(grid, 1)
	fmt.Println(grid)
}

func Test_shift(t *testing.T) {
	m, n := 3, 3
	var index = func(i, j, k int) (int, int) {
		a, b := k/n, k%n
		fmt.Printf("a=%v, b=%v\n", a, b)
		return (i - a + m) % m, (j - b + n) % n
	}
	i, j, k := 0, 0, 1
	x, y := index(i, j, k)
	fmt.Printf("(%v,%v) => (%v,%v)\n", i, j, x, y)

}

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans[i][j] = matrix[j][i]
		}
	}

	return ans
}

func Test_transpose(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	m := transpose(matrix)
	fmt.Println(m)
}

func longestCommonPrefix(strs []string) string {
	ss := make([]byte, 0)
	size := math.MaxInt
	for _, str := range strs {
		size = min(size, len(str))
	}

	for i := 0; i < size; i++ {
		var temp byte
		for j, str := range strs {
			if j == 0 {
				temp = str[i]
			} else {
				if str[i] != temp {
					return string(ss)
				}
			}
		}
		ss = append(ss, temp)
	}
	return string(ss)
}

func Test_longestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	ans := longestCommonPrefix(strs)
	fmt.Println(ans)
}
