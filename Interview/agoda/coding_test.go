package agoda

import (
	"fmt"
	"testing"
)

// 给定一个整数数组 numbers 和一个正整数 k，要求计算出数组 numbers 中，包含至少 k 对重复元素的所有连续子数组的个数。
func kSubArray(nums []int, k int) int {
	n := len(nums)
	ans := 0
	window := make(map[int]int)
	count := 0

	left, right := 0, 0
	for right < n {
		num := nums[right]
		right++

		window[num]++
		if window[num] == 2 {
			count++
		}
		if count == k {
			ans++
		}
		for left <= right && count >= k {
			d := nums[left]
			left++
			window[d]--
			if window[d] < 2 {
				count--
			}
		}
	}

	return ans
}

func Test_kSubArray(t *testing.T) {
	numbers := []int{0, 1, 0, 1, 0}
	k := 2
	ans := kSubArray(numbers, k)
	fmt.Println(ans)
}
