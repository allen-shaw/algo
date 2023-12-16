package nsums

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := math.MaxInt

	for i := 0; i < len(nums); i++ {
		sum := nums[i] + twoSumClosest(nums[i+1:], target-nums[i])
		if abs(target-sum) < abs(min) {
			min = target - sum
		}
	}

	return target - min
}

func twoSumClosest(nums []int, target int) int {
	min := math.MaxInt
	l, r := 0, len(nums)-1

	for l < r {
		left, right := nums[l], nums[r]
		sum := left + right
		if abs(target-sum) < abs(min) {
			min = target - sum
		}

		if sum <= target {
			l++
		} else if sum > target {
			r--
		}
	}

	return target - min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
