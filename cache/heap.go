package cache

type Heap interface {
	Init(arr []int)
	Push(x int)
	Pop() int
	Len() int
	Top() int
}

func NewHeap() Heap {
	return newHeap()
}

type heap struct {
	data []int
}

func newHeap() *heap {
	return &heap{}
}

func (h *heap) Init(arr []int) {
	h.data = arr
	n := len(arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		h.down(i)
	}
}

func (h *heap) Push(x int) {
	h.data = append(h.data, x)
	h.up(len(h.data) - 1)
}

func (h *heap) Pop() int {
	top := h.data[0]
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	h.data = h.data[:len(h.data)-1]
	h.down(0)

	return top
}

func (h *heap) Len() int {
	return len(h.data)
}

func (h *heap) Top() int {
	return h.data[0]
}

func (h *heap) up(idx int) {
	i := idx
	for {
		parent := (i - 1) >> 1
		if parent < 0 {
			break
		}
		if h.data[parent] <= h.data[i] {
			break
		}

		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *heap) down(idx int) {
	i := idx
	for {
		left := 2*i + 1
		right := 2*i + 2
		if left >= h.Len() {
			break
		}
		lchild := h.data[left]
		minChild := left
		if right < h.Len() && h.data[right] < lchild {
			minChild = right
		}
		if h.data[i] <= h.data[minChild] {
			break
		}

		h.data[i], h.data[minChild] = h.data[minChild], h.data[i]
		i = minChild
	}
}
