package leetcode

func longestConsecutive(nums []int) int {
	set := listToSet(nums)
	maxLength := 0
	for num := range set {
		if set[num-1] {
			continue
		}
		curLength := 1
		for {
			num++
			if _, ok := set[num]; !ok {
				break
			} else {
				curLength++
			}
		}
		maxLength = max(maxLength, curLength)
	}

	return maxLength
}

func listToSet(arr []int) map[int]bool {
	s := make(map[int]bool)
	for _, n := range arr {
		s[n] = true
	}
	return s
}
