package stack

type pair struct {
	count int
	str   []byte
}

func decodeString(s string) string {
	count := 0
	stack := NewStack[pair]()
	ans := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		c := s[i]

		if isNum(c) {
			count = 10*count + ctoNum(c)
		} else if c == '[' {
			stack.Push(pair{count: count, str: ans})
			count = 0
			ans = make([]byte, 0)
		} else if c == ']' {
			p := stack.Pop()
			ss := buildStr(p.count, ans)
			ans = append(p.str, ss...)
		} else {
			ans = append(ans, c)
		}
	}
	return string(ans)
}

func isNum(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func ctoNum(c byte) int {
	return int(c - '0')
}

func buildStr(num int, ss []byte) []byte {
	out := make([]byte, 0, num*len(ss))
	for i := 0; i < num; i++ {
		out = append(out, ss...)
	}
	return out
}
