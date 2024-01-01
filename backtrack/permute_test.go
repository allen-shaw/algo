package backtrack

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := permute(nums)
	t.Log(ans)
}

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	ans := permuteUnique(nums)
	fmt.Println(ans)
}
