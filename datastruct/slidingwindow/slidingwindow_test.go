package slidingwindow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		c := s1[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s2) {
		c := s2[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}

func Test_checkInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidboaoo"

	ok := checkInclusion(s1, s2)
	fmt.Println(ok)
}

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		c := t[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0

	ans := ""

	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		if need[c] == window[c] {
			valid++
		}

		for valid == len(need) {
			if ans == "" || right-left < len(ans) {
				ans = s[left:right]
				fmt.Println("update:", ans)
			}

			d := s[left]
			left++

			if window[d] == need[d] {
				valid--
			}
			window[d]--
		}
	}

	return ans
}

func Test_minWindow(t *testing.T) {
	ss := "ADOBECODEBANC"
	st := "ABC"
	ans := minWindow(ss, st)
	fmt.Println(ans)
}

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)

	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		c := p[i]
		need[c]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s) {
		c := s[right]
		window[c]++
		right++

		if window[c] == need[c] {
			valid++
		}

		if right-left >= len(p) {
			if valid == len(need) {
				ans = append(ans, left)
			}

			d := s[left]
			left++
			if window[d] == need[d] {
				valid--
			}

			window[d]--
		}
	}

	return ans
}

func Test_findAnagrams(t *testing.T) {
	cases := []struct {
		name  string
		input struct {
			s string
			p string
		}
		expect []int
	}{
		{"case1", struct {
			s string
			p string
		}{s: "abaacbabc", p: "abc"}, []int{3, 4, 6}},
		{"case2", struct {
			s string
			p string
		}{s: "cccccccbbbbbbbbbaaaaa", p: "abc"}, []int{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := findAnagrams(c.input.s, c.input.p)
			assert.Equal(t, c.expect, out)
		})
	}
}
