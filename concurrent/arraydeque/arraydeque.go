package arraydeque

import (
	"sync"
	"sync/atomic"
)

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
	data   []elem[E]
	head   int32
	tail   int32
	size   int32
	cap    int32
	gmu    sync.RWMutex
	noCopy noCopy
}

func newArrayDeque[E any](cap int) *arrayDeque[E] {
	ad := &arrayDeque[E]{}
	ad.data = make([]elem[E], cap)
	ad.head = 0
	ad.tail = 0
	ad.size = 0
	ad.cap = int32(cap)
	return ad
}

func (a *arrayDeque[E]) AddFirst(e E) {
	a.gmu.RLock()
	for {
		size := atomic.LoadInt32(&a.size)
		if size == a.cap {
			a.gmu.RUnlock()
			a.growth()
			a.AddFirst(e)
			return
		}

		head := atomic.LoadInt32(&a.head)
		newHead := (head - 1 + a.cap) % a.cap

		if atomic.CompareAndSwapInt32(&a.head, head, newHead) {
			a.data[newHead] = elem[E]{val: e}
			break
		}
	}

	atomic.AddInt32(&a.size, 1)
	a.gmu.Unlock()
}

func (a *arrayDeque[E]) AddLast(e E) {
	a.gmu.RLock()

	for {
		size := atomic.LoadInt32(&a.size)
		if size == a.cap {
			a.gmu.RUnlock()
			a.growth()
			a.AddLast(e)
			return
		}

		tail := atomic.LoadInt32(&a.tail)
		newTail := (tail + 1) % a.cap
		if atomic.CompareAndSwapInt32(&a.tail, tail, newTail) {
			a.data[tail] = elem[E]{val: e}
			break
		}
	}

	atomic.AddInt32(&a.size, 1)
	a.gmu.Unlock()
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
	size := atomic.LoadInt32(&a.size)
	return int(size)
}

func (a *arrayDeque[E]) growth() {
	a.gmu.Lock()

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

	a.gmu.Unlock()
}

type Iterator[E any] interface {
	HasNext() bool
	Next() E
}

type iterator[E any] struct {
	next   int32
	a      *arrayDeque[E]
	noCopy noCopy
}

// HasNext implements Iterator.
func (i *iterator[E]) HasNext() bool {
	return i.next < int32(i.a.Size())
}

// Next implements Iterator.
func (i *iterator[E]) Next() E {
	if !i.HasNext() {
		panic("arraydeque: iterator out of range")
	}
	head := atomic.LoadInt32(&i.a.head)
	idx := (head + i.next) % i.a.cap
	e := i.a.data[idx].val
	i.next++
	return e
}

func (a *arrayDeque[E]) Iterator() Iterator[E] {
	return &iterator[E]{next: 0, a: a}
}

type noCopy struct{}

func (noCopy) Lock() {}

func (noCopy) Unlock() {}
