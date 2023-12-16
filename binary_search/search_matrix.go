package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		n := matrix[i][j]

		if n == target {
			return true
		} else if n < target {
			i++
		} else if n > target {
			j--
		}
	}

	return false
}
