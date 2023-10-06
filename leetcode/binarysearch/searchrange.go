package binarysearch

func searchRange(nums []int, target int) []int {
	li := leftIndex(nums, target)
	ri := rightIndex(nums, target)
	return []int{li, ri}
}

func leftIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		m := nums[mid]
		if m == target {
			right = mid - 1
		} else if m < target {
			left = mid + 1
		} else if m > target {
			right = mid - 1
		}
	}

	if left == len(nums) || nums[right+1] != target {
		return -1
	}

	return right + 1
}

func rightIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		m := nums[mid]
		if m == target {
			left = mid + 1
		} else if m < target {
			left = mid + 1
		} else if m > target {
			right = mid - 1
		}
	}

	if right == -1 || nums[left-1] != target {
		return -1
	}

	return left - 1
}
