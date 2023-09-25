package leetcode

import "testing"

func TestPartition(t *testing.T) {
	h := buildList([]int{1, 4, 3, 2, 5, 2})
	h = partition(h, 3)
	printList(h)
}
