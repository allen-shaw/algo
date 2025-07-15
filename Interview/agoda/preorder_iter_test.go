package agoda

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Iterator struct {
	stack []*TreeNode
}

func NewIterator(root *TreeNode) Iterator {
	it := Iterator{
		stack: make([]*TreeNode, 0),
	}
	it.stack = append(it.stack, root)
	return it
}

func (it *Iterator) Next() int {
	n := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	if n.Right != nil {
		it.stack = append(it.stack, n.Right)
	}
	if n.Left != nil {
		it.stack = append(it.stack, n.Left)
	}

	return n.Val
}

func (it *Iterator) HasNext() bool {
	return len(it.stack) != 0
}
