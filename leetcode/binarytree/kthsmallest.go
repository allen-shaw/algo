package binarytree

type myTreeNode struct {
	val   int
	index int
	left  *myTreeNode
	right *myTreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var myRoot *myTreeNode
	sum := 0
	bstToSmallTree(root, &myRoot, &sum)
	return findKth(myRoot, k)
}

func bstToSmallTree(root *TreeNode, myRoot **myTreeNode, sum *int) {
	if root == nil {
		return
	}

	*myRoot = &myTreeNode{
		val: root.Val,
	}

	bstToSmallTree(root.Left, &(*myRoot).left, sum)
	*sum += 1
	(*myRoot).index += *sum
	bstToSmallTree(root.Right, &(*myRoot).right, sum)
}

func findKth(root *myTreeNode, k int) int {
	if root == nil {
		return -1
	}
	if root.index == k {
		return root.val
	} else if root.index < k {
		return findKth(root.right, k)
	} else {
		return findKth(root.left, k)
	}
}
