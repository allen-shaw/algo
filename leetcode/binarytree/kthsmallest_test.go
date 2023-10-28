package binarytree

import (
	"fmt"
	"testing"
)

func TestBstToSmallTree(t *testing.T) {
	arr := []int{3, 1, 4, -1, 2}
	root := buildTree(arr)

	var myRoot *myTreeNode
	sum := 0
	bstToSmallTree(root, &myRoot, &sum)

	printMyTree(myRoot)

	val := findKth(myRoot, 1)
	fmt.Println(val)
}

func printMyTree(root *myTreeNode) {
	q := newQueue[*myTreeNode]()
	q.Push(root)

	for !q.Empty() {
		size := q.Size()
		for i := 0; i < size; i++ {
			n := q.Pop()
			fmt.Printf("%v:%v ", n.val, n.index)
			if n.left != nil {
				q.Push(n.left)
			}
			if n.right != nil {
				q.Push(n.right)
			}
		}
		fmt.Println()
	}
}

func testChangePoint(p *TreeNode) {
	p = &TreeNode{Val: 100}
}
func TestChangePoint(t *testing.T) {
	var p *TreeNode
	testChangePoint(p)
	fmt.Println(p)
}
