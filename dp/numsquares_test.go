package dp

import (
	"fmt"
	"math"
	"testing"
)

func numSquares(num int) int {
	memo := make(map[int]int)

	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}

		if c, ok := memo[n]; ok {
			return c
		}

		ans := math.MaxInt
		for i := 1; i*i <= n; i++ {
			num := dp(n - i*i)
			if num == -1 {
				continue
			}
			ans = min(ans, num+1)
		}
		if ans == math.MaxInt {
			ans = -1
		}

		memo[n] = ans
		return ans
	}

	return dp(num)
}

func TestNumSquares(t *testing.T) {
	fmt.Println(numSquares(12))
}

func TestSqrt(t *testing.T) {
	n := 9
	fmt.Println(int(math.Sqrt(float64(n))))
}

func maxProduct(nums []int) int {
	ans := math.MinInt
	memo := make([][]int, len(nums))

	// dp 返回nums[0:i]的最大乘积和最小乘积
	var dp func(i int) (int, int)
	dp = func(i int) (int, int) {
		if i == 0 {
			ans = max(nums[0], ans)
			return nums[0], nums[0]
		}

		if len(memo[i]) != 0 {
			return memo[i][0], memo[i][1]
		}

		var (
			maxi, mini int
		)

		ma, mi := dp(i - 1)
		if nums[i] <= 0 {
			maxi = max(nums[i], mi*nums[i])
			mini = min(nums[i], ma*nums[i])
		} else {
			maxi = max(nums[i], ma*nums[i])
			mini = min(nums[i], mi*nums[i])
		}

		memo[i] = []int{maxi, mini}
		ans = max(ans, maxi)
		return maxi, mini
	}

	dp(len(nums) - 1)
	return ans
}

func TestMaxProduct(t *testing.T) {
	nums := []int{-2} // 2, 3, -2, 4}
	ans := maxProduct(nums)
	fmt.Println(ans)
}
