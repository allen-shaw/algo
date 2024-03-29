package intervals

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ans := merge(intervals)
	fmt.Println(ans)
}

func TestMerge2(t *testing.T) {
	intervals := [][]int{{1, 4}, {4, 5}}
	ans := merge(intervals)
	fmt.Println(ans)
}
