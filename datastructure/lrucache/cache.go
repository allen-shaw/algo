package lrucache

type Node struct {
	key        string
	value      string
	prev, next *Node
}

type LinkedList struct {
	head, tail *Node
}

func NewLinkedList() *LinkedList {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return &LinkedList{
		head: head,
		tail: tail,
	}
}

func (ll *LinkedList) MoveToHead(n *Node) {
	ll.Delete(n)
	ll.Insert(n)
}

func (ll *LinkedList) Delete(n *Node) {
	prev, next := n.prev, n.next
	next.prev = prev
	prev.next = next
}

func (ll *LinkedList) Insert(n *Node) {
	prev, next := ll.head, ll.head.next
	prev.next = n
	n.prev = prev
	next.prev = n
	n.next = next
}

func (ll *LinkedList) GetLast() *Node {
	return ll.tail.prev
}

type LRUCache struct {
	cap  int
	m    map[string]*Node
	ll   *LinkedList
	size int
}

func New(cap int) *LRUCache {
	cache := &LRUCache{
		cap: cap,
		m:   make(map[string]*Node),
		ll:  NewLinkedList(),
	}
	return cache
}

func (c *LRUCache) Put(key, value string) {
	node, ok := c.m[key]
	if ok {
		node.value = value
		c.ll.MoveToHead(node)
		return
	}

	if c.size == c.cap {
		last := c.ll.GetLast()
		c.delete(last)
	}

	node = &Node{key: key, value: value}
	c.ll.Insert(node)
	c.m[key] = node
	c.size++
}

func (c *LRUCache) Get(key string) string {
	node, ok := c.m[key]
	if ok {
		c.ll.MoveToHead(node)
		return node.value
	}

	return ""
}

func (c *LRUCache) delete(n *Node) {
	c.ll.Delete(n)
	delete(c.m, n.key)
	c.size--
}
