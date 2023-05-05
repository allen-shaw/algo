package tree

import (
	"fmt"
	"testing"
)

func buildTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	return root
}

func TestPreorderTraversal(t *testing.T) {
	root := buildTree()
	PreorderTraversal(root)

	fmt.Println("-----------")
	PreorderNonrecu(root)
}

func TestInorderTraversal(t *testing.T) {
	root := buildTree()
	InorderTraversal(root)

	fmt.Println("-----------")
	InorderNonrecu(root)
}

func TestPostorderTraversal(t *testing.T) {
	root := buildTree()
	PostorderTraversal(root)

	fmt.Println("-----------")
	PostorderNonrecu(root)
}

func TestBFS(t *testing.T) {
	root := buildTree()
	BFS(root)
}
