package flexport

import (
	"fmt"
	"math"
	"testing"
)

// 1.1 Minimum Sum Partition with Max Subarray Length
// Given an array nums and an integer k,
// partition the array into subarrays with a maximum length of k.
// Each subarray should take its maximum value to form a sum.
// Find the number of partition schemes that result in the minimum sum.
// For instance, for nums: [1, 5, 3, 4], k: 2, the output is 1, with the minimum sum partition being [1, 5] [3, 4].
// The method to solve this problem is dynamic programming.

func minSumPartitionWithMaxSubarrLength(nums []int, k int) int {
	n := len(nums)

	// dp(i) => [i, n)'s minSum, schemeNum
	var dp func(i int) (int, int)
	dp = func(i int) (int, int) {
		if i >= n {
			return 0, 0
		}

		maxVal := math.MinInt
		minSum := math.MaxInt
		schemeNum := 0

		for j := 0; j < k; j++ {
			idx := i + j
			if idx >= n {
				break
			}
			maxVal = max(maxVal, nums[idx])
			nextSum, nextNum := dp(idx + 1)
			totalSum := maxVal + nextSum
			if totalSum < minSum {
				minSum = totalSum
				schemeNum = 1
			} else if totalSum == minSum {
				schemeNum += nextNum
			}
		}

		return minSum, schemeNum
	}

	minSum, schemeNum := dp(0)
	fmt.Println("minSum:", minSum, " schemeNum:", schemeNum)
	return schemeNum
}

func Test_minSumPartitionWithMaxSubarrLength(t *testing.T) {
	nums := []int{1, 5, 3, 4}
	k := 2
	minSumPartitionWithMaxSubarrLength(nums, k)
}

// 1043. 分隔数组以得到最大和
func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	memo := make(map[int]int)

	// dp(i) -> [i, n) maxSum
	var dp func(i int) int
	dp = func(i int) int {
		if i >= n {
			return 0
		}

		if sum, ok := memo[i]; ok {
			return sum
		}

		maxVal := math.MinInt
		maxSum := 0

		for j := 0; j < k; j++ {
			idx := i + j
			if idx >= n {
				break
			}
			maxVal = max(maxVal, arr[idx])
			curSum := (j + 1) * maxVal
			totalSum := curSum + dp(i+j+1)
			maxSum = max(maxSum, totalSum)
		}

		memo[i] = maxSum
		return maxSum
	}

	return dp(0)
}

func Test_maxSumAfterPartitioning(t *testing.T) {
	arr := []int{1, 15, 7, 9, 2, 5, 10}
	// arr := []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}
	k := 3
	ans := maxSumAfterPartitioning(arr, k)
	fmt.Println(ans)
}
