package array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func TestRotate2(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	rotate(nums, k)
	fmt.Println(nums)
}
