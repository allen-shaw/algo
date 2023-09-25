package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	l, r := 0, 1
	set := make(map[byte]int)
	maxLength := 0
	set[s[l]] = l

	for r < len(s) {
		v, ok := set[s[r]]
		if ok && v >= l {
			maxLength = max(maxLength, r-l)
			delete(set, s[l])
			l = v + 1
		}
		set[s[r]] = r
		r++
	}

	maxLength = max(maxLength, r-l)
	return maxLength
}
