package backtrack

import (
	"fmt"
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

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	var bt func(path []int, index int)
	bt = func(path []int, index int) {
		if len(path) == k {
			ans = append(ans, clone(path))
			return
		}

		for i := index; i <= n; i++ {
			path = append(path, i)
			bt(path, i+1)
			path = path[:len(path)-1]
		}
	}
	bt(make([]int, 0), 1)
	return ans
}

func Test_combine(t *testing.T) {
	out := combine(4, 2)
	fmt.Println(out)
}
