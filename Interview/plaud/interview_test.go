package plaud

type Node struct {
	key   string
	value string
	prev  *Node
	next  *Node
}

type LinkedList struct {
	head, tail *Node
	size       int
}

func NewLinkedList() LinkedList {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head

	return LinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (l *LinkedList) moveToHead(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev

	l.InsertToHead(node)
}

func (l *LinkedList) InsertToHead(node *Node) {
	headNext := l.head.next
	l.head.next = node
	node.prev = l.head
	node.next = headNext
	headNext.prev = node
}

func (l *LinkedList) RemoveTail() *Node {
	toRemove := l.tail.prev
	pre := toRemove.prev
	pre.next = l.tail
	l.tail.prev = pre

	toRemove.prev = nil
	toRemove.next = nil
	return toRemove
}

type LRUCache struct {
	size int
	cap  int
	m    map[string]*Node
	list LinkedList
}

func NewCache(cap int) LRUCache {
	return LRUCache{
		size: 0,
		cap:  cap,
		m:    make(map[string]*Node),
		list: NewLinkedList(),
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	node, ok := c.m[key]
	if !ok {
		return "", false
	}
	c.list.moveToHead(node)
	return node.value, true
}

func (c *LRUCache) Set(key, val string) {
	node, ok := c.m[key]
	if ok {
		node.value = val
		c.list.moveToHead(node)
		return
	}

	if c.size == c.cap {
		tail := c.list.RemoveTail()
		k := tail.key
		delete(c.m, k)
		c.size--
	}

	node = &Node{key: key, value: val}
	c.m[key] = node
	c.list.InsertToHead(node)
	c.size++
}
