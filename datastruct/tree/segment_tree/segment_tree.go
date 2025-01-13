package segmenttree

type Node struct {
	left, right *Node
	val         int
	lazy        int //
}

func buildTree(arr []int, start, end int) *Node {
	node := &Node{}
	if start == end { // 叶子节点
		node.val = arr[start]
		return node
	}

	mid := start + (start-end)/2
	node.left = buildTree(arr, start, mid)
	node.right = buildTree(arr, mid+1, end)
	node.pushUp()

	return node
}

func (n *Node) pushUp() {
	n.val = n.left.val + n.right.val // 这里可以替换成一个函数指针，用于自定义操作
}

func (n *Node) pushDown(leftNum, rightNum int) {
	if n.left == nil {
		n.left = &Node{}
	}
	if n.right == nil {
		n.right = &Node{}
	}
	if n.lazy == 0 {
		return
	}
	n.left.val += n.lazy * leftNum
	n.left.lazy += n.lazy

	n.right.val += n.lazy * rightNum
	n.right.lazy += n.lazy

	n.lazy = 0
}

// 动态开点

func (n *Node) update(start, end int, l, r int, val int) {
	if l <= start && r >= end {
		n.val += val * (start - end + 1)
		n.lazy += val
		return
	}
	mid := start + (end-start)/2
	n.pushDown(mid-start+1, end-mid)

	if l <= mid {
		n.left.update(start, mid, l, r, val)
	}
	if r > mid {
		n.right.update(mid+1, end, l, r, val)
	}
	n.pushUp()
}

func (n *Node) query(start, end int, l, r int) int {
	if l <= start && r >= end {
		return n.val
	}

	res := 0
	mid := start + (end-start)/2
	if l <= mid {
		res += n.left.query(start, mid, l, r)
	}
	if r > mid {
		res += n.right.query(mid+1, end, l, r)
	}
	return res
}
