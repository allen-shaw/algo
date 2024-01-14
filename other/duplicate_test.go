package other

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findErrorNums(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}

	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			return []int{nums[i], i + 1}
		}
	}
	return nil
}

func TestFindErrorNums(t *testing.T) {
	var testsets = []struct {
		nums []int
		out  []int
	}{
		{
			[]int{1, 2, 2, 4},
			[]int{2, 3},
		},
	}

	for _, tt := range testsets {
		t.Run("findErrorNums", func(t *testing.T) {
			ans := findErrorNums(tt.nums)
			assert.Equal(t, tt.out, ans)
		})
	}
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	fmt.Println(nums)

	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func TestFindDisappearedNumbers(t *testing.T) {
	var testsets = []struct {
		nums []int
		out  []int
	}{
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
	}

	for _, tt := range testsets {
		t.Run("findDisappearedNumbers", func(t *testing.T) {
			ans := findDisappearedNumbers(tt.nums)
			assert.Equal(t, tt.out, ans)
		})
	}
}
