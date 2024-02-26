package backtrack

import (
	"fmt"
	"testing"
)

func findTargetSumWays(nums []int, target int) int {
	memo := make(map[string]int, 0)

	var dfs func(i, val int) int
	dfs = func(i, val int) int {
		if i == len(nums) {
			if val == 0 {
				return 1
			}
			return 0
		}

		key := fmt.Sprintf("%d:%d", i, val)
		if v, ok := memo[key]; ok {
			return v
		}

		// nums[i]为+				// nums[i]为-
		res := dfs(i+1, val-nums[i]) + dfs(i+1, val+nums[i])
		memo[key] = res
		return res
	}

	return dfs(0, target)
}

func TestFindTargetSumWays(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	ans := findTargetSumWays(nums, target)
	fmt.Println(ans)
}
