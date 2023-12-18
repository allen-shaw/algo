package doublepointer

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, n := range nums {
		p.Next = &ListNode{Val: n}
		p = p.Next
	}

	return dummy.Next
}

func (n *ListNode) String() string {
	var s string
	p := n
	for p != nil {
		s += fmt.Sprintf("%d->", p.Val)
		p = p.Next
	}

	return s
}
