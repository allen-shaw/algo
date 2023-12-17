package doublepointer

func longestPalindrome(s string) string {
	return longestPalindromeDoublePointer(s)
}

func longestPalindromeDoublePointer(s string) string {
	var ans string

	for i := 0; i < len(s); i++ {
		palindrome1 := tryPalindrome(s, i, i)
		palindrome2 := tryPalindrome(s, i, i+1)

		if len(palindrome1) > len(ans) {
			ans = palindrome1
		}
		if len(palindrome2) > len(ans) {
			ans = palindrome2
		}
	}
	return ans
}

func tryPalindrome(s string, left, right int) string {
	// log.Print(s, left, right)

	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return s[left+1 : right]
}
