package backtrack

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := permute(nums)
	t.Log(ans)
}

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	ans := permuteUnique(nums)
	fmt.Println(ans)
}

func TestPermute2(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := permute2(nums)
	t.Log(ans)
}

func permute2(nums []int) [][]int {
	visited := make([]bool, len(nums))
	ans := make([][]int, 0)
	var dfs func(idx int, path []int)
	dfs = func(idx int, path []int) {
		if idx == len(nums) {
			ans = append(ans, clone(path))
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			path = append(path, nums[i])
			visited[i] = true

			dfs(idx+1, path)

			path = path[:len(path)-1]
			visited[i] = false
		}
	}

	dfs(0, make([]int, 0))
	return ans
}

func clone(nums []int) []int {
	tgt := make([]int, len(nums))
	copy(tgt, nums)
	return tgt
}
