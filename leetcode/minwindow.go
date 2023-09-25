package leetcode

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range t {
		c := t[i]
		need[c]++
	}

	start, length := 0, -1
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++

		// 扩张窗口
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			// 更新长度
			if right-left < length || length < 0 {
				start = left
				length = right - left
			}

			// 收缩窗口
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

	if length == -1 {
		return ""
	} else {
		return s[start : start+length]
	}
}
