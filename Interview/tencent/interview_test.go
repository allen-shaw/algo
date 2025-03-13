package tencent

import (
	"container/heap"
	"fmt"
	"testing"
)

type Node struct {
	val  int
	next *Node
}

type hp []*Node

// Len implements heap.Interface.
func (h *hp) Len() int {
	return len(*h)
}

// Less implements heap.Interface.
func (h *hp) Less(i int, j int) bool {
	return (*h)[i].val < (*h)[j].val
}

// Pop implements heap.Interface.
func (h *hp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// Push implements heap.Interface.
func (h *hp) Push(x any) {
	*h = append(*h, x.(*Node))
}

// Swap implements heap.Interface.
func (h *hp) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func mergeLists(lists []*Node) *Node {
	h := hp{}

	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h)

	dummy := &Node{}
	cur := dummy

	for len(h) > 0 {
		n := heap.Pop(&h).(*Node)
		if n.next != nil {
			heap.Push(&h, n.next)
		}
		cur.next = n
		cur = cur.next
	}

	return dummy.next
}

func Test_mergeList(t *testing.T) {
	// [[1,4,5],[1,3,4],[2,6]]
	list1 := newList([]int{1, 4, 5})
	list2 := newList([]int{1, 3, 4})
	list3 := newList([]int{2, 6})

	lists := []*Node{list1, list2, list3}

	head := mergeLists(lists)
	printList(head)
}

func newList(nums []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, n := range nums {
		cur.next = &Node{val: n}
		cur = cur.next
	}
	return dummy.next
}

func printList(head *Node) {
	fmt.Println()
	for cur := head; cur != nil; cur = cur.next {
		fmt.Printf("%v->", cur.val)
	}
	fmt.Println()
}
