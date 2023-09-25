package leetcode

import (
	"container/heap"
)

type Heap struct {
	data []*ListNode
}

// Len implements heap.Interface.
func (h *Heap) Len() int {
	return len(h.data)
}

// Less implements heap.Interface.
func (h *Heap) Less(i int, j int) bool {
	return h.data[i].Val < h.data[j].Val
}

// Pop implements heap.Interface.
func (h *Heap) Pop() any {
	n := len(h.data)
	last := h.data[n-1]
	h.data = h.data[0 : n-1]
	return last
}

// Push implements heap.Interface.
func (h *Heap) Push(x any) {
	h.data = append(h.data, x.(*ListNode))
}

// Swap implements heap.Interface.
func (h *Heap) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &Heap{data: make([]*ListNode, 0)}
	heap.Init(h)

	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	dummy := &ListNode{}
	p := dummy
	for h.Len() > 0 {
		n := heap.Pop(h).(*ListNode)
		p.Next = n
		p = p.Next
		if n.Next != nil {
			heap.Push(h, n.Next)
		}
	}

	return dummy.Next
}
