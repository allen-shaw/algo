package backtrack

import "testing"

func TestPartition(t *testing.T) {
	s := "aab"
	ans := partition(s)
	t.Log(ans)
}

func TestSlice(t *testing.T) {
	s := "0123456789"
	t.Log(s[9:10])
}
