package other

import "testing"

func TestTrap(t *testing.T) {
	// h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	h := []int{4, 2, 0, 3, 2, 5}
	ans := trap(h)
	t.Log(ans)
}

func TestPreCalcMaxLeft(t *testing.T) {
	h := []int{4, 2, 0, 3, 2, 5}
	ans := preCalcMaxLeft(h)
	t.Log(ans)
}

func TestTrap2(t *testing.T) {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// h := []int{4, 2, 0, 3, 2, 5}
	ans := trap2(h)
	t.Log(ans)
}

func TestTrap3(t *testing.T) {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// h := []int{4, 2, 0, 3, 2, 5}
	ans := trap3(h)
	t.Log(ans)
}
