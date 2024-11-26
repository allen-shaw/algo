package backtrack

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)

	var bt func(path []int, index int)
	bt = func(path []int, index int) {
		ans = append(ans, clone(path))
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			bt(path, i+1)
			path = path[:len(path)-1]
		}
	}

	bt(make([]int, 0), 0)
	return ans
}

func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}
	out := subsets(nums)
	fmt.Println(out)
}

// 1, 2, 3
//   		 	[]
// 	[1] 		[2] 		[3]
// [1,2][1,3]	[2,3]
// [1,2,3]

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)

	var bt func(path []int, index int)
	bt = func(path []int, index int) {
		if len(path) == k {
			if sum(path) == n {
				ans = append(ans, clone(path))
			}
			return
		}

		for i := index; i < 10; i++ {
			path = append(path, i)
			bt(path, i+1)
			path = path[:len(path)-1]
		}
	}

	bt(make([]int, 0), 1)

	return ans
}

func Test_combinationSum3(t *testing.T) {
	cases := []struct {
		name   string
		k      int
		n      int
		expect [][]int
	}{
		{"case1", 3, 7, [][]int{{1, 2, 4}}},
		{"case2", 3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := combinationSum3(c.k, c.n)
			assert.ElementsMatch(t, c.expect, out)
		})
	}

}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	var bt func(path []int, index int)
	bt = func(path []int, index int) {
		s := sum(path)
		if s == target {
			ans = append(ans, clone(path))
			return
		}
		if s > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			bt(path, i)
			path = path[:len(path)-1]
		}
	}

	bt(make([]int, 0), 0)

	return ans
}

func Test_combinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7

	out := combinationSum(candidates, target)
	fmt.Println(out)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	used := make([]bool, len(candidates))

	var bt func(path []int, index int)
	bt = func(path []int, index int) {
		if s := sum(path); s >= target {
			if s == target {
				ans = append(ans, clone(path))
			}
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, candidates[i])
			bt(path, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	bt(make([]int, 0), 0)

	return ans
}

func Test_combinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	out := combinationSum2(candidates, target)
	fmt.Println(out)
}

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	used := make([]bool, len(nums))

	var bt func(path []int)
	bt = func(path []int) {
		if len(path) == len(nums) {
			ans = append(ans, clone(path))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}

			path = append(path, nums[i])
			used[i] = true
			bt(path)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	bt(make([]int, 0))
	return ans
}

func Test_permuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	out := permuteUnique(nums)
	fmt.Println(out)
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)
	used := make([]bool, len(nums))

	var bt func(path []int, index int)
	bt = func(path []int, index int) {
		fmt.Println("index: ", index, " path: ", path)
		ans = append(ans, clone(path))

		for i := index; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			bt(path, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	bt(make([]int, 0), 0)

	return ans
}

func Test_subsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	out := subsetsWithDup(nums)
	fmt.Println(out)
}

//  []
// 1, 		2
// 1,2		2,2
// 1,2,2
