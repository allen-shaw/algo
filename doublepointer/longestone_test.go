package doublepointer

import (
	"fmt"
	"testing"
)

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	count := 0
	ans := 0
	flag := false

	for right < len(nums) {
		n := nums[right]

		if n == 1 {
			count++
			right++
		} else { // n == 0
			if k > 0 {
				k--
				count++
				right++
			} else if k == 0 {
				flag = true
			}
		}
		if count > ans {
			fmt.Println("left:", left, " right:", right, " count:", count)
			ans = count
		}
		// ans = max(ans, count)

		for flag && left <= right {
			d := nums[left]
			if d == 0 {
				k++
				flag = false
			}
			count--
			left++
		}
	}

	return ans
}

func TestLongestOnes(t *testing.T) {
	var testsets = []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3},
		// 	   0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1
		//     0  1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16 17 18
	}

	for _, tt := range testsets {
		t.Run("longestOnes", func(t *testing.T) {
			ans := longestOnes(tt.nums, tt.k)
			fmt.Println(ans)
		})
	}
}
