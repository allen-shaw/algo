package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func firstMissingPositive(nums []int) int {
	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}

	for i := 0; i < len(nums); i++ {
		// nums[i]应该放到 nums[i]-1的位置上
		for nums[i] < len(nums) && nums[i] > 0 && nums[i] != nums[nums[i]-1] {
			swap(nums[i]-1, i)
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func TestFirstMissingPostive(t *testing.T) {
	var testsets = []struct {
		nums []int
		ans  int
	}{
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}

	for _, tt := range testsets {
		t.Run("firstMissingPositive", func(t *testing.T) {
			out := firstMissingPositive(tt.nums)
			assert.Equal(t, tt.ans, out)
		})
	}
}
