package dp

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -100
	}

	return doCoinChange(coins, amount, dp)
}

func doCoinChange(coins []int, amount int, dp []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if dp[amount] != -100 {
		return dp[amount]
	}

	ans := math.MaxInt
	for _, coin := range coins {
		subAns := doCoinChange(coins, amount-coin, dp)
		if subAns == -1 {
			continue
		}
		ans = min(subAns+1, ans)
	}

	dp[amount] = ans
	if ans == math.MaxInt {
		dp[amount] = -1
	}
	return ans
}

func min(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}
