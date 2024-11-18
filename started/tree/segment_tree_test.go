package tree

import (
	"fmt"
	"testing"
)

func TestSuffixMin(t *testing.T) {
	nums := []int{3, 1, 4, 2}
	suffixMin := make([]int, len(nums))

	len := len(nums)
	suffixMin[len-1] = nums[len-1]
	for i := len - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], nums[i])
	}

	fmt.Println(suffixMin)
}
