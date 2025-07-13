package agoda

import (
	"fmt"
	"testing"
)

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}

	fmt.Println(right)

	if nums[right+1] == target {
		return right + 1
	}
	return -1
}

func Test_leftBound(t *testing.T) {
	nums := []int{2, 3, 5, 7}
	ans := leftBound(nums, 4)
	fmt.Println(ans)
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}

	fmt.Println(left)
	if nums[left-1] == target {
		return left - 1
	}
	return -1
}

func Test_rightBound(t *testing.T) {
	// nums := []int{1, 2, 2, 2, 3}
	nums := []int{2, 3, 5, 7}
	ans := rightBound(nums, 4)
	fmt.Println(ans)
}
