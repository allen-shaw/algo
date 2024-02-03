package doublepointer

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	ans := removeDuplicates(nums)
	t.Log(ans)
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ans := removeDuplicates(nums)
	t.Log(ans)
}

func removeDuplicatesk(nums []int) int {
	return rdk(nums, 2)
}

func rdk(nums []int, k int) int {
	if len(nums) <= k {
		return len(nums)
	}

	slow, fast := k, k
	for fast < len(nums) {
		if nums[slow-k] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

func TestRemoveDuplicatesK(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	ans := removeDuplicatesk(nums)
	fmt.Println(nums, ans)

}
