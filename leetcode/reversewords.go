package leetcode

import (
	"strings"
)

func reverseWords(s string) string {
	rs := reverseWord([]byte(s))
	ss := strings.Split(rs, " ")

	ans := make([]string, 0)
	for _, substr := range ss {
		if strings.Trim(substr, " ") == "" {
			continue
		}
		ans = append(ans, reverseWord([]byte(substr)))
	}
	return strings.Join(ans, " ")
}

func reverseWord(s []byte) string {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return string(s)
}
