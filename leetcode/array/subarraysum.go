package array

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数
// 子数组是数组中元素的连续非空序列。
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	presum := buildPrefixSum(nums)
	m := make(map[int]int)
	ans := 0

	for _, s := range presum {
		if v, ok := m[s-k]; ok {
			ans += v
		}
		m[s]++
	}

	return ans
}

func buildPrefixSum(arr []int) []int {
	ps := make([]int, len(arr)+1)

	for i, n := range arr {
		ps[i+1] = ps[i] + n
	}
	return ps
}
