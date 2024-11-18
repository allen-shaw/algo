package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}

func Test_insertSort(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case1", []int{5, 1, 4, 5, 7, 2}, []int{1, 2, 4, 5, 5, 7}},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			insertSort(c.input)
			assert.Equal(t, c.expect, c.input)
		})
	}
}
