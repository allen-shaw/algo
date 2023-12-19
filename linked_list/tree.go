package linkedlist

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTreverse(root *TreeNode) []int {
	nums := make([]int, 0)
	inorderTreverse(root, &nums)
	return nums
}

func inorderTreverse(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}

	inorderTreverse(root.Left, out)
	*out = append(*out, root.Val)
	inorderTreverse(root.Right, out)
}
