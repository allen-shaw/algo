package nsums

import "testing"

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0

	ans := fourSum(nums, target)
	t.Log(ans)
}
