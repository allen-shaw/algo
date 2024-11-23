package backtrack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	used := make([]bool, len(nums))

	var bt func(path []int, used []bool)
	bt = func(path []int, used []bool) {
		if len(path) == len(nums) {
			ans = append(ans, clone(path))
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true
			bt(path, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	bt(make([]int, 0), used)
	return ans
}

func clone(src []int) []int {
	target := make([]int, len(src))
	copy(target, src)
	return target
}

func Test_permute(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{"case1", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := permute(c.input)
			assert.ElementsMatch(t, c.expect, out)
		})
	}
}
