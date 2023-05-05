package tree

import (
	"fmt"

	"github.com/allen-shaw/algo/tree/queue"
	"github.com/allen-shaw/algo/tree/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{Val: v}
}

func PreorderTraversal(n *TreeNode) {
	if n == nil {
		return
	}

	fmt.Println(n.Val)
	PreorderTraversal(n.Left)
	PreorderTraversal(n.Right)
}

func InorderTraversal(n *TreeNode) {
	if n == nil {
		return
	}

	InorderTraversal(n.Left)
	fmt.Println(n.Val)
	InorderTraversal(n.Right)
}

func PostorderTraversal(n *TreeNode) {
	if n == nil {
		return
	}

	PostorderTraversal(n.Left)
	PostorderTraversal(n.Right)
	fmt.Println(n.Val)
}

func PreorderNonrecu(n *TreeNode) {
	if n == nil {
		return
	}
	s := stack.New[*TreeNode]()
	s.Push(n)

	for !s.Empty() {
		n := s.Pop()
		fmt.Println(n.Val)
		if n.Right != nil {
			s.Push(n.Right)
		}
		if n.Left != nil {
			s.Push(n.Left)
		}
	}
}

func InorderNonrecu(n *TreeNode) {
	if n == nil {
		return
	}

	s := stack.New[*TreeNode]()

	pushLeftToStack := func(s *stack.Stack[*TreeNode], n *TreeNode) {
		for n != nil {
			s.Push(n)
			n = n.Left
		}
	}

	// 所有左节点入栈
	pushLeftToStack(s, n)

	for !s.Empty() {
		// 打印最左节点
		n := s.Pop()
		fmt.Println(n.Val)
		// 当前节点的右节点，左孩子全部入栈
		if n.Right != nil {
			pushLeftToStack(s, n.Right)
		}
	}
}

func PostorderNonrecu(n *TreeNode) {
	if n == nil {
		return
	}

	s1 := stack.New[*TreeNode]()
	s2 := stack.New[*TreeNode]()

	s1.Push(n)
	for !s1.Empty() {
		n := s1.Pop()
		s2.Push(n)

		// 头-> 左-> 右
		if n.Left != nil {
			s1.Push(n.Left)
		}
		if n.Right != nil {
			s1.Push(n.Right)
		}
	}

	for !s2.Empty() {
		fmt.Println(s2.Pop().Val)
	}
}

func BFS(n *TreeNode) {
	if n == nil {
		return
	}

	q := queue.New[*TreeNode]()
	q.Enqueue(n)

	for !q.Empty() {
		n := q.Dequeue()
		fmt.Println(n.Val)

		if n.Left != nil {
			q.Enqueue(n.Left)
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
		}
	}
}
