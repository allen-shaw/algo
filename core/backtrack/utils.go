package backtrack

func clone(src []int) []int {
	target := make([]int, len(src))
	copy(target, src)
	return target
}

// 0,1,2,3,...,9
func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
