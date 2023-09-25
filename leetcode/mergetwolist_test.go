package leetcode

import "testing"

func TestMergeTwoLists(t *testing.T) {
	// l1 = [1,2,4], l2 = [1,3,4]
	h1 := buildList([]int{1, 2, 4})
	h2 := buildList([]int{1, 3, 4})

	h := mergeTwoLists(h1, h2)
	printList(h)
}

func TestMergeTwoLists2(t *testing.T) {
	// l1 = [1,2,4], l2 = [1,3,4]
	h1 := buildList([]int{})
	h2 := buildList([]int{})

	h := mergeTwoLists(h1, h2)
	printList(h)
}

func TestMergeTwoLists3(t *testing.T) {
	// l1 = [1,2,4], l2 = [1,3,4]
	h1 := buildList([]int{})
	h2 := buildList([]int{0})

	h := mergeTwoLists(h1, h2)
	printList(h)
}
