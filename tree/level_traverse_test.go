package tree

import (
	"fmt"
	"testing"
)

type MAryToBinaryTree struct {
}

type NTreeNode struct {
	val     int
	chidren []*NTreeNode
}

func printNTree(root *NTreeNode) {
	q := make([]*NTreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		val := make([]int, 0)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			val = append(val, n.val)
			for _, c := range n.chidren {
				q = append(q, c)
			}
		}
		fmt.Println(val)
	}
}

func printBTree(root *TreeNode) {
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		val := make([]int, 0)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			val = append(val, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		fmt.Println(val)
	}
}

func (*MAryToBinaryTree) Encode(root *NTreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tdummy := &TreeNode{}
	tp := tdummy

	q := make([]*NTreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)

		dummy := &TreeNode{}
		p := dummy

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			p.Right = &TreeNode{Val: n.val}
			p = p.Right

			q = append(q, n.chidren...)
		}

		tp.Left = dummy.Right
		tp = tp.Left
	}

	return tdummy.Left
}

type nodePair struct {
	n *NTreeNode
	b *TreeNode
}

func (*MAryToBinaryTree) Decode(root *TreeNode) *NTreeNode {
	if root == nil {
		return nil
	}

	ntRoot := &NTreeNode{val: root.Val}

	q := make([]*nodePair, 0)
	q = append(q, &nodePair{ntRoot, root})

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			pair := q[0]
			n := pair.b
			q = q[1:]

			children := make([]*NTreeNode, 0)

			for p := n.Left; p != nil; p = p.Right {
				ntn := &NTreeNode{val: p.Val}
				q = append(q, &nodePair{ntn, p})
				children = append(children, ntn)
			}
			pair.n.chidren = children
		}
	}

	return ntRoot
}

func Test_MAryToBinaryTree(t *testing.T) {
	root := &NTreeNode{val: 1}
	root.chidren = []*NTreeNode{
		{val: 2, chidren: []*NTreeNode{{val: 5}, {val: 6}}},
		{val: 3},
		{val: 4, chidren: []*NTreeNode{{val: 7}, {val: 8}}},
	}
	fmt.Println("source:")
	printNTree(root)

	// encode
	obj := &MAryToBinaryTree{}
	broot := obj.Encode(root)

	fmt.Println("encode: binaryTree:")
	printBTree(broot)

	// decode
	root = obj.Decode(broot)

	fmt.Println("decode: NTree:")
	printNTree(root)

}
