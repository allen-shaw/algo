package doublepointer

import "testing"

func TestRemoveElements(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	ans := removeElement(nums, val)
	t.Log(ans)
}

func TestRemoveElements2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	ans := removeElement(nums, val)
	t.Log(ans)
}
