package other

import (
	"fmt"
	"math"
	"testing"
)

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i >= 0 {
		j := len(nums) - 1
		for ; j > i+1; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	k1, k2 := i+1, len(nums)-1
	for k1 < k2 {
		nums[k1], nums[k2] = nums[k2], nums[k1]
		k1++
		k2--
	}
}

func TestNextPermutation(t *testing.T) {
	// nums := []int{1, 7, 3, 6, 2, 0}
	// nums := []int{1, 7, 3, 6, 5, 1}
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextGreaterElement(n int) int {
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}

	fmt.Println(nums)

	i := 0
	for ; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			break
		}
	}

	if i == len(nums)-1 {
		return -1
	}

	j := 0
	for ; j < i; j++ {
		if nums[j] > nums[i+1] {
			break
		}
	}
	nums[j], nums[i+1] = nums[i+1], nums[j]

	for k1, k2 := 0, i; k1 < k2; k1, k2 = k1+1, k2-1 {
		nums[k1], nums[k2] = nums[k2], nums[k1]
	}

	var ans int64
	for i := len(nums) - 1; i >= 0; i-- {
		ans = ans*10 + int64(nums[i])
	}

	if ans > math.MaxInt32 {
		return -1
	}
	return int(ans)
}

func TestNextGreaterElement(t *testing.T) {
	n := 21
	ans := nextGreaterElement(n)
	fmt.Println(ans)
}
