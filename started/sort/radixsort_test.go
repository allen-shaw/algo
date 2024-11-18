package sort

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func radixSort(nums []int) {
	minNum, maxNum := 0, 0
	for _, num := range nums {
		minNum = int(math.Min(float64(minNum), float64(num)))
		maxNum = int(math.Max(float64(maxNum), float64(num)))
	}

	offset := 0
	if minNum < 0 {
		offset = -minNum
	}
	for i := range nums {
		nums[i] += offset
	}

	maxRadix := 0
	for maxNum > 0 {
		maxNum /= 10
		maxRadix++
	}

	for k := 0; k < maxRadix; k++ {
		countSort(nums, k)
	}

	for i := range nums {
		nums[i] -= offset
	}
}

func countSort(nums []int, k int) {
	count := make([]int, 10)

	for _, num := range nums {
		digit := (num / int(math.Pow10(k))) % 10
		count[digit]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		digit := (nums[i] / int(math.Pow10(k))) % 10
		sorted[count[digit]-1] = nums[i]
		count[digit]--
	}

	copy(nums, sorted)
}

func Test_countSort(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case1", []int{10, 2, 1, 21, 134, 20}, []int{10, 20, 1, 21, 2, 134}},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			countSort(c.input, 0)
			assert.Equal(t, c.expect, c.input)
		})
	}
}

func Test_radixSort(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case1", []int{10, 2, 1, 21, 134, 20}, []int{1, 2, 10, 20, 21, 134}},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			radixSort(c.input)
			assert.Equal(t, c.expect, c.input)
		})
	}
}
