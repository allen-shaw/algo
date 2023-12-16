package nsums

import (
	"math"
	"sort"
)

func twoSumSmallerThanK(nums []int, k int) int {
	ans := 0
	delta := math.MaxInt
	sort.Ints(nums)

	l, r := 0, len(nums)-1

	for l < r {
		left, right := nums[l], nums[r]
		sum := left + right

		if sum < k && abs(k-sum) < abs(delta) {
			delta = k - sum
			ans = sum
		}

		if sum < k {
			l++
		} else if sum >= k {
			r--
		}
	}

	if delta == math.MaxInt {
		return -1
	}
	return ans
}
