package linkedlist

import (
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	h := createLinkedList(1, 2, 3, 4, 5)
	printLinkedList(h)
}

// 反转单链表
func ReverseLindedList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var prev *Node
	cur := head
	next := cur.Next

	for cur != nil {
		cur.Next = prev
		prev = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}

	return prev
}

func TestReverseLindedList(t *testing.T) {
	head := createLinkedList(1, 2, 3, 4, 5)
	printLinkedList(head)

	head = ReverseLindedList(head)
	printLinkedList(head)
}

// 反转双向链表
func ReverseLindedList2(head *Node) *Node {
	return nil
}

// 打印链表公共部分

// 回文
