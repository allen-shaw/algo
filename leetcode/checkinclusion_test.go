package leetcode

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func TestCheckInclusion2(t *testing.T) {
	s1 := "ab"
	s2 := "eidboaoo"
	fmt.Println(checkInclusion(s1, s2))
}

func TestCheckInclusion3(t *testing.T) {
	s1 := "hello"
	s2 := "ooolleoooleh"
	fmt.Println(checkInclusion(s1, s2))
}
