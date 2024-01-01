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

type ocell struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	ans := 0
	q := make([]ocell, 0, len(grid)*len(grid[0]))
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				q = append(q, ocell{i, j})
				visited[i][j] = true
			}
		}
	}

	for len(q) > 0 {
		size := len(q)
		for k := 0; k < size; k++ {
			c := q[0]
			q = q[1:] // pop

			for i := 0; i < 4; i++ {
				x, y := c.x+dir[i], c.y+dir[i+1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && !visited[x][y] && grid[x][y] == 1 {
					grid[x][y] = 2
					q = append(q, ocell{x, y})
					visited[x][y] = true
				}
			}
		}
		ans++
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return ans - 1
}
