package nsums

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}

			if sum == target {
				return sum
			} else if sum < target {
				j++
			} else if sum > target {
				k--
			}
		}
	}

	return ans
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
