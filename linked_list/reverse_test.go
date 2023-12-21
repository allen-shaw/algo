package linkedlist

import "testing"

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
