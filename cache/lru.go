package cache

import "fmt"

type node struct {
	key   int
	value int
	next  *node
	prev  *node
}

type linkedlist struct {
	head *node
	tail *node
	size int
}

func newLinkedList() *linkedlist {
	ll := &linkedlist{}
	ll.head = &node{value: -1}
	ll.tail = &node{value: -2}
	ll.head.next = ll.tail
	ll.tail.prev = ll.head
	return ll
}

func (ll *linkedlist) Println() {
	p := ll.head
	for p != nil {
		fmt.Printf("%v->", p.value)
		p = p.next
	}
	fmt.Println()
}

func (ll *linkedlist) moveToHead(n *node) {
	prev := n.prev
	next := n.next

	prev.next = next
	next.prev = prev

	n.prev = nil
	n.next = nil

	temp := ll.head.next
	ll.head.next = n
	n.prev = ll.head
	n.next = temp
	temp.prev = n
}

func (ll *linkedlist) removeTail() {
	if ll.size == 0 {
		return
	}

	rt := ll.tail.prev
	temp := rt.prev
	temp.next = ll.tail
	ll.tail.prev = temp

	ll.size--
}

func (ll *linkedlist) insert(n *node) {
	temp := ll.head.next
	ll.head.next = n
	n.prev = ll.head

	n.next = temp
	temp.prev = n

	ll.size++
}

type LinkedHashMap struct {
	l    *linkedlist
	m    map[int]*node
	cap  int
	size int
}

func NewLinkedHashMap(cap int) *LinkedHashMap {
	lhm := &LinkedHashMap{
		l:    newLinkedList(),
		m:    make(map[int]*node, cap),
		cap:  cap,
		size: 0,
	}
	return lhm
}

func (m *LinkedHashMap) Set(k, v int) {
	n, ok := m.m[k]
	if ok {
		n.value = v
		m.moveToHead(n)
		return
	}

	// insert
	if m.size == m.cap {
		tail := m.l.tail.prev
		m.removeTail()
		delete(m.m, tail.key)
		m.size--
	}

	// do insert
	n = &node{key: k, value: v}
	m.insert(n)
	m.m[k] = n
	m.size++
}

func (m *LinkedHashMap) Get(k int) (int, bool) {
	n, ok := m.m[k]
	if !ok {
		return 0, false
	}
	m.moveToHead(n)
	return n.value, true
}

func (m *LinkedHashMap) Size() int {
	return len(m.m)
}

func (m *LinkedHashMap) moveToHead(n *node) {
	m.l.moveToHead(n)
}

func (m *LinkedHashMap) removeTail() {
	m.l.removeTail()
}

func (m *LinkedHashMap) insert(n *node) {
	m.l.insert(n)
}
