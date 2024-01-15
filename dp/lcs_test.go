package dp

import (
	"fmt"
	"testing"
)

func longestCommonSubsequence(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

func TestLCS(t *testing.T) {
	var testsets = []struct {
		s1, s2 string
		out    int
	}{
		{"abcde", "ace", 3},
	}

	for _, tt := range testsets {
		t.Run("lcs", func(t *testing.T) {
			ans := longestCommonSubsequence(tt.s1, tt.s2)
			fmt.Println(ans)
		})
	}
}
