package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

type MedianFinder struct {
	left  *Heap
	right *Heap
}

func Constructor() MedianFinder {
	mf := MedianFinder{
		left:  newMaxHeap(),
		right: newMinHeap(),
	}
	return mf
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Empty() || num <= mf.left.Top() {
		heap.Push(mf.left, num)
		if mf.left.Size() > mf.right.Size()+1 {
			t := heap.Pop(mf.left)
			heap.Push(mf.right, t)
		}
	} else {
		heap.Push(mf.right, num)
		if mf.left.Size() < mf.right.Size() {
			t := heap.Pop(mf.right)
			heap.Push(mf.left, t)
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Size() == mf.right.Size() {
		return float64(mf.left.Top()+mf.right.Top()) / 2
	}
	return float64(mf.left.Top())
}

type Heap struct {
	data []int
	less func(data []int, i, j int) bool
}

// Len implements heap.Interface.
func (h *Heap) Len() int {
	return len(h.data)
}

// Less implements heap.Interface.
func (h *Heap) Less(i int, j int) bool {
	return h.less(h.data, i, j)
}

// Pop implements heap.Interface.
func (h *Heap) Pop() any {
	x := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return x
}

func (h *Heap) Top() int {
	return h.data[0]
}

// Push implements heap.Interface.
func (h *Heap) Push(x any) {
	n := x.(int)
	h.data = append(h.data, n)
}

// Swap implements heap.Interface.
func (h *Heap) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) Empty() bool {
	return h.Size() == 0
}

func newHeap(less func(data []int, i, j int) bool) *Heap {
	h := &Heap{
		data: make([]int, 0),
		less: less,
	}
	heap.Init(h)
	return h
}

func newMinHeap() *Heap {
	lessFunc := func(data []int, i, j int) bool {
		return data[i] < data[j]
	}
	return newHeap(lessFunc)
}

func newMaxHeap() *Heap {
	lessFunc := func(data []int, i, j int) bool {
		return data[i] > data[j]
	}
	return newHeap(lessFunc)
}

func TestMedianFinder(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian(), mf.left, mf.right)
	mf.AddNum(3)
	fmt.Println(mf.FindMedian(), mf.left, mf.right)
}
