package zhipu

import "fmt"

type node struct {
	key  int
	val  int
	prev *node
	next *node
}

type linkedList struct {
	head, tail *node
	len        int
	cap        int
}

func newLinkedList(cap int) linkedList {
	ll := linkedList{}
	ll.head = &node{}
	ll.tail = &node{}
	ll.head.next = ll.tail
	ll.tail.prev = ll.head
	ll.cap = cap
	ll.len = 0
	return ll
}

func (ll *linkedList) append(n *node) {
	prev := ll.tail.prev
	prev.next = n
	n.prev = prev
	n.next = ll.tail
	ll.tail.prev = n
	ll.len++
}

func (ll *linkedList) moveToHead(n *node) {
	prev, next := n.prev, n.next
	prev.next = next
	next.prev = prev

	n.prev, n.next = nil, nil

	hnext := ll.head.next

	ll.head.next = n
	n.prev = ll.head
	n.next = hnext
	hnext.prev = n
}

func (ll *linkedList) popLastNode() *node {
	if ll.len == 0 {
		return nil
	}
	n := ll.tail.prev
	prev := n.prev
	prev.next = ll.tail
	ll.tail.prev = prev
	return n
}

func (ll *linkedList) print() {
	for cur := ll.head.next; cur != ll.tail; cur = cur.next {
		fmt.Printf("{k: %v, v: %v}->", cur.key, cur.val)
	}
	fmt.Printf("null\n")
}
