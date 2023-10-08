package binarysearch

func searchRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		m := nums[mid]
		if m == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			// left -> mid有序
			if target >= nums[left] && target < m {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// mid -> right有序
			if target > m && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
