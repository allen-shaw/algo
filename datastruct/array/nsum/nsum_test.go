package nsum

import (
	"fmt"
	"sort"
	"testing"
)

func twoSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	left, right := 0, len(nums)-1

	for left < right {
		for left > 0 && left < right && nums[left-1] == nums[left] {
			left++
		}
		for right < len(nums)-1 && left < right && nums[right+1] == nums[right] {
			right--
		}

		sum := nums[left] + nums[right]
		if sum == target {
			ans = append(ans, []int{nums[left], nums[right]})
			left++
			right--
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return ans
}

func Test_twoSumTarget(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3, 3}
	target := 4
	ans := twoSumTarget(nums, target)
	fmt.Println(ans)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var twoSum = func(nums []int, target, start int) [][]int {
		ans := make([][]int, 0)
		l, r := start, len(nums)-1
		for l < r {
			for l > start && l < r && nums[l] == nums[l-1] {
				l++
			}
			for r < len(nums)-1 && l < r && nums[r] == nums[r+1] {
				r--
			}
			if l >= r {
				break
			}
			sum := nums[l] + nums[r]
			if sum == target {
				ans = append(ans, []int{nums[l], nums[r]})
				l++
				r--
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
		return ans
	}

	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			continue
		}

		res := twoSum(nums, -nums[i], i+1)
		for j := range res {
			res[j] = append(res[j], nums[i])
		}
		fmt.Println("target:", nums[i], " list:", res)
		ans = append(ans, res...)
	}

	return ans
}

func Test_threeSum(t *testing.T) {
	nums := []int{-2, 0, 0, 2, 2}
	ans := threeSum(nums)
	fmt.Println(ans)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	var twoSum = func(nums []int, start, target int) [][]int {
		ans := make([][]int, 0)
		left, right := start, len(nums)-1
		for left < right {
			for left > start && left < right && nums[left] == nums[left]-1 {
				left++
			}
			for right < len(nums)-1 && left < right && nums[right] == nums[right+1] {
				right--
			}
			if left >= right {
				break
			}
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[left], nums[right]})
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
		return ans
	}

	var threeSum = func(nums []int, start, target int) [][]int {
		ans := make([][]int, 0)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			res := twoSum(nums, i+1, target-nums[i])
			for j := range res {
				res[j] = append(res[j], nums[i])
			}
			ans = append(ans, res...)
		}
		return ans
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		res := threeSum(nums, i+1, target-nums[i])
		for j := range res {
			res[j] = append(res[j], nums[i])
		}
		ans = append(ans, res...)
	}

	return ans
}

func Test_fourSum(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	target := 0
	out := fourSum(nums, target)
	fmt.Println(out)
}

func nSum(nums []int, target, n int) [][]int {
	sort.Ints(nums)

	var twoSum = func(nums []int, start, target int) [][]int {
		ans := make([][]int, 0)
		left, right := start, len(nums)-1
		for left < right {
			for left > start && left < right && nums[left] == nums[left]-1 {
				left++
			}
			for right < len(nums)-1 && left < right && nums[right] == nums[right+1] {
				right--
			}
			if left >= right {
				break
			}
			sum := nums[left] + nums[right]
			if sum == target {
				ans = append(ans, []int{nums[left], nums[right]})
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
		return ans
	}

	var nsum func(nums []int, n, start, target int) [][]int
	nsum = func(nums []int, n int, start int, target int) [][]int {
		ans := make([][]int, 0)
		if n == 2 {
			return twoSum(nums, start, target)
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			num := nums[i]
			res := nsum(nums, n-1, i+1, target-num)
			for k := range res {
				res[k] = append(res[k], num)
			}
			fmt.Println(n, "sum ", start, target, res)
			ans = append(ans, res...)
		}
		return ans
	}

	return nsum(nums, n, 0, target)
}

func Test_nSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	ans := nSum(nums, target, 4)
	fmt.Println(ans)
}
