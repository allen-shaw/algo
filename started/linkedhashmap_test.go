package started

import (
	"fmt"
	"testing"
)

type node[K comparable, V any] struct {
	key K
	val V

	prev, next *node[K, V]
}

type LinkedHashMap[K comparable, V any] struct {
	m          map[K]*node[K, V]
	head, tail *node[K, V]
}

func NewLinkedHashMap[K comparable, V any]() LinkedHashMap[K, V] {
	head := &node[K, V]{}
	tail := &node[K, V]{}
	head.next = tail
	tail.prev = head

	return LinkedHashMap[K, V]{
		m:    make(map[K]*node[K, V]),
		head: head,
		tail: tail,
	}
}

func (m *LinkedHashMap[K, V]) Put(key K, val V) {
	node := &node[K, V]{
		key: key,
		val: val,
	}

	m.addLastNode(node)
	m.m[key] = node
}

func (m *LinkedHashMap[K, V]) Get(key K) (v V, ok bool) {
	node, ok := m.m[key]
	if !ok {
		return
	}

	return node.val, true
}

func (m *LinkedHashMap[K, V]) Remove(key K) {
	node, ok := m.m[key]
	if !ok {
		return
	}

	delete(m.m, key)
	m.removeNode(node)
}

func (m *LinkedHashMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.m[key]
	return ok
}

func (m *LinkedHashMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Size())
	for p := m.head; p != m.tail; p = p.next {
		keys = append(keys, p.key)
	}
	return keys
}

func (m *LinkedHashMap[K, V]) Size() int {
	return len(m.m)
}

func (m *LinkedHashMap[K, V]) addLastNode(n *node[K, V]) {
	prev := m.tail.prev
	prev.next = n
	n.prev = prev
	n.next = m.tail
	m.tail.prev = n
}

func (m *LinkedHashMap[K, V]) removeNode(n *node[K, V]) {
	prev, next := n.prev, n.next

	prev.next = next
	next.prev = prev

	n.prev = nil
	n.next = nil
}

func TestLinkedHashMap(t *testing.T) {
	m := NewLinkedHashMap[string, int]()

	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	m.Put("d", 4)
	m.Put("e", 5)

	fmt.Println(m.Keys())
	m.Remove("c")
	fmt.Println(m.Keys())
}
