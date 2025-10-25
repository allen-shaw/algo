package agoda

import (
	"fmt"
	"iter"
	"testing"
)

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		data: make(map[T]struct{}),
	}
}

func (s *Set[T]) Get(val T) bool {
	_, ok := s.data[val]
	return ok
}

func (s *Set[T]) Set(val T) {
	s.data[val] = struct{}{}
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for val := range s.data {
			fmt.Printf("iter val %v\n", val)
			if !yield(val) {
				return
			}
		}
	}
}

func Test_Set(t *testing.T) {
	set := NewSet[string]()
	set.Set("1")
	set.Set("2")
	set.Set("3")

	for v := range set.All() {
		fmt.Printf("for All: %v\n", v)
	}
}

func Test_SetIter(t *testing.T) {
	set := NewSet[string]()
	set.Set("1")
	set.Set("2")
	set.Set("3")

	seq := set.All()

	next, stop := iter.Pull(seq)
	defer stop()
	for {
		val, ok := next()
		if !ok {
			break
		}
		fmt.Printf("iter out: %v\n", val)
	}
}
