package backtrack

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://blog.csdn.net/qq_36282995/article/details/126078742

func findMaxLessN(arr []int, n int) int {
	sort.Ints(arr) // 从小到大排序
	nums := parseN(n)

	return dfs(arr, nums, true, 0)
}

func parseN(n int) []int {
	res := make([]int, 0)
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	fmt.Println(res)
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

// arr:可使用数字的数组
// nums:目标数字对应的整数数组
// equal:之前是否每一位都使用了相等的值
// index:当前考虑到了第几位，从0开始
// 返回第index位后可组成的最大的数字，包括第index位
func dfs(arr []int, nums []int, equal bool, index int) int {
	if index == len(nums) {
		return 0
	}

	if equal {
		for i := len(arr) - 1; i >= 0; i-- {
			if index == len(nums)-1 { // 最后一位
				if arr[i] < nums[index] {
					temp := dfs(arr, nums, arr[i] == nums[index], index+1)
					if temp == -1 {
						continue
					}
					return arr[i]*int(math.Pow10(len(nums)-index-1)) + temp
				}
			} else {
				if arr[i] <= nums[index] {
					temp := dfs(arr, nums, arr[i] == nums[index], index+1)
					if temp == -1 {
						continue
					}
					return arr[i]*int(math.Pow10(len(nums)-index-1)) + temp
				}
			}
		}
		// 如果上面的for循环走完都没有return的话，那就说明这一层所有的数字都不满足要求
		// 如果是第一个数字都无法满足的话，那就直接从第二位开始，将preEq置为False递归
		if index == 0 {
			return dfs(arr, nums, false, index+1)
		}
		return -1
	} else {
		// 如果前面不是选用了相等的数字，那就直接每次挑最大的数字就行了
		return arr[len(arr)-1]*int(math.Pow10(len(nums)-index-1)) + dfs(arr, nums, false, index+1)
	}
}

func TestFindMaxLessN(t *testing.T) {
	var testsets = []struct {
		arr []int
		n   int
		out int
	}{
		{[]int{1, 2, 9, 4}, 2533, 2499},
		{[]int{1, 2, 5, 4}, 2543, 2542},
		{[]int{1, 2, 5, 4}, 2541, 2525},
		{[]int{1, 2, 9, 4}, 2111, 1999},
		{[]int{5, 9}, 5555, 999},
	}
	for _, tt := range testsets {
		t.Run("findMaxLessN", func(t *testing.T) {
			ans := findMaxLessN(tt.arr, tt.n)
			assert.Equal(t, tt.out, ans)
		})
	}
}
