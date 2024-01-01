package backtrack

import (
	"sort"
)

func permute(nums []int) [][]int {
	s := newPermuteSolution()
	return s.Permute(nums)
}

type permuteSolution struct {
	res [][]int
}

func newPermuteSolution() *permuteSolution {
	s := &permuteSolution{}
	return s
}

func (p *permuteSolution) Permute(nums []int) [][]int {
	path := NewLinkedList[int]()
	used := make([]bool, len(nums))

	p.backtrack(nums, path, used)
	return p.res
}

func (p *permuteSolution) backtrack(nums []int, path *LinkedList[int], used []bool) {
	if path.Size() == len(nums) {
		p.res = append(p.res, path.ToList())
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		path.PushBack(nums[i])
		used[i] = true

		p.backtrack(nums, path, used)

		path.RemoveLast()
		used[i] = false
	}
}

var (
	puans   [][]int
	putrack []int
	puused  []bool
)

func permuteUnique(nums []int) [][]int {
	puans = make([][]int, 0)
	putrack = make([]int, 0, len(nums))
	puused = make([]bool, len(nums))

	sort.Ints(nums)

	putrackback(nums)
	return puans
}

func putrackback(nums []int) {
	if len(putrack) == len(nums) {
		puans = append(puans, copyInts(putrack))
		return
	}

	for i := 0; i < len(nums); i++ {
		if puused[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !puused[i-1] {
			continue
		}

		puused[i] = true
		putrack = append(putrack, nums[i])

		putrackback(nums)

		putrack = putrack[:len(putrack)-1]
		puused[i] = false
	}
}
