package array

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans := maxSubArray(nums)
	fmt.Println(ans)
}

func TestMaxSubArray2(t *testing.T) {
	nums := []int{1}
	ans := maxSubArray(nums)
	fmt.Println(ans)
}

func TestMaxSubArray3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	ans := maxSubArray(nums)
	fmt.Println(ans)
}
