package monostack

import "testing"

func TestCanSeePersonCount(t *testing.T) {
	heights := []int{10, 6, 8, 5, 11, 9}
	ans := canSeePersonsCount(heights)
	t.Log(ans)
}
