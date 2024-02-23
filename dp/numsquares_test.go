package dp

import (
	"fmt"
	"math"
	"testing"
)

func numSquares(num int) int {
	memo := make(map[int]int)

	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}

		if c, ok := memo[n]; ok {
			return c
		}

		ans := math.MaxInt
		for i := 1; i*i <= n; i++ {
			num := dp(n - i*i)
			if num == -1 {
				continue
			}
			ans = min(ans, num+1)
		}
		if ans == math.MaxInt {
			ans = -1
		}

		memo[n] = ans
		return ans
	}

	return dp(num)
}

func TestNumSquares(t *testing.T) {
	fmt.Println(numSquares(12))
}

func TestSqrt(t *testing.T) {
	n := 9
	fmt.Println(int(math.Sqrt(float64(n))))
}
