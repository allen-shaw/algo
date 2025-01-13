package flexport

import (
	"fmt"
	"sort"
	"testing"
)

// 1.6 Schedule Ship Problem
// The interview question was similar to the 'merge interval' problem but adapted to scheduling ships.
// The task involved handling input/output on your own, and coding was done on a HackerRank whiteboard.
// A follow-up question was asked about how to arrange the schedule if multiple ships were involved.
func mergeInterval(intervals [][]int) [][]int {
	ans := make([][]int, 0, len(intervals))

	// sort by begin time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println("intervals:", intervals)

	for _, interval := range intervals {
		if len(ans) == 0 {
			ans = append(ans, interval)
		} else {
			lastIdx := len(ans) - 1
			// overlap
			if ans[lastIdx][1] >= interval[0] {
				ans[lastIdx][1] = max(ans[lastIdx][1], interval[1])
			} else {
				ans = append(ans, interval)
			}
			fmt.Println("ans:", ans)
		}
	}

	return ans
}

func Test_mergeInterval(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ans := mergeInterval(intervals)
	fmt.Println(ans)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	lefts, rights := make([][]int, 0), make([][]int, 0)

	st, end := newInterval[0], newInterval[1]
	for _, i := range intervals {
		ist, iend := i[0], i[1]
		if iend < st {
			lefts = append(lefts, i)
		} else if ist > end {
			rights = append(rights, i)
		} else {
			st = min(st, ist)
			end = max(end, iend)
		}
	}

	lefts = append(lefts, []int{st, end})
	lefts = append(lefts, rights...)
	return lefts
}

func Test_insert(t *testing.T) {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	out := insert(intervals, newInterval)
	fmt.Println(out)
}
