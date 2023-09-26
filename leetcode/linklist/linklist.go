package linklist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(arr []int) *ListNode {
	dummy := &ListNode{}
	p := dummy

	for _, n := range arr {
		p.Next = &ListNode{Val: n}
		p = p.Next
	}

	return dummy.Next
}

func buildCycleLink(arr []int, pos int) *ListNode {
	dummy := &ListNode{}
	var cycleNode *ListNode
	p := dummy

	for i, n := range arr {
		p.Next = &ListNode{Val: n}
		p = p.Next
		if i == pos {
			cycleNode = p
		}
	}

	p.Next = cycleNode
	return dummy.Next
}

func printListNode(n *ListNode) {
	if n == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(n.Val)
}

func buildIntersectedList(arr1, arr2 []int, pos1, pos2 int) (*ListNode, *ListNode) {
	var insertNode *ListNode
	dummyA, dummyB := &ListNode{}, &ListNode{}
	p1, p2 := dummyA, dummyB

	for i, n := range arr1 {
		p1.Next = &ListNode{Val: n}
		p1 = p1.Next
		if i == pos1 {
			insertNode = p1
		}
	}

	for i, n := range arr2 {
		p2.Next = &ListNode{Val: n}
		p2 = p2.Next

		if i == pos2-1 {
			p2.Next = insertNode
			break
		}
	}

	return dummyA.Next, dummyB.Next
}

func printList(h *ListNode) {
	if h == nil {
		fmt.Println("nil")
		return
	}

	p := h
	for p.Next != nil {
		fmt.Printf("%v -> ", p.Val)
		p = p.Next
	}
	fmt.Printf("%v", p.Val)
	fmt.Println()
}
