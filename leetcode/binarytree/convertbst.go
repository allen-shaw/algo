package binarytree

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convertbst(root, &sum)
	return root
}

func convertbst(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	convertbst(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	convertbst(root.Left, sum)
}
