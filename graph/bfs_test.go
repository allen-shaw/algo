package graph

import (
	"fmt"
	"testing"
)

func TestFindMinHeightTrees(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {0, 2}}
	// edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
	// edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}

	ans := findMinHeightTrees(n, edges)
	fmt.Println(ans)
}

func TestIsBipartite(t *testing.T) {
	// graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	ans := isBipartite(graph)
	fmt.Println(ans)
}
