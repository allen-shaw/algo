package errorset

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type hp []*ListNode

func (hp *hp) Len() int {
	return len(*hp)
}

func (hp *hp) Less(i, j int) bool {
	return (*hp)[i].Val < (*hp)[j].Val
}

func (hp *hp) Push(x any) {
	(*hp) = append((*hp), x.(*ListNode))
}

func (hp *hp) Pop() any {
	idx := hp.Len() - 1
	last := (*hp)[idx]
	(*hp) = (*hp)[:idx]
	return last
}

func (hp *hp) Swap(i, j int) {
	(*hp)[i], (*hp)[j] = (*hp)[j], (*hp)[i]
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	h := make(hp, 0)
	heap.Init(&h)

	for _, head := range lists {
		heap.Push(&h, head)
	}
	for h.Len() > 0 {
		n := heap.Pop(&h).(*ListNode)
		cur.Next = n
		cur = cur.Next
		if n.Next != nil {
			h.Push(n.Next)
		}
	}

	return dummy.Next
}
