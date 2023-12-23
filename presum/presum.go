package presum

func NewMatrix(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{
		preSum: NewMatrix(len(matrix)+1, len(matrix[0])+1),
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			nm.preSum[i+1][j+1] = nm.preSum[i+1][j] + nm.preSum[i][j+1] - nm.preSum[i][j] + matrix[i][j]
		}
	}

	return nm
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.preSum[row2+1][col2+1] - nm.preSum[row2+1][col1+1] - nm.preSum[row1+1][col2+1] + nm.preSum[row1+1][col1+1]
}
