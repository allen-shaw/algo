package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums1, 9))

	nums2 := []int{2, 3, 4}
	fmt.Println(twoSum(nums2, 6))

	nums3 := []int{-1, 0}
	fmt.Println(twoSum(nums3, -1))
}
