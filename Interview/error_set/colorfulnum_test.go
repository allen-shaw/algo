package errorset

import (
	"fmt"
	"testing"
)

func isColorful(num int) bool {
	nums := make([]int, 0)
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	n := len(nums)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	fmt.Println("nums: ", nums)

	set := make(map[int]struct{})

	var dfs func(idx int, path []int)
	ans := true
	dfs = func(idx int, path []int) {
		if !ans {
			return
		}
		prod := product(path)
		fmt.Println("idx:", idx, "path:", path, "prod:", prod)

		if _, ok := set[prod]; ok {
			ans = false
			return
		}
		set[prod] = struct{}{}

		for i := idx; i < n; i++ {
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, make([]int, 0, n))

	return ans
}

func product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prod := 1
	for _, num := range nums {
		prod *= num
	}
	return prod
}

func Test_isColorful(t *testing.T) {
	fmt.Println(isColorful(3245))
	fmt.Println(isColorful(326))
}
