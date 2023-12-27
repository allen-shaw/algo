package graph

import (
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	// graph := [][]int{{1, 2}, {3}, {3}, {}}
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	ans := allPathsSourceTarget(graph)
	t.Log(ans)
}

func TestFindCelebrity(t *testing.T) {
	graph := [][]int{
		{1, 1, 1, 0},
		{1, 1, 1, 1},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
	}

	knows := buildKnown(graph)
	ans := findCelebrity(4, knows)
	t.Log(ans)
}

func buildKnown(g [][]int) func(i, j int) bool {
	return func(i, j int) bool {
		return g[i][j] == 1
	}
}
