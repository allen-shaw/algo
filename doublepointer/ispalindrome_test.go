package doublepointer

import "testing"

func TestIspalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	ans := isPalindrome(s)
	t.Log(ans)
}

func TestIspalindrome2(t *testing.T) {
	s := "race a car"
	ans := isPalindrome(s)
	t.Log(ans)
}

func TestIspalindrome3(t *testing.T) {
	s := " "
	ans := isPalindrome(s)
	t.Log(ans)
}
