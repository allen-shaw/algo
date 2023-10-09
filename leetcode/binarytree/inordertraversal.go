package binarytree

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	traversal(root, &ans)
	return ans
}

func traversal(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, ans)
	*ans = append(*ans, root.Val)
	traversal(root.Right, ans)
}
