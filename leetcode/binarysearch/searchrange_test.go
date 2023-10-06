package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	idxs := searchRange(nums, target)
	fmt.Println(idxs)
}

func TestSearchRange2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	idxs := searchRange(nums, target)
	fmt.Println(idxs)
}

func TestSearchRange3(t *testing.T) {
	nums := []int{}
	target := 0
	idxs := searchRange(nums, target)
	fmt.Println(idxs)
}
