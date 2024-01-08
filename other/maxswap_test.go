package other

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximumSwap(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	fmt.Println("nums", nums)

	ans := 0
	n := len(nums)
	idx := make([]int, n)

	for i, j := 0, 0; i < n; i++ {
		if nums[i] > nums[j] {
			j = i
		}
		idx[i] = j
	}

	fmt.Println("idx", idx)

	for i := n - 1; i >= 0; i-- {
		if nums[idx[i]] != nums[i] {
			nums[idx[i]], nums[i] = nums[i], nums[idx[i]]
			break
		}
	}
	fmt.Println("nums", nums)

	for i := n - 1; i >= 0; i-- {
		ans = ans*10 + nums[i]
	}
	return ans
}

func TestMaximumSwap(t *testing.T) {
	var testsets = []struct {
		num int
		out int
	}{
		{2736, 7236},
	}

	for _, tt := range testsets {
		t.Run("maximumSwap", func(t *testing.T) {
			ans := maximumSwap(tt.num)
			assert.Equal(t, tt.out, ans)
		})
	}
}
