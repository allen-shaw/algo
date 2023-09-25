package leetcode

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ans := spiralOrder(matrix)
	fmt.Println(ans)

	// [1,2,3,6,9,8,7,4,5]
	// [1 2 3 6 9 8 7 4 5]
}
