package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepthWithBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	md := 1
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if n.Left == nil && n.Right == nil {
				return md
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		md++
	}

	return md
}

func minDepthWithDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minDepth := 0
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			if minDepth == 0 {
				minDepth = depth
			} else {
				minDepth = min(minDepth, depth)
			}
			return
		}

		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}

	dfs(root, 1)
	return minDepth
}
