package array

import (
	"fmt"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	l := len(nums1) - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[l] = nums1[i]
			i--
		} else {
			nums1[l] = nums2[j]
			j--
		}
		l--
	}

	for j >= 0 {
		nums1[l] = nums2[j]
		j--
		l--
	}
}

func TestMerge(t *testing.T) {
	var testsets = []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
		},
	}

	for _, tt := range testsets {
		t.Run("merge", func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			fmt.Println(tt.nums1)
		})
	}

}
