package dp

import (
	"fmt"
	"testing"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dp[i])
	}

	return ans
}

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	ans := lengthOfLIS(nums)
	fmt.Println(ans)
}
