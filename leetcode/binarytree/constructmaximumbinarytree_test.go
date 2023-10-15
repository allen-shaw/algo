package binarytree

import "testing"

func TestConstructMaximumBinaryTree(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	printTree(root)
}
