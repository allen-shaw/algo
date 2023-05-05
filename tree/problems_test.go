package tree

import (
	"fmt"
	"testing"

	"github.com/allen-shaw/algo/tree/queue"
)

// 求一颗二叉树最大宽度
func MaxWidth(n *TreeNode) int {
	if n == nil {
		return 0
	}

	levelMap := make(map[*TreeNode]int)
	maxWidth := 0
	curLvl := 1
	curLvlCnt := 0

	q := queue.New[*TreeNode]()
	q.Enqueue(n)
	levelMap[n] = 1

	for !q.Empty() {
		n := q.Dequeue()
		lvl := levelMap[n]

		if lvl != curLvl {
			// 到下一层了，进行数据结算
			maxWidth = maxInt(curLvlCnt, maxWidth)

			// 结算后
			curLvl = lvl
			curLvlCnt = 0
		}

		curLvlCnt++
		if n.Left != nil {
			q.Enqueue(n.Left)
			levelMap[n.Left] = curLvl + 1
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
			levelMap[n.Right] = curLvl + 1
		}
	}

	// 最后一层结算
	fmt.Println(curLvl, curLvlCnt)
	maxWidth = maxInt(curLvlCnt, maxWidth)

	return maxWidth
}

func buildTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Right.Left.Right = &TreeNode{Val: 8}
	return root
}

func TestMaxWidth(t *testing.T) {
	root := buildTree2()
	n := MaxWidth(root)
	fmt.Println("max width", n)
	n2 := MaxWidth2(root)
	fmt.Println("max width", n2)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxWidth2(h *TreeNode) int {
	if h == nil {
		return 0
	}

	maxWidth := 0
	curLvlCnt := 0
	var (
		curLvlEnd  *TreeNode
		nextLvlEnd *TreeNode
	)

	q := queue.New[*TreeNode]()
	q.Enqueue(h)
	curLvlEnd = h

	for !q.Empty() {
		n := q.Dequeue()
		curLvlCnt++
		if n.Left != nil {
			q.Enqueue(n.Left)
			nextLvlEnd = n.Left
		}
		if n.Right != nil {
			q.Enqueue(n.Right)
			nextLvlEnd = n.Right
		}

		if n == curLvlEnd {
			// 当前层到尾部,结算
			maxWidth = maxInt(curLvlCnt, maxWidth)

			curLvlCnt = 0
			curLvlEnd = nextLvlEnd
			nextLvlEnd = nil
		}
	}

	maxWidth = maxInt(curLvlCnt, maxWidth)
	return maxWidth
}

// 判断一棵二叉树是不是搜索二叉树
func IsBST(h *TreeNode) bool {
	_, _, ok := isBst(h)
	return ok
}

func isBst(n *TreeNode) (int, int, bool) {
	if n == nil {
		return 0, 0, true
	}

	lmin, lmax, lok := isBst(n.Left)
	rmin, rmax, rok := isBst(n.Right)

	ok := lok && rok && (lmax <= n.Val && n.Val <= rmin)
	min := minInts(lmin, rmin, n.Val)
	max := maxInts(lmax, rmax, n.Val)

	return min, max, ok
}

func TestIsBST(t *testing.T) {

}

// 判断一棵二叉树是不是完全二叉树

// 判断一棵二叉树是不是满二叉树

// 判断一棵二叉树是不是平衡二叉树
func IsBalanceTree(h *TreeNode) bool {
	_, ok := isBalanceTree(h)
	return ok
}

func isBalanceTree(n *TreeNode) (int, bool) {
	if n == nil {
		return 0, true
	}

	lh, lok := isBalanceTree(n.Left)
	rh, rok := isBalanceTree(n.Right)

	return maxInt(lh, rh) + 1, lok && rok && (absInt(lh-rh) <= 1)
}

func TestIsBalanceTree(t *testing.T) {

}

func absInt(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func maxInts(nums ...int) int {
	if len(nums) == 0 {
		panic("nil nums list")
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max > nums[i] {
			max = nums[i]
		}
	}
	return max
}

func minInts(nums ...int) int {
	if len(nums) == 0 {
		panic("nil nums list")
	}

	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if min < nums[i] {
			min = nums[i]
		}
	}
	return min
}

func TestForLoop(t *testing.T) {

	fn := func(i int) int {
		if i == 3 {
			return 1
		}
		return 5
	}

	for i := 0; i < fn(i); i++ {
		fmt.Println(i)
	}
}

const nilNum = -100

func TestLowestCommonAncestor(t *testing.T) {
	// root := buildATree(3, 5, 1, 6, 2, 0, 8, nilNum, nilNum, 7, 4)
	// lowestCommonAncestor1(root, )
}

func TestBuildATree(t *testing.T) {
	root := buildATree(3, 5, 1, 6, 2, 0, 8, nilNum, nilNum, 7, 4)
	//   	  3
	//  	5    1
	// 	   6 2  0 8
	//      7 4
	// 3 5 6 2 7 4 1 0 8
	PreorderTraversal(root)
}

func buildATree(nums ...int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(nums))
	for i, num := range nums {
		nodes[i] = &TreeNode{Val: num}
	}

	size := len(nodes)
	for i, node := range nodes {
		if node.Val == nilNum {
			continue
		}

		lcIdx := 2*i + 1
		rcIdx := 2*i + 2

		if lcIdx < size {
			lc := nodes[lcIdx]
			if lc.Val == nilNum {
				node.Left = nil
			} else {
				node.Left = lc
			}
		}

		if rcIdx < size {
			rc := nodes[rcIdx]
			if rc.Val == nilNum {
				node.Right = nil
			} else {
				node.Right = rc
			}
		}
	}

	return nodes[0]
}

// 求两个节点的最低公共节点
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	// 遍历求解所有的父节点
	m := make(map[*TreeNode]*TreeNode)
	set := make(map[*TreeNode]struct{})

	findParent(root, m)

	n := p
	for n != nil {
		set[n] = struct{}{}
		n = m[n]
	}

	n2 := q
	for {
		if a, ok := m[n2]; ok {
			return a
		}
		n2 = m[n2]
	}
}

func findParent(n *TreeNode, m map[*TreeNode]*TreeNode) {
	if n == nil {
		return
	}

	if n.Left != nil {
		m[n.Left] = n
		findParent(n.Left, m)
	}
	if n.Right != nil {
		m[n.Right] = n
		findParent(n.Right, m)
	}
}
