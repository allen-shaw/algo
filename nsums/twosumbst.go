package nsums

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return findTarget2(root, m, k)
}

func findTarget2(root *TreeNode, m map[int]struct{}, target int) bool {
	if root == nil {
		return false
	}

	if _, ok := m[target-root.Val]; ok {
		return true
	}
	m[root.Val] = struct{}{}

	return findTarget2(root.Left, m, target) || findTarget2(root.Right, m, target)
}
