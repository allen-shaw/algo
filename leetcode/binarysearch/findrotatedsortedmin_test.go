package binarysearch

import (
	"fmt"
	"testing"
)

func TestFindRotatedSortedMin(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	min := findRotatedSortedMin(nums)
	fmt.Println(min)
}
