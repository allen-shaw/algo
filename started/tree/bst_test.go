package tree

import (
	"fmt"
	"testing"
)

type TreeNode[T any] struct {
	left, right *TreeNode[T]
	val         T
}

func BFS2[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}

	q := make([]*TreeNode[T], 0)
	q = append(q, root)
	depth := 1

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:] // q.pop

			fmt.Printf("%v: %v\n", node.val, depth)
			if node.left != nil {
				q = append(q, node.left)
			}
			if node.right != nil {
				q = append(q, node.right)
			}
		}
		depth++
	}
}

type state[T any] struct {
	n     *TreeNode[T]
	depth int
}

func newState[T any](n *TreeNode[T], depth int) state[T] {
	return state[T]{n: n, depth: depth}
}

func BFS3[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}

	q := make([]state[T], 0)
	q = append(q, newState(root, 1))

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			s := q[0]
			q = q[1:]

			n := s.n
			fmt.Printf("%v: %v\n", n.val, s.depth)
			if n.left != nil {
				q = append(q, newState(n.left, s.depth+1))
			}
			if n.right != nil {
				q = append(q, newState(n.right, s.depth+1))
			}
		}
	}
}

func TestBFS(t *testing.T) {
	root := buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	BFS2(root)
	fmt.Println()
	BFS3(root)
}

func buildTree(arr []int) *TreeNode[int] {
	if len(arr) == 0 {
		return nil
	}

	nodes := make([]*TreeNode[int], len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		nodes[i] = &TreeNode[int]{val: arr[i]}
		if i < len(arr)/2 {
			if i*2+1 < len(arr) {
				nodes[i].left = nodes[i*2+1]
			}
			if i*2+2 < len(arr) {
				nodes[i].right = nodes[i*2+2]
			}
		}
	}

	return nodes[0]
}
