package binarytree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaxBTree(nums, 0, len(nums)-1)
}

func constructMaxBTree(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	max, index := findmax(nums, start, end)
	root := &TreeNode{Val: max}
	root.Left = constructMaxBTree(nums, start, index-1)
	root.Right = constructMaxBTree(nums, index+1, end)

	return root
}

func findmax(nums []int, start, end int) (int, int) {
	max := nums[start]
	idx := start

	for i := start; i <= end; i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}

	return max, idx
}
