package matrix

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	martix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(martix, 5))
	fmt.Println(searchMatrix(martix, 20))
}
