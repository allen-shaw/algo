package lru

import "fmt"

type LRUCache struct {
	index map[int]*node
	cache doubleList
	cap   int
	size  int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		index: map[int]*node{},
		cache: newDoubleList(),
		cap:   capacity,
		size:  0,
	}
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.index[key]
	if !ok {
		return -1
	}

	v := n.val
	c.cache.moveToTail(n)
	return v
}

func (c *LRUCache) Put(key int, value int) {
	n, ok := c.index[key]
	if ok {
		c.cache.delete(n)
		c.size--
	}
	n = &node{key: key, val: value}
	c.cache.append(n)
	c.index[key] = n
	c.size++

	c.print("middle")

	for c.size > c.cap {
		// remove first
		n = c.cache.removeFirst()
		fmt.Println("remove:", n.key, n.val)
		delete(c.index, n.key)
		c.size--
	}
}

func (c *LRUCache) print(s string) {
	fmt.Println("------", s, "------")
	for k, n := range c.index {
		fmt.Println("key:", k, "val:", n.val)
	}
	fmt.Println("--------- size:", c.size)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
