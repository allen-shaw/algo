package started

type RangeBuffer[T any] struct {
	arr   []T
	start int
	end   int
	size  int
	cap   int
}

func NewRangeBuffer[T any](cap int) RangeBuffer[T] {
	return RangeBuffer[T]{
		arr:   make([]T, cap),
		start: 0,
		end:   0,
		size:  0,
		cap:   cap,
	}
}

func (rb *RangeBuffer[T]) AddFirst(val T) {
	if rb.isFull() {
		rb.resize(rb.cap * 2)
	}

	rb.start = (rb.start - 1 + rb.cap) % rb.cap
	rb.arr[rb.start] = val
	rb.size++
}

func (rb *RangeBuffer[T]) RemoveFirst() {
	if rb.IsEmpty() {
		return
	}

	rb.start = (rb.start + 1) & rb.cap
	rb.size--
}

func (rb *RangeBuffer[T]) AddLast(val T) {
	if rb.isFull() {
		rb.resize(rb.cap * 2)
	}

	rb.arr[rb.end] = val
	rb.end = (rb.end + 1) & rb.cap
	rb.size++
}

func (rb *RangeBuffer[T]) RemoveLast(val T) {
	if rb.IsEmpty() {
		return
	}
	rb.end = (rb.end - 1 + rb.cap) % rb.cap
	rb.size--
}

func (rb *RangeBuffer[T]) GetFirst() T {
	if rb.IsEmpty() {
		panic("range buff is empty")
	}

	return rb.arr[rb.start]
}

func (rb *RangeBuffer[T]) GetLast() T {
	if rb.IsEmpty() {
		panic("range buff is empty")
	}

	index := (rb.end - 1 + rb.cap) % rb.cap
	return rb.arr[index]
}

func (rb *RangeBuffer[T]) resize(size int) {
	newArr := make([]T, size)

	for i := 0; i < rb.size; i++ {
		index := (i + rb.start) % rb.cap
		newArr[i] = rb.arr[index]
	}

	rb.arr = newArr
	rb.cap = size

	// update start and end
	rb.start = 0
	rb.end = rb.size
}

func (rb *RangeBuffer[T]) Size() int {
	return rb.size
}

func (rb *RangeBuffer[T]) isFull() bool {
	return rb.size == len(rb.arr)
}

func (rb *RangeBuffer[T]) IsEmpty() bool {
	return rb.size == 0
}
