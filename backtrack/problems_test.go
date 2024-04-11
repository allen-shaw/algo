package backtrack

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	bucket := make([]int, k)
	var dfs func(idx int) bool //
	dfs = func(idx int) bool {
		if idx == len(nums) {
			return true
		}

		for i := 0; i < k; i++ {
			if bucket[i]+nums[idx] > target {
				continue
			}
			if i > 0 && bucket[i] == bucket[i-1] {
				continue
			}

			bucket[i] += nums[idx]
			if dfs(idx + 1) {
				return true
			}
			bucket[i] -= nums[idx]
		}
		return false
	}

	return dfs(0)
}
