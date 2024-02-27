package dp

import (
	"fmt"
	"math"
	"testing"
)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == n-1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + matrix[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			}
		}
	}

	fmt.Println(dp)

	ans := math.MaxInt
	for j := 0; j < n; j++ {
		ans = min(ans, dp[n-1][j])
	}

	return ans
}

func TestMinFallingPathSum(t *testing.T) {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	ans := minFallingPathSum(matrix)
	fmt.Println(ans)
}
