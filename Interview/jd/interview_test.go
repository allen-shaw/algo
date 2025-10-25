package jd

import (
	"fmt"
	"testing"
)

// 给定不重复的数组candidates和目标target，返回candidates中所有能够组成target的集合，
// 元素可以重复使用。例子：candidates = [2, 3, 5],
// target = 11： 返回：[2, 2, 2, 2, 3], [2, 2, 2, 5], [2, 3, 3, 3], [3, 3, 5]

func test1(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	var dfs func(start int, sum int, path []int)
	dfs = func(start, sum int, path []int) {
		if sum > target || start == len(candidates) {
			return
		}
		if sum == target {
			ans = append(ans, clone(path))
			return
		}

		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i, sum, path)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0, make([]int, 0))
	return ans
}

func clone(src []int) []int {
	target := make([]int, len(src))
	copy(target, src)
	return target
}

func Test_test1(t *testing.T) {
	candidates := []int{2, 3, 5}
	ans := test1(candidates, 11)
	fmt.Println(ans)
}


// T10 2-2 SEE L6-2