package leetcode

// s1 的排列之一是 s2 的 子串 。
// s2 > s1
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	need, window := make(map[byte]int), make(map[byte]int)
	for i := range s1 {
		c := s1[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s2) {
		// 窗口扩张
		c := s2[right]
		right++

		// 处理
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 收缩窗口
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[left]
			left++

			if window[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}
