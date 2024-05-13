package heap

type Heap struct {
	data []int
	size int
	cap  int
}

func New(cap int) *Heap {
	h := new(Heap)
	h.data = make([]int, cap)
	h.size = 0
	h.cap = cap
	return h
}

func (h *Heap) Insert(x int) bool {
	if h.size == h.cap {
		return false
	}
	h.size++
	h.data[h.size] = x
	h.up(h.size)
	return true
}

func (h *Heap) GetMax() (int, bool) {
	if h.size == 0 {
		return 0, false
	}

	return h.data[0], true
}

func (h *Heap) PopMax() (int, bool) {
	if h.size == 0 {
		return 0, false
	}
	val := h.data[0]

	h.size--
	h.data[0] = h.data[h.size]
	h.down(0)

	return val, true
}

func (h *Heap) up(idx int) {
	for p := parent(idx); p >= 0 && h.data[p] < h.data[idx]; {
		h.data[p], h.data[idx] = h.data[idx], h.data[p]
		idx = p
	}
}

func (h *Heap) down(idx int) {
	for left(idx) < h.size {
		max := left(idx)
		if right(idx) < h.size && h.data[right(idx)] > h.data[max] {
			max = right(idx)
		}

		if h.data[idx] >= h.data[max] {
			break
		}

		h.data[max], h.data[idx] = h.data[idx], h.data[max]
		idx = max
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func left(idx int) int {
	return 2*idx + 1
}

func right(idx int) int {
	return 2*idx + 2
}
