package numarray2

import (
	"fmt"
	"testing"
)

type NumArray struct {
	nums []int
	tree []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{nums: make([]int, len(nums)), tree: make([]int, len(nums)+1)}
	for i, num := range nums {
		arr.Update(i, num)
	}
	return arr
}

func (a *NumArray) Update(index int, val int) {
	delta := val - a.nums[index]
	a.nums[index] = val

	i := index + 1
	for i < len(a.tree) {
		a.tree[i] += delta
		i += lowbit(i)
	}
}

func (a *NumArray) SumRange(left int, right int) int {
	return a.sum(right+1) - a.sum(left)
}

func (a *NumArray) sum(index int) int {
	sum := 0
	for index > 0 {
		sum += a.tree[index]
		index -= lowbit(index)
	}
	return sum
}

func lowbit(i int) int {
	return i & (-i)
}

func TestNumArray(t *testing.T) {
	nums := []int{1, 3, 5}
	narr := Constructor(nums)
	fmt.Println(narr)
	fmt.Println(narr.SumRange(0, 2))
	narr.Update(1, 2)
	fmt.Println(narr.SumRange(0, 2))
}
