package linklist

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	head := buildList([]int{1, 1, 2, 3, 3})
	head = deleteDuplicates(head)
	printList(head)
}
