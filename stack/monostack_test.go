package stack

import (
	"fmt"
	"testing"
)

func largestRectangleArea(heights []int) int {
	area := 0
	stack := make([]int, 0, len(heights)+2)

	// 增加前后哨兵
	heights = append(append([]int{0}, heights...), 0)
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		for heights[stack[len(stack)-1]] > heights[i] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] // stack pop
			area = max(area, h*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return area
}

func TestLargestRectangleArea(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	ans := largestRectangleArea(heights)
	fmt.Println(ans)
}
