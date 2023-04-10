package linkedlist

import "fmt"

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

func createLinkedList(nums ...int) *Node {
	if len(nums) == 0 {
		return nil
	}

	head := &Node{Value: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &Node{Value: nums[i]}
		cur = cur.Next
	}

	return head
}

func printLinkedList(h *Node) {
	cur := h
	for cur != nil {
		fmt.Printf("%v->", cur.Value)
		cur = cur.Next
	}
	fmt.Println("nil")
}
