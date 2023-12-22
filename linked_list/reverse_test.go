package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := buildLinkedList(1, 2, 3, 4, 5)
	head = reverseList(head)
	printLinkedList(head)
}

func TestReverseBetween(t *testing.T) {
	head := buildLinkedList(1, 2, 3, 4, 5)
	head = reverseBetween(head, 2, 4)
	printLinkedList(head)
}

func TestReverseBetween2(t *testing.T) {
	head := buildLinkedList(3, 5)
	head = reverseBetween(head, 1, 1)
	printLinkedList(head)
}

func TestReverseBetween22(t *testing.T) {
	head := buildLinkedList(1, 2, 3, 4, 5)
	head = reverseBetween2(head, 2, 4)
	printLinkedList(head)
}

func TestReverseKGroup(t *testing.T) {
	head := buildLinkedList(1, 2, 3, 4, 5)
	head = reverseKGroup(head, 2)

	// p := head
	// for i := 0; i < 6; i++ {
	// 	fmt.Println(p.Val)
	// 	p = p.Next
	// }

	printLinkedList(head)
}

func TestReverseTo(t *testing.T) {
	head := buildLinkedList(1, 2, 3)
	tail := head.Next.Next

	newHead := reverseTo(head, tail)
	p := newHead
	for i := 0; i < 6; i++ {
		if p == nil {
			break
		}
		fmt.Println(p.Val)
		p = p.Next
	}

}

func TestSwapPairs(t *testing.T) {
	head := buildLinkedList(1, 2, 3, 4, 5)
	head = swapPairs(head)
	printLinkedList(head)
}

func TestIsPalindrome(t *testing.T) {
	head := buildLinkedList(1, 2, 3, 2, 1)
	ans := isPalindrome(head)
	t.Log(ans)
}
