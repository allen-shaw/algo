package stack

import (
	"container/list"
)

type Stack[T any] struct {
	list *list.List
}

func New[T any]() *Stack[T] {
	s := &Stack[T]{
		list: list.New(),
	}
	return s
}

func (s *Stack[T]) Push(v T) {
	s.list.PushBack(v)
}

func (s *Stack[T]) Pop() (v T) {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		v = e.Value.(T)
	}
	return
}

func (s *Stack[T]) Peak() (v T) {
	e := s.list.Back()
	if e != nil {
		v = e.Value.(T)
	}
	return
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}
