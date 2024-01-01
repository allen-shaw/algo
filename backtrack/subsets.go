package backtrack

import "sort"

var (
	ans   [][]int
	track []int
)

func subsets(nums []int) [][]int {
	ssbacktrack(nums, 0)
	return ans
}

func ssbacktrack(nums []int, cur int) {
	t := make([]int, len(track))
	copy(t, track)
	ans = append(ans, t)

	for i := cur; i < len(nums); i++ {
		track = append(track, nums[i])
		ssbacktrack(nums, i+1)
		track = track[:len(track)-1]
	}
}

var (
	ssdans   [][]int
	ssdtrack []int
)

func subsetsWithDup(nums []int) [][]int {
	ssdans = make([][]int, 0)
	ssdtrack = make([]int, 0, len(nums))
	sort.Ints(nums)

	ssdtrackback(nums, 0)
	return ssdans
}

func ssdtrackback(nums []int, cur int) {
	ssdans = append(ssdans, copyInts(ssdtrack))

	for i := cur; i < len(nums); i++ {
		if i > cur && nums[i] == nums[i-1] {
			continue
		}

		ssdtrack = append(ssdtrack, nums[i])
		ssdtrackback(nums, i+1)
		ssdtrack = ssdtrack[:len(ssdtrack)-1]
	}
}

func copyInts(src []int) []int {
	traget := make([]int, len(src))
	copy(traget, src)
	return traget
}
