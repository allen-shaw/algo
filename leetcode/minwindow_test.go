package leetcode

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	str := "ADOBECODEBANC"
	st := "ABC"
	substr := minWindow(str, st)
	fmt.Println(substr)
}
