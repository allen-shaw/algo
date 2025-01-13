package flexport

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Example 1:
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"

// Example 2:
// Input: s = "3[a2[c]]"		// 3
// Output: "accaccacc"

// Example 3:
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"
func decodeString1(s string) string {
	stackStr := make([]string, 0)
	stackNum := make([]int, 0)

	curNum := 0
	curStr := make([]rune, 0)

	for _, c := range s {
		if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
			curStr = append(curStr, c)
		} else if '0' <= c && c <= '9' {
			curNum = curNum*10 + int(c-'0')
		} else if c == '[' {
			// store curnum and curstr
			stackNum = append(stackNum, curNum)
			curNum = 0
			stackStr = append(stackStr, string(curStr))
			curStr = make([]rune, 0)
		} else if c == ']' {
			times := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1] // pop

			lastStr := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]
			curStr = []rune(lastStr + buildStr(times, string(curStr)))
		}
	}

	return string(curStr)
}

func buildStr(times int, str string) string {
	ss := make([]string, 0, times)
	for i := 0; i < times; i++ {
		ss = append(ss, str)
	}
	return strings.Join(ss, "")
}

func Test_decodeString1(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	for _, c := range cases {
		out := decodeString1(c.input)
		assert.Equal(t, c.expected, out)
	}
}
