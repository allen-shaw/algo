package linklist

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	head := buildList([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome(head))
}

func TestIsPalindrome2(t *testing.T) {
	head := buildList([]int{1, 2, 3, 2, 1})
	fmt.Println(isPalindrome(head))
}

func TestIsPalindrome3(t *testing.T) {
	head := buildList([]int{1, 2, 4, 1})
	fmt.Println(isPalindrome(head))
}
func TestIsPalindrome4(t *testing.T) {
	head := buildList([]int{1, 2, 4, 3, 1})
	fmt.Println(isPalindrome(head))
}
