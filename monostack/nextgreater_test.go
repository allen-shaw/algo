package monostack

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

func TestSlice(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	t.Log(len(nums))

	nums = nums[0 : len(nums)-1]
	t.Log(len(nums))
}

func TestNextGreaterElement2(t *testing.T) {
	// nums1 := []int{1, 2, 1}
	nums1 := []int{1, 2, 3, 4, 3}
	ans := nextGreaterElements(nums1)
	fmt.Println(ans)
}

func TestDailyTemperatures(t *testing.T) {
	// ts := []int{30, 60, 90}
	// ts := []int{30, 40, 50, 60}
	ts := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(ts)
	t.Log(ans)
}
