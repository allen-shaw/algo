package agoda

import "math/rand/v2"

func findKthLargest(nums []int, k int) int {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	left, right := 0, len(nums)
	for left <= right {
		p := partition(nums, left, right)
		if p == k {
			return nums[p]
		} else if p < k {
			left = p + 1
		} else if p > k {
			right = p - 1
		}
	}

	return -1
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	i, j := start+1, end
	for i <= j {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] >= pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[j], nums[start] = nums[start], nums[j]
	return j
}
