package cache

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := newLinkedList()

	ll.insert(&node{key: 1, value: 1})
	ll.insert(&node{key: 2, value: 2})
	ll.insert(&node{key: 3, value: 3})

	ll.Println()

	ll.removeTail()
	ll.Println()

}

func TestLRU(t *testing.T) {
	lhm := NewLinkedHashMap(4)

	lhm.Set(1, 1)
	lhm.l.Println()

	fmt.Println("-----")
	lhm.Set(2, 2)
	lhm.l.Println()

	fmt.Println(lhm.Get(1))
	fmt.Println(lhm.Get(2))
	fmt.Println(lhm.Get(3))

	fmt.Println("size:", lhm.Size())
	lhm.l.Println()

	lhm.Set(3, 3)
	lhm.l.Println()

	lhm.Set(4, 4)
	lhm.l.Println()

	lhm.Set(5, 5)
	lhm.l.Println()

	fmt.Println(lhm.Get(1))
	fmt.Println(lhm.Get(2))
	fmt.Println(lhm.Get(3))
	fmt.Println(lhm.Get(4))
	fmt.Println(lhm.Get(5))

	fmt.Println("size:", lhm.Size())
}
