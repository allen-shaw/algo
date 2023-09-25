package leetcode

import (
	"testing"
)

func TestDetechCycle(t *testing.T) {
	h1 := buildCycleLink([]int{3, 2, 0, -4}, 1)
	printListNode(detectCycle(h1))
}
