package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestValidParentheses(s string) int {
	maxLen := 0

	stack := make([]int, 0, len(s)+1)
	stack = append(stack, -1)

	for i := range s {
		c := s[i]
		if c == '(' {
			stack = append(stack, i)
		} else { // c == ')'
			stack = stack[:len(stack)-1] // pop
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}

	return maxLen
}

func TestLongestValidParentheses(t *testing.T) {
	var testsets = []struct {
		s   string
		ans int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
		{"()(()", 2},
	}

	for _, tt := range testsets {
		t.Run("longestValidParentheses", func(t *testing.T) {
			res := longestValidParentheses(tt.s)
			assert.Equal(t, tt.ans, res)
		})
	}
}
