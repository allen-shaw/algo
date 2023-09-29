package array

func minWindow(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range t {
		c := t[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	minlen := -1
	minIndex := 0

	for right < len(s) {
		c := s[right]
		right++

		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			// 判断是否最小
			if right-left < minlen || minlen == -1 {
				minlen = right - left
				minIndex = left
			}

			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minlen == -1 {
		return ""
	} else {
		return s[minIndex : minIndex+minlen]
	}
}
