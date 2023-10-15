package binarytree

var indexmap = make(map[int]int)

func buildTreeInAndPostOrder(inorder []int, postorder []int) *TreeNode {
	for i, n := range inorder {
		indexmap[n] = i
	}
	return buildTreeInAndPost(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildTreeInAndPost(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	// fmt.Printf("instart %v, inend %v; poststart %v, postend %v \n", inStart, inEnd, postStart, postEnd)
	if inStart > inEnd {
		return nil
	}

	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}
	index := indexmap[rootVal]
	leftSize := index - inStart
	// rightSize := inEnd - index

	root.Left = buildTreeInAndPost(inorder, inStart, index-1, postorder, postStart, postStart+leftSize-1)
	root.Right = buildTreeInAndPost(inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1)

	return root
}
