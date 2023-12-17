package backtrack

type Node[T any] struct {
	Val  T
	Prev *Node[T]
	Next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	dammy := &Node[T]{}
	ll := &LinkedList[T]{
		head: dammy,
		tail: dammy,
	}
	return ll
}

func (l *LinkedList[T]) PushBack(x T) {
	n := &Node[T]{Val: x, Prev: l.tail}
	l.tail.Next = n
	l.tail = l.tail.Next
	l.size++
}

func (l *LinkedList[T]) RemoveLast() (x T) {
	x = l.tail.Val
	l.tail = l.tail.Prev
	l.tail.Next = nil
	l.size--
	return x
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) ToList() []T {
	list := make([]T, 0, l.size)

	p := l.head.Next
	for p != nil {
		list = append(list, p.Val)
		p = p.Next
	}

	return list
}
