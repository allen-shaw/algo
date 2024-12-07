package reverse

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var temp *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 0 {
		temp = head.Next
		return head
	}

	h := reverseN(head.Next, n-1)

	head.Next.Next = head
	head.Next = temp

	return h
}

func TestReverseN(t *testing.T) {
	head := toLinkedList([]int{1, 2, 3, 4, 5, 6})
	newHead := reverseN(head, 3)
	fmt.Println(toArray(newHead))
}

func toLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, num := range nums {
		p.Next = &ListNode{Val: num}
		p = p.Next
	}
	return dummy.Next
}

func toArray(head *ListNode) []int {
	arr := make([]int, 0)
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}
	return arr
}

// d 1, 2, 3, 4, 5    left:2 right:4
//
//	p
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	for i := 1; i < left; i++ {
		p = p.Next
	}

	var next *ListNode

	var reverseN func(head *ListNode, n int) *ListNode
	reverseN = func(head *ListNode, n int) *ListNode {
		if n == 0 {
			next = head.Next
			return head
		}

		last := reverseN(head.Next, n-1)
		head.Next.Next = head
		head.Next = next
		return last
	}

	p.Next = reverseN(p.Next, right-left)
	return dummy.Next
}

func Test_reverseBetween(t *testing.T) {
	head := toLinkedList([]int{1, 2, 3, 4, 5})
	head = reverseBetween(head, 2, 4)
	fmt.Println(toArray(head))
}
