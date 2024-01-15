package backtrack

import (
	"fmt"
	"testing"
)

var memo [][]int

func numTrees(n int) int {
	memo = make([][]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, n+1)
	}
	return ntdfs(1, n)
}

func ntdfs(left, right int) int {
	if left > right {
		return 1
	}

	if memo[left][right] != 0 {
		return memo[left][right]
	}

	ans := 0
	for i := left; i <= right; i++ {
		ans += ntdfs(left, i-1) * ntdfs(i+1, right)
	}
	memo[left][right] = ans
	return ans
}

func TestNumsTree(t *testing.T) {
	var testsets = []struct {
		n   int
		out int
	}{
		{3, 5},
	}

	for _, tt := range testsets {
		t.Run("numsTree", func(t *testing.T) {
			ans := numTrees(tt.n)
			fmt.Println(ans)
		})
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var gtans []*TreeNode

func generateTrees(n int) []*TreeNode {
	gtans = make([]*TreeNode, 0)
	gtdfs(1, n)
	return gtans
}

func gtdfs(left, right int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if left > right {
		res = append(res, nil)
		return res
	}

	for i := left; i <= right; i++ {
		leftTrees := gtdfs(left, i-1)
		rightTrees := gtdfs(i+1, right)
		for _, lt := range leftTrees {
			for _, rt := range rightTrees {
				res = append(res, &TreeNode{Val: i, Left: lt, Right: rt})
			}
		}
	}

	return res
}
