package linkedlist

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
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

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	m := make(map[int]int)
	for p := head; p != nil; p = p.Next {
		m[p.Val]++
	}

	p, q := dummy, head
	for q != nil {
		if m[q.Val] == 1 {
			if p.Next == q {
				p = p.Next
			} else {
				p.Next = q
				p = q
			}
		}
		q = q.Next
	}

	return dummy.Next
}

func Test_deleteDuplicatesUnsorted(t *testing.T) {
	testCases := []struct {
		name   string
		input  *ListNode
		expect []int
	}{
		{"case3", toLinkedList([]int{3, 2, 2, 1, 3, 2, 4}), []int{1, 4}},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			out := deleteDuplicatesUnsorted(c.input)
			assert.Equal(t, c.expect, toArray(out))
		})
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	hp := &Heap{}
	heap.Init(hp)

	for _, l := range lists {
		heap.Push(hp, l)
	}

	for hp.Len() > 0 {
		n := heap.Pop(hp).(*ListNode)
		p.Next = n
		p = p.Next
		next := n.Next
		n.Next = nil
		heap.Push(hp, next)
	}

	return dummy.Next
}

type Heap []*ListNode

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}
func (h *Heap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
