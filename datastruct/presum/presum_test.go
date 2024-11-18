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
