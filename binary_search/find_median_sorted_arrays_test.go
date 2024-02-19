package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 合并数组
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	// m == 0
	if m == 0 {
		if n%2 == 0 {
			return float64((nums2[n/2] + nums2[n/2-1])) / 2
		} else {
			return float64(nums2[n/2])
		}
	}

	// n == 0
	if n == 0 {
		if m%2 == 0 {
			return float64((nums1[m/2] + nums1[m/2-1])) / 2
		} else {
			return float64(nums1[m/2])
		}
	}

	// m != 0 && n != 0, 合并到nums
	nums := make([]int, m+n)

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		nums[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		nums[k] = nums2[j]
		j++
		k++
	}

	l := len(nums)
	if l%2 == 0 {
		return float64((nums[l/2-1] + nums[l/2])) / 2
	} else {
		return float64(nums[l%2])
	}
}

// 不合并数组
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	left, right := -1, -1

	for k := 0; k < (m+n)/2+1; k++ {
		left = right
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				right = nums1[i]
				i++
			} else {
				right = nums2[j]
				j++
			}
		} else if i < m {
			right = nums1[i]
			i++
		} else if j < n {
			right = nums2[j]
			j++
		}
	}

	if (m+n)%2 == 0 {
		return float64(left+right) / 2
	}
	return float64(right)
}

func TestFindMedianSortedArrays(t *testing.T) {
	var testsets = []struct {
		nums1 []int
		nums2 []int
		ans   float64
	}{
		{[]int{1, 3}, []int{2}, float64(2)},
		{[]int{1, 2}, []int{3, 4}, float64(2.5)},
	}

	// for _, tt := range testsets {
	// 	t.Run("findMedianSortedArrays1", func(t *testing.T) {
	// 		out := findMedianSortedArrays1(tt.nums1, tt.nums2)
	// 		assert.Equal(t, tt.ans, out)
	// 	})
	// }

	// for _, tt := range testsets {
	// 	t.Run("findMedianSortedArrays2", func(t *testing.T) {
	// 		out := findMedianSortedArrays2(tt.nums1, tt.nums2)
	// 		assert.Equal(t, tt.ans, out)
	// 	})
	// }

	for _, tt := range testsets {
		t.Run("findMedianSortedArrays3", func(t *testing.T) {
			out := findMedianSortedArrays3(tt.nums1, tt.nums2)
			assert.Equal(t, tt.ans, out)
		})
	}
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 0 {
		return float64(findKth(nums1, nums2, (m+n)/2)+findKth(nums1, nums2, (m+n)/2+1)) / 2
	}
	return float64(findKth(nums1, nums2, (m+n)/2+1))
}

func findKth(nums1, nums2 []int, k int) int {
	fmt.Println(nums1, nums2, k)
	m, n := len(nums1), len(nums2)
	// 保证nums1 长度小于 nums2
	if m > n {
		return findKth(nums2, nums1, k)
	}

	if m == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	i := k/2 - 1
	fmt.Println("k:", k, "i:", i)
	if i < m {
		if nums1[i] < nums2[i] {
			return findKth(nums1[i+1:], nums2, k-(i+1))
		} else {
			return findKth(nums1, nums2[i+1:], k-(i+1))
		}
	} else {
		if nums1[m-1] < nums2[i] {
			return nums2[k-1-m]
		} else {
			return findKth(nums1, nums2[i+1:], k-(i+1))
		}
	}
}

func TestFindKth(t *testing.T) {
	nums1 := []int{4}
	nums2 := []int{1, 2, 3, 5, 6}
	ans := findKth(nums1, nums2, 3)
	ans2 := findKth(nums1, nums2, 4)

	fmt.Println("ans:", ans)
	fmt.Println("ans2:", ans2)
}
