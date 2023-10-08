package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	ok := searchMatrix(matrix, target)
	fmt.Println(ok)
}

func TestSearchMatrix2(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 13
	ok := searchMatrix(matrix, target)
	fmt.Println(ok)
}

func TestValue(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	id := 1
	v := value(matrix, id)
	fmt.Println(v)
}
