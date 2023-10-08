package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	left, right := 0, len(matrix)*len(matrix[0])-1

	for left <= right {
		mid := left + (right-left)/2
		m := value(matrix, mid)

		// fmt.Printf("left=%v, right=%v, mid=%v, m=%v, target=%v \n", left, right, mid, m, target)

		if m == target {
			return true
		} else if m < target {
			left = mid + 1
		} else if m > target {
			right = mid - 1
		}
	}

	return false
}

func value(matrix [][]int, id int) int {
	width := len(matrix[0])
	row := id / width
	col := id % width
	return matrix[row][col]
}
