package binarytree

var (
	rightSideViewAns                []int
	rightSideViewDepth int
)

func rightSideView(root *TreeNode) []int {
	rightSideViewAns = make([]int, 0)
	depth = 0
	rightSideViewTraverse(root)
	return rightSideViewAns
}

func rightSideViewTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	rightSideViewDepth++
	if len(rightSideViewAns) < rightSideViewDepth {
		rightSideViewAns = append(rightSideViewAns, root.Val)
	}
	rightSideViewTraverse(root.Right)
	rightSideViewTraverse(root.Left)
	rightSideViewDepth--
}
