package started

import (
	"fmt"
	"testing"
)

type ArrayList[T any] struct {
	size int
	cap  int
	data []T
}

func NewArrayList[T any](cap int) ArrayList[T] {
	l := ArrayList[T]{
		size: 0,
		cap:  cap,
		data: make([]T, cap),
	}
	return l
}

func (l *ArrayList[T]) AddLast(e T) {
	if l.size == l.cap {
		l.resize(l.cap * 2)
	}
	l.data[l.size] = e
	l.size++
}

func (l *ArrayList[T]) Add(index int, e T) {
	l.checkPosIndex(index)
	if l.size == l.cap {
		l.resize(l.cap * 2)
	}

	for i := l.size - 1; i >= index; i-- {
		l.data[i+1] = l.data[i]
	}
	l.data[index] = e
	l.size++
}

func (l *ArrayList[T]) AddFirst(e T) {
	l.Add(0, e)
}

func (l *ArrayList[T]) RemoveLast() T {
	if l.size == 0 {
		panic("array empty")
	}

	l.size--
	val := l.data[l.size]
	if l.size <= l.cap/4 {
		l.resize(l.cap / 2)
	}
	return val
}

func (l *ArrayList[T]) Remove(index int) T {
	l.checkElemIndex(index)

	val := l.data[index]
	l.size--
	if l.size <= l.cap/4 {
		l.resize(l.cap / 2)
	}

	for i := index; i < l.size; i++ {
		l.data[i] = l.data[i+1]
	}

	return val
}

func (l *ArrayList[T]) RemoveFirst() {
	l.Remove(0)
}

func (l *ArrayList[T]) Get(index int) T {
	l.checkElemIndex(index)
	return l.data[index]
}

func (l *ArrayList[T]) Set(index int, e T) {
	l.checkElemIndex(index)
	l.data[index] = e
}

func (l *ArrayList[T]) Size() int {
	return l.size
}

func (l *ArrayList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *ArrayList[T]) resize(size int) {
	l.cap = size
	newData := make([]T, l.cap)
	copy(newData, l.data)
	l.data = newData
}

func (l *ArrayList[T]) checkPosIndex(index int) {
	if index < 0 || index > l.size {
		panic("index out of bound")
	}
}

func (l *ArrayList[T]) checkElemIndex(index int) {
	if index < 0 || index >= l.size {
		panic("index out of bound")
	}
}

func (l *ArrayList[T]) Display() {
	fmt.Printf("size = %v, cap = %v\n", l.size, l.cap)
	fmt.Println(l.data[:l.size])
}

func TestArrayList(t *testing.T) {
	arr := NewArrayList[int](3)

	for i := 0; i <= 5; i++ {
		arr.AddLast(i)
	}
	arr.Display()

	arr.Remove(3)
	arr.Display()

	arr.Add(1, 9)
	arr.Display()

	arr.AddFirst(100)
	arr.Display()

	val := arr.RemoveLast()
	fmt.Printf("remove last' %v\n", val)
	arr.Display()

	for i := 0; i < arr.Size(); i++ {
		fmt.Printf("%v ", arr.Get(i))
	}

}
