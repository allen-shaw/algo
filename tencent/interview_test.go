package tencent

import "errors"

type Ringbuffer struct {
	data []int
	head int
	tail int
	len  int
	cap  int
}

func NewRingBuffer(cap int) *Ringbuffer {
	rb := new(Ringbuffer)
	rb.data = make([]int, cap)
	rb.head = 0
	rb.tail = 0
	rb.len = 0
	rb.cap = cap

	return rb
}

func (rb *Ringbuffer) Push(x int) error {
	if rb.len == rb.cap {
		return errors.New("ringbuffer full")
	}

	rb.data[rb.tail] = x
	rb.tail++
	rb.tail = rb.tail % rb.cap
	rb.len++

	return nil
}

func (rb *Ringbuffer) Pop() (int, error) {
	if rb.len == 0 {
		return 0, errors.New("ringbuffer empty")
	}

	x := rb.data[rb.head]
	rb.head++
	rb.head = rb.head % rb.cap
	rb.len--

	return x, nil
}

// netty ?

// epoll -> net

// 1. cache miss 
// 2. 伪共享？