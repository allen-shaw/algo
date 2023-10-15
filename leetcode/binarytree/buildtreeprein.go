package binarytree

var indexMap = make(map[int]int)

func buildTreePreAndInOrder(preorder []int, inorder []int) *TreeNode {
	for i, num := range inorder {
		indexMap[num] = i
	}
	return buildTreePreAndIn(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func buildTreePreAndIn(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	index := indexMap[rootVal]

	leftSize := index - inStart
	// rightSize := inEnd - index

	root.Left = buildTreePreAndIn(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)
	root.Right = buildTreePreAndIn(preorder, preStart+1+leftSize, preEnd, inorder, index+1, inEnd)

	return root
}
