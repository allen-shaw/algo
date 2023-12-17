package doublepointer

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	ans := longestPalindrome(s)
	t.Log(ans)
}

func TestSlice(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := nums[3:6] // 3,4,5
	t.Log(s)
}

func TestLongestPalindrome2(t *testing.T) {
	s := "a"
	ans := longestPalindrome(s)
	t.Log(ans)
}


