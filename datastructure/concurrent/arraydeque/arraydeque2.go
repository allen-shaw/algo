package arraydeque

type arrayDequeChain[E any] struct {
	arr        []E
	head       int32
	tail       int32
	cap        int
	prev, next *arrayDequeChain[E]
}

func (ac *arrayDequeChain[E]) Reset() {
	ac.arr = make([]E, ac.cap)
	ac.head = 0
	ac.tail = 0
}

// func (ac *arrayDequeChain[E]

// type arrayDeque2[E any] struct {
// 	head *arrayDequeChain[E]
// 	tail *arrayDequeChain[E]
// }
