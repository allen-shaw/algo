package graph

var (
	// visited []bool
	onPath []int
	ans    [][]int
)

func allPathsSourceTarget(graph [][]int) [][]int {
	// visited = make([]bool, len(graph))
	traverse(graph, 0, len(graph)-1)
	return ans
}

func traverse(g [][]int, start, end int) {
	onPath = append(onPath, start)

	if start == end {
		temp := make([]int, len(onPath))
		copy(temp, onPath)
		ans = append(ans, temp)
		onPath = onPath[:len(onPath)-1]
		return
	}

	neighborhoods := g[start]
	for _, n := range neighborhoods {
		traverse(g, n, end)
	}

	onPath = onPath[:len(onPath)-1]
}

func findCelebrity(n int, knows func(i, j int) bool) int {
	cand := 0
	for other := 1; other < n; other++ {
		if knows(cand, other) || !knows(other, cand) {
			cand = other
		}
	}

	for i := 0; i < n; i++ {
		if i == cand {
			continue
		}
		if !knows(i, cand) || knows(cand, i) {
			return -1
		}
	}

	return cand
}
