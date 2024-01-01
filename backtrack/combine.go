package backtrack

import "sort"

var (
	cans   [][]int
	ctrack []int
)

func combine(n int, k int) [][]int {
	cans = make([][]int, 0)
	ctrack = make([]int, 0, k)

	cbacktrack(n, 1, k)
	return cans
}

func cbacktrack(n, cur, k int) {
	if len(ctrack) == k {
		t := make([]int, k)
		copy(t, ctrack)
		cans = append(cans, t)
		return
	}

	for i := cur; i <= n; i++ {
		ctrack = append(ctrack, i)
		cbacktrack(n, i+1, k)
		ctrack = ctrack[:len(ctrack)-1]
	}
}

var (
	csans   [][]int
	cstrack []int
	cssum   int
)

func combinationSum2(candidates []int, target int) [][]int {
	csans = make([][]int, 0)
	cstrack = make([]int, 0, len(candidates))
	cssum = 0

	sort.Ints(candidates)
	cstrackback(candidates, 0, target)
	return csans
}

func cstrackback(nums []int, cur int, target int) {
	if cssum == target {
		csans = append(csans, copyInts(cstrack))
		return
	}
	if cssum > target {
		return
	}

	for i := cur; i < len(nums); i++ {
		if i > cur && nums[i] == nums[i-1] {
			continue
		}

		cstrack = append(cstrack, nums[i])
		cssum += nums[i]

		cstrackback(nums, i+1, target)

		cstrack = cstrack[:len(cstrack)-1]
		cssum -= nums[i]
	}
}
