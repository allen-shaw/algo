package grab

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MaximumSum(arr []int) int {
	m := make(map[string][]int)
	for _, num := range arr {
		first, last := parseNum(num)
		key := fmt.Sprintf("%v:%v", first, last)
		m[key] = append(m[key], num)
	}
	ans := -1
	for _, nums := range m {
		if len(nums) < 2 {
			continue
		}
		sort.Ints(nums)
		ans = max(ans, nums[len(nums)-1]+nums[len(nums)-2])
	}
	return ans
}

func parseNum(num int) (int, int) {
	last := num % 10
	first := num
	for num > 0 {
		first = num
		num /= 10
	}
	return first, last
}

func Test_parseNum(t *testing.T) {
	num := 109
	fmt.Println(parseNum(num))
}

func TestMaximumSum(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{130, 191, 200, 10},
			output: 140,
		},
		{
			input:  []int{405, 45, 300, 300},
			output: 600,
		},
		{
			input:  []int{50, 222, 49, 52, 25},
			output: -1,
		},
		{
			input:  []int{30, 909, 3190, 99, 3990, 9009},
			output: 9918,
		},
	}

	for _, tc := range testCases {
		ans := MaximumSum(tc.input)
		assert.Equal(t, tc.output, ans)
	}
}
