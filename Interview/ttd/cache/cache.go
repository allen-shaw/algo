package cache

import "sync"

// support LRU/MRU
type item[T any] struct {
	key string
	val T
}

type shard[T any] struct {
	mu sync.RWMutex
	m  map[string]*item[T]
}

func newShard[T any]() *shard[T] {
	return &shard[T]{
		m: make(map[string]*item[T]),
	}
}

type Eviction interface {
}

type KVCache[T any] struct {
	gmu      sync.RWMutex
	shards   []*shard[T]
	eviction Eviction
	cap      int
	size     int
}

func NewKVCache[T any](cap int, evict Eviction) *KVCache[T] {
	shards := make([]*shard[T], 32)
	for i := range shards {
		shards[i] = newShard[T]()
	}

	return &KVCache[T]{
		shards:   shards,
		eviction: evict,
		cap:      cap,
		size:     0,
	}
}

func (c *KVCache[T]) Get(key string) (T, error) {
	
}

func (c *KVCache[T]) Set(key string) error {

}
