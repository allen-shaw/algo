package ringbuffer

import (
	"fmt"
	"testing"
)

type RingBuffer struct {
	buff []byte
	size int
	cap  int
	head int
	tail int
}

func NewRingBuffer(cap int) RingBuffer {
	rb := RingBuffer{}
	rb.buff = make([]byte, cap)
	rb.size = 0
	rb.cap = cap
	rb.head = 0
	rb.tail = 0

	return rb
}

func (rb *RingBuffer) isFull() bool {
	return rb.size == rb.cap
}

func (rb *RingBuffer) isEmpty() bool {
	return rb.size == 0
}

func (rb *RingBuffer) Put(v byte) bool {
	if rb.isFull() {
		return false
	}

	idx := rb.head
	rb.buff[idx] = v

	rb.head++
	rb.head %= rb.cap
	rb.size++

	return true
}

func (rb *RingBuffer) Pop() (byte, bool) {
	if rb.isEmpty() {
		return 0, false
	}

	idx := rb.tail
	v := rb.buff[idx]

	rb.tail++
	rb.tail %= rb.cap
	rb.size--

	return v, true
}

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(3)
	fmt.Println(rb.Put(1))
	fmt.Println(rb.Put(2))
	fmt.Println(rb.Put(3)) // true
	fmt.Println(rb.Put(4)) // false

	fmt.Println(rb.Pop())  // 1, true
	fmt.Println(rb.Put(5)) // true

	fmt.Println(rb.buff)
}
