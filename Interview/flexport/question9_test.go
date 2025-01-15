package flexport

import (
	"fmt"
	"testing"
)

func knapsack(W int, N int, wt, val []int) int {
	// 第i个物品，背包容量为c时的最大价值
	var dp func(i, c int) int
	dp = func(i, c int) int {
		if i == 0 || c == 0 {
			return 0
		}
		if wt[i-1] > c {
			return dp(i-1, c)
		}
		return max(dp(i-1, c), dp(i-1, c-wt[i-1])+val[i-1])
	}
	return dp(N, W)
}

func Test_knapsack(t *testing.T) {
	W, N := 4, 3
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}

	fmt.Println(knapsack(W, N, wt, val))
}

func minCost(W int, wt []int, val []int) int {
	// 第i个物品，容量为c的情况下，最小花费
	var dp func(i int, c int) int
	dp = func(i, c int) int {
		if c == 0 {
			// 全都在当地购买
			return sumArray(val...)
		}
		if i == 0 {
			return 0
		}

		if wt[i-1] > c {
			return dp(i-1, c)
		}
		return min(dp(i-1, c), dp(i-1, c-wt[i-1])-val[i-1])
	}

	n := len(wt)
	return dp(n, W)
}

func sumArray(nums ...int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func TestMinCosr(t *testing.T) {
	W := 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}

	fmt.Println(minCost(W, wt, val))
}
