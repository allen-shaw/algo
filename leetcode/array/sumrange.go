package array

type NumArray struct {
	preSum []int
}

func NewNumArray(nums []int) NumArray {
	na := NumArray{
		preSum: make([]int, len(nums)+1),
	}

	for i, num := range nums {
		na.preSum[i+1] = na.preSum[i] + num
	}

	return na
}

func (a *NumArray) SumRange(left int, right int) int {
	return a.preSum[right+1] - a.preSum[left]
}
