package nsums

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)

	t.Log(res)
}

func TestTwoSumTarget(t *testing.T) {
	nums := []int{0, 1, 1, 2}
	sort.Ints(nums)
	t.Log(nums)

	res := twoSumTarget(nums, 2)
	t.Log(res)
}

func TestThreeSumTarget(t *testing.T) {
	nums := []int{-2, 0, 1, 1, 2}
	sort.Ints(nums)
	t.Log(nums)

	res := threeSumTarget(nums, 0)
	t.Log(res)

}
