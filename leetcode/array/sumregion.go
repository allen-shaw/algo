package array

type NumMatrix struct {
	preSum [][]int
}

func NewNumMatrix(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix)+1)
	preSum[0] = make([]int, len(matrix[0])+1)

	for i, row := range matrix {
		preSum[i+1] = make([]int, len(matrix[0])+1)
		for j, num := range row {
			preSum[i+1][j+1] = preSum[i][j+1] + preSum[i+1][j] - preSum[i][j] + num
		}
	}

	return NumMatrix{
		preSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row2+1][col1] - this.preSum[row1][col2+1] + this.preSum[row1][col1]
}
