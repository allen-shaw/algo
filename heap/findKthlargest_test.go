package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	var testsets = []struct {
		nums []int
		k    int
		out  int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{-1, 2, 0}, 2, 0},
	}

	for _, tt := range testsets {
		t.Run("findKthLargest", func(t *testing.T) {
			ans := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.out, ans)
		})
	}
}

func TestPartition(t *testing.T) {
	var testsets = []struct {
		nums []int
		out  int
	}{
		{[]int{5, 3, 6, 7, 2}, 2},
		{[]int{5, 3, 3, 3, 6}, 3},
		{[]int{4, 5, 5, 5, 2, 3, 6}, 0},
		{[]int{3, 2, 1, 5, 6, 4}, 0},
		{[]int{3, 3, 3, 3, 3, 3, 3, 3, 3}, 0},
	}
	for _, tt := range testsets {
		t.Run("partition", func(t *testing.T) {
			ans := partition(tt.nums, 0, len(tt.nums)-1)
			fmt.Println(ans, tt.nums)
		})
	}
}

func TestFindKthLargestQ(t *testing.T) {
	var testsets = []struct {
		nums []int
		k    int
		out  int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{-1, 2, 0}, 2, 0},
		{[]int{3, 3, 3, 3, 3, 3, 3, 3, 3}, 1, 3},
	}

	for _, tt := range testsets {
		t.Run("findKthLargest", func(t *testing.T) {
			ans := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.out, ans)
		})
	}
}

func TestQSort(t *testing.T) {
	var testsets = []struct {
		nums []int
		k    int
		out  int
	}{
		// {[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		// {[]int{-1, 2, 0}, 2, 0},
		// {[]int{3, 3, 3, 3, 3, 3, 3, 3, 3}, 1, 3},
	}
	for _, tt := range testsets {
		t.Run("qsort", func(t *testing.T) {
			qosrt(tt.nums, tt.k, 0, len(tt.nums)-1)
			fmt.Println(tt.nums)
		})
	}
}

func qosrt(nums []int, k, lo, hi int) {
	if lo > hi {
		return
	}
	fmt.Printf("nums: %v, k: %v, lo: %v, hi: %v p:", nums, k, lo, hi)
	p := partition(nums, lo, hi)
	fmt.Println(p)
	if p <= k {
		qosrt(nums, k, p+1, hi)
	} else {
		qosrt(nums, k, lo, p-1)
	}
}
