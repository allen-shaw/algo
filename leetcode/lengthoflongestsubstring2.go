package leetcode

func lengthOfLongestSubstring2(s string) int {
	window := make(map[byte]int)

	maxLength := 0
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			if window[d] > 0 {
				window[d]--
			}
		}
		maxLength = max(right-left, maxLength)
	}

	return maxLength
}
