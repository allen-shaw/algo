package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(numbers []int, target int) []int {
	// sort.Ints(numbers)

	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		}
		if numbers[left]+numbers[right] > target {
			right--
		}
	}

	return nil
}

func Test_twoSum(t *testing.T) {
	cases := []struct {
		name  string
		input struct {
			nums   []int
			target int
		}
		expect []int
	}{
		{"case1", struct {
			nums   []int
			target int
		}{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := twoSum(c.input.nums, c.input.target)
			assert.Equal(t, c.expect, out)
		})
	}
}

func removeDuplicates(nums []int) int {
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func Test_removeDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	last := removeDuplicates(nums)
	fmt.Println(nums[:last])
}

// 0, 1, 3, 0, 4, 0, 4, 2
//                s     f

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func Test_removeElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	last := removeElement(nums, val)
	fmt.Println(nums[:last])
}

func reverseWords(s string) string {
	ss := []byte(s)
	reverse(ss, 0, len(s)-1)
	left := 0
	for right := 0; right < len(ss); right++ {
		if ss[right] == ' ' {
			reverse(ss, left, right-1)
			left = right + 1
		}
	}
	return string(ss)
}

func reverse(ss []byte, i, j int) {
	for l, r := i, j; l < r; l, r = l+1, r-1 {
		ss[l], ss[r] = ss[r], ss[l]
	}
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	ans := make([]int, 0, m*n)
	up, down, left, right := 0, m-1, 0, n-1

	for len(ans) < m*n {
		if up <= down {
			for i := left; i <= right; i++ {
				ans = append(ans, matrix[up][i])
			}
			up++
		}
		if left <= right {
			for i := up; i <= down; i++ {
				ans = append(ans, matrix[i][right])
			}
			right--
		}
		if up <= down {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[down][i])
			}
			down--
		}
		if left <= right {
			for i := down; i >= up; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}

	return ans
}

func Test_spiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	out := spiralOrder(matrix)
	fmt.Println(out)
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	up, down, left, right := 0, n-1, 0, n-1

	i := 1
	for i <= n*n {
		if up <= down {
			for j := left; j <= right; j++ {
				matrix[up][j] = i
				i++
			}
			fmt.Println(up, down, left, right, "M1", matrix)
			up++
		}
		if left <= right {
			for j := up; j <= down; j++ {
				matrix[j][right] = i
				i++
			}
			fmt.Println(up, down, left, right, "M2", matrix)
			right--
		}
		if up <= down {
			for j := right; j >= left; j-- {
				matrix[down][j] = i
				i++
			}
			fmt.Println(up, down, left, right, "M3", matrix)
			down--
		}
		if left <= right {
			for j := down; j >= up; j-- {
				matrix[j][left] = i
				i++
			}
			fmt.Println(up, down, left, right, "M4", matrix)
			left++
		}
	}

	return matrix
}

func Test_generateMatrix(t *testing.T) {
	matrix := generateMatrix(4)
	for _, line := range matrix {
		fmt.Println(line)
	}
}
