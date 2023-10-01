package stack

import (
	"fmt"
	"testing"
)

func TestGenMonotonicStack(t *testing.T) {
	nums := []int{2, 1, 2, 4, 3}
	res := genMonotonicStack(nums)
	fmt.Println(res)
}
