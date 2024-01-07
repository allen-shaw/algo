package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minAddToMakeValid(s string) int {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(':
			stack = append(stack, c)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1] // pop
			} else {
				stack = append(stack, c)
			}
		}
	}

	return len(stack)
}

func TestMinAddToMakeValid(t *testing.T) {
	var testcases = []struct {
		s   string
		out int
	}{
		{"())", 1},
		{"(((", 3},
	}

	for _, tt := range testcases {
		t.Run("minAddToMakeValid", func(t *testing.T) {
			ans := minAddToMakeValid(tt.s)
			assert.Equal(t, tt.out, ans)
		})
	}
}

func minInsertions(s string) int {
	ans := 0  // 需要添加的'('或')'的数量
	left := 0 // 多余的'('的数量

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			if i+1 < len(s) && s[i+1] == ')' {
				i++ // 直接跳过s[i+1],如果前面有'('则可以直接抵消，否则需要补充一个'()
			} else {
				ans++ // 需要补一个')'
			}
			if left > 0 {
				left-- // 如果前面有一个'(',抵消掉一个多余'('
			} else {
				ans++ // 前面没有多余的'(', 需要补一个'('
			}
		}
	}

	ans += left * 2
	return ans
}

func TestMinInsertions(t *testing.T) {
	var testsets = []struct {
		s   string
		out int
	}{
		{"(()))", 1},
		{"())", 0},
		{"))())(", 3},
		{"((((((", 12},
		{")))))))", 5},
		{"(()))(()))()())))", 4},
		{"((())))))", 0},
	}

	for _, tt := range testsets {
		t.Run("minInsertions", func(t *testing.T) {
			ans := minInsertions2(tt.s)
			assert.Equal(t, tt.out, ans)
		})
	}

}
func minInsertions2(s string) int {
	res, needRight := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			needRight += 2
			if needRight%2 == 1 {
				res++
				needRight--
			}
		}

		if s[i] == ')' {
			needRight--
			if needRight < 0 {
				res++
				needRight = 1
			}
		}
	}

	return res + needRight
}
