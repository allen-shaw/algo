package doublepointer

import "testing"

func TestRemoveZero(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	removeZeroes(nums)
	t.Log(nums)
}
