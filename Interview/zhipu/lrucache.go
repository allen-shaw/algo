package zhipu

import (
	"fmt"
)

type LRUCache struct {
	m    map[int]*node
	list linkedList
	cap  int
	size int
}

func NewLRUCache(cap int) LRUCache {
	c := LRUCache{}
	c.m = make(map[int]*node, cap)
	c.list = newLinkedList(cap)
	c.cap = cap
	c.size = 0
	return c
}

func (c *LRUCache) Get(key int) (int, bool) {
	n, ok := c.m[key]
	if !ok {
		return 0, false
	}
	v := n.val
	c.list.moveToHead(n)

	return v, true
}

func (c *LRUCache) Set(key, val int) {
	n, ok := c.m[key]
	if ok {
		n.val = val
		c.list.moveToHead(n)
		return
	}
	n = &node{key: key, val: val}
	c.m[key] = n
	c.size++
	c.list.append(n)
	c.list.moveToHead(n)

	if c.size > c.cap {
		ln := c.list.popLastNode()
		delete(c.m, ln.key)
		c.size--
	}
}

func (c *LRUCache) print() {
	fmt.Println("the map")
	for k, v := range c.m {
		fmt.Printf("key: %v -> val: %v\n", k, v.val)
	}
	fmt.Println("the list")
	c.list.print()
}
