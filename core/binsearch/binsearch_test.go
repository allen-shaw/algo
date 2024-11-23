package binsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

func Test_binarySearch(t *testing.T) {
	cases := []struct {
		name  string
		input struct {
			nums   []int
			target int
		}
		expect int
	}{
		{"case1", struct {
			nums   []int
			target int
		}{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{"case2", struct {
			nums   []int
			target int
		}{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := binarySearch(c.input.nums, c.input.target)
			assert.Equal(t, c.expect, out)
		})
	}
}

func leftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if nums[right+1] == target {
		return right + 1
	}

	return -1
}

func Test_leftBound(t *testing.T) {
	cases := []struct {
		name  string
		input struct {
			nums   []int
			target int
		}
		expect int
	}{
		{"case1", struct {
			nums   []int
			target int
		}{[]int{5, 7, 7, 8, 8, 10}, 8}, 3},
		{"case2", struct {
			nums   []int
			target int
		}{[]int{5, 7, 7, 8, 8, 10}, 6}, -1},
		{"case3", struct {
			nums   []int
			target int
		}{[]int{}, 0}, -1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := leftBound(c.input.nums, c.input.target)
			assert.Equal(t, c.expect, out)
		})
	}
}

func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if nums[left-1] == target {
		return left - 1
	}

	return -1
}

func Test_rightBound(t *testing.T) {
	cases := []struct {
		name  string
		input struct {
			nums   []int
			target int
		}
		expect int
	}{
		{"case1", struct {
			nums   []int
			target int
		}{[]int{5, 7, 7, 8, 8, 10}, 8}, 4},
		{"case2", struct {
			nums   []int
			target int
		}{[]int{5, 7, 7, 8, 8, 10}, 6}, -1},
		{"case3", struct {
			nums   []int
			target int
		}{[]int{}, 0}, -1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := rightBound(c.input.nums, c.input.target)
			assert.Equal(t, c.expect, out)
		})
	}
}
