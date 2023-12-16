package nsums

import "testing"

func TestThreeSumSmaller(t *testing.T) {
	nums := []int{-2, 0, 1, 3}
	target := 2
	ans := threeSumSmaller(nums, target)
	t.Log(ans)
}
