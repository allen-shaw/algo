package linklist

import "testing"

func TestSortList(t *testing.T) {
	head := buildList([]int{4, 2, 1, 3})
	head = sortList(head)
	printList(head)
}

func TestSortList2(t *testing.T) {
	head := buildList([]int{-1, 5, 3, 4, 0})
	head = sortList(head)
	printList(head)
}
