package graph

import "testing"

func TestCanFinish(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{{1, 1}}
	ans := canFinish(numCourses, prerequisites)
	t.Log(ans)
}

func TestBuildGraph(t *testing.T) {
	edges := [][]int{{1, 1}, {0, 2}}
	ans := buildGraph(3, edges)
	t.Log(ans)
}

func TestFindOrder(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}

	// numCourses := 4
	// prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	ans := findOrder(numCourses, prerequisites)
	t.Log(ans)

}
