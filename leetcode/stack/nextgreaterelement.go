package stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	glist := genMonotonicStack(nums2)
	m := make(map[int]int, len(glist))
	for i, num := range glist {
		k := nums2[i]
		m[k] = num
	}

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		ans[i] = m[num]
	}
	return ans
}
