package doublepointer

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	ans := removeDuplicates(nums)
	t.Log(ans)
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ans := removeDuplicates(nums)
	t.Log(ans)
}
