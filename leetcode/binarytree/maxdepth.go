package binarytree

func maxDepth(root *TreeNode) int {
	maxDepthTraversal(root)
	return maxdepth
}

var (
	depth    = 0
	maxdepth = 0
)

func maxDepthTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	depth++
	maxdepth = max(depth, maxdepth)
	maxDepthTraversal(root.Left)
	maxDepthTraversal(root.Right)
	depth--
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
