package stack

func isValid(s string) bool {
	stack := NewStack[byte]()
	for i := range s {
		ss := s[i]
		if stack.Empty() {
			stack.Push(ss)
			continue
		}

		t := stack.Peek()
		switch ss {
		case ')':
			if t == '(' {
				stack.Pop()
				continue
			}
		case ']':
			if t == '[' {
				stack.Pop()
				continue
			}
		case '}':
			if t == '{' {
				stack.Pop()
				continue
			}
		}
		//
		stack.Push(ss)
	}

	return stack.Empty()
}
