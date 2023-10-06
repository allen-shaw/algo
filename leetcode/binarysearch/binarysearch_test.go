package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	idx := binarySearch(nums, target)
	fmt.Println(idx)
}

func TestBinarySearch2(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	idx := binarySearch(nums, target)
	fmt.Println(idx)
}

func TestBSLeftBound(t *testing.T) {
	nums := []int{1, 3, 3, 5, 5, 6}
	target := 5
	idx := binarySearch(nums, target)
	idx2 := bsLeftBound(nums, target)
	fmt.Println(idx, idx2)
}

func TestBSLeftBound2(t *testing.T) {
	nums := []int{1, 3, 3, 5, 5, 6}
	target := 2
	idx := binarySearch(nums, target)
	idx2 := bsLeftBound(nums, target)
	fmt.Println(idx, idx2)
}

func TestBSRightBound(t *testing.T) {
	nums := []int{1, 3, 3, 5, 5, 5, 6}
	target := 5
	idx := binarySearch(nums, target)
	idx2 := bsLeftBound(nums, target)
	idx3 := bsRightBound(nums, target)
	fmt.Println(idx, idx2, idx3)
}

func TestBSRightBound2(t *testing.T) {
	nums := []int{1, 3, 3, 5, 5, 6}
	target := 2
	idx := binarySearch(nums, target)
	idx2 := bsLeftBound(nums, target)
	idx3 := bsRightBound(nums, target)
	fmt.Println(idx, idx2, idx3)
}
