package sort

import "testing"

func TestLargestNumber(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	ans := largestNumber(nums) //9534330
	t.Log(ans)
}

func TestLargestNumber2(t *testing.T) {
	nums := []int{0, 0}
	ans := largestNumber(nums) //0
	t.Log(ans)
}
