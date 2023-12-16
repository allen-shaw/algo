package nsums

func twoSumBsts(root1, root2 *TreeNode, k int) bool {
	nums1 := InorderTravverse(root1)
	nums2 := InorderTravverse(root2)

	l, r := 0, len(nums2)-1

	for l < len(nums1) && r >= 0 {
		sum := nums1[l] + nums2[r]
		if sum == k {
			return true
		} else if sum < k {
			l++
		} else if sum > k {
			r--
		}
	}

	return false
}
