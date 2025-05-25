package zoom

import (
	"fmt"
	"testing"
)

// 4 4 5 6 1 2 3
func search(nums []int, target int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}
	leftIdx := findLeftIndex(nums, target)
	if leftIdx == -1 {
		return -1, -1
	}
	rightIdx := findRightIndex(nums, target)
	return leftIdx, rightIdx
}

func findLeftIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1

	pos := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			pos = mid
			right = mid - 1
		} else {
			if nums[left] == nums[mid] && nums[mid] == nums[right] {
				if nums[left] == target {
					return left
				}
				left++
				right--
			} else if nums[left] <= nums[mid] {
				// 左边有序
				if nums[left] <= target && target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				// 右边有序
				if nums[mid] < target && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	return pos
}

func findRightIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1
	pos := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			pos = mid
			// 找右边界，因此继续往右找
			left = mid + 1
		} else {
			if nums[left] == nums[mid] && nums[mid] == nums[right] {
				if nums[right] == target {
					return right
				}
				left++
				right--
			} else if nums[left] <= nums[mid] {
				// 左侧有序
				if nums[left] <= target && target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				// 右侧有序
				if nums[mid] < target && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}

	return pos
}

func TestSearch(t *testing.T) {
	nums1 := []int{6, 6, 7, 0, 1, 2, 4, 4, 5}
	target1 := 4
	fmt.Println(search(nums1, target1))
}
