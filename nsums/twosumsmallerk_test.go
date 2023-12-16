package nsums

import "testing"

func TestTwoSumSmallerThanK(t *testing.T) {
	nums := []int{34, 23, 1, 24, 75, 33, 54, 8}
	k := 60

	ans := twoSumSmallerThanK(nums, k)
	t.Log(ans)
}

func TestTwoSumSmallerThanK2(t *testing.T) {
	nums := []int{10, 20, 30}
	k := 15

	ans := twoSumSmallerThanK(nums, k)
	t.Log(ans)
}
