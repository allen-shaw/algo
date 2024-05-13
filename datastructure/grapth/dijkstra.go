package grapth

import "math"

type State struct {
	id   int
	dist int
}

func dijkstra(graph [][]int, start int) []int {
	v := len(graph)
	dists := make([]int, v)
	for i := 0; i < len(dists); i++ {
		dists[i] = math.MaxInt
	}

	dists[start] = 0
	q := []State{{start, 0}}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur.dist > dists[cur.id] {
			continue
		}

		for _, n := range graph[cur.id] {
			distToNext := dists[cur.id] + graph[cur.id][n]
			if dists[n] > distToNext {
				dists[n] = distToNext
				q = append(q, State{n, distToNext})
			}
		}
	}

	return dists
}
