package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New[int]()

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	for !q.Empty() {
		v := q.Dequeue()
		fmt.Printf("%v ", v)
	}
}
