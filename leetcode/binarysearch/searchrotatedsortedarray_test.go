package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearchRotatedSortedArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	ans := searchRotatedSortedArray(nums, target)
	fmt.Println(ans)
}

func TestSearchRotatedSortedArray2(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	ans := searchRotatedSortedArray(nums, target)
	fmt.Println(ans)
}

func TestSearchRotatedSortedArray3(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 1, 2, 3}
	target := 8
	ans := searchRotatedSortedArray(nums, target)
	fmt.Println(ans)
}
