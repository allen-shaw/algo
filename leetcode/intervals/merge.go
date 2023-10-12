package intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		first, second := intervals[i], intervals[j]
		if first[0] != second[0] {
			return first[0] < second[0]
		}
		return first[1] > second[1]
	})

	ans := make([][]int, 0)
	ans = append(ans, intervals[0])

	i, j := 0, 1
	for i < len(ans) && j < len(intervals) {
		first, second := ans[i], intervals[j]
		if first[1] >= second[1] {
			j++
		} else if first[1] < second[0] {
			ans = append(ans, second)
			i++
		} else if first[1] >= second[0] && first[1] < second[1] {
			ans = ans[:len(ans)-1]
			ans = append(ans, []int{first[0], second[1]})
			j++
		}
	}

	return ans
}
