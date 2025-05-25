package errorset

import (
	"fmt"
	"math"
	"testing"
)

func minWindow(s string, t string) string {
	need := make(map[byte]int, 0)
	for i := range t {
		need[s[i]]++
	}

	window := make(map[byte]int)
	left, right := 0, 0

	valid := 0
	start, length := 0, math.MaxInt

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
			if right-left < length {
				length = right - left
				start = left
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

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}

func Test_subString(t *testing.T) {
	str := "0123456789"
	fmt.Println(str[2:3])
}
