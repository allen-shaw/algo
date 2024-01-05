package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := make([]*Node, 0)
	q = append(q, root)

	dummy := &Node{}
	pre := dummy

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:] // pop

			pre.Next = node
			pre = pre.Next

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		pre = dummy
	}

	return root
}
