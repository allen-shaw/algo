package doubleheap

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type room struct {
	endtime int
	id      int
}

type UsingHeap []room

// Len implements heap.Interface.
func (h UsingHeap) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h UsingHeap) Less(i int, j int) bool {
	return h[i].endtime < h[j].endtime
}

// Pop implements heap.Interface.
func (h *UsingHeap) Pop() any {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

// Push implements heap.Interface.
func (h *UsingHeap) Push(x any) {
	(*h) = append((*h), x.(room))
}

// Swap implements heap.Interface.
func (h UsingHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

type IdleHeap []int

// Len implements heap.Interface.
func (h IdleHeap) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h IdleHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Pop implements heap.Interface.
func (h *IdleHeap) Pop() any {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

// Push implements heap.Interface.
func (h *IdleHeap) Push(x any) {
	(*h) = append((*h), x.(int))
}

// Swap implements heap.Interface.
func (h IdleHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// roomId

func mostBooked(n int, meetings [][]int) int {
	cnt := make([]int, n)
	using, idle := UsingHeap{}, IdleHeap{}

	for i := 0; i < n; i++ {
		idle = append(idle, i)
	}
	heap.Init(&idle)
	heap.Init(&using)

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, m := range meetings {
		st, end := m[0], m[1]

		// clean using heap
		for len(using) > 0 && using[0].endtime <= st {
			r := heap.Pop(&using).(room)
			heap.Push(&idle, r.id)
		}

		var ro room

		if idle.Len() == 0 {
			r := heap.Pop(&using).(room)
			end += r.endtime - st
			ro.endtime = end
			ro.id = r.id
		} else {
			rid := heap.Pop(&idle).(int)
			ro.endtime = end
			ro.id = rid
		}

		heap.Push(&using, ro)
		cnt[ro.id]++
	}

	maxCnt, maxIdx := 0, 0
	for i, c := range cnt {
		if c > maxCnt {
			maxCnt = c
			maxIdx = i
		}
	}
	return maxIdx
}

func Test_mostBooked(t *testing.T) {
	n := 4
	meetings := [][]int{{10, 11}, {13, 15}, {9, 19}, {0, 12}, {12, 20}}

	fmt.Println(mostBooked(n, meetings))
}
