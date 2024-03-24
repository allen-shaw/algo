package tree

import (
	"fmt"
	"testing"
)

type MAryToBinaryTree struct {
}

type NTreeNode struct {
	val     int
	chidren []*NTreeNode
}

func printNTree(root *NTreeNode) {
	q := make([]*NTreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		val := make([]int, 0)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			val = append(val, n.val)
			for _, c := range n.chidren {
				q = append(q, c)
			}
		}
		fmt.Println(val)
	}
}

func printBTree(root *TreeNode) {
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)
		val := make([]int, 0)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			val = append(val, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		fmt.Println(val)
	}
}

func (*MAryToBinaryTree) Encode(root *NTreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tdummy := &TreeNode{}
	tp := tdummy

	q := make([]*NTreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		size := len(q)

		dummy := &TreeNode{}
		p := dummy

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			p.Right = &TreeNode{Val: n.val}
			p = p.Right

			q = append(q, n.chidren...)
		}

		tp.Left = dummy.Right
		tp = tp.Left
	}

	return tdummy.Left
}

type nodePair struct {
	n *NTreeNode
	b *TreeNode
}

func (*MAryToBinaryTree) Decode(root *TreeNode) *NTreeNode {
	if root == nil {
		return nil
	}

	ntRoot := &NTreeNode{val: root.Val}

	q := make([]*nodePair, 0)
	q = append(q, &nodePair{ntRoot, root})

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			pair := q[0]
			n := pair.b
			q = q[1:]

			children := make([]*NTreeNode, 0)

			for p := n.Left; p != nil; p = p.Right {
				ntn := &NTreeNode{val: p.Val}
				q = append(q, &nodePair{ntn, p})
				children = append(children, ntn)
			}
			pair.n.chidren = children
		}
	}

	return ntRoot
}

func Test_MAryToBinaryTree(t *testing.T) {
	root := &NTreeNode{val: 1}
	root.chidren = []*NTreeNode{
		{val: 2, chidren: []*NTreeNode{{val: 5}, {val: 6}}},
		{val: 3},
		{val: 4, chidren: []*NTreeNode{{val: 7}, {val: 8}}},
	}
	fmt.Println("source:")
	printNTree(root)

	// encode
	obj := &MAryToBinaryTree{}
	broot := obj.Encode(root)

	fmt.Println("encode: binaryTree:")
	printBTree(broot)

	// decode
	root = obj.Decode(broot)

	fmt.Println("decode: NTree:")
	printNTree(root)

}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	m := make(map[*TreeNode]*TreeNode) // child -> parent

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			dfs(root.Left)
			m[root.Left] = root
		}
		if root.Right != nil {
			dfs(root.Right)
			m[root.Right] = root
		}
	}

	dfs(root)

	q := []*TreeNode{target}
	lvl := 0
	ans := make([]int, 0)
	visited := make(map[*TreeNode]bool)

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if visited[n] {
				continue
			}
			visited[n] = true

			if lvl == k {
				ans = append(ans, n.Val)
			}
			if lvl > k {
				return ans
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			if p, ok := m[n]; ok {
				q = append(q, p)
			}
		}

		lvl++
	}
	return ans
}

func findClosestLeaf(root *TreeNode, k int) int {
	parents := make(map[*TreeNode]*TreeNode)
	var kNode *TreeNode

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Val == k {
			kNode = root
		}

		if root.Left != nil {
			dfs(root.Left)
			parents[root.Left] = root
		}
		if root.Right != nil {
			dfs(root.Right)
			parents[root.Right] = root
		}
	}
	visited := make(map[*TreeNode]bool)
	q := []*TreeNode{kNode}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if n.Left == nil && n.Right == nil {
				return n.Val
			}

			if visited[n] {
				continue
			}
			visited[n] = true

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			if p, ok := parents[n]; ok {
				q = append(q, p)
			}
		}
	}

	return 0
}

func killProcess(pid, ppid []int, kill int) []int {
	pids := make(map[int][]int, 0)

	for i, parentid := range ppid {
		processid := pid[i]
		pids[parentid] = append(pids[parentid], processid)
	}

	ans := make([]int, 0)
	q := []int{kill}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		ans = append(ans, p)
		if len(pids[p]) > 0 {
			q = append(q, pids[p]...)
		}
	}

	return ans
}

func Test_killProcess(t *testing.T) {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	kill := 5

	ans := killProcess(pid, ppid, kill)
	fmt.Println(ans)
}
