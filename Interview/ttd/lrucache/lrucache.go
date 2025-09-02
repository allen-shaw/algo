package lrucache

import (
	"container/list"
	"sync"
)

type item struct {
	key string
	val string
}

type LRUCache struct {
	mu  sync.RWMutex
	m   map[string]*list.Element
	l   *list.List
	cap int
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		m:   make(map[string]*list.Element),
		l:   list.New(),
		cap: cap,
	}
}

func (c *LRUCache) Get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.m[key]
	if !ok {
		return ""
	}

	c.l.MoveToFront(elem)
	return elem.Value.(*item).val
}

func (c *LRUCache) Set(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.m[key]
	if ok {
		elem.Value.(*item).val = val
		c.l.MoveToFront(elem)
		return
	}

	sz := len(c.m)
	if sz == c.cap {
		last := c.l.Back()
		c.l.Remove(last)
		delete(c.m, last.Value.(*item).key)
	}

	it := &item{key: key, val: val}
	elem = c.l.PushFront(it)
	c.m[key] = elem
}
