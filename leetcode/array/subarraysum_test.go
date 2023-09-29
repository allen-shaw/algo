package array

import (
	"fmt"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 2

	fmt.Println(subarraySum(nums, k))
}

func TestSubarraySum2(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 3

	fmt.Println(subarraySum(nums, k))
}
