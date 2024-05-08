package arraydeque

import (
	"fmt"
	"sync"
	"testing"
)

func TestArrayDeque(t *testing.T) {
	deque := newArrayDeque[int](4)
	deque.AddLast(1)
	deque.AddLast(2)
	deque.AddLast(3)
	t.Log(deque.data)
	t.Log(deque.Size())
	t.Log(deque.GetFirst())
	t.Log(deque.GetLast())
	t.Log(deque.PollFirst())
	t.Log(deque.PollLast())
	t.Log(deque.data, deque.Size())
	t.Log(deque.IsEmpty())
}

func TestArrayDeque_Growth(t *testing.T) {
	deque := newArrayDeque[int](4)
	deque.AddLast(1)
	deque.AddLast(2)
	deque.AddLast(3)
	t.Log(deque.data, deque.Size())

	deque.AddLast(4)
	deque.AddLast(5)
	deque.AddLast(6)
	t.Log(deque.data, deque.Size())

	deque.AddFirst(7)
	deque.AddFirst(8)
	t.Log(deque.data, deque.Size())

	it := deque.Iterator()
	for it.HasNext() {
		fmt.Print(it.Next(), " ")
		// 8 7 1 2 3 4 5 6
	}
	t.Log("")

	fmt.Println(deque)
	deque.growth()
	fmt.Println(deque)

	deque.AddLast(9)
	deque.AddFirst(10)
	t.Log(deque.data)

	it = deque.Iterator()
	for it.HasNext() {
		fmt.Print(it.Next(), " ")
		// 10 8 7 1 2 3 4 5 6 9
	}
}

func TestMutexReenter(t *testing.T) {
	var mu sync.RWMutex

	mu.RLock()
	fmt.Println("get lock success")
	mu.Lock()
	fmt.Println("get lock success again")

}
