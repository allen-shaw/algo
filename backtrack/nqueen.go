package backtrack

func solveNQueens(n int) [][]string {
	return newNQueenSolution().solveNQueens(n)
}

type nQueenSolution struct {
	res [][]string
}

func newNQueenSolution() *nQueenSolution {
	s := &nQueenSolution{}
	s.res = make([][]string, 0)
	return s
}

func (s *nQueenSolution) solveNQueens(n int) [][]string {
	board := NewMatrix[byte](n, '.')
	row := 0

	s.backtrack(board, n, row)
	return s.res
}

func (s *nQueenSolution) backtrack(board [][]byte, n int, row int) {
	if row == n {
		s.res = append(s.res, s.toRes(board))
		return
	}

	for col := 0; col < n; col++ {
		if !s.isValid(board, row, col) {
			continue
		}

		board[row][col] = 'Q'
		s.backtrack(board, n, row+1)
		board[row][col] = '.'
	}
}

func (s *nQueenSolution) isValid(board [][]byte, row, col int) bool {
	// 检查同一列
	for i := 0; i <= row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检查左上角对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	// 检查右上角对角线
	for i, j := row-1, col+1; i >= 0 && j < len(board[i]); {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	return true
}

func (s *nQueenSolution) toRes(board [][]byte) []string {
	res := make([]string, 0, len(board))

	for i := 0; i < len(board); i++ {
		res = append(res, string(board[i]))
	}

	return res
}
