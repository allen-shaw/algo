package leetcode

func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range p {
		c := p[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	ans := make([]int, 0)

	for right < len(s) {
		c := s[right]
		right++

		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				ans = append(ans, left)
			}

			d := s[left]
			left++

			if window[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return ans
}
