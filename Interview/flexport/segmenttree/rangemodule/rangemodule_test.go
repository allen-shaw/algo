package rangemodule

import (
	"fmt"
	"testing"
)

type RangeModule struct {
	root *Node
}

func Constructor() RangeModule {
	start, end := 0, 20
	return RangeModule{
		root: &Node{start: start, end: end, mid: (start + end) / 2},
	}
}

func (m *RangeModule) AddRange(left int, right int) {
	m.root.Update(left, right-1, 1)
}

func (m *RangeModule) QueryRange(left int, right int) bool {
	return m.root.Query(left, right-1) == (right - left)
}

func (m *RangeModule) RemoveRange(left int, right int) {
	m.root.Update(left, right-1, -1)
}

type Node struct {
	val             int
	start, end, mid int
	lazy            int
	left, right     *Node
}

func (n *Node) Update(left, right, val int) {
	if left <= n.start && n.end <= right {
		n.val = 0
		if val == 1 {
			n.val = (n.end - n.start + 1)
		}
		n.lazy = val
		return
	}

	n.pushDown()
	if left <= n.mid {
		n.left.Update(left, right, val)
	}
	if right > n.mid {
		n.right.Update(left, right, val)
	}

	n.val = n.left.val + n.right.val
}

func (n *Node) Query(left, right int) int {
	if left <= n.start && n.end <= right {
		return n.val
	}

	n.pushDown()
	res := 0
	if left <= n.mid {
		res += n.left.Query(left, right)
	}
	if right > n.mid {
		res += n.right.Query(left, right)
	}
	return res
}

func (n *Node) pushDown() {
	if n.left == nil {
		n.left = &Node{start: n.start, end: n.mid, mid: (n.start + n.mid) / 2}
	}
	if n.right == nil {
		n.right = &Node{start: n.mid + 1, end: n.end, mid: (n.mid + 1 + n.end) / 2}
	}

	if n.lazy == 0 {
		return
	}

	n.left.val = 0
	n.right.val = 0
	if n.lazy == 1 {
		n.left.val = n.left.end - n.left.start + 1
		n.right.val = n.right.end - n.right.start + 1
	}

	n.left.lazy = n.lazy
	n.right.lazy = n.lazy

	n.lazy = 0
}

func TestRangeModule(t *testing.T) {
	rangeModule := Constructor()
	rangeModule.AddRange(10, 20)
	rangeModule.RemoveRange(14, 16)
	fmt.Println(rangeModule.QueryRange(10, 14)) //  true
	fmt.Println(rangeModule.QueryRange(13, 15)) // 返回 false（未跟踪区间 [13, 15) 中像 14, 14.03, 14.17 这样的数字）
	fmt.Println(rangeModule.QueryRange(16, 17)) //  返回 true （尽管执行了删除操作，区间 [16, 17) 中的数字 16 仍然会被跟踪）
}
