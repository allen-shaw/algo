package backtrack

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := subsets(nums)
	fmt.Println(ans)
}

func TestSubsetsDup(t *testing.T) {
	nums := []int{1, 2, 2}
	ans := subsetsWithDup(nums)
	fmt.Println(ans)
}
