package nsums

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 1. sort nums
	sort.Ints(nums)

	return threeSumTarget(nums, 0)
}

func threeSumTarget(nums []int, target int) [][]int {
	res := make([][]int, 0)
	// for loop nums
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		t := target - n

		// call 2sum
		rs := twoSumTarget(nums[i+1:], t)
		for _, r := range rs {
			fmt.Println(n, r)
			res = append(res, append(r, n))
		}

		for i < len(nums)-1 && nums[i+1] == nums[i] {
			i++
		}
	}

	return res
}

func twoSumTarget(nums []int, target int) [][]int {
	res := make([][]int, 0)

	left, right := 0, len(nums)-1

	for left < right {
		nleft, nright := nums[left], nums[right]
		sum := nleft + nright

		if sum == target {
			res = append(res, []int{nleft, nright})
			for left < right && nums[left] == nleft {
				left++
			}
			for left < right && nums[right] == nright {
				right--
			}
		} else if sum < target {
			for left < right && nums[left] == nleft {
				left++
			}
		} else if sum > target {
			for left < right && nums[right] == nright {
				right--
			}
		}
	}

	return res
}
