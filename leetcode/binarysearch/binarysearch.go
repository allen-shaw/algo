package binarysearch

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		num := nums[mid]
		if num == target {
			return mid
		} else if num < target {
			left = mid + 1
		} else if num > target {
			right = mid - 1
		}
	}

	return -1
}

func bsLeftBound(nums []int, target int) int {
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

	if right == len(nums) || nums[right+1] != target {
		return -1
	}
	return right + 1
}

func bsRightBound(nums []int, target int) int {
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
