package matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		num := matrix[i][j]
		if num == target {
			return true
		} else if num < target {
			i++
		} else if num > target {
			j--
		}
	}

	return false
}
