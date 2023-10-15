package binarytree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lcaTraversal(root, p, q)
}

func lcaTraversal(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lcaTraversal(root.Left, p, q)
	right := lcaTraversal(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}
