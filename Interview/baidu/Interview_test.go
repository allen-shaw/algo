package baidu

import (
	"math"
	"sort"
)

func threeSum(nums []int, sum int) int {
	sort.Ints(nums)
	n := len(nums)

	minAbs := math.MaxInt
	maSum := 0

	for i, num := range nums {
		for j, k := i+1, n-1; j < k; {
			s := num + nums[j] + nums[k]
			if s == sum {
				return s
			} else if s < sum {
				j++
			} else if s > sum {
				k--
			}
			if abs(s-sum) < minAbs {
				minAbs = abs(s - sum)
				maSum = s
			}
		}
	}

	return maSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
