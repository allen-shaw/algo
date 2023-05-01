package tree

import (
	"fmt"
	"testing"
)

func buildTree() *Node {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Right = &Node{Val: 6}

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
