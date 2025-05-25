package errorset

import (
	"fmt"
	"testing"
)

func getPermutation(n int, k int) string {
	factorical := calculateFactorial(n)
	path := make([]byte, 0, n)
	dfs(0, n, k, make([]bool, n+1), path, factorical)
	return string(path)
}

func calculateFactorial(n int) []int {
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i < n+1; i++ {
		factorial[i] = factorial[i-1] * i
	}
	return factorial
}

func dfs(idx, n, k int, used []bool, path []byte, factorial []int) {
	if idx == n {
		return
	}

	cnt := factorial[n-1-idx] // 剩余的阶乘的数量
	for i := 1; i <= n; i++ {
		if used[i] {
			continue
		}
		if cnt < k {
			k = k - cnt
			continue // 剪枝，如果选择i，剩下的数形成的 排列 数量<k，肯定得不到第k个组合
		}
		path = append(path, byte(i+'0'))
		fmt.Println("path:", path)
		used[i] = true
		dfs(idx+1, n, k, used, path, factorial)

		// 不可以回溯（重置变量），没必要回头，因为已经找到所在的树枝，回溯之后后续的 排列 肯定大于k
		return
	}
}

func Test_calculateFactorial(t *testing.T) {
	ans := calculateFactorial(9)
	fmt.Println(ans)
}

func Test_getPermutation(t *testing.T) {
	fmt.Println(getPermutation(3, 3))
}
