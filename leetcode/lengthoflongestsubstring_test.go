package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring2(str))
	fmt.Println(lengthOfLongestSubstring2("bbbbb"))
	fmt.Println(lengthOfLongestSubstring2("pwwkew"))
}
