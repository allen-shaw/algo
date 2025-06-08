package zoom

func lowestCommonAncestor4(root *TreeNode, nodes []*TreeNode) *TreeNode {
	set := make(map[*TreeNode]bool)
	for _, node := range nodes {
		set[node] = true
	}
	var dfs func(root *TreeNode, nodes []*TreeNode) *TreeNode
	dfs = func(root *TreeNode, nodes []*TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if set[root] {
			return root
		}
		left := dfs(root.left, nodes)
		right := dfs(root.right, nodes)

		if left != nil && right != nil {
			return root
		}
		if left != nil {
			return left
		}
		return right
	}
	return dfs(root, nodes)
}
