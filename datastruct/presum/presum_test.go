package presum

import (
	"fmt"
	"testing"
)

func Test_presum(t *testing.T) {
	nums := []int{3, 5, 2, -2, 4, 1}

	sum1, sum2 := make([]int, len(nums)+2), make([]int, len(nums)+2)
	for i := 1; i <= len(nums); i++ {
		sum1[i] = sum1[i-1] + nums[i-1]
	}
	// for i := len(nums) - 1; i >= 0; i-- {
	// 	j := len(nums) - 1 - i
	// 	sum2[j+1] = sum2[j] + nums[i]
	// }
	for i := len(nums); i >= 1; i-- {
		sum2[i] = sum2[i+1] + nums[i-1]
	}
	// sum2[6] = sum2[7] + sum2[5]

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2: ", sum2)
}

func Test_copy(t *testing.T) {
	nums1 := make([]int, 1)
	nums2 := []int{1}

	copy(nums1[:0], nums2[:])
}

type NumArray struct {
	presum []int
}

func Constructor(nums []int) NumArray {
	presum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		presum[i+1] = presum[i] + nums[i]
	}
	return NumArray{
		presum: presum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.presum[right+1] - this.presum[left]
}

func Test_NumArray(t *testing.T) {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	// 0 -2 -2 1 -4 -2 -3
	fmt.Println(na.presum)
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
}

func longestWPI(hours []int) int {
	for i := range hours {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}

	ans := 0
	// 转换为求和大于0的最长子数组
	m := make(map[int]int)
	presum := make([]int, len(hours)+1)
	for i := range hours {
		presum[i+1] = presum[i] + hours[i]
	}
	for i := range presum {
		s := presum[i]
		if _, ok := m[s]; !ok {
			m[s] = i
		}

		if s > 0 {
			ans = max(ans, i)
		} else {
			j, ok := m[s-1]
			if ok {
				ans = max(ans, i-j)
			}
		}
	}

	return ans
}

func Test_longestWPI(t *testing.T) {
	hours := []int{6, 6, 9}
	//        [  -1,-1, 1]
	// presum [0,-1,-2,-1]
	//         0 1   2  3
	fmt.Println(longestWPI(hours))
}
