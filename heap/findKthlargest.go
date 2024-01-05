package heap

import (
	"container/heap"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	// return findKthLargestHeap(nums, k)
	return findKthLargestQSort(nums, k)
}

func findKthLargestHeap(nums []int, k int) int {
	h := &hp{}

	for i := 0; i < len(nums); i++ {
		if len(*h) < k {
			heap.Push(h, nums[i])
		} else if len(*h) == k && (*h)[0] < nums[i] {
			heap.Push(h, nums[i])
		}
	}
	return (*h)[0]
}

type hp []int

// Len implements heap.Interface.
func (h *hp) Len() int {
	return len(*h)
}

// 小顶堆
func (h *hp) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Pop implements heap.Interface.
func (h *hp) Pop() any {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

// Push implements heap.Interface.
func (h *hp) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements heap.Interface.
func (h hp) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func findKthLargestQSort(nums []int, k int) int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	idx := r.Intn(len(nums))
	nums[idx], nums[0] = nums[0], nums[idx]

	k = len(nums) - k
	return findKthLargestQ(nums, k, 0, len(nums)-1)
}

func findKthLargestQ(nums []int, k, lo, hi int) int {
	pivot := partition(nums, lo, hi)
	if pivot == k {
		return nums[pivot]
	} else if pivot > k {
		return findKthLargestQ(nums, k, lo, pivot-1)
	} else if pivot < k {
		return findKthLargestQ(nums, k, pivot+1, hi)
	}
	return -1
}

func partition(nums []int, lo, hi int) int {
	p := nums[lo]
	i, j := lo+1, hi

	// [lo, i) p (j, hi]
	for i <= j {
		for i < hi && nums[i] <= p {
			i++
		}
		for lo < j && nums[j] >= p {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
