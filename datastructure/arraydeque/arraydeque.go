package arraydeque

type ArrayDeque[E any] interface {
	AddLast(E)
	AddFirst(E)
	GetFirst() E
	GetLast() E
	PollFirst() E
	PollLast() E
	IsEmpty() bool
	Size() int
	Iterator() Iterator[E]
}

const (
	defCap = 8
)

func New[E any]() ArrayDeque[E] {
	return newArrayDeque[E](defCap)
}

type elem[E any] struct {
	val E
}

type arrayDeque[E any] struct {
	data []elem[E]
	head int
	tail int
	size int
	cap  int
}

func newArrayDeque[E any](cap int) *arrayDeque[E] {
	ad := &arrayDeque[E]{}
	ad.data = make([]elem[E], cap)
	ad.head = 0
	ad.tail = 0
	ad.size = 0
	ad.cap = cap
	return ad
}

func (a *arrayDeque[E]) AddFirst(e E) {
	if a.size == a.cap {
		a.growth()
		a.AddFirst(e)
		return
	}

	newHead := (a.head - 1 + a.cap) % a.cap
	a.data[newHead] = elem[E]{val: e}
	a.size++
	a.head = newHead
}

func (a *arrayDeque[E]) AddLast(e E) {
	if a.size == a.cap {
		a.growth()
		a.AddLast(e)
		return
	}

	newTail := (a.tail + 1) % a.cap
	a.data[a.tail] = elem[E]{val: e}
	a.size++
	a.tail = newTail
}

func (a *arrayDeque[E]) GetFirst() E {
	if a.size == 0 {
		panic("arraydeque: empty deque")
	}

	return a.data[a.head].val
}

func (a *arrayDeque[E]) GetLast() E {
	if a.size == 0 {
		panic("arraydeque: empty deque")
	}

	tail := (a.tail - 1 + a.cap) % a.cap
	return a.data[tail].val
}

func (a *arrayDeque[E]) IsEmpty() bool {
	return a.size == 0
}

func (a *arrayDeque[E]) PollFirst() E {
	if a.size == 0 {
		panic("arraydeque: empty deque")
	}
	head := a.head
	e := a.data[head]
	a.head = (head + 1) % a.cap
	a.size--

	a.data[head] = elem[E]{}
	return e.val
}

func (a *arrayDeque[E]) PollLast() E {
	if a.size == 0 {
		panic("arraydeque: empty deque")
	}
	tail := (a.tail - 1 + a.cap) % a.cap
	e := a.data[tail]
	a.tail = tail
	a.size--

	a.data[tail] = elem[E]{}
	return e.val
}

func (a *arrayDeque[E]) Size() int {
	return a.size
}

func (a *arrayDeque[E]) growth() {
	newCap := a.cap * 2
	if a.cap > 1024 {
		newCap = a.cap + a.cap/4 // cap growth 25%
	}

	newData := make([]elem[E], newCap)
	if a.head < a.tail {
		copy(newData, a.data)
	} else {
		hsize := a.cap - a.head
		copy(newData, a.data[a.head:])
		copy(newData[hsize:], a.data[:a.tail])
	}

	a.data = newData
	a.cap = newCap
	a.head = 0
	a.tail = a.size
}

type Iterator[E any] interface {
	HasNext() bool
	Next() E
}

type iterator[E any] struct {
	next int
	a    *arrayDeque[E]
}

// HasNext implements Iterator.
func (i *iterator[E]) HasNext() bool {
	return i.next < i.a.size
}

// Next implements Iterator.
func (i *iterator[E]) Next() E {
	if !i.HasNext() {
		panic("arraydeque: iterator out of range")
	}
	idx := (i.a.head + i.next) % i.a.cap
	e := i.a.data[idx].val
	i.next++
	return e
}

func (a *arrayDeque[E]) Iterator() Iterator[E] {
	return &iterator[E]{next: 0, a: a}
}
