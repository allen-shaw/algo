package leetcode

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0
	for l < r {
		area := calcArea(height, l, r)
		maxArea = max(area, maxArea)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}

func calcArea(nums []int, l, r int) int {
	h := min(nums[l], nums[r])
	return h * (r - l)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
