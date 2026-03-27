package coding

import (
	"fmt"
	"testing"
)

func expandFromCenter(s string, left, right int) []string {
	i, j := left, right
	result := make([]string, 0)

	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		result = append(result, s[i:j+1])
		i--
		j++
	}

	return result
}

func subString(s string) []string {
	ans := make([]string, 0)

	for i := 0; i < len(s); i++ {
		ans = append(ans, expandFromCenter(s, i, i)...)
		ans = append(ans, expandFromCenter(s, i, i+1)...)
	}
	return ans
}

func Test_subString(t *testing.T) {
	s := "abcbad"
	ans := subString(s)
	fmt.Println(ans)
}
