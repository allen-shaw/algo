package linkedlist

import "testing"

func TestSortedListToBST(t *testing.T) {
	head := buildLinkedList(-10, -3, 0, 5, 9)
	bst := sortedListToBST(head)
	ans := InorderTreverse(bst)
	t.Log(ans)
}
