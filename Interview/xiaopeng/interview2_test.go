package xiaopeng

import (
	"fmt"
	"strconv"
	"testing"
)

// digits = ["1","3","5","7"], n = 100
// 1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.

func atMostNGivenDigitSet(digits []string, k int) int {
	n := len(digits)
	ans := make([]int, 0)
	var dfs func(num int)
	dfs = func(num int) {
		if num > k {
			return
		}
		if num != 0 {
			ans = append(ans, num)
		}
		for i := 0; i < n; i++ {
			val, _ := strconv.Atoi(digits[i])
			num = num*10 + val
			dfs(num)
			num = (num - val) / 10
		}
	}

	dfs(0)
	fmt.Println(ans, len(ans))
	return len(ans)
}

func Test_atMostNGivenDigitSet(t *testing.T) {
	atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100)
}
