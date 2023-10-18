package lru

import (
	"fmt"
	"testing"
)

// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4

func TestLRUCache(t *testing.T) {
	lruCache := NewLRUCache(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1))
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2))
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(3))
	fmt.Println(lruCache.Get(4))
}

// ["LRUCache","put","put","put","put","get","get"]
// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
func TestLRUCache2(t *testing.T) {
	lruCache := NewLRUCache(2)
	lruCache.Put(2, 1)
	lruCache.print("put 2 1")

	lruCache.Put(1, 1)
	lruCache.print("put 1 1")

	lruCache.Put(2, 3)
	lruCache.print("put 2 3")

	lruCache.Put(4, 1)
	lruCache.print("put 4 1")

	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))
}
