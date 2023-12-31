package graph

import (
	"fmt"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	g := buildUndirectedGraph(n, edges)
	fmt.Println(g)

	degrees := make([]int, n)
	for _, edge := range edges {
		degrees[edge[0]]++
		degrees[edge[1]]++
	}

	// find all leaves
	leaves := make([]int, 0)
	for i, d := range degrees {
		if d == 1 {
			leaves = append(leaves, i)
		}
	}

	fmt.Println("leaves:", leaves)

	pre := make([]int, 2)
	for len(leaves) > 1 {
		size := len(leaves)
		if size == 2 {
			copy(pre, leaves)
		}

		for i := 0; i < size; i++ {
			v := leaves[0]
			leaves = leaves[1:] // pop

			fmt.Println("v:", v, g[v])
			for _, n := range g[v] {
				degrees[n]--
				if degrees[n] == 1 {
					leaves = append(leaves, n)
				}
			}
		}
	}

	if len(leaves) == 1 {
		return leaves
	}

	return pre
}

func buildUndirectedGraph(n int, edges [][]int) [][]int {
	g := make([][]int, n)

	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}

	return g
}

func isBipartite(graph [][]int) bool {
	visited := make([]bool, len(graph))
	color := make([]bool, len(graph))

	q := make([]int, 0)
	q = append(q, 0)
	visited[0] = true
	color[0] = true

	for len(q) > 0 {
		v := q[0]
		q = q[1:] // pop

		edges := graph[v]
		for _, n := range edges {
			if visited[n] {
				if color[n] == color[v] {
					return false
				}
			} else {
				q = append(q, n)
				visited[n] = true
				color[n] = !color[v]
			}
		}
	}

	return true
}

var (
	visited []bool
	color   []bool
)

func possibleBipartition(n int, dislikes [][]int) bool {
	g := buildUndirectedGraph(n, dislikes)
	visited = make([]bool, len(g))
	color = make([]bool, len(g))

	for v := range g {
		if !visited[v] {
			ok := doPossibleBipartition(g, v)
			if !ok {
				return false
			}
		}
	}

	return true
}

func doPossibleBipartition(g [][]int, v int) bool {
	q := make([]int, 0, len(g))

	q = append(q, v)
	visited[v] = true
	color[v] = true

	for len(q) > 0 {
		n := q[0]
		q = q[1:] // pop

		neighberhood := g[n]
		for _, ne := range neighberhood {
			if visited[ne] {
				if color[ne] == color[n] {
					return false
				}
			} else {
				q = append(q, ne)
				visited[ne] = true
				color[ne] = !color[n]
			}
		}
	}

	return true
}
