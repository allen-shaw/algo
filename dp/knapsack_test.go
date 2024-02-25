package dp

import (
	"fmt"
	"testing"
)

func knapsack(W, N int, wt, val []int) int {
	memo := make([][]int, N)
	for i := range memo {
		memo[i] = make([]int, W+1)
	}

	var dp func(i, w int) int
	dp = func(i, w int) int {
		if w <= 0 || i < 0 {
			return 0
		}

		if memo[i][w] != 0 {
			return memo[i][w]
		}

		var ans int
		if wt[i] > w {
			ans = dp(i-1, w)
		} else {
			ans = max(dp(i-1, w-wt[i])+val[i], dp(i-1, w))
		}
		memo[i][w] = ans

		return ans
	}
	return dp(N-1, W)
}

func TestKnapsack(t *testing.T) {
	N := 3
	W := 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}

	ans := knapsack(W, N, wt, val)
	fmt.Println(ans)
}

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}

	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, sum/2+1)
	}
	// dp函数表示nums[:i],能否刚好凑成和为v
	var dp func(i, v int) bool
	dp = func(i, v int) bool {
		if v <= 0 {
			// 因为nums都为正整数，所以不可能为0
			return false
		}
		if i == 0 {
			return nums[0] == v
		}

		if memo[i][v] != 0 {
			return memo[i][v] == 1
		}

		ans := dp(i-1, v) || dp(i-1, v-nums[i])

		if ans {
			memo[i][v] = 1
		} else {
			memo[i][v] = -1
		}
		return ans
	}

	return dp(len(nums)-1, sum/2)
}

func TestCanPartition(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	ans := canPartition(nums)
	fmt.Println(ans)
}

// 完全背包问题
func change(amount int, coins []int) int {
	memo := make([][]int, len(coins)+1)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(i, amount int) int
	dp = func(i, amount int) int {
		if amount <= 0 {
			return 1
		}
		if i < 0 {
			return 0
		}

		if memo[i][amount] != -1 {
			return memo[i][amount]
		}

		var ans int
		if amount < coins[i] {
			ans = dp(i-1, amount)
		} else {
			ans = dp(i-1, amount) + dp(i, amount-coins[i])
		}

		memo[i][amount] = ans
		return ans
	}

	return dp(len(coins)-1, amount)
}

func TestChanege(t *testing.T) {
	amount := 5
	coins := []int{1, 2, 5}
	ans := change(amount, coins)
	fmt.Println(ans)
}
