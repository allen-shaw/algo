package nsums

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTravverse(n *TreeNode) []int {
	ans := make([]int, 0)
	inorderTravverse(n, &ans)
	return ans
}

func inorderTravverse(n *TreeNode, nums *[]int) {
	if n == nil {
		return
	}

	inorderTravverse(n.Left, nums)
	*nums = append(*nums, n.Val)
	inorderTravverse(n.Right, nums)
}

func buildTree(nums []int) *TreeNode {
	q := newQueue[**TreeNode]()
	root := &TreeNode{Val: nums[0]}
	q.Push(&root)

	i := 0
	for !q.Empty() && i < len(nums) {
		n := q.Pop()
		if nums[i] == math.MinInt {
			continue
		}
		*n = &TreeNode{Val: nums[i]}
		q.Push(&(*n).Left)
		q.Push(&(*n).Right)
		i++
	}

	return root
}

type Queue[T any] struct {
	data []T
}

func newQueue[T any]() *Queue[T] {
	q := &Queue[T]{}
	q.data = make([]T, 0)
	return q
}

func (q *Queue[T]) Push(x T) {
	q.data = append(q.data, x)
}

func (q *Queue[T]) Pop() (x T) {
	x = q.data[0]
	q.data = q.data[1:]
	return x
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}
