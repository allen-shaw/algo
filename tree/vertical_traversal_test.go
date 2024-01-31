package tree

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

var list [][]int

func verticalTraversal(root *TreeNode) [][]int {
	list = make([][]int, 0)
	vtdfs(root, 0, 0)
	sort.Slice(list, func(i, j int) bool {
		if list[i][0] == list[j][0] {
			return list[i][1] < list[j][1]
		}
		return list[i][0] < list[j][0]
	})

	ans := merge(list)
	return ans
}

func merge(list [][]int) [][]int {
	ans := make([][]int, 0)
	temp := make([]int, 0)
	lastx := math.MinInt
	for _, elem := range list {
		if lastx == elem[0] {
			temp = append(temp, elem[2])
		} else {
			if len(temp) != 0 {
				ans = append(ans, copylist(temp))
				temp = make([]int, 0)
			}
			temp = append(temp, elem[2])
			lastx = elem[0]
		}
	}
	ans = append(ans, temp)
	return ans
}

func vtdfs(root *TreeNode, x, y int) {
	if root == nil {
		return
	}

	list = append(list, []int{x, y, root.Val})
	vtdfs(root.Left, x-1, y+1)
	vtdfs(root.Right, x+1, y+1)
}

func copylist(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func TestMerge(t *testing.T) {
	// list := [][]int{{-1, 1, 9}, {0, 0, 3}, {0, 2, 15}, {1, 1, 20}, {2, 2, 7}}
	list := [][]int{{-2, 2, 4}, {-1, 1, 2}, {0, 0, 1}, {1, 1, 3}, {0, 2, 5}, {0, 2, 6}, {2, 2, 7}}
	out := merge(list)
	fmt.Println(out)
}
