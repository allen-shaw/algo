package zoom

type TreeNode struct {
	val         int
	left, right *TreeNode
}

// 给定一棵二叉树的根节点 root，返回给定节点 p 和 q 的最近公共祖先（LCA）节点。
// 如果 p 或 q 之一 不存在 于该二叉树中，返回 null。树中的每个节点值都是互不相同的。
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode
	var dfs func(root *TreeNode, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		leftFound := dfs(root.left, p, q)
		rightFound := dfs(root.right, p, q)

		if leftFound != nil && rightFound != nil {
			lca = root
			return root
		}
		if (leftFound != nil || rightFound != nil) && (root == p || root == q) {
			lca = root
			return root
		}

		if leftFound != nil {
			return leftFound
		}
		return rightFound
	}

	dfs(root, p, q)
	return lca
}
