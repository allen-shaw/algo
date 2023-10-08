package binarysearch

import "math"

func findRotatedSortedMin(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	if nums[left] <= nums[right] {
		return nums[left]
	}

	ans := math.MaxInt

	for left <= right {
		mid := left + (right-left)/2
		if nums[left] <= nums[mid] {
			// left 到 mid 有序
			ans = min(nums[left], ans)
			left = mid + 1
		} else {
			// mid 到 right 有序
			ans = min(nums[mid], ans)
			right = mid - 1
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
