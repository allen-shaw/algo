package array

import (
	"fmt"
	"testing"
)

func TestRemoveZero(t *testing.T) {
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums1)
	fmt.Println(nums1)

	nums2 := []int{0}
	moveZeroes(nums2)
	fmt.Println(nums2)
}
