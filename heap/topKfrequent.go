package heap

import (
	"fmt"
	"math/rand"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, n := range nums {
		m[n]++
	}

	datas := make([][]int, 0, len(m))
	for n, f := range m {
		datas = append(datas, []int{n, f})
	}

	fmt.Println("datas:", datas)

	lo, hi := 0, len(datas)-1
	for lo < hi {
		p := tkfParition(datas, lo, hi)
		if p == k {
			break
		} else if p < k {
			lo = p + 1
		} else { // p > k
			hi = p - 1
		}
	}

	fmt.Println("after data:", datas)

	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, datas[i][0])
	}

	return ans
}

func tkfParition(nums [][]int, lo, hi int) int {
	r := rand.Intn(hi-lo) + lo
	nums[lo], nums[r] = nums[r], nums[lo]

	p := nums[lo][1]

	i, j := lo, hi
	for i <= j {
		for i <= hi && nums[i][1] >= p {
			i++
		}
		for j >= lo && nums[j][1] < p {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
