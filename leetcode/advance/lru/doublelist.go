package lru


type node struct {
	prev *node
	next *node
	val  int
	key  int
}

type doubleList struct {
	head *node
	tail *node
}

func newDoubleList() doubleList {
	dummy := &node{}
	return doubleList{
		head: dummy,
		tail: dummy,
	}
}

func (dl *doubleList) moveToTail(n *node) {
	dl.delete(n)
	dl.append(n)
}

func (dl *doubleList) delete(n *node) {
	n.prev.next = n.next
	if n == dl.tail {
		dl.tail = dl.tail.prev
	} else {
		n.next.prev = n.prev
	}
	n.next, n.prev = nil, nil
}

func (dl *doubleList) append(n *node) {
	dl.tail.next = n
	n.prev = dl.tail
	dl.tail = n
}


func (dl *doubleList) removeFirst() *node {
	n := dl.head.next
	dl.head.next = n.next
	dl.head.next.prev = dl.head

	return n
}