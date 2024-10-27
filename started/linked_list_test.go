package started

type listNode[T any] struct {
	prev, next *listNode[T]
	elem       T
}

func newListNode[T any](e T) *listNode[T] {
	return &listNode[T]{
		elem: e,
	}
}

type LinkedList[T any] struct {
	head, tail *listNode[T]
	size       int
}

func NewLinkedLIst[T any]() LinkedList[T] {
	head, tail := &listNode[T]{}, &listNode[T]{}
	head.next = tail
	tail.prev = head
	return LinkedList[T]{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (l *LinkedList[T]) AddLast(e T) {
	node := newListNode(e)

	temp := l.tail.prev

	temp.next = node
	node.prev = temp

	l.tail.prev = node
	node.next = l.tail

	l.size++
}

func (l *LinkedList[T]) AddFirst(e T) {
	node := newListNode(e)

	temp := l.head.next

	l.head.next = node
	node.prev = l.head

	node.next = temp
	temp.prev = node

	l.size++
}

func (l *LinkedList[T]) Add(index int, e T) {
	l.checkPosIndex(index)

	node := newListNode(e)

	p := l.getNode(index)
	prev := p.prev

	prev.next = node
	node.prev = prev

	node.next = p
	p.prev = node

	l.size++
}

func (l *LinkedList[T]) getNode(index int) *listNode[T] {
	l.checkPosIndex(index)

	p := l.head.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p
}

func (l *LinkedList[T]) checkPosIndex(index int) {
	if index < 0 || index > l.size {
		panic("index out of bound")
	}
}
