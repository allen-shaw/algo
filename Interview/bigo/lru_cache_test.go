package bigo

import (
	"fmt"
	"testing"
)

type Node struct {
	key        int
	val        int
	prev, next *Node
}

type LinkedList struct {
	head, tail *Node
}

func newLinkedList() LinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LinkedList{
		head: head,
		tail: tail,
	}
}

func (l *LinkedList) insertToHead(n *Node) {
	hnext := l.head.next
	l.head.next = n
	n.prev = l.head
	hnext.prev = n
	n.next = hnext
}

func (l *LinkedList) getAndRemoveLast() *Node {
	if l.tail.prev == l.head {
		return nil
	}

	n := l.tail.prev
	prev := n.prev
	prev.next = l.tail
	l.tail.prev = prev

	n.prev = nil
	n.next = nil
	return n
}

func (l *LinkedList) moveToHead(n *Node) {
	hnext := l.head.next
	l.head.next = n
	n.prev = l.head
	hnext.prev = n
	n.next = hnext
}

type LRUCache struct {
	m    map[int]*Node
	list LinkedList
	cap  int
}

func NewCache(cap int) LRUCache {
	c := LRUCache{
		m:    make(map[int]*Node),
		list: newLinkedList(),
		cap:  cap,
	}
	return c
}

func (c *LRUCache) Put(key, val int) bool {
	n, ok := c.m[key]
	if ok {
		c.list.moveToHead(n)
		n.val = val
		return true
	}

	if len(c.m) == c.cap {
		n := c.list.getAndRemoveLast()
		delete(c.m, n.key)
	}
	n = &Node{key: key, val: val}
	c.m[key] = n
	c.list.insertToHead(n)
	return true
}

func (c *LRUCache) Get(key int) (int, bool) {
	n, ok := c.m[key]
	if !ok {
		return 0, false
	}

	c.list.moveToHead(n)
	return n.val, true
}

//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // 返回  1
// cache.put(3, 3);    // 该操作会使得key=2 作废
// cache.get(2);       // 返回 -1 (未找到)
// cache.put(4, 4);    // 该操作会使得key=1 作废
// cache.get(1);       // 返回 -1 (未找到)
// cache.get(3);       // 返回  3
// cache.get(4);       // 返回  4

func Test_LRUCache(t *testing.T) {
	c := NewCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(1))

}
