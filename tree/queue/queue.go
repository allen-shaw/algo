package queue

import "container/list"

type Queue[T any] struct {
	list *list.List
}

func New[T any]() *Queue[T] {
	q := &Queue[T]{
		list: list.New(),
	}
	return q
}

func (q *Queue[T]) Enqueue(e T) {
	q.list.PushBack(e)
}

func (q *Queue[T]) Dequeue() (e T) {
	v := q.list.Front()
	if v != nil {
		e = v.Value.(T)
		q.list.Remove(v)
	}
	return
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}
