package leetcode

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	h1 := buildCycleLink([]int{3, 2, 0, -4}, 1)
	fmt.Println(hasCycle(h1))
}
