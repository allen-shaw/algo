package array

import "math"

func maxSubArray(nums []int) int {
	window, maxSum := 0, math.MinInt

	left, right := 0, 0

	for right < len(nums) {
		num := nums[right]
		right++

		window += num
		maxSum = max(maxSum, window)

		for left < right && window < 0 {
			window -= nums[left]
			left++
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
