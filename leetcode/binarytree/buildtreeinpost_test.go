package binarytree

import "testing"

func TestBuildTreeInAndPostOrder(t *testing.T) {
	//inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTreeInAndPostOrder(inorder, postorder)
	printTree(root)
}
