package linkedqueue

import (
	"sync/atomic"
)

type Node struct {
	val  int
	next atomic.Pointer[Node]
}

func (n *Node) Next() *Node {
	return n.next.Load()
}

type LinkedQueue struct {
	head, tail atomic.Pointer[Node]
}

func New() *LinkedQueue {
	q := &LinkedQueue{}

	dummy := &Node{}
	q.head.Store(dummy)
	q.tail.Store(dummy)

	return q
}

func (q *LinkedQueue) Push(x int) {
	node := &Node{val: x}

	for {
		tail := q.tail.Load()
		next := tail.Next()

		if tail != q.tail.Load() {
			continue
		}

		if next != nil {
			q.tail.CompareAndSwap(tail, next)
			continue
		}

		if tail.next.CompareAndSwap(nil, node) {
			q.tail.CompareAndSwap(tail, node)
			break
		}
	}
}

func (q *LinkedQueue) Pop() (int, bool) {
	for {
		head := q.head.Load()
		tail := q.tail.Load()
		next := head.Next()

		if head != q.head.Load() {
			continue
		}

		if head == tail && next == nil {
			return 0, false
		}

		if head == tail && next != nil {
			q.tail.CompareAndSwap(tail, next)
			continue
		}

		if q.head.CompareAndSwap(head, next) {
			return next.val, true
		}
	}
}
