package martix

import (
	"fmt"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 1
	ans := shiftGrid(grid, k)
	fmt.Println(ans)
}

func TestDiagonalSort(t *testing.T) {
	mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	mat = diagonalSort(mat)
	fmt.Println(mat)
}

func TestTranspose(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	matrix = transpose(matrix)
	fmt.Println(matrix)
}

func TestGenerateMatrix(t *testing.T) {
	matrix := generateMatrix(4)
	fmt.Println(matrix)
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	ans := spiralOrder(matrix)
	fmt.Println(ans)
	// 1,2,3,4,8,12,11,10,9,5,6,7
}

func TestRotate(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}
