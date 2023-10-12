package array

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		a := intervals[i]
		b := intervals[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1] > b[1]
	})

	fmt.Println("after sorted:", intervals)

	del := 0
	i, j := 0, 1
	for j < len(intervals) {
		first := intervals[i]
		next := intervals[j]

		if first[0] <= next[0] && first[1] >= next[1] {
			del++
			j++
		} else {
			i = j
			j++
		}
	}

	return len(intervals) - del
}
