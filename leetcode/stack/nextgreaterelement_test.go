package stack

import (
	"fmt"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	ans := nextGreaterElement(nums1, nums2)
	fmt.Println(ans)
}
