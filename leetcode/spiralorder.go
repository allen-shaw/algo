package leetcode

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix)-1, len(matrix[0])-1
	up, low, left, right := 0, m, 0, n
	ans := make([]int, 0)
	for len(ans) < len(matrix)*len(matrix[0]) {
		if up <= low {
			for i := left; i <= right; i++ {
				ans = append(ans, matrix[up][i])
			}
			up++
		}

		if left <= right {
			for i := up; i <= low; i++ {
				ans = append(ans, matrix[i][right])
			}
			right--
		}

		if up <= low {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[low][i])
			}
			low--
		}

		if left <= right {
			for i := low; i >= up; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}

	return ans
}
