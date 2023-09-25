package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
func TestMaxArea2(t *testing.T) {
	nums := []int{1, 1}
	fmt.Println(maxArea(nums))
}
