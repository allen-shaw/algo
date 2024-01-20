package heap

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testsets := []struct {
		nums []int
		k    int
		ans  []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		// {[]int{1}, 1, []int{1}},
	}

	for _, tt := range testsets {
		t.Run("topKFrequent", func(t *testing.T) {
			ans := topKFrequent(tt.nums, tt.k)
			fmt.Println(ans)
		})
	}
}

func TestTKParition(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	m := make(map[int]int, 0)
	for _, n := range nums {
		m[n]++
	}

	datas := make([][]int, 0, len(m))
	for n, f := range m {
		datas = append(datas, []int{n, f})
	}

	p := tkfParition(datas, 0, len(datas)-1)
	fmt.Println(datas, p)
}
