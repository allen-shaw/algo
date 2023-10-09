package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(arr []int) *TreeNode {
	nodes := make([]*TreeNode, len(arr))
	for i := range arr {
		n := arr[i]
		if n != -1 {
			nodes[i] = &TreeNode{Val: n}
		}
	}

	for i := 0; 2*i+2 < len(nodes); i++ {
		if nodes[i] != nil {
			nodes[i].Left = nodes[2*i+1]
			nodes[i].Right = nodes[2*i+2]
		}
	}

	return nodes[0]
}

type queue[T any] struct {
	data []T
}

func newQueue[T any]() *queue[T] {
	return &queue[T]{
		data: make([]T, 0),
	}
}

func (q *queue[T]) Push(x T) {
	q.data = append(q.data, x)
}

func (q *queue[T]) Pop() T {
	x := q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *queue[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *queue[T]) Size() int {
	return len(q.data)
}

func printTree(root *TreeNode) {
	q := newQueue[*TreeNode]()
	q.Push(root)

	for !q.Empty() {
		size := q.Size()
		for i := 0; i < size; i++ {
			n := q.Pop()
			fmt.Printf("%v ", n.Val)
			if n.Left != nil {
				q.Push(n.Left)
			}
			if n.Right != nil {
				q.Push(n.Right)
			}
		}
		fmt.Println()
	}
}
