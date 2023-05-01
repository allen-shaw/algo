package tree

import (
	"fmt"
	"testing"

	"github.com/allen-shaw/algo/tree/queue"
)

// 求一颗二叉树最大宽度
func MaxWidth(n *Node) int {
	if n == nil {
		return 0
	}

	levelMap := make(map[*Node]int)
	maxWidth := 0
	curLvl := 1
	curLvlCnt := 0

	q := queue.New[*Node]()
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

func buildTree2() *Node {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}
	root.Right.Left = &Node{Val: 6}
	root.Left.Right.Left = &Node{Val: 7}
	root.Right.Left.Right = &Node{Val: 8}
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

func MaxWidth2(h *Node) int {
	if h == nil {
		return 0
	}

	maxWidth := 0
	curLvlCnt := 0
	var (
		curLvlEnd  *Node
		nextLvlEnd *Node
	)

	q := queue.New[*Node]()
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
func IsBST(h *Node) bool {

}

func isBst(n *Node) (int, int, bool) {
	if n == nil {
		return 0, 0, true
	}

	lmin, lmax, lok := isBst(n.Left)
	rmin, rmax, rok := isBst(n.Right)

	ok := lok && rok && (lmax <= n.Val && n.Val <= rmin)
	min := 
	return
}

func TestIsBST(t *testing.T) {

}

// 判断一棵二叉树是不是完全二叉树

// 判断一棵二叉树是不是满二叉树

// 判断一棵二叉树是不是平衡二叉树
func IsBalanceTree(h *Node) bool {
	_, ok := isBalanceTree(h)
	return ok
}

func isBalanceTree(n *Node) (int, bool) {
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
