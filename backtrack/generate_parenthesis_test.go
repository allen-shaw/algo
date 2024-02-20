package backtrack

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) []string {
	var (
		str string
		res = make([]string, 0)
	)

	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left == 0 && right == 0 {
			res = append(res, str)
			return
		}
		if left > right {
			return
		}

		if left > 0 {
			str += "("
			dfs(left-1, right)
			str = str[:len(str)-1]
		}
		if right > 0 {
			str += ")"
			dfs(left, right-1)
			str = str[:len(str)-1]
		}
	}

	dfs(n, n)
	return res
}

func TestGenerateParenthesis(t *testing.T) {
	ans := generateParenthesis(3)
	fmt.Println(ans)
}
