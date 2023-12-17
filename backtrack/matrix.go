package backtrack

func NewMatrix[T any](size int, init T) [][]T {
	m := make([][]T, size)
	for i := 0; i < size; i++ {
		m[i] = make([]T, size)
		for j := 0; j < size; j++ {
			m[i][j] = init
		}
	}
	return m
}
