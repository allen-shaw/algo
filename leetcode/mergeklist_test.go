package leetcode

import "testing"

func TestMergeKLists(t *testing.T) {
	h1 := buildList([]int{1, 4, 5})
	h2 := buildList([]int{1, 3, 4})
	h3 := buildList([]int{2, 6})

	h := mergeKLists([]*ListNode{h1, h2, h3})
	printList(h)

}

func TestMergeKLists2(t *testing.T) {
	h1 := buildList([]int{})

	h := mergeKLists([]*ListNode{h1})
	printList(h)

}

func TestMergeKLists3(t *testing.T) {
	h := mergeKLists([]*ListNode{})
	printList(h)

}
