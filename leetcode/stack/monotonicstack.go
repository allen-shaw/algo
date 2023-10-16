package stack

func genMonotonicStack(arr []int) []int {
	s := NewStack[int]()
	res := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]

		for !s.Empty() && s.Peek() <= num {
			s.Pop()
		}

		res[i] = -1
		if !s.Empty() {
			res[i] = s.Peek()
		}
		s.Push(num)
	}

	return res
}
