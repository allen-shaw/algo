package leetcode

import (
	"sort"
)

func threeSum(nums []int, k int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := make([][]int, 0)

	for i, num := range nums {
		if i != 0 && num == nums[i-1] {
			continue
		}

		numlists := twoSum(nums, k-num, i)
		if len(numlists) != 0 {
			for i := range numlists {
				numlists[i] = append(numlists[i], num)
			}
			ans = append(ans, numlists...)
		}
	}

	return ans
}

func twoSum(nums []int, k int, curIdx int) [][]int {
	ans := make([][]int, 0)
	i, j := curIdx+1, len(nums)-1
	for i < j {
		if i != curIdx+1 && nums[i] == nums[i-1] {
			i++
			continue
		}
		if j != len(nums)-1 && nums[j] == nums[j+1] {
			j--
			continue
		}

		if nums[i]+nums[j] < k {
			i++
		} else if nums[i]+nums[j] > k {
			j--
		} else {
			res := []int{nums[i], nums[j]}
			ans = append(ans, res)
			i++
			j--
		}
	}
	// fmt.Println("k:", k, "n0", curIdx, nums[curIdx], "ans", ans)
	return ans
}
