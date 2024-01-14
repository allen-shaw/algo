package backtrack

// https://blog.csdn.net/qq_36282995/article/details/126078742
import (
	"sort"
	"strconv"
)

func findMaxLessN(arr []int, n int) int {
	sort.Ints(arr) // 从小到大排序
	nums := parseN(n)

	return dfs(arr, nums, 0, false)
}

func parseN(n int) []int {
	s := strconv.Itoa(n)
	res := make([]int, len(s))
	for i, c := range s {
		res[i] = int(c - '0')
	}
	return res
}

func dfs() int {

}
