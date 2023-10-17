package array

import (
	"fmt"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	ans1 := productExceptSelf(nums1)
	fmt.Println(ans1)

	nums2 := []int{-1, 1, 0, -3, 3}
	ans2 := productExceptSelf(nums2)
	fmt.Println(ans2)

}
