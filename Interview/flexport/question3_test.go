package flexport

import (
	"fmt"
	"testing"
)

// 1.3 Container Loading Problem
// Given a container of size 1000 and a set of items with specific sizes, such as 50, 80, 100,
// find any combination of items that exactly fills the container.
// The solution should add up to the size of the container.
// Describe how you would approach this problem, and if you used depth-first search (DFS),
// explain why and how you implemented it.

func containerLoading(size int, items []int) [][]int {
	result := make([][]int, 0)

	var dfs func(list []int, sum int)
	dfs = func(list []int, sum int) {
		fmt.Println("list: ", list, "sum: ", sum)
		if sum == size {
			result = append(result, clone(list))
			return
		}
		if sum > size {
			return
		}

		for _, item := range items {
			list = append(list, item)
			dfs(list, sum+item)
			list = list[:len(list)-1]
		}
	}

	dfs(make([]int, 0), 0)
	return result
}

func clone(src []int) []int {
	target := make([]int, len(src))
	copy(target, src)
	return target
}

func Test_containerLoading(t *testing.T) {
	result := containerLoading(100, []int{50, 80, 100})
	fmt.Println(result)
}
