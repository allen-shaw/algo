package started

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

type anode[K comparable, V any] struct {
	key   K
	val   V
	index int
}

type ArrayHashMap[K comparable, V any] struct {
	m   map[K]*anode[K, V]
	arr []*anode[K, V]
}

func NewArrayHashMap[K comparable, V any]() ArrayHashMap[K, V] {
	return ArrayHashMap[K, V]{
		m:   make(map[K]*anode[K, V]),
		arr: make([]*anode[K, V], 0),
	}
}

func (m *ArrayHashMap[K, V]) Get(key K) (v V, ok bool) {
	node, ok := m.m[key]
	if !ok {
		return
	}
	return node.val, true
}

func (m *ArrayHashMap[K, V]) Put(key K, val V) {
	node, ok := m.m[key]
	if !ok {
		m.add(key, val)
		return
	}
	node.val = val
}

func (m *ArrayHashMap[K, V]) add(key K, val V) {
	node := &anode[K, V]{
		key:   key,
		val:   val,
		index: len(m.arr),
	}
	m.m[key] = node
	m.arr = append(m.arr, node)
}

func (m *ArrayHashMap[K, V]) Remove(key K) {
	node, ok := m.m[key]
	if !ok {
		return
	}
	delete(m.m, key)
	m.arr[node.index], m.arr[len(m.arr)-1] = m.arr[len(m.arr)-1], m.arr[node.index]
	m.arr[node.index].index = node.index
	m.arr = m.arr[:len(m.arr)-1]
}

func (m *ArrayHashMap[K, V]) RandomKey() K {
	index := rand.Int32N(int32(m.Size()))
	return m.arr[index].key
}

func (m *ArrayHashMap[K, V]) Size() int {
	return len(m.m)
}

func TestArrayHashMap(t *testing.T) {
	m := NewArrayHashMap[int, int]()

	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	m.Put(5, 5)
	m.Put(6, 6)

	for i, e := range m.arr {
		fmt.Println(i, " : ", e)
	}

	fmt.Println(m.Get(1))
	fmt.Println(m.RandomKey())

	m.Remove(4)
	for i, e := range m.arr {
		fmt.Println(i, " : ", e)
	}
	m.Remove(5)
	for i, e := range m.arr {
		fmt.Println(i, " : ", e)
	}
	fmt.Println(m.RandomKey())
	fmt.Println(m.RandomKey())
}
