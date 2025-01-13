package recentcounter

import (
	"fmt"
	"testing"
)

type RecentCounter struct {
	root *Node
}

func Constructor() RecentCounter {
	n := 6000 // 1e9
	return RecentCounter{
		root: &Node{start: 0, end: n},
	}
}

func (c *RecentCounter) Ping(t int) int {
	c.root.Update(t, t, 1)
	left := max(t-3000, 0)
	return c.root.Query(left, t)
}

type Node struct {
	val         int
	left, right *Node
	lazy        int
	start, end  int
}

func (n *Node) Update(left, right, val int) {
	if left <= n.start && right >= n.end {
		n.val += val * (n.end - n.start + 1)
		n.lazy += val
		return
	}

	n.pushDown()

	mid := (n.start + n.end) / 2
	if left <= mid {
		n.left.Update(left, right, val)
	}
	if right > mid {
		n.right.Update(left, right, val)
	}

	n.pushUp()
}

func (n *Node) pushUp() {
	n.val = n.left.val + n.right.val
}

func (n *Node) pushDown() {
	mid := (n.start + n.end) / 2
	if n.left == nil {
		n.left = &Node{start: n.start, end: mid}
	}
	if n.right == nil {
		n.right = &Node{start: mid + 1, end: n.end}
	}

	if n.lazy == 0 {
		return
	}

	leftCount := mid - n.start + 1
	rightCount := n.end - mid

	n.left.lazy += n.lazy
	n.right.lazy += n.lazy

	n.left.val += n.lazy * leftCount
	n.right.val += n.lazy * rightCount

	n.lazy = 0
}

func (n *Node) Query(left, right int) int {
	if left <= n.start && n.end <= right {
		if n.val != 0 {
			fmt.Println("return:", left, right, " node:", n.start, n.end, " val:", n.val)
		}
		return n.val
	}

	n.pushDown()

	res := 0
	mid := (n.end + n.start) / 2
	if left <= mid {
		res += n.left.Query(left, right)
	}
	if right > mid {
		res += n.right.Query(left, right)
	}
	if res != 0 {
		fmt.Println("query:", left, right, " node:", n.start, n.end, " res:", res)
	}
	return res
}

func TestRecentCounter(t *testing.T) {
	rc := Constructor()
	fmt.Println(rc.Ping(2196))
	fmt.Println(rc.Ping(3938))
	fmt.Println(rc.Ping(4723))
	fmt.Println(rc.Ping(4775))
	fmt.Println(rc.Ping(5952))

	printTree(rc.root, 0)
}

func printTree(n *Node, level int) {
	if n == nil {
		return
	}
	for range level {
		fmt.Print(" ")
	}
	fmt.Printf("[%d,%d] val=%d lazy=%d\n", n.start, n.end, n.val, n.lazy)
	printTree(n.left, level+1)
	printTree(n.right, level+1)
}
