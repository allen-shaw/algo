package kthsmallest

import "math/rand/v2"

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k

	left, right := 0, len(nums)-1

	for left <= right {
		pivot := partition(nums, left, right)
		if pivot == k {
			return nums[pivot]
		} else if pivot < k {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}

func partition(nums []int, left, right int) int {
	p := left + rand.IntN(right-left+1)
	pivot := nums[p]
    nums[p], nums[left] = nums[left], nums[p]

	i, j := left+1, right

	for i <= j {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

}
