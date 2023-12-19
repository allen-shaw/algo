package linkedlist

import (
	"testing"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB

	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	return p1
}

func TestGetIntersectionNode(t *testing.T) {
	a1 := &ListNode{Val: 4}
	a2 := &ListNode{Val: 1}
	b1 := &ListNode{Val: 5}
	b2 := &ListNode{Val: 6}
	b3 := &ListNode{Val: 1}
	c1 := &ListNode{Val: 8}
	c2 := &ListNode{Val: 4}
	c3 := &ListNode{Val: 5}

	a1.Next = a2
	a2.Next = c1
	c1.Next = c2
	c2.Next = c3
	b1.Next = b2
	b2.Next = b3
	b3.Next = c1

	i := getIntersectionNode(a1, b1)
	t.Log(i)
}

func TestGetIntersectionNode2(t *testing.T) {
	a1 := &ListNode{Val: 4}
	a2 := &ListNode{Val: 1}
	b1 := &ListNode{Val: 5}
	b2 := &ListNode{Val: 6}
	b3 := &ListNode{Val: 1}
	c1 := &ListNode{Val: 8}
	c2 := &ListNode{Val: 4}
	c3 := &ListNode{Val: 5}

	a1.Next = a2
	a2.Next = c1
	c1.Next = c2
	c2.Next = c3
	b1.Next = b2
	b2.Next = b3
	b3.Next = c1

	i := getIntersectionNode(a1, b1)
	t.Log(i)
}
