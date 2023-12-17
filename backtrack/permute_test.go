package backtrack

import "testing"

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := permute(nums)
	t.Log(ans)
}
