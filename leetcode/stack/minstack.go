package stack

type MinStack struct {
	data *Stack[int]
	min  *Stack[int]
}

func NewMinStack() MinStack {
	return MinStack{
		data: NewStack[int](),
		min:  NewStack[int](),
	}
}

func (s *MinStack) Push(val int) {
	s.data.Push(val)
	if s.min.Empty() || val <= s.min.Peek() {
		s.min.Push(val)
	}
}

func (s *MinStack) Pop() {
	val := s.data.Pop()
	if s.min.Peek() == val {
		s.min.Pop()
	}
}

func (s *MinStack) Top() int {
	return s.data.Peek()
}

func (s *MinStack) GetMin() int {
	return s.min.Peek()
}
