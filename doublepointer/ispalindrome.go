package doublepointer

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isDigestAlpha(s[l]) {
			l++
		}

		for l < r && !isDigestAlpha(s[r]) {
			r--
		}

		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}

	return true
}

func isDigestAlpha(c byte) bool {
	return c >= '0' && c <= '9' ||
		c >= 'a' && c <= 'z'
}
