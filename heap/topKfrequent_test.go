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
		{[]int{1}, 1, []int{1}},
	}

	for _, tt := range testsets {
		t.Run("topKFrequent", func(t *testing.T) {
			ans := topKFrequent(tt.nums, tt.k)
			fmt.Println(ans)
		})
	}
}
