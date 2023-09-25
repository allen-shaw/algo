package leetcode

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	h := buildList([]int{1, 2, 3, 4, 5})
	h = removeNthFromEnd(h, 2)
	printList(h)
}

func TestRemoveNthFromEnd2(t *testing.T) {
	h := buildList([]int{1})
	h = removeNthFromEnd(h, 1)
	printList(h)
}

func TestRemoveNthFromEnd3(t *testing.T) {
	h := buildList([]int{1, 2})
	h = removeNthFromEnd(h, 1)
	printList(h)
}
