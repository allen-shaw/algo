package recursion

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Printf("%v ", head.Val)
	PrintList(head.Next)
}

func AddLast(head *ListNode, val int) *ListNode {
	if head == nil {
		return &ListNode{Val: val}
	}

	head.Next = AddLast(head.Next, val)
	return head
}

func RemoveLast(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	head.Next = RemoveLast(head.Next)
	return head
}

func Size(head *ListNode) int {
	if head == nil {
		return 0
	}
	return Size(head.Next) + 1
}

func PrintListReserved(head *ListNode) {
	if head == nil {
		return
	}

	PrintListReserved(head.Next)
	fmt.Printf("%v ", head.Val)
}

func buildList(size int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for i := 0; i < size; i++ {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}
	return dummy.Next
}

func TestAddList(t *testing.T) {
	head := buildList(5)
	PrintList(head)
	fmt.Println()
	PrintListReserved(head)
	fmt.Println()
	fmt.Println(Size(head))

	RemoveLast(head)
	PrintList(head)
	fmt.Println()
	AddLast(head, 6)
	PrintList(head)
}
