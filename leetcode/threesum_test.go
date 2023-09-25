package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	arr := []int{-1, 0, 1, 2, -1, -4}
	out := threeSum(arr, 0)
	for _, o := range out {
		fmt.Println(o)
	}
}
func TestThreeSum2(t *testing.T) {
	arr := []int{0, 1, 1}
	out := threeSum(arr, 0)
	for _, o := range out {
		fmt.Println(o)
	}
}

func TestThreeSum3(t *testing.T) {
	arr := []int{0, 0, 0}
	out := threeSum(arr, 0)
	for _, o := range out {
		fmt.Println(o)
	}
}

func TestTwoSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// out := twoSum(nums, -1, 0)
	// for _, o := range out {
	// 	fmt.Println(o)
	// }

	// out = twoSum(nums, -2, 0)
	// for _, o := range out {
	// 	fmt.Println(o)
	// }

	out := twoSum(nums, 1, 1)
	for _, o := range out {
		fmt.Println(o)
	}
}
