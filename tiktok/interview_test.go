package tiktok

import (
	"fmt"
	"testing"
)

// 给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。

// 返回这 两个区间列表的交集 。

// 形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。

// 两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。

// 示例：
// 输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
// 输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

func TestInterview(t *testing.T) {
	a := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	b := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	ans := intersection(a, b)
	fmt.Println(ans)
}

func intersection(a, b [][]int) [][]int {
	ans := make([][]int, 0)

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		lo := max(a[i][0], b[j][0])
		hi := min(a[i][1], b[j][1])
		if lo <= hi {
			ans = append(ans, []int{lo, hi})
		}

		if a[i][1] < b[j][1] {
			i++
		} else {
			j++
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
