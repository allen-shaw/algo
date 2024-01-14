package heap

import (
	"fmt"
)

var (
	topK []int
)

func topKFrequent(nums []int, k int) []int {
	topK = make([]int, 0, k)
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		m[n]++
	}

	freqs := make([][]int, 0, len(m))
	for i, f := range m {
		freqs = append(freqs, []int{i, f})
	}

	lo, hi := 0, len(freqs)-1
	for lo < hi {
		p := partitionFreq(freqs, lo, hi)
		if p <= k {
			for i := lo; i <= p; i++ {
				topK = append(topK, freqs[i][0])
			}
			lo = p + 1
		} else {
			hi = p - 1
		}
	}

	return topK
}

func topKFrequentQSort(nums [][]int, k, lo, hi int) {
	fmt.Println("nums:", nums, "k:", k, "lo:", lo, "hi", hi)
	if lo >= hi {
		return
	}

	p := partitionFreq(nums, lo, hi)
	fmt.Println("nums:", nums, "p:", p, "k:", k)
	if p <= k {
		for i := lo; i <= p; i++ {
			topK = append(topK, nums[i][0])
		}
		topKFrequentQSort(nums, k-(p-lo)+1, p+1, hi)
	} else {
		topKFrequentQSort(nums, k, lo, p-1)
	}

}

func partitionFreq(nums [][]int, lo, hi int) int {
	// 	r := rand.Intn(hi-lo) + lo
	// 	nums[lo], nums[r] = nums[r], nums[lo]

	p := nums[lo][1]
	i, j := lo+1, hi

	for i <= j {
		for i < hi && nums[i][1] >= p {
			i++
		}
		for j > lo && nums[j][1] <= p {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j], nums[lo] = nums[lo], nums[j]
	return j
}
