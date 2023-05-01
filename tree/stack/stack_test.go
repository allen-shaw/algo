package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := New[int]()

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	for i := 0; i < 5; i++ {
		v := s.Pop()
		fmt.Println(v)
	}
}
