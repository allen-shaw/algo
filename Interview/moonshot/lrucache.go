package moonshot

type node[T any] struct {
	key        int
	value      T
	prev, next *node[T]
}

func (c *LRUCache[T]) Set(key int, value T) {
	if n, ok := c.m[key]; ok {
		n.value = value
		c.list.MoveToHead(n)
		return
	}
	n := &node[T]{key: key, value: value}
	c.m[key] = n
	c.list.InsertToHead(n)
	if len(c.m) > c.cap {
		delete(c.m, c.list.tail.key)
		c.list.Remove(c.list.tail)
	}
}

func (c *LRUCache[T]) Get(key int) (T, bool) {
	if n, ok := c.m[key]; ok {
		c.list.MoveToHead(n)
		return n.value, true
	}
	var val T
	return val, false
}

func (c *LRUCache[T]) Evict() T {
	panic("unimplemented")
}

type LRUCache[T any] struct {
	m    map[int]*node[T]
	list LinkedList[T]
	cap  int
}

func NewLRUCache[T any](M int) *LRUCache[T] {
	return &LRUCache[T]{
		m:    make(map[int]*node[T], M),
		list: NewLinkedList[T](),
		cap:  M,
	}
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
}

func NewLinkedList[T any]() LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{}

	head.next = tail
	tail.prev = head

	return LinkedList[T]{
		head: head,
		tail: tail,
	}
}

func (l LinkedList[T]) Remove(tail *node[T]) {
	prev, next := tail.prev, tail.next
	prev.next = next
	next.prev = prev
}

func (l *LinkedList[T]) MoveToHead(n *node[T]) {
	prev, next := n.prev, n.next
	n.prev = nil
	n.next = nil

	prev.next = next
	next.prev = prev

	hnext := l.head.next
	l.head.next = n
	n.prev = l.head
	hnext.prev = n
	n.next = hnext
}

func (l *LinkedList[T]) InsertToHead(n *node[T]) {
	prev, next := l.head, l.head.next
	prev.next = n
	n.prev = prev
	n.next = next
	next.prev = n
}
