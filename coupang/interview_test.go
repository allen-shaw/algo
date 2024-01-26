package coupang

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 最长子序列
// 8 5 3 6 2 1 4 3
func longestSubsequence(nums []int) int {
	m := make(map[int]struct{})

	for _, n := range nums {
		m[n] = struct{}{}
	}

	ans := 0
	for _, n := range nums {
		if _, ok := m[n-1]; ok {
			continue
		}

		length := 0
		for k := n; ; k++ {
			if _, ok := m[k]; ok {
				length++
			} else {
				ans = max(ans, length)
				break
			}
		}
	}

	return ans
}

func TestLS(t *testing.T) {
	var testsets = []struct {
		nums []int
		out  int
	}{
		{[]int{8, 5, 3, 6, 2, 1, 4, 3}, 6},
	}

	for _, tt := range testsets {
		t.Run("ls", func(t *testing.T) {
			ans := longestSubsequence(tt.nums)
			assert.Equal(t, tt.out, ans)
		})
	}

}

// n * n
// [1,2,3]		[7,4,1]
// [4,5,6]		[8,5,2]
// [7,8,9]		[9,6,3]
func rotateMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		j, k := 0, len(matrix[i])-1
		for j < k {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
			j++
			k--
		}
	}
}

func TestRotateMatrix(t *testing.T) {
	var testsets = []struct {
		matrix [][]int
		out    [][]int
	}{
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
		},
	}

	for _, tt := range testsets {
		t.Run("rotateMatrix", func(t *testing.T) {
			rotateMatrix(tt.matrix)
			assert.Equal(t, tt.out, tt.matrix)
		})
	}
}

func print10() {
	n := 0

	var wg sync.WaitGroup
	chanA := make(chan struct{})
	chanB := make(chan struct{})

	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			_, ok := <-chanA
			if !ok {
				return
			}

			n++
			println(n)
			chanB <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_, ok := <-chanB
			if !ok {
				return
			}

			n--
			println(n)
			chanA <- struct{}{}
		}
	}()

	chanA <- struct{}{}
	wg.Wait()
}

func TestPrintln10(t *testing.T) {

}

// local cache Key Level TTL
type LocalCache[K comparable, V any] interface {
	Set(key K, val V, ttl time.Time) bool
	Get(key K) (V, bool)
	Insert(key K, val V, ttl time.Time) bool
	Delete(key K) bool
	Size() int
}

func NewLocalCache[K comparable, V any](cap int) LocalCache[K, V] {
	return newLocalCache[K, V](cap)
}

type node[K comparable, V any] struct {
	prev *node[K, V]
	next *node[K, V]
	key  K
	val  V
	ttl  time.Time
}

type sortedList[K comparable, V any] struct {
	head *node[K, V]
	tail *node[K, V]
	size int
}

func newSortedList[K comparable, V any]() *sortedList[K, V] {
	l := &sortedList[K, V]{}
	l.head = &node[K, V]{}
	l.tail = &node[K, V]{}

	return l
}

func (l *sortedList[K, V]) Head() *node[K, V] {

}

func (l *sortedList[K, V]) Insert(n *node[K, V]) {

}

func (l *sortedList[K, V]) Delete(n *node[K, V]) {

}

func (l *sortedList[K, V]) Pop() {

}

type localCache[K comparable, V any] struct {
	m   map[K]*node[K, V]
	l   *sortedList[K, V]
	mu  sync.RWMutex
	cap int
}

func newLocalCache[K comparable, V any](cap int) *localCache[K, V] {
	c := &localCache[K, V]{}
	c.m = make(map[K]*node[K, V], cap)
	c.l = newSortedList[K, V]()
	c.cap = cap
	return c
}

// Get implements LocalCache.
func (c *localCache[K, V]) Get(key K) (val V, ok bool) {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()

	c.check()

	var n *node[K, V]
	n, ok = c.m[key]
	if !ok {
		return val, ok
	}

	if now.Sub(n.ttl) > 0 {
		delete(c.m, key)
		c.l.Delete(n)
		return val, false
	}

	return n.val, true
}

// Insert implements LocalCache.
func (c *localCache[K, V]) Insert(key K, val V, ttl time.Time) bool {
	now := time.Now()
	if now.Sub(ttl) > 0 {
		return false
	}

	n := &node[K, V]{key: key, val: val, ttl: ttl}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.check()

	n2, ok := c.m[key]
	if ok {
		if now.Sub(n2.ttl) < 0 {
			return false
		}
	}

	c.l.Delete(n2)
	c.m[key] = n
	c.l.Insert(n)
	return true
}

// Set implements LocalCache.
func (c *localCache[K, V]) Set(key K, val V, ttl time.Time) bool {
	now := time.Now()
	if now.Sub(ttl) > 0 {
		return false
	}

	n := &node[K, V]{key: key, val: val, ttl: ttl}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.check()
	n2, ok := c.m[key]
	if ok {
		c.l.Delete(n2)
	}
	c.m[key] = n
	c.l.Insert(n)
	return true
}

func (c *localCache[K, V]) Delete(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.check()

	n, ok := c.m[key]
	if !ok {
		return false
	}

	delete(c.m, key)
	c.l.Delete(n)
	return true
}

func (c *localCache[K, V]) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.m)
}

func (c *localCache[K, V]) check() {
	n := c.l.Head()
	if n == nil {
		return
	}
	now := time.Now()
	if now.Sub(n.ttl) > 0 {
		c.l.Pop()
	}
	key := n.key
	delete(c.m, key)
}
