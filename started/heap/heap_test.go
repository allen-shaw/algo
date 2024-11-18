package heap

type Heap struct {
	data []int
	size int
}

func NewHeap(cap int) Heap {
	return Heap{
		data: make([]int, cap),
		size: 0,
	}
}

func (h *Heap) Push(x int) bool {
	if h.Full() {
		return false
	}

	h.data[h.size] = x
	h.swin(h.size)
	h.size++

	return true
}

func (h *Heap) Pop() (int, bool) {
	if h.Empty() {
		return 0, false
	}

	x := h.data[0]
	h.data[0], h.data[h.size-1] = h.data[h.size-1], h.data[0]
	h.size--

	h.sink(0)
	return x, true
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Full() bool {
	return h.size == len(h.data)
}

func (h *Heap) Empty() bool {
	return h.size == 0
}

func (h *Heap) swin(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] < h.data[index] {
			h.data[parent], h.data[index] = h.data[index], h.data[parent]
			index = parent
		} else {
			break
		}
	}
}

func (h *Heap) sink(index int) {
	for (index*2 + 2) < h.size {
		left := index*2 + 1
		right := index*2 + 2

		min := left
		if right < h.size && h.data[left] > h.data[right] {
			min = right
		}

		if h.data[min] < h.data[index] {
			h.data[min], h.data[index] = h.data[index], h.data[min]
			index = min
		} else {
			break
		}
	}
}
