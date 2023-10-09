package binarytree

import "testing"

func TestBuildTree(t *testing.T) {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	tree := buildTree(nums)
	printTree(tree)
}
