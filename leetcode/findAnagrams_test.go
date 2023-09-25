package leetcode

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	ans := findAnagrams(s, p)
	fmt.Println(ans)
}

func TestFindAnagrams2(t *testing.T) {
	s := "abab"
	p := "ab"
	ans := findAnagrams(s, p)
	fmt.Println(ans)
}
