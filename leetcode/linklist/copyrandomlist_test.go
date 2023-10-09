package linklist

import (
	"fmt"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	head := buildRandomList([][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
	printRandomList(head)
	newHead := copyRandomList(head)
	printRandomList(newHead)
}

func buildRandomList(arr [][]int) *Node {
	prefix := &Node{}
	cur := prefix
	nodes := make([]*Node, len(arr))

	for i, node := range arr {
		n := &Node{Val: node[0]}
		nodes[i] = n
		cur.Next = n
		cur = cur.Next
	}

	for i, node := range arr {
		if node[1] != -1 {
			nodes[i].Random = nodes[node[1]]
		}
	}

	return prefix.Next
}

func printRandomList(head *Node) {
	cur := head
	for ; cur.Next != nil; cur = cur.Next {
		fmt.Printf("[%v, %v] ->", cur.Val, cur.Random)
	}
	fmt.Printf("[%v, %v]", cur.Val, cur.Random)
	fmt.Println()
}
