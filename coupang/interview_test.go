package coupang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 最长子序列
// 8 5 3 6 2 1 4 3
func longestSubsequence(nums []int) int {
	m := make(map[int]struct{})

	for _, n := range nums {
		m[n] = struct{}{}
	}

	ans := 0
	for _, n := range nums {
		if _, ok := m[n-1]; ok {
			continue
		}

		length := 0
		for k := n; ; k++ {
			if _, ok := m[k]; ok {
				length++
			} else {
				ans = max(ans, length)
				break
			}
		}
	}

	return ans
}

func TestLS(t *testing.T) {
	var testsets = []struct {
		nums []int
		out  int
	}{
		{[]int{8, 5, 3, 6, 2, 1, 4, 3}, 6},
	}

	for _, tt := range testsets {
		t.Run("ls", func(t *testing.T) {
			ans := longestSubsequence(tt.nums)
			assert.Equal(t, tt.out, ans)
		})
	}

}

// n * n
// [1,2,3]		[7,4,1]
// [4,5,6]		[8,5,2]
// [7,8,9]		[9,6,3]
func rotateMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		j, k := 0, len(matrix[i])-1
		for j < k {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
			j++
			k--
		}
	}
}

func TestRotateMatrix(t *testing.T) {
	var testsets = []struct {
		matrix [][]int
		out    [][]int
	}{
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
		},
	}

	for _, tt := range testsets {
		t.Run("rotateMatrix", func(t *testing.T) {
			rotateMatrix(tt.matrix)
			assert.Equal(t, tt.out, tt.matrix)
		})
	}
}
