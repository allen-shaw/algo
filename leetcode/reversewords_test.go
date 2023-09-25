package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s1 := "the sky is blue"
	s2 := "  hello world  "
	s3 := "a good   example"

	fmt.Println(reverseWords(s1))
	fmt.Println(reverseWords(s2))
	fmt.Println(reverseWords(s3))
}
