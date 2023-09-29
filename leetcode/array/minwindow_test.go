package array

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	st := "ABC"
	ans := minWindow(s, st)
	fmt.Println(ans)
}
