package leetcode

func setZeroes(matrix [][]int) {
	row0, col0 := false, false
	h, w := len(matrix), len(matrix[0])

	for i := 0; i < h; i++ {
		val := matrix[0][i]
		if val == 0 {
			col0 = true
			break
		}
	}
	for i := 0; i < w; i++ {
		val := matrix[0][i]
		if val == 0 {
			row0 = true
			break
		}
	}
	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < w; i++ {
		val := matrix[0][i]
		if val == 0 {
			for j := 0; j < h; j++ {
				matrix[j][i] = 0
			}
		}
	}
	for i := 1; i < h; i++ {
		val := matrix[i][0]
		if val == 0 {
			for j := 0; j < w; j++ {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for i := 0; i < w; i++ {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for i := 0; i < h; i++ {
			matrix[i][0] = 0
		}
	}
}
