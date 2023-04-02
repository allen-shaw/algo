package heap

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errHeapEmpty = errors.New("heap empty")
	errHeapFull  = errors.New("heap full")
)

// 小顶堆
type Heap struct {
	nums []int
	cap  int
	size int
}

func NewHeap(cap int) *Heap {
	return &Heap{
		nums: make([]int, cap),
		size: 0,
		cap:  cap,
	}
}

func (h *Heap) isEmpty() bool {
	return h.size == 0
}

func (h *Heap) isFull() bool {
	return h.size == h.cap
}

func (h *Heap) Insert(v int) error {
	if h.isFull() {
		return errHeapFull
	}
	heapInsert(h.nums, h.size, v)
	h.size++
	return nil
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Pop() (int, error) {
	if h.isEmpty() {
		return 0, errHeapEmpty
	}
	res := h.nums[0]
	h.size--

	h.nums[0], h.nums[h.size] = h.nums[h.size], h.nums[0]
	heapify(h.nums, 0, h.size-1)

	return res, nil
}

func heapInsert(arr []int, size, v int) {
	arr[size] = v

	i := size
	for arr[i] < arr[(i-1)/2] {
		// swap arr[i], arr[i].parent
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
		i = (i - 1) / 2
	}
}

func heapify(arr []int, l, r int) {
	i := 0
	for {
		lc := 2*i + 1
		smaller := lc
		if lc >= r {
			// 没有子节点
			break
		}
		if lc+1 <= r && arr[lc+1] < arr[lc] {
			// 有右孩子，且右孩子小于左孩子
			smaller = lc + 1
		}
		// 父节点与较小的孩子交换
		arr[i], arr[smaller] = arr[smaller], arr[i]
		i = smaller
	}
	// 越界，return
}

func TestHeap(t *testing.T) {
	h := NewHeap(5)
	// 3, 7, 4, 8, 1
	h.Insert(3)
	h.Insert(7)
	h.Insert(4)
	h.Insert(8)
	h.Insert(1)

	fmt.Println(h.nums)

	err := h.Insert(100)
	assert.Equal(t, errHeapFull, err)

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

	_, err = h.Pop()
	assert.Equal(t, errHeapEmpty, err)

}
