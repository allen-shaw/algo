package nsums

import (
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)

	ans := 0

	for i := 0; i < len(nums); i++ {
		ans += twoSumSmaller(nums[i+1:], target-nums[i])
	}

	return ans
}

func twoSumSmaller(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := 0

	for l < r {
		left, right := nums[l], nums[r]
		sum := left + right

		if sum < target {
			ans += r - l
			l++
		} else if sum >= target {
			r--
		}
	}

	// fmt.Println(nums, target, ans)
	return ans
}
