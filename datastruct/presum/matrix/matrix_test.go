package matrix

import (
	"fmt"
	"testing"
)

type NumMatrix struct {
	presum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	presum := make([][]int, m+1)
	for i := range presum {
		presum[i] = make([]int, n+1)
	}
	for i := 1; i < len(presum); i++ {
		for j := 1; j < len(presum[i]); j++ {
			presum[i][j] = presum[i-1][j] + presum[i][j-1] - presum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	return NumMatrix{presum: presum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.presum[row2+1][col2+1] - this.presum[row1][col2+1] - this.presum[row2+1][col1] + this.presum[row1][col1]
}

// 3, 0, 1, 4, 2
// 5, 6, 3, 2, 1
// 1, 2, 0, 1, 5
// 4, 1, 0, 1, 7
// 1, 0, 3, 0, 5
func TestNumMatrix(t *testing.T) {
	nums := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	nm := Constructor(nums)
	fmt.Println(nm.presum)
	fmt.Println(nm.SumRegion(2, 1, 4, 3)) // 17 38
}

// 3  3  4   8   10
// 8  14 18  24  27
// 9  17 21  28  36
// 13 22 26  34  49
// 14 23 30  38  58
