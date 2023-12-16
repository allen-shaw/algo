package nsums

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	return nSum(nums, target, 4)
}

func nSum(nums []int, target int, n int) [][]int {
	if n == 2 {
		return twoSumTarget1(nums, target)
	}

	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		t := target - num

		res := nSum(nums[i+1:], t, n-1)
		for _, r := range res {
			ans = append(ans, append(r, num))
		}

		for i < len(nums)-1 && nums[i+1] == num {
			i++
		}
	}

	return ans
}

func twoSumTarget1(nums []int, target int) [][]int {
	ans := make([][]int, 0)

	l, r := 0, len(nums)-1

	for l < r {
		left, right := nums[l], nums[r]
		sum := left + right

		if sum == target {
			ans = append(ans, []int{left, right})
			for l < r && nums[l] == left {
				l++
			}
			for l < r && nums[r] == right {
				r--
			}
		} else if sum < target {
			for l < r && nums[l] == left {
				l++
			}
		} else if sum > target {
			for l < r && nums[r] == right {
				r--
			}
		}
	}

	return ans
}
