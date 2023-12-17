package backtrack

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
