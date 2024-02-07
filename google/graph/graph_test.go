package graph

import (
	"fmt"
	"testing"
)

func maximumDetonation(bombs [][]int) int {

	g := make([][]int, len(bombs))
	// 构建有向图 a -> b : a能引爆b
	for a, bomba := range bombs {
		for b, bombb := range bombs {
			if isConnect(bomba, bombb) {
				g[a] = append(g[a], b)
			}
		}
	}

	ans := 0
	for n := range g {
		visited := make([]bool, len(g))
		res := 0
		dfs(g, n, visited, &res)
		ans = max(ans, res)
	}

	return ans
}

func isConnect(a, b []int) bool {
	xa, ya := uint64(a[0]), uint64(a[1])
	r := uint64(a[2])
	xb, yb := uint64(b[0]), uint64(b[1])

	return (xa-xb)*(xa-xb)+(ya-yb)*(ya-yb) <= r*r
}

func dfs(g [][]int, v int, visited []bool, res *int) {
	if visited[v] {
		return
	}

	(*res)++
	visited[v] = true
	for _, n := range g[v] {
		dfs(g, n, visited, res)
	}
}

func TestMaximumDetonation(t *testing.T) {
	bombs := [][]int{{2, 1, 3}, {6, 1, 4}}
	ans := maximumDetonation(bombs)
	fmt.Println(ans)
}
