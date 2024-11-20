package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(numbers []int, target int) []int {
	// sort.Ints(numbers)

	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left+1, right+1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		}
		if numbers[left]+numbers[right] > target {
			right--
		}
	}

	return nil
}

func Test_twoSum(t *testing.T) {
	cases := []struct {
		name  string
		input struct {
			nums   []int
			target int
		}
		expect []int
	}{
		{"case1", struct {
			nums   []int
			target int
		}{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := twoSum(c.input.nums, c.input.target)
			assert.Equal(t, c.expect, out)
		})
	}
}
