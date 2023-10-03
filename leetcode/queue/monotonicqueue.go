package queue

type MonotonicQueue struct {
	maxq []int
	popped int
	size   int
}

func NewMonotonicQueue(n int) *MonotonicQueue {
	return &MonotonicQueue{
		maxq: make([]int, 0, n),
	}
}

func (q *MonotonicQueue) Push(e int) {
	for !q.Empty() && q.first() < e {
		q.popped++
		q.pop()
	}

	q.maxq = append(q.maxq, e)
	q.size++
}

func (q *MonotonicQueue) Pop() {
	q.size--
	if q.popped > 0 {
		q.popped--
		return
	}
	q.pop()
}

func (q *MonotonicQueue) Size() int {
	return q.size
}

func (q *MonotonicQueue) Max() int {
	return q.first()
}

func (q *MonotonicQueue) Min() int {
	panic("todo")
}

func (q *MonotonicQueue) Empty() bool {
	return len(q.maxq) == 0
}

func (q *MonotonicQueue) first() int {
	return q.maxq[0]
}
func (q *MonotonicQueue) pop() {
	q.maxq = q.maxq[1:]
}
