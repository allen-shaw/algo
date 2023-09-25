package leetcode

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	a, b := buildIntersectedList([]int{4, 1, 8, 4, 5}, []int{5, 6, 1, 8, 4, 5}, 2, 3)
	printListNode(getIntersectionNode(a, b))
}
