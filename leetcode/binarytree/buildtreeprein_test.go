package binarytree

import "testing"

func TestBuildTreePreAndInOrder(t *testing.T) {
	// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTreePreAndInOrder(preorder, inorder)
	printTree(root)
}
