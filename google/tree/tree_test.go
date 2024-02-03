package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ans []*TreeNode
)

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	ans = make([]*TreeNode, 0)
	toDelete := make(map[int]struct{})
	for _, n := range to_delete {
		toDelete[n] = struct{}{}
	}

	doDelNodes(root, true, toDelete)
	return ans
}

func doDelNodes(root *TreeNode, isRoot bool, toDelete map[int]struct{}) *TreeNode {
	if root == nil {
		return nil
	}

	_, ok := toDelete[root.Val]

	root.Left = doDelNodes(root.Left, ok, toDelete)
	root.Right = doDelNodes(root.Right, ok, toDelete)

	if ok {
		return nil
	}

	if isRoot {
		ans = append(ans, root)
	}
	return root
}
