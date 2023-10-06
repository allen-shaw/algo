package queue

type elem struct {
	val    int
	popped int
}

type MonotonicQueue struct {
	maxq []*elem
	size int
}

func NewMonotonicQueue(n int) *MonotonicQueue {
	return &MonotonicQueue{
		maxq: make([]*elem, 0, n),
	}
}

func (q *MonotonicQueue) Push(e int) {
	// fmt.Println("Push:", e)
	n := 0
	for !q.Empty() && q.last().val < e {
		n += q.last().popped + 1
		q.popLast()
	}

	q.maxq = append(q.maxq, &elem{val: e, popped: n})
	q.size++
	// fmt.Println("After Push:", q.maxq)
}

func (q *MonotonicQueue) Pop() {
	// defer func() {
	// 	fmt.Println("After Pop:", q.maxq)
	// }()

	q.size--
	e := q.first()
	if e.popped > 0 {
		e.popped--
		return
	}
	q.pop()

}

func (q *MonotonicQueue) Size() int {
	return q.size
}

func (q *MonotonicQueue) Max() int {
	return q.first().val
}

func (q *MonotonicQueue) Min() int {
	panic("todo")
}

func (q *MonotonicQueue) Empty() bool {
	return len(q.maxq) == 0
}

func (q *MonotonicQueue) first() *elem {
	return q.maxq[0]
}
func (q *MonotonicQueue) pop() {
	q.maxq = q.maxq[1:]
}
func (q *MonotonicQueue) last() *elem {
	return q.maxq[len(q.maxq)-1]
}
func (q *MonotonicQueue) popLast() {
	q.maxq = q.maxq[:len(q.maxq)-1]
}
