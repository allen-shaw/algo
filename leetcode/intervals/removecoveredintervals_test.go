package intervals

import (
	"fmt"
	"testing"
)

func TestRemoveCoveredIntervals(t *testing.T) {
	intervals := [][]int{{1, 4}, {3, 6}, {2, 8}}
	ans := removeCoveredIntervals(intervals)
	fmt.Println(ans)
}

// [[34335,39239],[15875,91969],[29673,66453],[53548,69161],[40618,93111]]
func TestRemoveCoveredIntervals2(t *testing.T) {
	intervals := [][]int{{34335, 39239}, {15875, 91969}, {29673, 66453}, {53548, 69161}, {40618, 93111}}
	ans := removeCoveredIntervals(intervals)
	fmt.Println(ans)
}
