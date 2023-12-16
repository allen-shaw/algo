package nsums

import "testing"

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1

	ans := threeSumClosest(nums, target)
	t.Log(ans)
}

func TestThreeSumClosest2(t *testing.T) {
	nums := []int{0, 1, 2}
	target := 3

	ans := threeSumClosest(nums, target)
	t.Log(ans)
}
